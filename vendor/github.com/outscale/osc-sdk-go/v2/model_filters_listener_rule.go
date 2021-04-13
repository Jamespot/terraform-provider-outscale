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

// FiltersListenerRule One or more filters.
type FiltersListenerRule struct {
	// The names of the listener rules.
	ListenerRuleNames *[]string `json:"ListenerRuleNames,omitempty"`
}

// NewFiltersListenerRule instantiates a new FiltersListenerRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersListenerRule() *FiltersListenerRule {
	this := FiltersListenerRule{}
	return &this
}

// NewFiltersListenerRuleWithDefaults instantiates a new FiltersListenerRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersListenerRuleWithDefaults() *FiltersListenerRule {
	this := FiltersListenerRule{}
	return &this
}

// GetListenerRuleNames returns the ListenerRuleNames field value if set, zero value otherwise.
func (o *FiltersListenerRule) GetListenerRuleNames() []string {
	if o == nil || o.ListenerRuleNames == nil {
		var ret []string
		return ret
	}
	return *o.ListenerRuleNames
}

// GetListenerRuleNamesOk returns a tuple with the ListenerRuleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersListenerRule) GetListenerRuleNamesOk() (*[]string, bool) {
	if o == nil || o.ListenerRuleNames == nil {
		return nil, false
	}
	return o.ListenerRuleNames, true
}

// HasListenerRuleNames returns a boolean if a field has been set.
func (o *FiltersListenerRule) HasListenerRuleNames() bool {
	if o != nil && o.ListenerRuleNames != nil {
		return true
	}

	return false
}

// SetListenerRuleNames gets a reference to the given []string and assigns it to the ListenerRuleNames field.
func (o *FiltersListenerRule) SetListenerRuleNames(v []string) {
	o.ListenerRuleNames = &v
}

func (o FiltersListenerRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListenerRuleNames != nil {
		toSerialize["ListenerRuleNames"] = o.ListenerRuleNames
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersListenerRule struct {
	value *FiltersListenerRule
	isSet bool
}

func (v NullableFiltersListenerRule) Get() *FiltersListenerRule {
	return v.value
}

func (v *NullableFiltersListenerRule) Set(val *FiltersListenerRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersListenerRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersListenerRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersListenerRule(val *FiltersListenerRule) *NullableFiltersListenerRule {
	return &NullableFiltersListenerRule{value: val, isSet: true}
}

func (v NullableFiltersListenerRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersListenerRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}