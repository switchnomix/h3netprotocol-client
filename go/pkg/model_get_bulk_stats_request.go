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

// checks if the GetBulkStatsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBulkStatsRequest{}

// GetBulkStatsRequest struct for GetBulkStatsRequest
type GetBulkStatsRequest struct {
	Interfaces []GetBulkStatsRequestInterfacesInner `json:"interfaces"`
	StatsTypes []string `json:"statsTypes,omitempty"`
}

type _GetBulkStatsRequest GetBulkStatsRequest

// NewGetBulkStatsRequest instantiates a new GetBulkStatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBulkStatsRequest(interfaces []GetBulkStatsRequestInterfacesInner) *GetBulkStatsRequest {
	this := GetBulkStatsRequest{}
	this.Interfaces = interfaces
	return &this
}

// NewGetBulkStatsRequestWithDefaults instantiates a new GetBulkStatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBulkStatsRequestWithDefaults() *GetBulkStatsRequest {
	this := GetBulkStatsRequest{}
	return &this
}

// GetInterfaces returns the Interfaces field value
func (o *GetBulkStatsRequest) GetInterfaces() []GetBulkStatsRequestInterfacesInner {
	if o == nil {
		var ret []GetBulkStatsRequestInterfacesInner
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *GetBulkStatsRequest) GetInterfacesOk() ([]GetBulkStatsRequestInterfacesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// SetInterfaces sets field value
func (o *GetBulkStatsRequest) SetInterfaces(v []GetBulkStatsRequestInterfacesInner) {
	o.Interfaces = v
}

// GetStatsTypes returns the StatsTypes field value if set, zero value otherwise.
func (o *GetBulkStatsRequest) GetStatsTypes() []string {
	if o == nil || IsNil(o.StatsTypes) {
		var ret []string
		return ret
	}
	return o.StatsTypes
}

// GetStatsTypesOk returns a tuple with the StatsTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBulkStatsRequest) GetStatsTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.StatsTypes) {
		return nil, false
	}
	return o.StatsTypes, true
}

// HasStatsTypes returns a boolean if a field has been set.
func (o *GetBulkStatsRequest) HasStatsTypes() bool {
	if o != nil && !IsNil(o.StatsTypes) {
		return true
	}

	return false
}

// SetStatsTypes gets a reference to the given []string and assigns it to the StatsTypes field.
func (o *GetBulkStatsRequest) SetStatsTypes(v []string) {
	o.StatsTypes = v
}

func (o GetBulkStatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBulkStatsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interfaces"] = o.Interfaces
	if !IsNil(o.StatsTypes) {
		toSerialize["statsTypes"] = o.StatsTypes
	}
	return toSerialize, nil
}

func (o *GetBulkStatsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interfaces",
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

	varGetBulkStatsRequest := _GetBulkStatsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBulkStatsRequest)

	if err != nil {
		return err
	}

	*o = GetBulkStatsRequest(varGetBulkStatsRequest)

	return err
}

type NullableGetBulkStatsRequest struct {
	value *GetBulkStatsRequest
	isSet bool
}

func (v NullableGetBulkStatsRequest) Get() *GetBulkStatsRequest {
	return v.value
}

func (v *NullableGetBulkStatsRequest) Set(val *GetBulkStatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBulkStatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBulkStatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBulkStatsRequest(val *GetBulkStatsRequest) *NullableGetBulkStatsRequest {
	return &NullableGetBulkStatsRequest{value: val, isSet: true}
}

func (v NullableGetBulkStatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBulkStatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


