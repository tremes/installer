// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSesGetter has a method to return a DNSInterface.
// A group's client should implement this interface.
type DNSesGetter interface {
	DNSes() DNSInterface
}

// DNSInterface has methods to work with DNS resources.
type DNSInterface interface {
	Create(*v1.DNS) (*v1.DNS, error)
	Update(*v1.DNS) (*v1.DNS, error)
	UpdateStatus(*v1.DNS) (*v1.DNS, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.DNS, error)
	List(opts meta_v1.ListOptions) (*v1.DNSList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DNS, err error)
	DNSExpansion
}

// dNSes implements DNSInterface
type dNSes struct {
	client rest.Interface
}

// newDNSes returns a DNSes
func newDNSes(c *ConfigV1Client) *dNSes {
	return &dNSes{
		client: c.RESTClient(),
	}
}

// Get takes name of the dNS, and returns the corresponding dNS object, and an error if there is any.
func (c *dNSes) Get(name string, options meta_v1.GetOptions) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Get().
		Resource("dnses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSes that match those selectors.
func (c *dNSes) List(opts meta_v1.ListOptions) (result *v1.DNSList, err error) {
	result = &v1.DNSList{}
	err = c.client.Get().
		Resource("dnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSes.
func (c *dNSes) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("dnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a dNS and creates it.  Returns the server's representation of the dNS, and an error, if there is any.
func (c *dNSes) Create(dNS *v1.DNS) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Post().
		Resource("dnses").
		Body(dNS).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dNS and updates it. Returns the server's representation of the dNS, and an error, if there is any.
func (c *dNSes) Update(dNS *v1.DNS) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Put().
		Resource("dnses").
		Name(dNS.Name).
		Body(dNS).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dNSes) UpdateStatus(dNS *v1.DNS) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Put().
		Resource("dnses").
		Name(dNS.Name).
		SubResource("status").
		Body(dNS).
		Do().
		Into(result)
	return
}

// Delete takes name of the dNS and deletes it. Returns an error if one occurs.
func (c *dNSes) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dnses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSes) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("dnses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dNS.
func (c *dNSes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Patch(pt).
		Resource("dnses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}