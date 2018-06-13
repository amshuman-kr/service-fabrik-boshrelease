//TODO copyright header
package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup/v1alpha1"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BOSHBackupRestoresGetter has a method to return a BOSHBackupRestoreInterface.
// A group's client should implement this interface.
type BOSHBackupRestoresGetter interface {
	BOSHBackupRestores(namespace string) BOSHBackupRestoreInterface
}

// BOSHBackupRestoreInterface has methods to work with BOSHBackupRestore resources.
type BOSHBackupRestoreInterface interface {
	Create(*v1alpha1.BOSHBackupRestore) (*v1alpha1.BOSHBackupRestore, error)
	Update(*v1alpha1.BOSHBackupRestore) (*v1alpha1.BOSHBackupRestore, error)
	UpdateStatus(*v1alpha1.BOSHBackupRestore) (*v1alpha1.BOSHBackupRestore, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BOSHBackupRestore, error)
	List(opts v1.ListOptions) (*v1alpha1.BOSHBackupRestoreList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BOSHBackupRestore, err error)
	BOSHBackupRestoreExpansion
}

// bOSHBackupRestores implements BOSHBackupRestoreInterface
type bOSHBackupRestores struct {
	client rest.Interface
	ns     string
}

// newBOSHBackupRestores returns a BOSHBackupRestores
func newBOSHBackupRestores(c *BackupV1alpha1Client, namespace string) *bOSHBackupRestores {
	return &bOSHBackupRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bOSHBackupRestore, and returns the corresponding bOSHBackupRestore object, and an error if there is any.
func (c *bOSHBackupRestores) Get(name string, options v1.GetOptions) (result *v1alpha1.BOSHBackupRestore, err error) {
	result = &v1alpha1.BOSHBackupRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BOSHBackupRestores that match those selectors.
func (c *bOSHBackupRestores) List(opts v1.ListOptions) (result *v1alpha1.BOSHBackupRestoreList, err error) {
	result = &v1alpha1.BOSHBackupRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bOSHBackupRestores.
func (c *bOSHBackupRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a bOSHBackupRestore and creates it.  Returns the server's representation of the bOSHBackupRestore, and an error, if there is any.
func (c *bOSHBackupRestores) Create(bOSHBackupRestore *v1alpha1.BOSHBackupRestore) (result *v1alpha1.BOSHBackupRestore, err error) {
	result = &v1alpha1.BOSHBackupRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		Body(bOSHBackupRestore).
		Do().
		Into(result)
	return
}

// Update takes the representation of a bOSHBackupRestore and updates it. Returns the server's representation of the bOSHBackupRestore, and an error, if there is any.
func (c *bOSHBackupRestores) Update(bOSHBackupRestore *v1alpha1.BOSHBackupRestore) (result *v1alpha1.BOSHBackupRestore, err error) {
	result = &v1alpha1.BOSHBackupRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		Name(bOSHBackupRestore.Name).
		Body(bOSHBackupRestore).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *bOSHBackupRestores) UpdateStatus(bOSHBackupRestore *v1alpha1.BOSHBackupRestore) (result *v1alpha1.BOSHBackupRestore, err error) {
	result = &v1alpha1.BOSHBackupRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		Name(bOSHBackupRestore.Name).
		SubResource("status").
		Body(bOSHBackupRestore).
		Do().
		Into(result)
	return
}

// Delete takes name of the bOSHBackupRestore and deletes it. Returns an error if one occurs.
func (c *bOSHBackupRestores) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bOSHBackupRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("boshbackuprestores").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched bOSHBackupRestore.
func (c *bOSHBackupRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BOSHBackupRestore, err error) {
	result = &v1alpha1.BOSHBackupRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("boshbackuprestores").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
