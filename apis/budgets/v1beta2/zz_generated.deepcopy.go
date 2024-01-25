//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoAdjustDataInitParameters) DeepCopyInto(out *AutoAdjustDataInitParameters) {
	*out = *in
	if in.AutoAdjustType != nil {
		in, out := &in.AutoAdjustType, &out.AutoAdjustType
		*out = new(string)
		**out = **in
	}
	if in.HistoricalOptions != nil {
		in, out := &in.HistoricalOptions, &out.HistoricalOptions
		*out = make([]HistoricalOptionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoAdjustDataInitParameters.
func (in *AutoAdjustDataInitParameters) DeepCopy() *AutoAdjustDataInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutoAdjustDataInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoAdjustDataObservation) DeepCopyInto(out *AutoAdjustDataObservation) {
	*out = *in
	if in.AutoAdjustType != nil {
		in, out := &in.AutoAdjustType, &out.AutoAdjustType
		*out = new(string)
		**out = **in
	}
	if in.HistoricalOptions != nil {
		in, out := &in.HistoricalOptions, &out.HistoricalOptions
		*out = make([]HistoricalOptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastAutoAdjustTime != nil {
		in, out := &in.LastAutoAdjustTime, &out.LastAutoAdjustTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoAdjustDataObservation.
func (in *AutoAdjustDataObservation) DeepCopy() *AutoAdjustDataObservation {
	if in == nil {
		return nil
	}
	out := new(AutoAdjustDataObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoAdjustDataParameters) DeepCopyInto(out *AutoAdjustDataParameters) {
	*out = *in
	if in.AutoAdjustType != nil {
		in, out := &in.AutoAdjustType, &out.AutoAdjustType
		*out = new(string)
		**out = **in
	}
	if in.HistoricalOptions != nil {
		in, out := &in.HistoricalOptions, &out.HistoricalOptions
		*out = make([]HistoricalOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoAdjustDataParameters.
func (in *AutoAdjustDataParameters) DeepCopy() *AutoAdjustDataParameters {
	if in == nil {
		return nil
	}
	out := new(AutoAdjustDataParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Budget) DeepCopyInto(out *Budget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Budget.
func (in *Budget) DeepCopy() *Budget {
	if in == nil {
		return nil
	}
	out := new(Budget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Budget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetInitParameters) DeepCopyInto(out *BudgetInitParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AutoAdjustData != nil {
		in, out := &in.AutoAdjustData, &out.AutoAdjustData
		*out = make([]AutoAdjustDataInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BudgetType != nil {
		in, out := &in.BudgetType, &out.BudgetType
		*out = new(string)
		**out = **in
	}
	if in.CostFilter != nil {
		in, out := &in.CostFilter, &out.CostFilter
		*out = make([]CostFilterInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CostTypes != nil {
		in, out := &in.CostTypes, &out.CostTypes
		*out = make([]CostTypesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LimitAmount != nil {
		in, out := &in.LimitAmount, &out.LimitAmount
		*out = new(string)
		**out = **in
	}
	if in.LimitUnit != nil {
		in, out := &in.LimitUnit, &out.LimitUnit
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]NotificationInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlannedLimit != nil {
		in, out := &in.PlannedLimit, &out.PlannedLimit
		*out = make([]PlannedLimitInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimePeriodEnd != nil {
		in, out := &in.TimePeriodEnd, &out.TimePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodStart != nil {
		in, out := &in.TimePeriodStart, &out.TimePeriodStart
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetInitParameters.
func (in *BudgetInitParameters) DeepCopy() *BudgetInitParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetList) DeepCopyInto(out *BudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Budget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetList.
func (in *BudgetList) DeepCopy() *BudgetList {
	if in == nil {
		return nil
	}
	out := new(BudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetObservation) DeepCopyInto(out *BudgetObservation) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AutoAdjustData != nil {
		in, out := &in.AutoAdjustData, &out.AutoAdjustData
		*out = make([]AutoAdjustDataObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BudgetType != nil {
		in, out := &in.BudgetType, &out.BudgetType
		*out = new(string)
		**out = **in
	}
	if in.CostFilter != nil {
		in, out := &in.CostFilter, &out.CostFilter
		*out = make([]CostFilterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CostTypes != nil {
		in, out := &in.CostTypes, &out.CostTypes
		*out = make([]CostTypesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LimitAmount != nil {
		in, out := &in.LimitAmount, &out.LimitAmount
		*out = new(string)
		**out = **in
	}
	if in.LimitUnit != nil {
		in, out := &in.LimitUnit, &out.LimitUnit
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]NotificationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlannedLimit != nil {
		in, out := &in.PlannedLimit, &out.PlannedLimit
		*out = make([]PlannedLimitObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimePeriodEnd != nil {
		in, out := &in.TimePeriodEnd, &out.TimePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodStart != nil {
		in, out := &in.TimePeriodStart, &out.TimePeriodStart
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetObservation.
func (in *BudgetObservation) DeepCopy() *BudgetObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetParameters) DeepCopyInto(out *BudgetParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.AutoAdjustData != nil {
		in, out := &in.AutoAdjustData, &out.AutoAdjustData
		*out = make([]AutoAdjustDataParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BudgetType != nil {
		in, out := &in.BudgetType, &out.BudgetType
		*out = new(string)
		**out = **in
	}
	if in.CostFilter != nil {
		in, out := &in.CostFilter, &out.CostFilter
		*out = make([]CostFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CostTypes != nil {
		in, out := &in.CostTypes, &out.CostTypes
		*out = make([]CostTypesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LimitAmount != nil {
		in, out := &in.LimitAmount, &out.LimitAmount
		*out = new(string)
		**out = **in
	}
	if in.LimitUnit != nil {
		in, out := &in.LimitUnit, &out.LimitUnit
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]NotificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlannedLimit != nil {
		in, out := &in.PlannedLimit, &out.PlannedLimit
		*out = make([]PlannedLimitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodEnd != nil {
		in, out := &in.TimePeriodEnd, &out.TimePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodStart != nil {
		in, out := &in.TimePeriodStart, &out.TimePeriodStart
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetParameters.
func (in *BudgetParameters) DeepCopy() *BudgetParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetSpec) DeepCopyInto(out *BudgetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetSpec.
func (in *BudgetSpec) DeepCopy() *BudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetStatus) DeepCopyInto(out *BudgetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetStatus.
func (in *BudgetStatus) DeepCopy() *BudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostFilterInitParameters) DeepCopyInto(out *CostFilterInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostFilterInitParameters.
func (in *CostFilterInitParameters) DeepCopy() *CostFilterInitParameters {
	if in == nil {
		return nil
	}
	out := new(CostFilterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostFilterObservation) DeepCopyInto(out *CostFilterObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostFilterObservation.
func (in *CostFilterObservation) DeepCopy() *CostFilterObservation {
	if in == nil {
		return nil
	}
	out := new(CostFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostFilterParameters) DeepCopyInto(out *CostFilterParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostFilterParameters.
func (in *CostFilterParameters) DeepCopy() *CostFilterParameters {
	if in == nil {
		return nil
	}
	out := new(CostFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostTypesInitParameters) DeepCopyInto(out *CostTypesInitParameters) {
	*out = *in
	if in.IncludeCredit != nil {
		in, out := &in.IncludeCredit, &out.IncludeCredit
		*out = new(bool)
		**out = **in
	}
	if in.IncludeDiscount != nil {
		in, out := &in.IncludeDiscount, &out.IncludeDiscount
		*out = new(bool)
		**out = **in
	}
	if in.IncludeOtherSubscription != nil {
		in, out := &in.IncludeOtherSubscription, &out.IncludeOtherSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRecurring != nil {
		in, out := &in.IncludeRecurring, &out.IncludeRecurring
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRefund != nil {
		in, out := &in.IncludeRefund, &out.IncludeRefund
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSubscription != nil {
		in, out := &in.IncludeSubscription, &out.IncludeSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSupport != nil {
		in, out := &in.IncludeSupport, &out.IncludeSupport
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTax != nil {
		in, out := &in.IncludeTax, &out.IncludeTax
		*out = new(bool)
		**out = **in
	}
	if in.IncludeUpfront != nil {
		in, out := &in.IncludeUpfront, &out.IncludeUpfront
		*out = new(bool)
		**out = **in
	}
	if in.UseAmortized != nil {
		in, out := &in.UseAmortized, &out.UseAmortized
		*out = new(bool)
		**out = **in
	}
	if in.UseBlended != nil {
		in, out := &in.UseBlended, &out.UseBlended
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostTypesInitParameters.
func (in *CostTypesInitParameters) DeepCopy() *CostTypesInitParameters {
	if in == nil {
		return nil
	}
	out := new(CostTypesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostTypesObservation) DeepCopyInto(out *CostTypesObservation) {
	*out = *in
	if in.IncludeCredit != nil {
		in, out := &in.IncludeCredit, &out.IncludeCredit
		*out = new(bool)
		**out = **in
	}
	if in.IncludeDiscount != nil {
		in, out := &in.IncludeDiscount, &out.IncludeDiscount
		*out = new(bool)
		**out = **in
	}
	if in.IncludeOtherSubscription != nil {
		in, out := &in.IncludeOtherSubscription, &out.IncludeOtherSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRecurring != nil {
		in, out := &in.IncludeRecurring, &out.IncludeRecurring
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRefund != nil {
		in, out := &in.IncludeRefund, &out.IncludeRefund
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSubscription != nil {
		in, out := &in.IncludeSubscription, &out.IncludeSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSupport != nil {
		in, out := &in.IncludeSupport, &out.IncludeSupport
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTax != nil {
		in, out := &in.IncludeTax, &out.IncludeTax
		*out = new(bool)
		**out = **in
	}
	if in.IncludeUpfront != nil {
		in, out := &in.IncludeUpfront, &out.IncludeUpfront
		*out = new(bool)
		**out = **in
	}
	if in.UseAmortized != nil {
		in, out := &in.UseAmortized, &out.UseAmortized
		*out = new(bool)
		**out = **in
	}
	if in.UseBlended != nil {
		in, out := &in.UseBlended, &out.UseBlended
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostTypesObservation.
func (in *CostTypesObservation) DeepCopy() *CostTypesObservation {
	if in == nil {
		return nil
	}
	out := new(CostTypesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostTypesParameters) DeepCopyInto(out *CostTypesParameters) {
	*out = *in
	if in.IncludeCredit != nil {
		in, out := &in.IncludeCredit, &out.IncludeCredit
		*out = new(bool)
		**out = **in
	}
	if in.IncludeDiscount != nil {
		in, out := &in.IncludeDiscount, &out.IncludeDiscount
		*out = new(bool)
		**out = **in
	}
	if in.IncludeOtherSubscription != nil {
		in, out := &in.IncludeOtherSubscription, &out.IncludeOtherSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRecurring != nil {
		in, out := &in.IncludeRecurring, &out.IncludeRecurring
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRefund != nil {
		in, out := &in.IncludeRefund, &out.IncludeRefund
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSubscription != nil {
		in, out := &in.IncludeSubscription, &out.IncludeSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSupport != nil {
		in, out := &in.IncludeSupport, &out.IncludeSupport
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTax != nil {
		in, out := &in.IncludeTax, &out.IncludeTax
		*out = new(bool)
		**out = **in
	}
	if in.IncludeUpfront != nil {
		in, out := &in.IncludeUpfront, &out.IncludeUpfront
		*out = new(bool)
		**out = **in
	}
	if in.UseAmortized != nil {
		in, out := &in.UseAmortized, &out.UseAmortized
		*out = new(bool)
		**out = **in
	}
	if in.UseBlended != nil {
		in, out := &in.UseBlended, &out.UseBlended
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostTypesParameters.
func (in *CostTypesParameters) DeepCopy() *CostTypesParameters {
	if in == nil {
		return nil
	}
	out := new(CostTypesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoricalOptionsInitParameters) DeepCopyInto(out *HistoricalOptionsInitParameters) {
	*out = *in
	if in.BudgetAdjustmentPeriod != nil {
		in, out := &in.BudgetAdjustmentPeriod, &out.BudgetAdjustmentPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoricalOptionsInitParameters.
func (in *HistoricalOptionsInitParameters) DeepCopy() *HistoricalOptionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(HistoricalOptionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoricalOptionsObservation) DeepCopyInto(out *HistoricalOptionsObservation) {
	*out = *in
	if in.BudgetAdjustmentPeriod != nil {
		in, out := &in.BudgetAdjustmentPeriod, &out.BudgetAdjustmentPeriod
		*out = new(float64)
		**out = **in
	}
	if in.LookbackAvailablePeriods != nil {
		in, out := &in.LookbackAvailablePeriods, &out.LookbackAvailablePeriods
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoricalOptionsObservation.
func (in *HistoricalOptionsObservation) DeepCopy() *HistoricalOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(HistoricalOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoricalOptionsParameters) DeepCopyInto(out *HistoricalOptionsParameters) {
	*out = *in
	if in.BudgetAdjustmentPeriod != nil {
		in, out := &in.BudgetAdjustmentPeriod, &out.BudgetAdjustmentPeriod
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoricalOptionsParameters.
func (in *HistoricalOptionsParameters) DeepCopy() *HistoricalOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(HistoricalOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationInitParameters) DeepCopyInto(out *NotificationInitParameters) {
	*out = *in
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.NotificationType != nil {
		in, out := &in.NotificationType, &out.NotificationType
		*out = new(string)
		**out = **in
	}
	if in.SubscriberEmailAddresses != nil {
		in, out := &in.SubscriberEmailAddresses, &out.SubscriberEmailAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubscriberSnsTopicArns != nil {
		in, out := &in.SubscriberSnsTopicArns, &out.SubscriberSnsTopicArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationInitParameters.
func (in *NotificationInitParameters) DeepCopy() *NotificationInitParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationObservation) DeepCopyInto(out *NotificationObservation) {
	*out = *in
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.NotificationType != nil {
		in, out := &in.NotificationType, &out.NotificationType
		*out = new(string)
		**out = **in
	}
	if in.SubscriberEmailAddresses != nil {
		in, out := &in.SubscriberEmailAddresses, &out.SubscriberEmailAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubscriberSnsTopicArns != nil {
		in, out := &in.SubscriberSnsTopicArns, &out.SubscriberSnsTopicArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationObservation.
func (in *NotificationObservation) DeepCopy() *NotificationObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationParameters) DeepCopyInto(out *NotificationParameters) {
	*out = *in
	if in.ComparisonOperator != nil {
		in, out := &in.ComparisonOperator, &out.ComparisonOperator
		*out = new(string)
		**out = **in
	}
	if in.NotificationType != nil {
		in, out := &in.NotificationType, &out.NotificationType
		*out = new(string)
		**out = **in
	}
	if in.SubscriberEmailAddresses != nil {
		in, out := &in.SubscriberEmailAddresses, &out.SubscriberEmailAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubscriberSnsTopicArns != nil {
		in, out := &in.SubscriberSnsTopicArns, &out.SubscriberSnsTopicArns
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdType != nil {
		in, out := &in.ThresholdType, &out.ThresholdType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationParameters.
func (in *NotificationParameters) DeepCopy() *NotificationParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlannedLimitInitParameters) DeepCopyInto(out *PlannedLimitInitParameters) {
	*out = *in
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlannedLimitInitParameters.
func (in *PlannedLimitInitParameters) DeepCopy() *PlannedLimitInitParameters {
	if in == nil {
		return nil
	}
	out := new(PlannedLimitInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlannedLimitObservation) DeepCopyInto(out *PlannedLimitObservation) {
	*out = *in
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlannedLimitObservation.
func (in *PlannedLimitObservation) DeepCopy() *PlannedLimitObservation {
	if in == nil {
		return nil
	}
	out := new(PlannedLimitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlannedLimitParameters) DeepCopyInto(out *PlannedLimitParameters) {
	*out = *in
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(string)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlannedLimitParameters.
func (in *PlannedLimitParameters) DeepCopy() *PlannedLimitParameters {
	if in == nil {
		return nil
	}
	out := new(PlannedLimitParameters)
	in.DeepCopyInto(out)
	return out
}
