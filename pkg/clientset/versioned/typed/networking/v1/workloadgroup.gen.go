// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "istio.io/client-go/pkg/apis/networking/v1"
	networkingv1 "istio.io/client-go/pkg/applyconfiguration/networking/v1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkloadGroupsGetter has a method to return a WorkloadGroupInterface.
// A group's client should implement this interface.
type WorkloadGroupsGetter interface {
	WorkloadGroups(namespace string) WorkloadGroupInterface
}

// WorkloadGroupInterface has methods to work with WorkloadGroup resources.
type WorkloadGroupInterface interface {
	Create(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.CreateOptions) (*v1.WorkloadGroup, error)
	Update(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.UpdateOptions) (*v1.WorkloadGroup, error)
	UpdateStatus(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.UpdateOptions) (*v1.WorkloadGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.WorkloadGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.WorkloadGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.WorkloadGroup, err error)
	Apply(ctx context.Context, workloadGroup *networkingv1.WorkloadGroupApplyConfiguration, opts metav1.ApplyOptions) (result *v1.WorkloadGroup, err error)
	ApplyStatus(ctx context.Context, workloadGroup *networkingv1.WorkloadGroupApplyConfiguration, opts metav1.ApplyOptions) (result *v1.WorkloadGroup, err error)
	WorkloadGroupExpansion
}

// workloadGroups implements WorkloadGroupInterface
type workloadGroups struct {
	client rest.Interface
	ns     string
}

// newWorkloadGroups returns a WorkloadGroups
func newWorkloadGroups(c *NetworkingV1Client, namespace string) *workloadGroups {
	return &workloadGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workloadGroup, and returns the corresponding workloadGroup object, and an error if there is any.
func (c *workloadGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.WorkloadGroup, err error) {
	result = &v1.WorkloadGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkloadGroups that match those selectors.
func (c *workloadGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.WorkloadGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.WorkloadGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloadgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workloadGroups.
func (c *workloadGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workloadgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workloadGroup and creates it.  Returns the server's representation of the workloadGroup, and an error, if there is any.
func (c *workloadGroups) Create(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.CreateOptions) (result *v1.WorkloadGroup, err error) {
	result = &v1.WorkloadGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workloadgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workloadGroup and updates it. Returns the server's representation of the workloadGroup, and an error, if there is any.
func (c *workloadGroups) Update(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.UpdateOptions) (result *v1.WorkloadGroup, err error) {
	result = &v1.WorkloadGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(workloadGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workloadGroups) UpdateStatus(ctx context.Context, workloadGroup *v1.WorkloadGroup, opts metav1.UpdateOptions) (result *v1.WorkloadGroup, err error) {
	result = &v1.WorkloadGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(workloadGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workloadGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workloadGroup and deletes it. Returns an error if one occurs.
func (c *workloadGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workloadGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloadgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workloadGroup.
func (c *workloadGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.WorkloadGroup, err error) {
	result = &v1.WorkloadGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied workloadGroup.
func (c *workloadGroups) Apply(ctx context.Context, workloadGroup *networkingv1.WorkloadGroupApplyConfiguration, opts metav1.ApplyOptions) (result *v1.WorkloadGroup, err error) {
	if workloadGroup == nil {
		return nil, fmt.Errorf("workloadGroup provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(workloadGroup)
	if err != nil {
		return nil, err
	}
	name := workloadGroup.Name
	if name == nil {
		return nil, fmt.Errorf("workloadGroup.Name must be provided to Apply")
	}
	result = &v1.WorkloadGroup{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *workloadGroups) ApplyStatus(ctx context.Context, workloadGroup *networkingv1.WorkloadGroupApplyConfiguration, opts metav1.ApplyOptions) (result *v1.WorkloadGroup, err error) {
	if workloadGroup == nil {
		return nil, fmt.Errorf("workloadGroup provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(workloadGroup)
	if err != nil {
		return nil, err
	}

	name := workloadGroup.Name
	if name == nil {
		return nil, fmt.Errorf("workloadGroup.Name must be provided to Apply")
	}

	result = &v1.WorkloadGroup{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("workloadgroups").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
