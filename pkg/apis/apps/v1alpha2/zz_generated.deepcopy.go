//go:build !ignore_autogenerated
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

package v1alpha2

import (
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectiveDeployment) DeepCopyInto(out *SelectiveDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectiveDeployment.
func (in *SelectiveDeployment) DeepCopy() *SelectiveDeployment {
	if in == nil {
		return nil
	}
	out := new(SelectiveDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelectiveDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectiveDeploymentList) DeepCopyInto(out *SelectiveDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SelectiveDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectiveDeploymentList.
func (in *SelectiveDeploymentList) DeepCopy() *SelectiveDeploymentList {
	if in == nil {
		return nil
	}
	out := new(SelectiveDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelectiveDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectiveDeploymentSpec) DeepCopyInto(out *SelectiveDeploymentSpec) {
	*out = *in
	in.Workloads.DeepCopyInto(&out.Workloads)
	if in.ClusterAffinity != nil {
		in, out := &in.ClusterAffinity, &out.ClusterAffinity
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectiveDeploymentSpec.
func (in *SelectiveDeploymentSpec) DeepCopy() *SelectiveDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(SelectiveDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectiveDeploymentStatus) DeepCopyInto(out *SelectiveDeploymentStatus) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(map[string]WorkloadClusterStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectiveDeploymentStatus.
func (in *SelectiveDeploymentStatus) DeepCopy() *SelectiveDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(SelectiveDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadClusterStatus) DeepCopyInto(out *WorkloadClusterStatus) {
	*out = *in
	in.Workloads.DeepCopyInto(&out.Workloads)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadClusterStatus.
func (in *WorkloadClusterStatus) DeepCopy() *WorkloadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DaemonSet != nil {
		in, out := &in.DaemonSet, &out.DaemonSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CronJob != nil {
		in, out := &in.CronJob, &out.CronJob
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workloads) DeepCopyInto(out *Workloads) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = make([]appsv1.Deployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DaemonSet != nil {
		in, out := &in.DaemonSet, &out.DaemonSet
		*out = make([]appsv1.DaemonSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = make([]appsv1.StatefulSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Job != nil {
		in, out := &in.Job, &out.Job
		*out = make([]batchv1.Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CronJob != nil {
		in, out := &in.CronJob, &out.CronJob
		*out = make([]batchv1.CronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workloads.
func (in *Workloads) DeepCopy() *Workloads {
	if in == nil {
		return nil
	}
	out := new(Workloads)
	in.DeepCopyInto(out)
	return out
}
