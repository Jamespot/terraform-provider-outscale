/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// UpdateSnapshotRequest struct for UpdateSnapshotRequest
type UpdateSnapshotRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun                    *bool                         `json:"DryRun,omitempty"`
	PermissionsToCreateVolume PermissionsOnResourceCreation `json:"PermissionsToCreateVolume"`
	// The ID of the snapshot.
	SnapshotId string `json:"SnapshotId"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateSnapshotRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateSnapshotRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateSnapshotRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field value
func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolume() PermissionsOnResourceCreation {
	if o == nil {
		var ret PermissionsOnResourceCreation
		return ret
	}

	return o.PermissionsToCreateVolume
}

// SetPermissionsToCreateVolume sets field value
func (o *UpdateSnapshotRequest) SetPermissionsToCreateVolume(v PermissionsOnResourceCreation) {
	o.PermissionsToCreateVolume = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *UpdateSnapshotRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// SetSnapshotId sets field value
func (o *UpdateSnapshotRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

type NullableUpdateSnapshotRequest struct {
	Value        UpdateSnapshotRequest
	ExplicitNull bool
}

func (v NullableUpdateSnapshotRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateSnapshotRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}