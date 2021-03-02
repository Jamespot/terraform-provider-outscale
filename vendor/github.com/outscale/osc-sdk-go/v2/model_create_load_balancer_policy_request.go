/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateLoadBalancerPolicyRequest struct for CreateLoadBalancerPolicyRequest
type CreateLoadBalancerPolicyRequest struct {
	// The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie.
	CookieName *string `json:"CookieName,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the load balancer for which you want to create a policy.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-).
	PolicyName string `json:"PolicyName"`
	// The type of stickiness policy you want to create: `app` or `load_balancer`.
	PolicyType string `json:"PolicyType"`
}

// NewCreateLoadBalancerPolicyRequest instantiates a new CreateLoadBalancerPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerPolicyRequest(loadBalancerName string, policyName string, policyType string) *CreateLoadBalancerPolicyRequest {
	this := CreateLoadBalancerPolicyRequest{}
	this.LoadBalancerName = loadBalancerName
	this.PolicyName = policyName
	this.PolicyType = policyType
	return &this
}

// NewCreateLoadBalancerPolicyRequestWithDefaults instantiates a new CreateLoadBalancerPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerPolicyRequestWithDefaults() *CreateLoadBalancerPolicyRequest {
	this := CreateLoadBalancerPolicyRequest{}
	return &this
}

// GetCookieName returns the CookieName field value if set, zero value otherwise.
func (o *CreateLoadBalancerPolicyRequest) GetCookieName() string {
	if o == nil || o.CookieName == nil {
		var ret string
		return ret
	}
	return *o.CookieName
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyRequest) GetCookieNameOk() (*string, bool) {
	if o == nil || o.CookieName == nil {
		return nil, false
	}
	return o.CookieName, true
}

// HasCookieName returns a boolean if a field has been set.
func (o *CreateLoadBalancerPolicyRequest) HasCookieName() bool {
	if o != nil && o.CookieName != nil {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given string and assigns it to the CookieName field.
func (o *CreateLoadBalancerPolicyRequest) SetCookieName(v string) {
	o.CookieName = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateLoadBalancerPolicyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateLoadBalancerPolicyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateLoadBalancerPolicyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *CreateLoadBalancerPolicyRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

// GetPolicyName returns the PolicyName field value
func (o *CreateLoadBalancerPolicyRequest) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyRequest) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *CreateLoadBalancerPolicyRequest) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetPolicyType returns the PolicyType field value
func (o *CreateLoadBalancerPolicyRequest) GetPolicyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerPolicyRequest) GetPolicyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value
func (o *CreateLoadBalancerPolicyRequest) SetPolicyType(v string) {
	o.PolicyType = v
}

func (o CreateLoadBalancerPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CookieName != nil {
		toSerialize["CookieName"] = o.CookieName
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if true {
		toSerialize["PolicyName"] = o.PolicyName
	}
	if true {
		toSerialize["PolicyType"] = o.PolicyType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerPolicyRequest struct {
	value *CreateLoadBalancerPolicyRequest
	isSet bool
}

func (v NullableCreateLoadBalancerPolicyRequest) Get() *CreateLoadBalancerPolicyRequest {
	return v.value
}

func (v *NullableCreateLoadBalancerPolicyRequest) Set(val *CreateLoadBalancerPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerPolicyRequest(val *CreateLoadBalancerPolicyRequest) *NullableCreateLoadBalancerPolicyRequest {
	return &NullableCreateLoadBalancerPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
