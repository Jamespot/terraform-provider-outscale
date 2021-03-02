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

// ReadListenerRulesRequest struct for ReadListenerRulesRequest
type ReadListenerRulesRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool                `json:"DryRun,omitempty"`
	Filters *FiltersListenerRule `json:"Filters,omitempty"`
}

// NewReadListenerRulesRequest instantiates a new ReadListenerRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadListenerRulesRequest() *ReadListenerRulesRequest {
	this := ReadListenerRulesRequest{}
	return &this
}

// NewReadListenerRulesRequestWithDefaults instantiates a new ReadListenerRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadListenerRulesRequestWithDefaults() *ReadListenerRulesRequest {
	this := ReadListenerRulesRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadListenerRulesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadListenerRulesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadListenerRulesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadListenerRulesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadListenerRulesRequest) GetFilters() FiltersListenerRule {
	if o == nil || o.Filters == nil {
		var ret FiltersListenerRule
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadListenerRulesRequest) GetFiltersOk() (*FiltersListenerRule, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadListenerRulesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersListenerRule and assigns it to the Filters field.
func (o *ReadListenerRulesRequest) SetFilters(v FiltersListenerRule) {
	o.Filters = &v
}

func (o ReadListenerRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadListenerRulesRequest struct {
	value *ReadListenerRulesRequest
	isSet bool
}

func (v NullableReadListenerRulesRequest) Get() *ReadListenerRulesRequest {
	return v.value
}

func (v *NullableReadListenerRulesRequest) Set(val *ReadListenerRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadListenerRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadListenerRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadListenerRulesRequest(val *ReadListenerRulesRequest) *NullableReadListenerRulesRequest {
	return &NullableReadListenerRulesRequest{value: val, isSet: true}
}

func (v NullableReadListenerRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadListenerRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
