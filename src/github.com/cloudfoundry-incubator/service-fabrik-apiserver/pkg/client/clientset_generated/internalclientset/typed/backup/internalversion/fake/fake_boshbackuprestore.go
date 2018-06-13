//TODO copyright header
package fake

import (
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBOSHBackupRestores implements BOSHBackupRestoreInterface
type FakeBOSHBackupRestores struct {
	Fake *FakeBackup
	ns   string
}

var boshbackuprestoresResource = schema.GroupVersionResource{Group: "backup.servicefabrik.io", Version: "", Resource: "boshbackuprestores"}

var boshbackuprestoresKind = schema.GroupVersionKind{Group: "backup.servicefabrik.io", Version: "", Kind: "BOSHBackupRestore"}

// Get takes name of the bOSHBackupRestore, and returns the corresponding bOSHBackupRestore object, and an error if there is any.
func (c *FakeBOSHBackupRestores) Get(name string, options v1.GetOptions) (result *backup.BOSHBackupRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(boshbackuprestoresResource, c.ns, name), &backup.BOSHBackupRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.BOSHBackupRestore), err
}

// List takes label and field selectors, and returns the list of BOSHBackupRestores that match those selectors.
func (c *FakeBOSHBackupRestores) List(opts v1.ListOptions) (result *backup.BOSHBackupRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(boshbackuprestoresResource, boshbackuprestoresKind, c.ns, opts), &backup.BOSHBackupRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &backup.BOSHBackupRestoreList{}
	for _, item := range obj.(*backup.BOSHBackupRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bOSHBackupRestores.
func (c *FakeBOSHBackupRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(boshbackuprestoresResource, c.ns, opts))

}

// Create takes the representation of a bOSHBackupRestore and creates it.  Returns the server's representation of the bOSHBackupRestore, and an error, if there is any.
func (c *FakeBOSHBackupRestores) Create(bOSHBackupRestore *backup.BOSHBackupRestore) (result *backup.BOSHBackupRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(boshbackuprestoresResource, c.ns, bOSHBackupRestore), &backup.BOSHBackupRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.BOSHBackupRestore), err
}

// Update takes the representation of a bOSHBackupRestore and updates it. Returns the server's representation of the bOSHBackupRestore, and an error, if there is any.
func (c *FakeBOSHBackupRestores) Update(bOSHBackupRestore *backup.BOSHBackupRestore) (result *backup.BOSHBackupRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(boshbackuprestoresResource, c.ns, bOSHBackupRestore), &backup.BOSHBackupRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.BOSHBackupRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBOSHBackupRestores) UpdateStatus(bOSHBackupRestore *backup.BOSHBackupRestore) (*backup.BOSHBackupRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(boshbackuprestoresResource, "status", c.ns, bOSHBackupRestore), &backup.BOSHBackupRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.BOSHBackupRestore), err
}

// Delete takes name of the bOSHBackupRestore and deletes it. Returns an error if one occurs.
func (c *FakeBOSHBackupRestores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(boshbackuprestoresResource, c.ns, name), &backup.BOSHBackupRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBOSHBackupRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(boshbackuprestoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &backup.BOSHBackupRestoreList{})
	return err
}

// Patch applies the patch and returns the patched bOSHBackupRestore.
func (c *FakeBOSHBackupRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *backup.BOSHBackupRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(boshbackuprestoresResource, c.ns, name, data, subresources...), &backup.BOSHBackupRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.BOSHBackupRestore), err
}
