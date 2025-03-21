/*
H3Net Protocol API

H3Net Protocol schema with various messages and configurations.

API version: 1.0.0
Contact: h3netprotocol@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package h3netclient

import (
	"encoding/json"
)

// checks if the SchemasDeviceSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasDeviceSettings{}

// SchemasDeviceSettings struct for SchemasDeviceSettings
type SchemasDeviceSettings struct {
	// Unique identifier for network device
	DeviceId *string `json:"deviceId,omitempty"`
}

// NewSchemasDeviceSettings instantiates a new SchemasDeviceSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasDeviceSettings() *SchemasDeviceSettings {
	this := SchemasDeviceSettings{}
	return &this
}

// NewSchemasDeviceSettingsWithDefaults instantiates a new SchemasDeviceSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasDeviceSettingsWithDefaults() *SchemasDeviceSettings {
	this := SchemasDeviceSettings{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *SchemasDeviceSettings) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasDeviceSettings) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *SchemasDeviceSettings) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *SchemasDeviceSettings) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o SchemasDeviceSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasDeviceSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	return toSerialize, nil
}

type NullableSchemasDeviceSettings struct {
	value *SchemasDeviceSettings
	isSet bool
}

func (v NullableSchemasDeviceSettings) Get() *SchemasDeviceSettings {
	return v.value
}

func (v *NullableSchemasDeviceSettings) Set(val *SchemasDeviceSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasDeviceSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasDeviceSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasDeviceSettings(val *SchemasDeviceSettings) *NullableSchemasDeviceSettings {
	return &NullableSchemasDeviceSettings{value: val, isSet: true}
}

func (v NullableSchemasDeviceSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasDeviceSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


