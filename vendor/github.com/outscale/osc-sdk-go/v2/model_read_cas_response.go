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

// ReadCasResponse struct for ReadCasResponse
type ReadCasResponse struct {
	// Information about one or more CAs.
	Cas             *[]Ca            `json:"Cas,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadCasResponse instantiates a new ReadCasResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadCasResponse() *ReadCasResponse {
	this := ReadCasResponse{}
	return &this
}

// NewReadCasResponseWithDefaults instantiates a new ReadCasResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadCasResponseWithDefaults() *ReadCasResponse {
	this := ReadCasResponse{}
	return &this
}

// GetCas returns the Cas field value if set, zero value otherwise.
func (o *ReadCasResponse) GetCas() []Ca {
	if o == nil || o.Cas == nil {
		var ret []Ca
		return ret
	}
	return *o.Cas
}

// GetCasOk returns a tuple with the Cas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCasResponse) GetCasOk() (*[]Ca, bool) {
	if o == nil || o.Cas == nil {
		return nil, false
	}
	return o.Cas, true
}

// HasCas returns a boolean if a field has been set.
func (o *ReadCasResponse) HasCas() bool {
	if o != nil && o.Cas != nil {
		return true
	}

	return false
}

// SetCas gets a reference to the given []Ca and assigns it to the Cas field.
func (o *ReadCasResponse) SetCas(v []Ca) {
	o.Cas = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadCasResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCasResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadCasResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadCasResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadCasResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cas != nil {
		toSerialize["Cas"] = o.Cas
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadCasResponse struct {
	value *ReadCasResponse
	isSet bool
}

func (v NullableReadCasResponse) Get() *ReadCasResponse {
	return v.value
}

func (v *NullableReadCasResponse) Set(val *ReadCasResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadCasResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadCasResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadCasResponse(val *ReadCasResponse) *NullableReadCasResponse {
	return &NullableReadCasResponse{value: val, isSet: true}
}

func (v NullableReadCasResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadCasResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
