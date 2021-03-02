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

// CreateSnapshotExportTaskRequest struct for CreateSnapshotExportTaskRequest
type CreateSnapshotExportTaskRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun    *bool     `json:"DryRun,omitempty"`
	OsuExport OsuExport `json:"OsuExport"`
	// The ID of the snapshot to export.
	SnapshotId string `json:"SnapshotId"`
}

// NewCreateSnapshotExportTaskRequest instantiates a new CreateSnapshotExportTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotExportTaskRequest(osuExport OsuExport, snapshotId string) *CreateSnapshotExportTaskRequest {
	this := CreateSnapshotExportTaskRequest{}
	this.OsuExport = osuExport
	this.SnapshotId = snapshotId
	return &this
}

// NewCreateSnapshotExportTaskRequestWithDefaults instantiates a new CreateSnapshotExportTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotExportTaskRequestWithDefaults() *CreateSnapshotExportTaskRequest {
	this := CreateSnapshotExportTaskRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateSnapshotExportTaskRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotExportTaskRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateSnapshotExportTaskRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateSnapshotExportTaskRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetOsuExport returns the OsuExport field value
func (o *CreateSnapshotExportTaskRequest) GetOsuExport() OsuExport {
	if o == nil {
		var ret OsuExport
		return ret
	}

	return o.OsuExport
}

// GetOsuExportOk returns a tuple with the OsuExport field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotExportTaskRequest) GetOsuExportOk() (*OsuExport, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsuExport, true
}

// SetOsuExport sets field value
func (o *CreateSnapshotExportTaskRequest) SetOsuExport(v OsuExport) {
	o.OsuExport = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *CreateSnapshotExportTaskRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotExportTaskRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *CreateSnapshotExportTaskRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

func (o CreateSnapshotExportTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["OsuExport"] = o.OsuExport
	}
	if true {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSnapshotExportTaskRequest struct {
	value *CreateSnapshotExportTaskRequest
	isSet bool
}

func (v NullableCreateSnapshotExportTaskRequest) Get() *CreateSnapshotExportTaskRequest {
	return v.value
}

func (v *NullableCreateSnapshotExportTaskRequest) Set(val *CreateSnapshotExportTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotExportTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotExportTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotExportTaskRequest(val *CreateSnapshotExportTaskRequest) *NullableCreateSnapshotExportTaskRequest {
	return &NullableCreateSnapshotExportTaskRequest{value: val, isSet: true}
}

func (v NullableCreateSnapshotExportTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotExportTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
