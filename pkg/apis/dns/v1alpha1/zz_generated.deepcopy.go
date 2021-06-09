// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSAnnotation) DeepCopyInto(out *DNSAnnotation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSAnnotation.
func (in *DNSAnnotation) DeepCopy() *DNSAnnotation {
	if in == nil {
		return nil
	}
	out := new(DNSAnnotation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSAnnotation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSAnnotationList) DeepCopyInto(out *DNSAnnotationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSAnnotation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSAnnotationList.
func (in *DNSAnnotationList) DeepCopy() *DNSAnnotationList {
	if in == nil {
		return nil
	}
	out := new(DNSAnnotationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSAnnotationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSAnnotationSpec) DeepCopyInto(out *DNSAnnotationSpec) {
	*out = *in
	out.ResourceRef = in.ResourceRef
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSAnnotationSpec.
func (in *DNSAnnotationSpec) DeepCopy() *DNSAnnotationSpec {
	if in == nil {
		return nil
	}
	out := new(DNSAnnotationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSAnnotationStatus) DeepCopyInto(out *DNSAnnotationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSAnnotationStatus.
func (in *DNSAnnotationStatus) DeepCopy() *DNSAnnotationStatus {
	if in == nil {
		return nil
	}
	out := new(DNSAnnotationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSBaseStatus) DeepCopyInto(out *DNSBaseStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.LastUptimeTime != nil {
		in, out := &in.LastUptimeTime, &out.LastUptimeTime
		*out = (*in).DeepCopy()
	}
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSBaseStatus.
func (in *DNSBaseStatus) DeepCopy() *DNSBaseStatus {
	if in == nil {
		return nil
	}
	out := new(DNSBaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntry) DeepCopyInto(out *DNSEntry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntry.
func (in *DNSEntry) DeepCopy() *DNSEntry {
	if in == nil {
		return nil
	}
	out := new(DNSEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSEntry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntryList) DeepCopyInto(out *DNSEntryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntryList.
func (in *DNSEntryList) DeepCopy() *DNSEntryList {
	if in == nil {
		return nil
	}
	out := new(DNSEntryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSEntryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntrySpec) DeepCopyInto(out *DNSEntrySpec) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(EntryReference)
		**out = **in
	}
	if in.OwnerId != nil {
		in, out := &in.OwnerId, &out.OwnerId
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int64)
		**out = **in
	}
	if in.CNameLookupInterval != nil {
		in, out := &in.CNameLookupInterval, &out.CNameLookupInterval
		*out = new(int64)
		**out = **in
	}
	if in.Text != nil {
		in, out := &in.Text, &out.Text
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntrySpec.
func (in *DNSEntrySpec) DeepCopy() *DNSEntrySpec {
	if in == nil {
		return nil
	}
	out := new(DNSEntrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSEntryStatus) DeepCopyInto(out *DNSEntryStatus) {
	*out = *in
	in.DNSBaseStatus.DeepCopyInto(&out.DNSBaseStatus)
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSEntryStatus.
func (in *DNSEntryStatus) DeepCopy() *DNSEntryStatus {
	if in == nil {
		return nil
	}
	out := new(DNSEntryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSLock) DeepCopyInto(out *DNSLock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSLock.
func (in *DNSLock) DeepCopy() *DNSLock {
	if in == nil {
		return nil
	}
	out := new(DNSLock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSLock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSLockList) DeepCopyInto(out *DNSLockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSLock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSLockList.
func (in *DNSLockList) DeepCopy() *DNSLockList {
	if in == nil {
		return nil
	}
	out := new(DNSLockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSLockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSLockSpec) DeepCopyInto(out *DNSLockSpec) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int64)
		**out = **in
	}
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSLockSpec.
func (in *DNSLockSpec) DeepCopy() *DNSLockSpec {
	if in == nil {
		return nil
	}
	out := new(DNSLockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSOwner) DeepCopyInto(out *DNSOwner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOwner.
func (in *DNSOwner) DeepCopy() *DNSOwner {
	if in == nil {
		return nil
	}
	out := new(DNSOwner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSOwner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSOwnerList) DeepCopyInto(out *DNSOwnerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSOwner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOwnerList.
func (in *DNSOwnerList) DeepCopy() *DNSOwnerList {
	if in == nil {
		return nil
	}
	out := new(DNSOwnerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSOwnerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSOwnerSpec) DeepCopyInto(out *DNSOwnerSpec) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(bool)
		**out = **in
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOwnerSpec.
func (in *DNSOwnerSpec) DeepCopy() *DNSOwnerSpec {
	if in == nil {
		return nil
	}
	out := new(DNSOwnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSOwnerStatus) DeepCopyInto(out *DNSOwnerStatus) {
	*out = *in
	in.Entries.DeepCopyInto(&out.Entries)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOwnerStatus.
func (in *DNSOwnerStatus) DeepCopy() *DNSOwnerStatus {
	if in == nil {
		return nil
	}
	out := new(DNSOwnerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSOwnerStatusEntries) DeepCopyInto(out *DNSOwnerStatusEntries) {
	*out = *in
	if in.ByType != nil {
		in, out := &in.ByType, &out.ByType
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOwnerStatusEntries.
func (in *DNSOwnerStatusEntries) DeepCopy() *DNSOwnerStatusEntries {
	if in == nil {
		return nil
	}
	out := new(DNSOwnerStatusEntries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProvider) DeepCopyInto(out *DNSProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProvider.
func (in *DNSProvider) DeepCopy() *DNSProvider {
	if in == nil {
		return nil
	}
	out := new(DNSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderList) DeepCopyInto(out *DNSProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderList.
func (in *DNSProviderList) DeepCopy() *DNSProviderList {
	if in == nil {
		return nil
	}
	out := new(DNSProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderSpec) DeepCopyInto(out *DNSProviderSpec) {
	*out = *in
	if in.ProviderConfig != nil {
		in, out := &in.ProviderConfig, &out.ProviderConfig
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = new(DNSSelection)
		(*in).DeepCopyInto(*out)
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = new(DNSSelection)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultTTL != nil {
		in, out := &in.DefaultTTL, &out.DefaultTTL
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderSpec.
func (in *DNSProviderSpec) DeepCopy() *DNSProviderSpec {
	if in == nil {
		return nil
	}
	out := new(DNSProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSProviderStatus) DeepCopyInto(out *DNSProviderStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.LastUptimeTime != nil {
		in, out := &in.LastUptimeTime, &out.LastUptimeTime
		*out = (*in).DeepCopy()
	}
	in.Domains.DeepCopyInto(&out.Domains)
	in.Zones.DeepCopyInto(&out.Zones)
	if in.DefaultTTL != nil {
		in, out := &in.DefaultTTL, &out.DefaultTTL
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSProviderStatus.
func (in *DNSProviderStatus) DeepCopy() *DNSProviderStatus {
	if in == nil {
		return nil
	}
	out := new(DNSProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSelection) DeepCopyInto(out *DNSSelection) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSelection.
func (in *DNSSelection) DeepCopy() *DNSSelection {
	if in == nil {
		return nil
	}
	out := new(DNSSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSelectionStatus) DeepCopyInto(out *DNSSelectionStatus) {
	*out = *in
	if in.Included != nil {
		in, out := &in.Included, &out.Included
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Excluded != nil {
		in, out := &in.Excluded, &out.Excluded
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSelectionStatus.
func (in *DNSSelectionStatus) DeepCopy() *DNSSelectionStatus {
	if in == nil {
		return nil
	}
	out := new(DNSSelectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntryReference) DeepCopyInto(out *EntryReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntryReference.
func (in *EntryReference) DeepCopy() *EntryReference {
	if in == nil {
		return nil
	}
	out := new(EntryReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}
