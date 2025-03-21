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

// checks if the NetworkDiscoveryResponseMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDiscoveryResponseMessage{}

// NetworkDiscoveryResponseMessage struct for NetworkDiscoveryResponseMessage
type NetworkDiscoveryResponseMessage struct {
	Type string `json:"type"`
	// ISO 8601 formatted timestamp with timezone
	Timestamp time.Time `json:"timestamp"`
	// Schema version
	SchemaVersion string `json:"schemaVersion"`
	DiscoveryType string `json:"discoveryType"`
	DiscoveredDevices []DiscoveredNeighbor `json:"discoveredDevices"`
	Topology *NetworkDiscoveryResponseMessageAllOfTopology `json:"topology,omitempty"`
}

type _NetworkDiscoveryResponseMessage NetworkDiscoveryResponseMessage

// NewNetworkDiscoveryResponseMessage instantiates a new NetworkDiscoveryResponseMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDiscoveryResponseMessage(type_ string, timestamp time.Time, schemaVersion string, discoveryType string, discoveredDevices []DiscoveredNeighbor) *NetworkDiscoveryResponseMessage {
	this := NetworkDiscoveryResponseMessage{}
	this.Type = type_
	this.Timestamp = timestamp
	this.SchemaVersion = schemaVersion
	this.DiscoveryType = discoveryType
	this.DiscoveredDevices = discoveredDevices
	return &this
}

// NewNetworkDiscoveryResponseMessageWithDefaults instantiates a new NetworkDiscoveryResponseMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDiscoveryResponseMessageWithDefaults() *NetworkDiscoveryResponseMessage {
	this := NetworkDiscoveryResponseMessage{}
	return &this
}

// GetType returns the Type field value
func (o *NetworkDiscoveryResponseMessage) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkDiscoveryResponseMessage) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *NetworkDiscoveryResponseMessage) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *NetworkDiscoveryResponseMessage) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *NetworkDiscoveryResponseMessage) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *NetworkDiscoveryResponseMessage) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetDiscoveryType returns the DiscoveryType field value
func (o *NetworkDiscoveryResponseMessage) GetDiscoveryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscoveryType
}

// GetDiscoveryTypeOk returns a tuple with the DiscoveryType field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetDiscoveryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscoveryType, true
}

// SetDiscoveryType sets field value
func (o *NetworkDiscoveryResponseMessage) SetDiscoveryType(v string) {
	o.DiscoveryType = v
}

// GetDiscoveredDevices returns the DiscoveredDevices field value
func (o *NetworkDiscoveryResponseMessage) GetDiscoveredDevices() []DiscoveredNeighbor {
	if o == nil {
		var ret []DiscoveredNeighbor
		return ret
	}

	return o.DiscoveredDevices
}

// GetDiscoveredDevicesOk returns a tuple with the DiscoveredDevices field value
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetDiscoveredDevicesOk() ([]DiscoveredNeighbor, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoveredDevices, true
}

// SetDiscoveredDevices sets field value
func (o *NetworkDiscoveryResponseMessage) SetDiscoveredDevices(v []DiscoveredNeighbor) {
	o.DiscoveredDevices = v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *NetworkDiscoveryResponseMessage) GetTopology() NetworkDiscoveryResponseMessageAllOfTopology {
	if o == nil || IsNil(o.Topology) {
		var ret NetworkDiscoveryResponseMessageAllOfTopology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryResponseMessage) GetTopologyOk() (*NetworkDiscoveryResponseMessageAllOfTopology, bool) {
	if o == nil || IsNil(o.Topology) {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *NetworkDiscoveryResponseMessage) HasTopology() bool {
	if o != nil && !IsNil(o.Topology) {
		return true
	}

	return false
}

// SetTopology gets a reference to the given NetworkDiscoveryResponseMessageAllOfTopology and assigns it to the Topology field.
func (o *NetworkDiscoveryResponseMessage) SetTopology(v NetworkDiscoveryResponseMessageAllOfTopology) {
	o.Topology = &v
}

func (o NetworkDiscoveryResponseMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDiscoveryResponseMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["schemaVersion"] = o.SchemaVersion
	toSerialize["discoveryType"] = o.DiscoveryType
	toSerialize["discoveredDevices"] = o.DiscoveredDevices
	if !IsNil(o.Topology) {
		toSerialize["topology"] = o.Topology
	}
	return toSerialize, nil
}

func (o *NetworkDiscoveryResponseMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timestamp",
		"schemaVersion",
		"discoveryType",
		"discoveredDevices",
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

	varNetworkDiscoveryResponseMessage := _NetworkDiscoveryResponseMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkDiscoveryResponseMessage)

	if err != nil {
		return err
	}

	*o = NetworkDiscoveryResponseMessage(varNetworkDiscoveryResponseMessage)

	return err
}

type NullableNetworkDiscoveryResponseMessage struct {
	value *NetworkDiscoveryResponseMessage
	isSet bool
}

func (v NullableNetworkDiscoveryResponseMessage) Get() *NetworkDiscoveryResponseMessage {
	return v.value
}

func (v *NullableNetworkDiscoveryResponseMessage) Set(val *NetworkDiscoveryResponseMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDiscoveryResponseMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDiscoveryResponseMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDiscoveryResponseMessage(val *NetworkDiscoveryResponseMessage) *NullableNetworkDiscoveryResponseMessage {
	return &NullableNetworkDiscoveryResponseMessage{value: val, isSet: true}
}

func (v NullableNetworkDiscoveryResponseMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDiscoveryResponseMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


