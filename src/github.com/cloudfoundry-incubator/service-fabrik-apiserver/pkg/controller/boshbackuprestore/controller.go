
//TODO copyright header


package boshbackuprestore

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/backup/v1alpha1"
)

// +controller:group=backup,version=v1alpha1,kind=BOSHBackupRestore,resource=boshbackuprestores
type BOSHBackupRestoreControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about BOSHBackupRestore
	lister listers.BOSHBackupRestoreLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *BOSHBackupRestoreControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing boshbackuprestores labels
	c.lister = arguments.GetSharedInformers().Factory.Backup().V1alpha1().BOSHBackupRestores().Lister()
}

// Reconcile handles enqueued messages
func (c *BOSHBackupRestoreControllerImpl) Reconcile(u *v1alpha1.BOSHBackupRestore) error {
	// Implement controller logic here
	log.Printf("Running reconcile BOSHBackupRestore for %s\n", u.Name)
	return nil
}

func (c *BOSHBackupRestoreControllerImpl) Get(namespace, name string) (*v1alpha1.BOSHBackupRestore, error) {
	return c.lister.BOSHBackupRestores(namespace).Get(name)
}
