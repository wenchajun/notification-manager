// +build !ignore_autogenerated

/*


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

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfig) DeepCopyInto(out *EmailConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfig.
func (in *EmailConfig) DeepCopy() *EmailConfig {
	if in == nil {
		return nil
	}
	out := new(EmailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfigList) DeepCopyInto(out *EmailConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmailConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfigList.
func (in *EmailConfigList) DeepCopy() *EmailConfigList {
	if in == nil {
		return nil
	}
	out := new(EmailConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfigSpec) DeepCopyInto(out *EmailConfigSpec) {
	*out = *in
	if in.Hello != nil {
		in, out := &in.Hello, &out.Hello
		*out = new(string)
		**out = **in
	}
	if in.AuthUsername != nil {
		in, out := &in.AuthUsername, &out.AuthUsername
		*out = new(string)
		**out = **in
	}
	if in.AuthIdentify != nil {
		in, out := &in.AuthIdentify, &out.AuthIdentify
		*out = new(string)
		**out = **in
	}
	if in.AuthPassword != nil {
		in, out := &in.AuthPassword, &out.AuthPassword
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthSecret != nil {
		in, out := &in.AuthSecret, &out.AuthSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfigSpec.
func (in *EmailConfigSpec) DeepCopy() *EmailConfigSpec {
	if in == nil {
		return nil
	}
	out := new(EmailConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailConfigStatus) DeepCopyInto(out *EmailConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailConfigStatus.
func (in *EmailConfigStatus) DeepCopy() *EmailConfigStatus {
	if in == nil {
		return nil
	}
	out := new(EmailConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiver) DeepCopyInto(out *EmailReceiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiver.
func (in *EmailReceiver) DeepCopy() *EmailReceiver {
	if in == nil {
		return nil
	}
	out := new(EmailReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailReceiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiverList) DeepCopyInto(out *EmailReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmailReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiverList.
func (in *EmailReceiverList) DeepCopy() *EmailReceiverList {
	if in == nil {
		return nil
	}
	out := new(EmailReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmailReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiverSpec) DeepCopyInto(out *EmailReceiverSpec) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EmailConfigSelector != nil {
		in, out := &in.EmailConfigSelector, &out.EmailConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiverSpec.
func (in *EmailReceiverSpec) DeepCopy() *EmailReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(EmailReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmailReceiverStatus) DeepCopyInto(out *EmailReceiverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmailReceiverStatus.
func (in *EmailReceiverStatus) DeepCopy() *EmailReceiverStatus {
	if in == nil {
		return nil
	}
	out := new(EmailReceiverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalSpec) DeepCopyInto(out *GlobalSpec) {
	*out = *in
	if in.EmailConfigSelector != nil {
		in, out := &in.EmailConfigSelector, &out.EmailConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WechatConfigSelector != nil {
		in, out := &in.WechatConfigSelector, &out.WechatConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SlackConfigSelector != nil {
		in, out := &in.SlackConfigSelector, &out.SlackConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookConfigSelector != nil {
		in, out := &in.WebhookConfigSelector, &out.WebhookConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalSpec.
func (in *GlobalSpec) DeepCopy() *GlobalSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationManager) DeepCopyInto(out *NotificationManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationManager.
func (in *NotificationManager) DeepCopy() *NotificationManager {
	if in == nil {
		return nil
	}
	out := new(NotificationManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotificationManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationManagerList) DeepCopyInto(out *NotificationManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NotificationManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationManagerList.
func (in *NotificationManagerList) DeepCopy() *NotificationManagerList {
	if in == nil {
		return nil
	}
	out := new(NotificationManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotificationManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationManagerSpec) DeepCopyInto(out *NotificationManagerSpec) {
	*out = *in
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(GlobalSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Receivers != nil {
		in, out := &in.Receivers, &out.Receivers
		*out = new(ReceiversSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationManagerSpec.
func (in *NotificationManagerSpec) DeepCopy() *NotificationManagerSpec {
	if in == nil {
		return nil
	}
	out := new(NotificationManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationManagerStatus) DeepCopyInto(out *NotificationManagerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationManagerStatus.
func (in *NotificationManagerStatus) DeepCopy() *NotificationManagerStatus {
	if in == nil {
		return nil
	}
	out := new(NotificationManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Receiver) DeepCopyInto(out *Receiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Receiver.
func (in *Receiver) DeepCopy() *Receiver {
	if in == nil {
		return nil
	}
	out := new(Receiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Receiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiverList) DeepCopyInto(out *ReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Receiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiverList.
func (in *ReceiverList) DeepCopy() *ReceiverList {
	if in == nil {
		return nil
	}
	out := new(ReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiverSpec) DeepCopyInto(out *ReceiverSpec) {
	*out = *in
	if in.EmailReceiverSelector != nil {
		in, out := &in.EmailReceiverSelector, &out.EmailReceiverSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WechatReceiverSelector != nil {
		in, out := &in.WechatReceiverSelector, &out.WechatReceiverSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SlackReceiverSelector != nil {
		in, out := &in.SlackReceiverSelector, &out.SlackReceiverSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.WebhookReceiverSelector != nil {
		in, out := &in.WebhookReceiverSelector, &out.WebhookReceiverSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiverSpec.
func (in *ReceiverSpec) DeepCopy() *ReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(ReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiverStatus) DeepCopyInto(out *ReceiverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiverStatus.
func (in *ReceiverStatus) DeepCopy() *ReceiverStatus {
	if in == nil {
		return nil
	}
	out := new(ReceiverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiversSpec) DeepCopyInto(out *ReceiversSpec) {
	*out = *in
	if in.ReceiverSelector != nil {
		in, out := &in.ReceiverSelector, &out.ReceiverSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiversSpec.
func (in *ReceiversSpec) DeepCopy() *ReceiversSpec {
	if in == nil {
		return nil
	}
	out := new(ReceiversSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfig) DeepCopyInto(out *SlackConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfig.
func (in *SlackConfig) DeepCopy() *SlackConfig {
	if in == nil {
		return nil
	}
	out := new(SlackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfigList) DeepCopyInto(out *SlackConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlackConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfigList.
func (in *SlackConfigList) DeepCopy() *SlackConfigList {
	if in == nil {
		return nil
	}
	out := new(SlackConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfigSpec) DeepCopyInto(out *SlackConfigSpec) {
	*out = *in
	if in.SlackApiUrlSecret != nil {
		in, out := &in.SlackApiUrlSecret, &out.SlackApiUrlSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SlackApiUrl != nil {
		in, out := &in.SlackApiUrl, &out.SlackApiUrl
		*out = new(string)
		**out = **in
	}
	if in.HttpConfigSelector != nil {
		in, out := &in.HttpConfigSelector, &out.HttpConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfigSpec.
func (in *SlackConfigSpec) DeepCopy() *SlackConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SlackConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackConfigStatus) DeepCopyInto(out *SlackConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackConfigStatus.
func (in *SlackConfigStatus) DeepCopy() *SlackConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SlackConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackReceiver) DeepCopyInto(out *SlackReceiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackReceiver.
func (in *SlackReceiver) DeepCopy() *SlackReceiver {
	if in == nil {
		return nil
	}
	out := new(SlackReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackReceiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackReceiverList) DeepCopyInto(out *SlackReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlackReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackReceiverList.
func (in *SlackReceiverList) DeepCopy() *SlackReceiverList {
	if in == nil {
		return nil
	}
	out := new(SlackReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlackReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackReceiverSpec) DeepCopyInto(out *SlackReceiverSpec) {
	*out = *in
	if in.SlackConfigSelector != nil {
		in, out := &in.SlackConfigSelector, &out.SlackConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackReceiverSpec.
func (in *SlackReceiverSpec) DeepCopy() *SlackReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(SlackReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlackReceiverStatus) DeepCopyInto(out *SlackReceiverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlackReceiverStatus.
func (in *SlackReceiverStatus) DeepCopy() *SlackReceiverStatus {
	if in == nil {
		return nil
	}
	out := new(SlackReceiverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigList) DeepCopyInto(out *WebhookConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebhookConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigList.
func (in *WebhookConfigList) DeepCopy() *WebhookConfigList {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigSpec) DeepCopyInto(out *WebhookConfigSpec) {
	*out = *in
	if in.UrlSecret != nil {
		in, out := &in.UrlSecret, &out.UrlSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.HttpConfigSelector != nil {
		in, out := &in.HttpConfigSelector, &out.HttpConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigSpec.
func (in *WebhookConfigSpec) DeepCopy() *WebhookConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfigStatus) DeepCopyInto(out *WebhookConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfigStatus.
func (in *WebhookConfigStatus) DeepCopy() *WebhookConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiver) DeepCopyInto(out *WebhookReceiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiver.
func (in *WebhookReceiver) DeepCopy() *WebhookReceiver {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookReceiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiverList) DeepCopyInto(out *WebhookReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WebhookReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiverList.
func (in *WebhookReceiverList) DeepCopy() *WebhookReceiverList {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiverSpec) DeepCopyInto(out *WebhookReceiverSpec) {
	*out = *in
	if in.WebhookConfigSelector != nil {
		in, out := &in.WebhookConfigSelector, &out.WebhookConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiverSpec.
func (in *WebhookReceiverSpec) DeepCopy() *WebhookReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookReceiverStatus) DeepCopyInto(out *WebhookReceiverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookReceiverStatus.
func (in *WebhookReceiverStatus) DeepCopy() *WebhookReceiverStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookReceiverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatConfig) DeepCopyInto(out *WechatConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatConfig.
func (in *WechatConfig) DeepCopy() *WechatConfig {
	if in == nil {
		return nil
	}
	out := new(WechatConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WechatConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatConfigList) DeepCopyInto(out *WechatConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WechatConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatConfigList.
func (in *WechatConfigList) DeepCopy() *WechatConfigList {
	if in == nil {
		return nil
	}
	out := new(WechatConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WechatConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatConfigSpec) DeepCopyInto(out *WechatConfigSpec) {
	*out = *in
	if in.WechatApiSecret != nil {
		in, out := &in.WechatApiSecret, &out.WechatApiSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatConfigSpec.
func (in *WechatConfigSpec) DeepCopy() *WechatConfigSpec {
	if in == nil {
		return nil
	}
	out := new(WechatConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatConfigStatus) DeepCopyInto(out *WechatConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatConfigStatus.
func (in *WechatConfigStatus) DeepCopy() *WechatConfigStatus {
	if in == nil {
		return nil
	}
	out := new(WechatConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatReceiver) DeepCopyInto(out *WechatReceiver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatReceiver.
func (in *WechatReceiver) DeepCopy() *WechatReceiver {
	if in == nil {
		return nil
	}
	out := new(WechatReceiver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WechatReceiver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatReceiverList) DeepCopyInto(out *WechatReceiverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WechatReceiver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatReceiverList.
func (in *WechatReceiverList) DeepCopy() *WechatReceiverList {
	if in == nil {
		return nil
	}
	out := new(WechatReceiverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WechatReceiverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatReceiverSpec) DeepCopyInto(out *WechatReceiverSpec) {
	*out = *in
	if in.WechatConfigSelector != nil {
		in, out := &in.WechatConfigSelector, &out.WechatConfigSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatReceiverSpec.
func (in *WechatReceiverSpec) DeepCopy() *WechatReceiverSpec {
	if in == nil {
		return nil
	}
	out := new(WechatReceiverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WechatReceiverStatus) DeepCopyInto(out *WechatReceiverStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WechatReceiverStatus.
func (in *WechatReceiverStatus) DeepCopy() *WechatReceiverStatus {
	if in == nil {
		return nil
	}
	out := new(WechatReceiverStatus)
	in.DeepCopyInto(out)
	return out
}