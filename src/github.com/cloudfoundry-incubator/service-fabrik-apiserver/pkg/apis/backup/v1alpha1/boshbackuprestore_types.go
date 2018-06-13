//TODO copyright header

package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BOSHBackupRestore
// +k8s:openapi-gen=true
// +resource:path=boshbackuprestores,strategy=BOSHBackupRestoreStrategy
type BOSHBackupRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BOSHBackupRestoreSpec   `json:"spec,omitempty"`
	Status BOSHBackupRestoreStatus `json:"status,omitempty"`
}

// BOSHBackupRestoreSpec defines the desired state of BOSHBackupRestore
type BOSHBackupRestoreSpec struct {
	Options string `json:"options,omitempty"`
}

// BOSHBackupRestoreStatus defines the observed state of BOSHBackupRestore
type BOSHBackupRestoreStatus struct {
	State string `json:"state,omitempty"`
}

// Validate checks that an instance of BOSHBackupRestore is well formed
func (BOSHBackupRestoreStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*backup.BOSHBackupRestore)
	log.Printf("Validating fields for BOSHBackupRestore %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default BOSHBackupRestore field values
func (BOSHBackupRestoreSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*BOSHBackupRestore)
	// set default field values here
	log.Printf("Defaulting fields for BOSHBackupRestore %s\n", obj.Name)
}
