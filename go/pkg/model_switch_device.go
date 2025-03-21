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
	"bytes"
	"fmt"
)

// checks if the SwitchDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchDevice{}

// SwitchDevice struct for SwitchDevice
type SwitchDevice struct {
	DeviceId string `json:"deviceId"`
	DeviceModel *string `json:"deviceModel,omitempty"`
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	NetworkInterfaces []SchemasNetworkInterface `json:"networkInterfaces,omitempty"`
	DeviceConfig *SwitchConfiguration `json:"deviceConfig,omitempty"`
}

type _SwitchDevice SwitchDevice

// NewSwitchDevice instantiates a new SwitchDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchDevice(deviceId string) *SwitchDevice {
	this := SwitchDevice{}
	this.DeviceId = deviceId
	return &this
}

// NewSwitchDeviceWithDefaults instantiates a new SwitchDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchDeviceWithDefaults() *SwitchDevice {
	this := SwitchDevice{}
	return &this
}

// GetDeviceId returns the DeviceId field value
func (o *SwitchDevice) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *SwitchDevice) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *SwitchDevice) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *SwitchDevice) GetDeviceModel() string {
	if o == nil || IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchDevice) GetDeviceModelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *SwitchDevice) HasDeviceModel() bool {
	if o != nil && !IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *SwitchDevice) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *SwitchDevice) GetFirmwareVersion() string {
	if o == nil || IsNil(o.FirmwareVersion) {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchDevice) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FirmwareVersion) {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *SwitchDevice) HasFirmwareVersion() bool {
	if o != nil && !IsNil(o.FirmwareVersion) {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *SwitchDevice) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *SwitchDevice) GetNetworkInterfaces() []SchemasNetworkInterface {
	if o == nil || IsNil(o.NetworkInterfaces) {
		var ret []SchemasNetworkInterface
		return ret
	}
	return o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchDevice) GetNetworkInterfacesOk() ([]SchemasNetworkInterface, bool) {
	if o == nil || IsNil(o.NetworkInterfaces) {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *SwitchDevice) HasNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NetworkInterfaces) {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []SchemasNetworkInterface and assigns it to the NetworkInterfaces field.
func (o *SwitchDevice) SetNetworkInterfaces(v []SchemasNetworkInterface) {
	o.NetworkInterfaces = v
}

// GetDeviceConfig returns the DeviceConfig field value if set, zero value otherwise.
func (o *SwitchDevice) GetDeviceConfig() SwitchConfiguration {
	if o == nil || IsNil(o.DeviceConfig) {
		var ret SwitchConfiguration
		return ret
	}
	return *o.DeviceConfig
}

// GetDeviceConfigOk returns a tuple with the DeviceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchDevice) GetDeviceConfigOk() (*SwitchConfiguration, bool) {
	if o == nil || IsNil(o.DeviceConfig) {
		return nil, false
	}
	return o.DeviceConfig, true
}

// HasDeviceConfig returns a boolean if a field has been set.
func (o *SwitchDevice) HasDeviceConfig() bool {
	if o != nil && !IsNil(o.DeviceConfig) {
		return true
	}

	return false
}

// SetDeviceConfig gets a reference to the given SwitchConfiguration and assigns it to the DeviceConfig field.
func (o *SwitchDevice) SetDeviceConfig(v SwitchConfiguration) {
	o.DeviceConfig = &v
}

func (o SwitchDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceId"] = o.DeviceId
	if !IsNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !IsNil(o.FirmwareVersion) {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if !IsNil(o.NetworkInterfaces) {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if !IsNil(o.DeviceConfig) {
		toSerialize["deviceConfig"] = o.DeviceConfig
	}
	return toSerialize, nil
}

func (o *SwitchDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSwitchDevice := _SwitchDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSwitchDevice)

	if err != nil {
		return err
	}

	*o = SwitchDevice(varSwitchDevice)

	return err
}

type NullableSwitchDevice struct {
	value *SwitchDevice
	isSet bool
}

func (v NullableSwitchDevice) Get() *SwitchDevice {
	return v.value
}

func (v *NullableSwitchDevice) Set(val *SwitchDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchDevice(val *SwitchDevice) *NullableSwitchDevice {
	return &NullableSwitchDevice{value: val, isSet: true}
}

func (v NullableSwitchDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


