// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/client-go/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Binding, InType: reflect.TypeOf(&Binding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingCondition, InType: reflect.TypeOf(&BindingCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingList, InType: reflect.TypeOf(&BindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingSpec, InType: reflect.TypeOf(&BindingSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BindingStatus, InType: reflect.TypeOf(&BindingStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Broker, InType: reflect.TypeOf(&Broker{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerCondition, InType: reflect.TypeOf(&BrokerCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerList, InType: reflect.TypeOf(&BrokerList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerSpec, InType: reflect.TypeOf(&BrokerSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_BrokerStatus, InType: reflect.TypeOf(&BrokerStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_Instance, InType: reflect.TypeOf(&Instance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceCondition, InType: reflect.TypeOf(&InstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceList, InType: reflect.TypeOf(&InstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceSpec, InType: reflect.TypeOf(&InstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_InstanceStatus, InType: reflect.TypeOf(&InstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceClass, InType: reflect.TypeOf(&ServiceClass{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServiceClassList, InType: reflect.TypeOf(&ServiceClassList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_ServicePlan, InType: reflect.TypeOf(&ServicePlan{})},
	)
}

func DeepCopy_v1alpha1_Binding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Binding)
		out := out.(*Binding)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_BindingSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_BindingStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BindingCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingCondition)
		out := out.(*BindingCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1alpha1_BindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingList)
		out := out.(*BindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Binding, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Binding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BindingSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingSpec)
		out := out.(*BindingSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.Checksum != nil {
			in, out := &in.Checksum, &out.Checksum
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BindingStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BindingStatus)
		out := out.(*BindingStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BindingCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_BindingCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_Broker(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Broker)
		out := out.(*Broker)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_BrokerSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_BrokerStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BrokerCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerCondition)
		out := out.(*BrokerCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1alpha1_BrokerList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerList)
		out := out.(*BrokerList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Broker, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Broker(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BrokerSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerSpec)
		out := out.(*BrokerSpec)
		*out = *in
		if in.AuthSecret != nil {
			in, out := &in.AuthSecret, &out.AuthSecret
			*out = new(api_v1.ObjectReference)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_BrokerStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerStatus)
		out := out.(*BrokerStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BrokerCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_BrokerCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_Instance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Instance)
		out := out.(*Instance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_InstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1alpha1_InstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_InstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceCondition)
		out := out.(*InstanceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1alpha1_InstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceList)
		out := out.(*InstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Instance, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_Instance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_InstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceSpec)
		out := out.(*InstanceSpec)
		*out = *in
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		if in.Checksum != nil {
			in, out := &in.Checksum, &out.Checksum
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_InstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*InstanceStatus)
		out := out.(*InstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]InstanceCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_InstanceCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.LastOperation != nil {
			in, out := &in.LastOperation, &out.LastOperation
			*out = new(string)
			**out = **in
		}
		if in.DashboardURL != nil {
			in, out := &in.DashboardURL, &out.DashboardURL
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ServiceClass(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClass)
		out := out.(*ServiceClass)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Plans != nil {
			in, out := &in.Plans, &out.Plans
			*out = make([]ServicePlan, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServicePlan(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.OSBTags != nil {
			in, out := &in.OSBTags, &out.OSBTags
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.OSBRequires != nil {
			in, out := &in.OSBRequires, &out.OSBRequires
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.OSBMetadata != nil {
			in, out := &in.OSBMetadata, &out.OSBMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ServiceClassList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceClassList)
		out := out.(*ServiceClassList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ServiceClass, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_ServiceClass(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_ServicePlan(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServicePlan)
		out := out.(*ServicePlan)
		*out = *in
		if in.OSBMetadata != nil {
			in, out := &in.OSBMetadata, &out.OSBMetadata
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*runtime.RawExtension)
			}
		}
		return nil
	}
}
