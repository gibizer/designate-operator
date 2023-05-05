//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Designate) DeepCopyInto(out *Designate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Designate.
func (in *Designate) DeepCopy() *Designate {
	if in == nil {
		return nil
	}
	out := new(Designate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Designate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAPI) DeepCopyInto(out *DesignateAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAPI.
func (in *DesignateAPI) DeepCopy() *DesignateAPI {
	if in == nil {
		return nil
	}
	out := new(DesignateAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAPIList) DeepCopyInto(out *DesignateAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAPIList.
func (in *DesignateAPIList) DeepCopy() *DesignateAPIList {
	if in == nil {
		return nil
	}
	out := new(DesignateAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAPISpec) DeepCopyInto(out *DesignateAPISpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateAPITemplate.DeepCopyInto(&out.DesignateAPITemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAPISpec.
func (in *DesignateAPISpec) DeepCopy() *DesignateAPISpec {
	if in == nil {
		return nil
	}
	out := new(DesignateAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAPIStatus) DeepCopyInto(out *DesignateAPIStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAPIStatus.
func (in *DesignateAPIStatus) DeepCopy() *DesignateAPIStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAPITemplate) DeepCopyInto(out *DesignateAPITemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAPITemplate.
func (in *DesignateAPITemplate) DeepCopy() *DesignateAPITemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateAPITemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAgent) DeepCopyInto(out *DesignateAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAgent.
func (in *DesignateAgent) DeepCopy() *DesignateAgent {
	if in == nil {
		return nil
	}
	out := new(DesignateAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAgentList) DeepCopyInto(out *DesignateAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAgentList.
func (in *DesignateAgentList) DeepCopy() *DesignateAgentList {
	if in == nil {
		return nil
	}
	out := new(DesignateAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAgentSpec) DeepCopyInto(out *DesignateAgentSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateAgentTemplate.DeepCopyInto(&out.DesignateAgentTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAgentSpec.
func (in *DesignateAgentSpec) DeepCopy() *DesignateAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAgentStatus) DeepCopyInto(out *DesignateAgentStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAgentStatus.
func (in *DesignateAgentStatus) DeepCopy() *DesignateAgentStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateAgentTemplate) DeepCopyInto(out *DesignateAgentTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateAgentTemplate.
func (in *DesignateAgentTemplate) DeepCopy() *DesignateAgentTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateAgentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateCentral) DeepCopyInto(out *DesignateCentral) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateCentral.
func (in *DesignateCentral) DeepCopy() *DesignateCentral {
	if in == nil {
		return nil
	}
	out := new(DesignateCentral)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateCentral) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateCentralList) DeepCopyInto(out *DesignateCentralList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateCentral, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateCentralList.
func (in *DesignateCentralList) DeepCopy() *DesignateCentralList {
	if in == nil {
		return nil
	}
	out := new(DesignateCentralList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateCentralList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateCentralSpec) DeepCopyInto(out *DesignateCentralSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateAgentTemplate.DeepCopyInto(&out.DesignateAgentTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateCentralSpec.
func (in *DesignateCentralSpec) DeepCopy() *DesignateCentralSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateCentralSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateCentralStatus) DeepCopyInto(out *DesignateCentralStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateCentralStatus.
func (in *DesignateCentralStatus) DeepCopy() *DesignateCentralStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateCentralStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateCentralTemplate) DeepCopyInto(out *DesignateCentralTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateCentralTemplate.
func (in *DesignateCentralTemplate) DeepCopy() *DesignateCentralTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateCentralTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateDebug) DeepCopyInto(out *DesignateDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateDebug.
func (in *DesignateDebug) DeepCopy() *DesignateDebug {
	if in == nil {
		return nil
	}
	out := new(DesignateDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateList) DeepCopyInto(out *DesignateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Designate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateList.
func (in *DesignateList) DeepCopy() *DesignateList {
	if in == nil {
		return nil
	}
	out := new(DesignateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateMdns) DeepCopyInto(out *DesignateMdns) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateMdns.
func (in *DesignateMdns) DeepCopy() *DesignateMdns {
	if in == nil {
		return nil
	}
	out := new(DesignateMdns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateMdns) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateMdnsList) DeepCopyInto(out *DesignateMdnsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateMdns, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateMdnsList.
func (in *DesignateMdnsList) DeepCopy() *DesignateMdnsList {
	if in == nil {
		return nil
	}
	out := new(DesignateMdnsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateMdnsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateMdnsSpec) DeepCopyInto(out *DesignateMdnsSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateMdnsTemplate.DeepCopyInto(&out.DesignateMdnsTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateMdnsSpec.
func (in *DesignateMdnsSpec) DeepCopy() *DesignateMdnsSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateMdnsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateMdnsStatus) DeepCopyInto(out *DesignateMdnsStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateMdnsStatus.
func (in *DesignateMdnsStatus) DeepCopy() *DesignateMdnsStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateMdnsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateMdnsTemplate) DeepCopyInto(out *DesignateMdnsTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateMdnsTemplate.
func (in *DesignateMdnsTemplate) DeepCopy() *DesignateMdnsTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateMdnsTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateProducer) DeepCopyInto(out *DesignateProducer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateProducer.
func (in *DesignateProducer) DeepCopy() *DesignateProducer {
	if in == nil {
		return nil
	}
	out := new(DesignateProducer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateProducer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateProducerList) DeepCopyInto(out *DesignateProducerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateProducer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateProducerList.
func (in *DesignateProducerList) DeepCopy() *DesignateProducerList {
	if in == nil {
		return nil
	}
	out := new(DesignateProducerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateProducerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateProducerSpec) DeepCopyInto(out *DesignateProducerSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateProducerTemplate.DeepCopyInto(&out.DesignateProducerTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateProducerSpec.
func (in *DesignateProducerSpec) DeepCopy() *DesignateProducerSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateProducerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateProducerStatus) DeepCopyInto(out *DesignateProducerStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateProducerStatus.
func (in *DesignateProducerStatus) DeepCopy() *DesignateProducerStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateProducerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateProducerTemplate) DeepCopyInto(out *DesignateProducerTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateProducerTemplate.
func (in *DesignateProducerTemplate) DeepCopy() *DesignateProducerTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateProducerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateServiceDebug) DeepCopyInto(out *DesignateServiceDebug) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateServiceDebug.
func (in *DesignateServiceDebug) DeepCopy() *DesignateServiceDebug {
	if in == nil {
		return nil
	}
	out := new(DesignateServiceDebug)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateServiceTemplate) DeepCopyInto(out *DesignateServiceTemplate) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Debug = in.Debug
	if in.CustomServiceConfigSecrets != nil {
		in, out := &in.CustomServiceConfigSecrets, &out.CustomServiceConfigSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateServiceTemplate.
func (in *DesignateServiceTemplate) DeepCopy() *DesignateServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSink) DeepCopyInto(out *DesignateSink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSink.
func (in *DesignateSink) DeepCopy() *DesignateSink {
	if in == nil {
		return nil
	}
	out := new(DesignateSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateSink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSinkList) DeepCopyInto(out *DesignateSinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateSink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSinkList.
func (in *DesignateSinkList) DeepCopy() *DesignateSinkList {
	if in == nil {
		return nil
	}
	out := new(DesignateSinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateSinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSinkSpec) DeepCopyInto(out *DesignateSinkSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateSinkTemplate.DeepCopyInto(&out.DesignateSinkTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSinkSpec.
func (in *DesignateSinkSpec) DeepCopy() *DesignateSinkSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateSinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSinkStatus) DeepCopyInto(out *DesignateSinkStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSinkStatus.
func (in *DesignateSinkStatus) DeepCopy() *DesignateSinkStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateSinkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSinkTemplate) DeepCopyInto(out *DesignateSinkTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSinkTemplate.
func (in *DesignateSinkTemplate) DeepCopy() *DesignateSinkTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateSinkTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateSpec) DeepCopyInto(out *DesignateSpec) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	out.Debug = in.Debug
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.DesignateAPI.DeepCopyInto(&out.DesignateAPI)
	in.DesignateCentral.DeepCopyInto(&out.DesignateCentral)
	in.DesignateSink.DeepCopyInto(&out.DesignateSink)
	in.DesignateWorker.DeepCopyInto(&out.DesignateWorker)
	in.DesignateMdns.DeepCopyInto(&out.DesignateMdns)
	in.DesignateProducer.DeepCopyInto(&out.DesignateProducer)
	in.DesignateAgent.DeepCopyInto(&out.DesignateAgent)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateSpec.
func (in *DesignateSpec) DeepCopy() *DesignateSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateStatus) DeepCopyInto(out *DesignateStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateStatus.
func (in *DesignateStatus) DeepCopy() *DesignateStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateTemplate) DeepCopyInto(out *DesignateTemplate) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateTemplate.
func (in *DesignateTemplate) DeepCopy() *DesignateTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateWorker) DeepCopyInto(out *DesignateWorker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateWorker.
func (in *DesignateWorker) DeepCopy() *DesignateWorker {
	if in == nil {
		return nil
	}
	out := new(DesignateWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateWorker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateWorkerList) DeepCopyInto(out *DesignateWorkerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesignateWorker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateWorkerList.
func (in *DesignateWorkerList) DeepCopy() *DesignateWorkerList {
	if in == nil {
		return nil
	}
	out := new(DesignateWorkerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesignateWorkerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateWorkerSpec) DeepCopyInto(out *DesignateWorkerSpec) {
	*out = *in
	out.DesignateTemplate = in.DesignateTemplate
	in.DesignateWorkerTemplate.DeepCopyInto(&out.DesignateWorkerTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateWorkerSpec.
func (in *DesignateWorkerSpec) DeepCopy() *DesignateWorkerSpec {
	if in == nil {
		return nil
	}
	out := new(DesignateWorkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateWorkerStatus) DeepCopyInto(out *DesignateWorkerStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateWorkerStatus.
func (in *DesignateWorkerStatus) DeepCopy() *DesignateWorkerStatus {
	if in == nil {
		return nil
	}
	out := new(DesignateWorkerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesignateWorkerTemplate) DeepCopyInto(out *DesignateWorkerTemplate) {
	*out = *in
	in.DesignateServiceTemplate.DeepCopyInto(&out.DesignateServiceTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesignateWorkerTemplate.
func (in *DesignateWorkerTemplate) DeepCopy() *DesignateWorkerTemplate {
	if in == nil {
		return nil
	}
	out := new(DesignateWorkerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}
