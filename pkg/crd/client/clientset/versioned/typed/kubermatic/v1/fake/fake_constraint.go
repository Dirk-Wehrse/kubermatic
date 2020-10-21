// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConstraints implements ConstraintInterface
type FakeConstraints struct {
	Fake *FakeKubermaticV1
	ns   string
}

var constraintsResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "constraints"}

var constraintsKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "Constraint"}

// Get takes name of the constraint, and returns the corresponding constraint object, and an error if there is any.
func (c *FakeConstraints) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.Constraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(constraintsResource, c.ns, name), &kubermaticv1.Constraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Constraint), err
}

// List takes label and field selectors, and returns the list of Constraints that match those selectors.
func (c *FakeConstraints) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.ConstraintList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(constraintsResource, constraintsKind, c.ns, opts), &kubermaticv1.ConstraintList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ConstraintList{ListMeta: obj.(*kubermaticv1.ConstraintList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ConstraintList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested constraints.
func (c *FakeConstraints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(constraintsResource, c.ns, opts))

}

// Create takes the representation of a constraint and creates it.  Returns the server's representation of the constraint, and an error, if there is any.
func (c *FakeConstraints) Create(ctx context.Context, constraint *kubermaticv1.Constraint, opts v1.CreateOptions) (result *kubermaticv1.Constraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(constraintsResource, c.ns, constraint), &kubermaticv1.Constraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Constraint), err
}

// Update takes the representation of a constraint and updates it. Returns the server's representation of the constraint, and an error, if there is any.
func (c *FakeConstraints) Update(ctx context.Context, constraint *kubermaticv1.Constraint, opts v1.UpdateOptions) (result *kubermaticv1.Constraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(constraintsResource, c.ns, constraint), &kubermaticv1.Constraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Constraint), err
}

// Delete takes name of the constraint and deletes it. Returns an error if one occurs.
func (c *FakeConstraints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(constraintsResource, c.ns, name), &kubermaticv1.Constraint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConstraints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(constraintsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ConstraintList{})
	return err
}

// Patch applies the patch and returns the patched constraint.
func (c *FakeConstraints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.Constraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(constraintsResource, c.ns, name, pt, data, subresources...), &kubermaticv1.Constraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.Constraint), err
}
