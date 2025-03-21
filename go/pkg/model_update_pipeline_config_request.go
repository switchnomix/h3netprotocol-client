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
	"time"
	"bytes"
	"fmt"
)

// checks if the UpdatePipelineConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePipelineConfigRequest{}

// UpdatePipelineConfigRequest struct for UpdatePipelineConfigRequest
type UpdatePipelineConfigRequest struct {
	Type string `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	// Client-generated request ID for tracking
	RequestId string `json:"requestId"`
	SwitchId string `json:"switchId"`
	Tables []MatchActionTable `json:"tables"`
	Vlans []VLAN `json:"vlans,omitempty"`
	Acls []ACL `json:"acls,omitempty"`
}

type _UpdatePipelineConfigRequest UpdatePipelineConfigRequest

// NewUpdatePipelineConfigRequest instantiates a new UpdatePipelineConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePipelineConfigRequest(type_ string, timestamp time.Time, requestId string, switchId string, tables []MatchActionTable) *UpdatePipelineConfigRequest {
	this := UpdatePipelineConfigRequest{}
	this.Type = type_
	this.Timestamp = timestamp
	this.RequestId = requestId
	this.SwitchId = switchId
	this.Tables = tables
	return &this
}

// NewUpdatePipelineConfigRequestWithDefaults instantiates a new UpdatePipelineConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePipelineConfigRequestWithDefaults() *UpdatePipelineConfigRequest {
	this := UpdatePipelineConfigRequest{}
	return &this
}

// GetType returns the Type field value
func (o *UpdatePipelineConfigRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdatePipelineConfigRequest) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *UpdatePipelineConfigRequest) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UpdatePipelineConfigRequest) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetRequestId returns the RequestId field value
func (o *UpdatePipelineConfigRequest) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *UpdatePipelineConfigRequest) SetRequestId(v string) {
	o.RequestId = v
}

// GetSwitchId returns the SwitchId field value
func (o *UpdatePipelineConfigRequest) GetSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetSwitchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SwitchId, true
}

// SetSwitchId sets field value
func (o *UpdatePipelineConfigRequest) SetSwitchId(v string) {
	o.SwitchId = v
}

// GetTables returns the Tables field value
func (o *UpdatePipelineConfigRequest) GetTables() []MatchActionTable {
	if o == nil {
		var ret []MatchActionTable
		return ret
	}

	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetTablesOk() ([]MatchActionTable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tables, true
}

// SetTables sets field value
func (o *UpdatePipelineConfigRequest) SetTables(v []MatchActionTable) {
	o.Tables = v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *UpdatePipelineConfigRequest) GetVlans() []VLAN {
	if o == nil || IsNil(o.Vlans) {
		var ret []VLAN
		return ret
	}
	return o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetVlansOk() ([]VLAN, bool) {
	if o == nil || IsNil(o.Vlans) {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *UpdatePipelineConfigRequest) HasVlans() bool {
	if o != nil && !IsNil(o.Vlans) {
		return true
	}

	return false
}

// SetVlans gets a reference to the given []VLAN and assigns it to the Vlans field.
func (o *UpdatePipelineConfigRequest) SetVlans(v []VLAN) {
	o.Vlans = v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *UpdatePipelineConfigRequest) GetAcls() []ACL {
	if o == nil || IsNil(o.Acls) {
		var ret []ACL
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineConfigRequest) GetAclsOk() ([]ACL, bool) {
	if o == nil || IsNil(o.Acls) {
		return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *UpdatePipelineConfigRequest) HasAcls() bool {
	if o != nil && !IsNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []ACL and assigns it to the Acls field.
func (o *UpdatePipelineConfigRequest) SetAcls(v []ACL) {
	o.Acls = v
}

func (o UpdatePipelineConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePipelineConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["requestId"] = o.RequestId
	toSerialize["switchId"] = o.SwitchId
	toSerialize["tables"] = o.Tables
	if !IsNil(o.Vlans) {
		toSerialize["vlans"] = o.Vlans
	}
	if !IsNil(o.Acls) {
		toSerialize["acls"] = o.Acls
	}
	return toSerialize, nil
}

func (o *UpdatePipelineConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timestamp",
		"requestId",
		"switchId",
		"tables",
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

	varUpdatePipelineConfigRequest := _UpdatePipelineConfigRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdatePipelineConfigRequest)

	if err != nil {
		return err
	}

	*o = UpdatePipelineConfigRequest(varUpdatePipelineConfigRequest)

	return err
}

type NullableUpdatePipelineConfigRequest struct {
	value *UpdatePipelineConfigRequest
	isSet bool
}

func (v NullableUpdatePipelineConfigRequest) Get() *UpdatePipelineConfigRequest {
	return v.value
}

func (v *NullableUpdatePipelineConfigRequest) Set(val *UpdatePipelineConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePipelineConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePipelineConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePipelineConfigRequest(val *UpdatePipelineConfigRequest) *NullableUpdatePipelineConfigRequest {
	return &NullableUpdatePipelineConfigRequest{value: val, isSet: true}
}

func (v NullableUpdatePipelineConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePipelineConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


