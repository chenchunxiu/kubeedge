// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

package config

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extender) DeepCopyInto(out *Extender) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(ExtenderTLSConfig)
		(*in).DeepCopyInto(*out)
	}
	out.HTTPTimeout = in.HTTPTimeout
	if in.ManagedResources != nil {
		in, out := &in.ManagedResources, &out.ManagedResources
		*out = make([]ExtenderManagedResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extender.
func (in *Extender) DeepCopy() *Extender {
	if in == nil {
		return nil
	}
	out := new(Extender)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderManagedResource) DeepCopyInto(out *ExtenderManagedResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderManagedResource.
func (in *ExtenderManagedResource) DeepCopy() *ExtenderManagedResource {
	if in == nil {
		return nil
	}
	out := new(ExtenderManagedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtenderTLSConfig) DeepCopyInto(out *ExtenderTLSConfig) {
	*out = *in
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtenderTLSConfig.
func (in *ExtenderTLSConfig) DeepCopy() *ExtenderTLSConfig {
	if in == nil {
		return nil
	}
	out := new(ExtenderTLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterPodAffinityArgs) DeepCopyInto(out *InterPodAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterPodAffinityArgs.
func (in *InterPodAffinityArgs) DeepCopy() *InterPodAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(InterPodAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterPodAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AlgorithmSource.DeepCopyInto(&out.AlgorithmSource)
	out.LeaderElection = in.LeaderElection
	out.ClientConnection = in.ClientConnection
	out.DebuggingConfiguration = in.DebuggingConfiguration
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]KubeSchedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extenders != nil {
		in, out := &in.Extenders, &out.Extenders
		*out = make([]Extender, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerProfile) DeepCopyInto(out *KubeSchedulerProfile) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(Plugins)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerProfile.
func (in *KubeSchedulerProfile) DeepCopy() *KubeSchedulerProfile {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelPreference) DeepCopyInto(out *LabelPreference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelPreference.
func (in *LabelPreference) DeepCopy() *LabelPreference {
	if in == nil {
		return nil
	}
	out := new(LabelPreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelsPresence) DeepCopyInto(out *LabelsPresence) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelsPresence.
func (in *LabelsPresence) DeepCopy() *LabelsPresence {
	if in == nil {
		return nil
	}
	out := new(LabelsPresence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLabelArgs) DeepCopyInto(out *NodeLabelArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.PresentLabels != nil {
		in, out := &in.PresentLabels, &out.PresentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AbsentLabels != nil {
		in, out := &in.AbsentLabels, &out.AbsentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PresentLabelsPreference != nil {
		in, out := &in.PresentLabelsPreference, &out.PresentLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AbsentLabelsPreference != nil {
		in, out := &in.AbsentLabelsPreference, &out.AbsentLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLabelArgs.
func (in *NodeLabelArgs) DeepCopy() *NodeLabelArgs {
	if in == nil {
		return nil
	}
	out := new(NodeLabelArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeLabelArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesFitArgs) DeepCopyInto(out *NodeResourcesFitArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.IgnoredResources != nil {
		in, out := &in.IgnoredResources, &out.IgnoredResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IgnoredResourceGroups != nil {
		in, out := &in.IgnoredResourceGroups, &out.IgnoredResourceGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesFitArgs.
func (in *NodeResourcesFitArgs) DeepCopy() *NodeResourcesFitArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesFitArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesFitArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopyInto(out *NodeResourcesLeastAllocatedArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesLeastAllocatedArgs.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopy() *NodeResourcesLeastAllocatedArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesLeastAllocatedArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesLeastAllocatedArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResourcesMostAllocatedArgs) DeepCopyInto(out *NodeResourcesMostAllocatedArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResourcesMostAllocatedArgs.
func (in *NodeResourcesMostAllocatedArgs) DeepCopy() *NodeResourcesMostAllocatedArgs {
	if in == nil {
		return nil
	}
	out := new(NodeResourcesMostAllocatedArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeResourcesMostAllocatedArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	if in.Args != nil {
		out.Args = in.Args.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	if in.QueueSort != nil {
		in, out := &in.QueueSort, &out.QueueSort
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreFilter != nil {
		in, out := &in.PreFilter, &out.PreFilter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PostFilter != nil {
		in, out := &in.PostFilter, &out.PostFilter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreScore != nil {
		in, out := &in.PreScore, &out.PreScore
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Reserve != nil {
		in, out := &in.Reserve, &out.Reserve
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Permit != nil {
		in, out := &in.Permit, &out.Permit
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreBind != nil {
		in, out := &in.PreBind, &out.PreBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Bind != nil {
		in, out := &in.Bind, &out.Bind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PostBind != nil {
		in, out := &in.PostBind, &out.PostBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTopologySpreadArgs) DeepCopyInto(out *PodTopologySpreadArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.DefaultConstraints != nil {
		in, out := &in.DefaultConstraints, &out.DefaultConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTopologySpreadArgs.
func (in *PodTopologySpreadArgs) DeepCopy() *PodTopologySpreadArgs {
	if in == nil {
		return nil
	}
	out := new(PodTopologySpreadArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodTopologySpreadArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Predicates != nil {
		in, out := &in.Predicates, &out.Predicates
		*out = make([]PredicatePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priorities != nil {
		in, out := &in.Priorities, &out.Priorities
		*out = make([]PriorityPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extenders != nil {
		in, out := &in.Extenders, &out.Extenders
		*out = make([]Extender, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicateArgument) DeepCopyInto(out *PredicateArgument) {
	*out = *in
	if in.ServiceAffinity != nil {
		in, out := &in.ServiceAffinity, &out.ServiceAffinity
		*out = new(ServiceAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelsPresence != nil {
		in, out := &in.LabelsPresence, &out.LabelsPresence
		*out = new(LabelsPresence)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredicateArgument.
func (in *PredicateArgument) DeepCopy() *PredicateArgument {
	if in == nil {
		return nil
	}
	out := new(PredicateArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredicatePolicy) DeepCopyInto(out *PredicatePolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		*out = new(PredicateArgument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredicatePolicy.
func (in *PredicatePolicy) DeepCopy() *PredicatePolicy {
	if in == nil {
		return nil
	}
	out := new(PredicatePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityArgument) DeepCopyInto(out *PriorityArgument) {
	*out = *in
	if in.ServiceAntiAffinity != nil {
		in, out := &in.ServiceAntiAffinity, &out.ServiceAntiAffinity
		*out = new(ServiceAntiAffinity)
		**out = **in
	}
	if in.LabelPreference != nil {
		in, out := &in.LabelPreference, &out.LabelPreference
		*out = new(LabelPreference)
		**out = **in
	}
	if in.RequestedToCapacityRatioArguments != nil {
		in, out := &in.RequestedToCapacityRatioArguments, &out.RequestedToCapacityRatioArguments
		*out = new(RequestedToCapacityRatioArguments)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityArgument.
func (in *PriorityArgument) DeepCopy() *PriorityArgument {
	if in == nil {
		return nil
	}
	out := new(PriorityArgument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityPolicy) DeepCopyInto(out *PriorityPolicy) {
	*out = *in
	if in.Argument != nil {
		in, out := &in.Argument, &out.Argument
		*out = new(PriorityArgument)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityPolicy.
func (in *PriorityPolicy) DeepCopy() *PriorityPolicy {
	if in == nil {
		return nil
	}
	out := new(PriorityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedToCapacityRatioArgs) DeepCopyInto(out *RequestedToCapacityRatioArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedToCapacityRatioArgs.
func (in *RequestedToCapacityRatioArgs) DeepCopy() *RequestedToCapacityRatioArgs {
	if in == nil {
		return nil
	}
	out := new(RequestedToCapacityRatioArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RequestedToCapacityRatioArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestedToCapacityRatioArguments) DeepCopyInto(out *RequestedToCapacityRatioArguments) {
	*out = *in
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = make([]UtilizationShapePoint, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestedToCapacityRatioArguments.
func (in *RequestedToCapacityRatioArguments) DeepCopy() *RequestedToCapacityRatioArguments {
	if in == nil {
		return nil
	}
	out := new(RequestedToCapacityRatioArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerAlgorithmSource) DeepCopyInto(out *SchedulerAlgorithmSource) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(SchedulerPolicySource)
		(*in).DeepCopyInto(*out)
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerAlgorithmSource.
func (in *SchedulerAlgorithmSource) DeepCopy() *SchedulerAlgorithmSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerAlgorithmSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyConfigMapSource) DeepCopyInto(out *SchedulerPolicyConfigMapSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyConfigMapSource.
func (in *SchedulerPolicyConfigMapSource) DeepCopy() *SchedulerPolicyConfigMapSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyConfigMapSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyFileSource) DeepCopyInto(out *SchedulerPolicyFileSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyFileSource.
func (in *SchedulerPolicyFileSource) DeepCopy() *SchedulerPolicyFileSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicySource) DeepCopyInto(out *SchedulerPolicySource) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(SchedulerPolicyFileSource)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(SchedulerPolicyConfigMapSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicySource.
func (in *SchedulerPolicySource) DeepCopy() *SchedulerPolicySource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAffinity) DeepCopyInto(out *ServiceAffinity) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAffinity.
func (in *ServiceAffinity) DeepCopy() *ServiceAffinity {
	if in == nil {
		return nil
	}
	out := new(ServiceAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAffinityArgs) DeepCopyInto(out *ServiceAffinityArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.AffinityLabels != nil {
		in, out := &in.AffinityLabels, &out.AffinityLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AntiAffinityLabelsPreference != nil {
		in, out := &in.AntiAffinityLabelsPreference, &out.AntiAffinityLabelsPreference
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAffinityArgs.
func (in *ServiceAffinityArgs) DeepCopy() *ServiceAffinityArgs {
	if in == nil {
		return nil
	}
	out := new(ServiceAffinityArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAffinityArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAntiAffinity) DeepCopyInto(out *ServiceAntiAffinity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAntiAffinity.
func (in *ServiceAntiAffinity) DeepCopy() *ServiceAntiAffinity {
	if in == nil {
		return nil
	}
	out := new(ServiceAntiAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UtilizationShapePoint) DeepCopyInto(out *UtilizationShapePoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UtilizationShapePoint.
func (in *UtilizationShapePoint) DeepCopy() *UtilizationShapePoint {
	if in == nil {
		return nil
	}
	out := new(UtilizationShapePoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBindingArgs) DeepCopyInto(out *VolumeBindingArgs) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBindingArgs.
func (in *VolumeBindingArgs) DeepCopy() *VolumeBindingArgs {
	if in == nil {
		return nil
	}
	out := new(VolumeBindingArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBindingArgs) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
