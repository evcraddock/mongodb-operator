package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type MongoDbBackupSpec struct {
	BackupLocation string `json:"backupLocation"`
	MongoDbUri 	   string `json:"mongoDbUri"`
	SecretKey      string `json:"secretKey"`
}

// mongoDbBackupStatus defines the observed state of mongoDbBackup
type MongoDbBackupStatus struct {
	Successful bool `json:"successful,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDbBackup is the Schema for the mongodbbackups API
// +k8s:openapi-gen=true
type MongoDbBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MongoDbBackupSpec   `json:"spec,omitempty"`
	Status MongoDbBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongoDbBackupList contains a list of MongoDbBackup
type MongoDbBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDbBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongoDbBackup{}, &MongoDbBackupList{})
}
