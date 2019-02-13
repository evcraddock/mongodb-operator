package mondodbbackup

import (
	mongodbv1alpha1 "github.com/evcraddock/mongodb-operator/pkg/apis/mongodb/v1alpha1"
	batch "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//BackupJob backup job
type BackupJob struct {
	Job *batch.Job
}

// NewBackupJob new backup job
func NewBackupJob(backup *mongodbv1alpha1.MondoDbBackup) BackupJob {
	job := createBatchJob(backup)
	return BackupJob{
		job,
	}
}

//IsCompleted checks for JobComplete status
func (j *BackupJob) IsCompleted(job *batch.Job) bool {
	for _, c := range job.Status.Conditions {
		if c.Type == batch.JobComplete {
			return true
		}
	}

	return false
}

func createBatchJob(backup *mongodbv1alpha1.MondoDbBackup) *batch.Job {
	return &batch.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      backup.Name + "-mongodb-backup",
			Namespace: backup.Namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Job",
			APIVersion: "v1",
		},
		Spec: batch.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "mongodump",
							Image:           "zoov/mongodb-gcs-backup:latest",
							ImagePullPolicy: corev1.PullIfNotPresent,
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "gcs-secrets",
									MountPath: "/secrets/gcp",
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "GCS_BUCKET",
									Value: backup.Spec.BackupLocation,
								},
								{
									Name:  "GCS_KEY_FILE_PATH",
									Value: "/secrets/gcp/key.json",
								},
								{
									Name:  "MONGODB_HOST",
									Value: backup.Spec.MongoDbHost,
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyOnFailure,
					Volumes: []corev1.Volume{
						{
							Name: "gcs-secrets",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: backup.Spec.SecretKey,
								},
							},
						},
					},
				},
			},
		},
	}
}
