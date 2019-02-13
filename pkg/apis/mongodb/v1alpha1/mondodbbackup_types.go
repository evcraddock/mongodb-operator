package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MondoDbBackupSpec defines the desired state of MondoDbBackup
type MondoDbBackupSpec struct {
	BackupLocation string `json:"backupLocation"`
	MongoDbHost    string `json:"mongoDbHost"`
	SecretKey      string `json:"secretKey"`
}

// MondoDbBackupStatus defines the observed state of MondoDbBackup
type MondoDbBackupStatus struct {
	Successful bool `json:"successful,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MondoDbBackup is the Schema for the mondodbbackups API
// +k8s:openapi-gen=true
type MondoDbBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MondoDbBackupSpec   `json:"spec,omitempty"`
	Status MondoDbBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MondoDbBackupList contains a list of MondoDbBackup
type MondoDbBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MondoDbBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MondoDbBackup{}, &MondoDbBackupList{})
}
