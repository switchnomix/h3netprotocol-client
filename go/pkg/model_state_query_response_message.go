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

// checks if the StateQueryResponseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateQueryResponseMessage{}

// StateQueryResponseMessage struct for StateQueryResponseMessage
type StateQueryResponseMessage struct {
	Type string `json:"type"`
	// ISO 8601 formatted timestamp with timezone
	Timestamp time.Time `json:"timestamp"`
	// Schema version
	SchemaVersion string `json:"schemaVersion"`
	Device Device `json:"device"`
	Query StateQuery `json:"query"`
	Data map[string]interface{} `json:"data"`
}

type _StateQueryResponseMessage StateQueryResponseMessage

// NewStateQueryResponseMessage instantiates a new StateQueryResponseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateQueryResponseMessage(type_ string, timestamp time.Time, schemaVersion string, device Device, query StateQuery, data map[string]interface{}) *StateQueryResponseMessage {
	this := StateQueryResponseMessage{}
	this.Type = type_
	this.Timestamp = timestamp
	this.SchemaVersion = schemaVersion
	this.Device = device
	this.Query = query
	this.Data = data
	return &this
}

// NewStateQueryResponseMessageWithDefaults instantiates a new StateQueryResponseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateQueryResponseMessageWithDefaults() *StateQueryResponseMessage {
	this := StateQueryResponseMessage{}
	return &this
}

// GetType returns the Type field value
func (o *StateQueryResponseMessage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StateQueryResponseMessage) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *StateQueryResponseMessage) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *StateQueryResponseMessage) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *StateQueryResponseMessage) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *StateQueryResponseMessage) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetDevice returns the Device field value
func (o *StateQueryResponseMessage) GetDevice() Device {
	if o == nil {
		var ret Device
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetDeviceOk() (*Device, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *StateQueryResponseMessage) SetDevice(v Device) {
	o.Device = v
}

// GetQuery returns the Query field value
func (o *StateQueryResponseMessage) GetQuery() StateQuery {
	if o == nil {
		var ret StateQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetQueryOk() (*StateQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *StateQueryResponseMessage) SetQuery(v StateQuery) {
	o.Query = v
}

// GetData returns the Data field value
func (o *StateQueryResponseMessage) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StateQueryResponseMessage) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *StateQueryResponseMessage) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o StateQueryResponseMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateQueryResponseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["schemaVersion"] = o.SchemaVersion
	toSerialize["device"] = o.Device
	toSerialize["query"] = o.Query
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *StateQueryResponseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timestamp",
		"schemaVersion",
		"device",
		"query",
		"data",
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

	varStateQueryResponseMessage := _StateQueryResponseMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStateQueryResponseMessage)

	if err != nil {
		return err
	}

	*o = StateQueryResponseMessage(varStateQueryResponseMessage)

	return err
}

type NullableStateQueryResponseMessage struct {
	value *StateQueryResponseMessage
	isSet bool
}

func (v NullableStateQueryResponseMessage) Get() *StateQueryResponseMessage {
	return v.value
}

func (v *NullableStateQueryResponseMessage) Set(val *StateQueryResponseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableStateQueryResponseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableStateQueryResponseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateQueryResponseMessage(val *StateQueryResponseMessage) *NullableStateQueryResponseMessage {
	return &NullableStateQueryResponseMessage{value: val, isSet: true}
}

func (v NullableStateQueryResponseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateQueryResponseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


