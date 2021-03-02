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

// Volume Information about the volume.
type Volume struct {
	// The number of I/O operations per second (IOPS):<br /> - For `io1` volumes, the number of provisioned IOPS<br /> - For `gp2` volumes, the baseline performance of the volume
	Iops *int32 `json:"Iops,omitempty"`
	// Information about your volume attachment.
	LinkedVolumes *[]LinkedVolume `json:"LinkedVolumes,omitempty"`
	// The size of the volume, in gibibytes (GiB).
	Size *int32 `json:"Size,omitempty"`
	// The snapshot from which the volume was created.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The state of the volume (`creating` \\| `available` \\| `in-use` \\| `updating` \\| `deleting` \\| `error`).
	State *string `json:"State,omitempty"`
	// The Subregion in which the volume was created.
	SubregionName *string `json:"SubregionName,omitempty"`
	// One or more tags associated with the volume.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the volume.
	VolumeId *string `json:"VolumeId,omitempty"`
	// The type of the volume (`standard` \\| `gp2` \\| `io1`).
	VolumeType *string `json:"VolumeType,omitempty"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume() *Volume {
	this := Volume{}
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *Volume) GetIops() int32 {
	if o == nil || o.Iops == nil {
		var ret int32
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIopsOk() (*int32, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *Volume) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int32 and assigns it to the Iops field.
func (o *Volume) SetIops(v int32) {
	o.Iops = &v
}

// GetLinkedVolumes returns the LinkedVolumes field value if set, zero value otherwise.
func (o *Volume) GetLinkedVolumes() []LinkedVolume {
	if o == nil || o.LinkedVolumes == nil {
		var ret []LinkedVolume
		return ret
	}
	return *o.LinkedVolumes
}

// GetLinkedVolumesOk returns a tuple with the LinkedVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetLinkedVolumesOk() (*[]LinkedVolume, bool) {
	if o == nil || o.LinkedVolumes == nil {
		return nil, false
	}
	return o.LinkedVolumes, true
}

// HasLinkedVolumes returns a boolean if a field has been set.
func (o *Volume) HasLinkedVolumes() bool {
	if o != nil && o.LinkedVolumes != nil {
		return true
	}

	return false
}

// SetLinkedVolumes gets a reference to the given []LinkedVolume and assigns it to the LinkedVolumes field.
func (o *Volume) SetLinkedVolumes(v []LinkedVolume) {
	o.LinkedVolumes = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Volume) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Volume) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Volume) SetSize(v int32) {
	o.Size = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *Volume) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *Volume) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *Volume) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Volume) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Volume) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Volume) SetState(v string) {
	o.State = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *Volume) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *Volume) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *Volume) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Volume) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Volume) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Volume) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *Volume) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *Volume) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *Volume) SetVolumeId(v string) {
	o.VolumeId = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *Volume) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *Volume) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *Volume) SetVolumeType(v string) {
	o.VolumeType = &v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.LinkedVolumes != nil {
		toSerialize["LinkedVolumes"] = o.LinkedVolumes
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.SnapshotId != nil {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	return json.Marshal(toSerialize)
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
