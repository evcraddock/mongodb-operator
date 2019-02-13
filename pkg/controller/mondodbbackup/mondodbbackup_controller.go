package mondodbbackup

import (
	"context"

	mongodbv1alpha1 "github.com/evcraddock/mongodb-operator/pkg/apis/mongodb/v1alpha1"
	batch "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mondodbbackup")

// Add creates a new MondoDbBackup Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMondoDbBackup{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("mondodbbackup-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &mongodbv1alpha1.MondoDbBackup{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &batch.Job{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &mongodbv1alpha1.MondoDbBackup{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMondoDbBackup{}

// ReconcileMondoDbBackup reconciles a MondoDbBackup object
type ReconcileMondoDbBackup struct {
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MondoDbBackup object and makes changes based on the state read
// and what is in the MondoDbBackup.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMondoDbBackup) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	backup := new(mongodbv1alpha1.MondoDbBackup)

	err := r.client.Get(context.TODO(), request.NamespacedName, backup)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	if backup.Status.Successful {
		return reconcile.Result{}, nil
	}

	backupJob := NewBackupJob(backup)

	found := &batch.Job{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: backupJob.Job.Name, Namespace: backupJob.Job.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Info("creating backup job", "name", backupJob.Job.Name)
		err = r.client.Create(context.TODO(), backupJob.Job)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	ok := backupJob.IsCompleted(found)
	if ok {
		backup.Status.Successful = true
		err = r.client.Status().Update(context.TODO(), backup)

		if err != nil {
			return reconcile.Result{}, err
		}

		log.Info("backup job finished", "name", backup.Name)
		return reconcile.Result{}, nil

	}

	return reconcile.Result{Requeue: true}, nil
}
