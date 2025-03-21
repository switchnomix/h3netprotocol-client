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

// checks if the MatchActionTable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchActionTable{}

// MatchActionTable struct for MatchActionTable
type MatchActionTable struct {
	Type *string `json:"type,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Schema version
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	// Device identifier
	DeviceId string `json:"deviceId"`
	// Unique table identifier for this device
	TableId int32 `json:"tableId"`
	// Human-readable table name
	Name string `json:"name"`
	// ID of the next table in the chain
	NextTableId *int32 `json:"nextTableId,omitempty"`
	// Request identifier for tracking
	RequestId *string `json:"requestId,omitempty"`
	FlowEntries []FlowEntry `json:"flowEntries,omitempty"`
}

type _MatchActionTable MatchActionTable

// NewMatchActionTable instantiates a new MatchActionTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchActionTable(deviceId string, tableId int32, name string) *MatchActionTable {
	this := MatchActionTable{}
	this.DeviceId = deviceId
	this.TableId = tableId
	this.Name = name
	return &this
}

// NewMatchActionTableWithDefaults instantiates a new MatchActionTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchActionTableWithDefaults() *MatchActionTable {
	this := MatchActionTable{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MatchActionTable) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MatchActionTable) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MatchActionTable) SetType(v string) {
	o.Type = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MatchActionTable) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MatchActionTable) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MatchActionTable) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *MatchActionTable) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *MatchActionTable) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *MatchActionTable) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetDeviceId returns the DeviceId field value
func (o *MatchActionTable) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *MatchActionTable) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetTableId returns the TableId field value
func (o *MatchActionTable) GetTableId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetTableIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableId, true
}

// SetTableId sets field value
func (o *MatchActionTable) SetTableId(v int32) {
	o.TableId = v
}

// GetName returns the Name field value
func (o *MatchActionTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MatchActionTable) SetName(v string) {
	o.Name = v
}

// GetNextTableId returns the NextTableId field value if set, zero value otherwise.
func (o *MatchActionTable) GetNextTableId() int32 {
	if o == nil || IsNil(o.NextTableId) {
		var ret int32
		return ret
	}
	return *o.NextTableId
}

// GetNextTableIdOk returns a tuple with the NextTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetNextTableIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NextTableId) {
		return nil, false
	}
	return o.NextTableId, true
}

// HasNextTableId returns a boolean if a field has been set.
func (o *MatchActionTable) HasNextTableId() bool {
	if o != nil && !IsNil(o.NextTableId) {
		return true
	}

	return false
}

// SetNextTableId gets a reference to the given int32 and assigns it to the NextTableId field.
func (o *MatchActionTable) SetNextTableId(v int32) {
	o.NextTableId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *MatchActionTable) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *MatchActionTable) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *MatchActionTable) SetRequestId(v string) {
	o.RequestId = &v
}

// GetFlowEntries returns the FlowEntries field value if set, zero value otherwise.
func (o *MatchActionTable) GetFlowEntries() []FlowEntry {
	if o == nil || IsNil(o.FlowEntries) {
		var ret []FlowEntry
		return ret
	}
	return o.FlowEntries
}

// GetFlowEntriesOk returns a tuple with the FlowEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchActionTable) GetFlowEntriesOk() ([]FlowEntry, bool) {
	if o == nil || IsNil(o.FlowEntries) {
		return nil, false
	}
	return o.FlowEntries, true
}

// HasFlowEntries returns a boolean if a field has been set.
func (o *MatchActionTable) HasFlowEntries() bool {
	if o != nil && !IsNil(o.FlowEntries) {
		return true
	}

	return false
}

// SetFlowEntries gets a reference to the given []FlowEntry and assigns it to the FlowEntries field.
func (o *MatchActionTable) SetFlowEntries(v []FlowEntry) {
	o.FlowEntries = v
}

func (o MatchActionTable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchActionTable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["tableId"] = o.TableId
	toSerialize["name"] = o.Name
	if !IsNil(o.NextTableId) {
		toSerialize["nextTableId"] = o.NextTableId
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.FlowEntries) {
		toSerialize["flowEntries"] = o.FlowEntries
	}
	return toSerialize, nil
}

func (o *MatchActionTable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deviceId",
		"tableId",
		"name",
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

	varMatchActionTable := _MatchActionTable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatchActionTable)

	if err != nil {
		return err
	}

	*o = MatchActionTable(varMatchActionTable)

	return err
}

type NullableMatchActionTable struct {
	value *MatchActionTable
	isSet bool
}

func (v NullableMatchActionTable) Get() *MatchActionTable {
	return v.value
}

func (v *NullableMatchActionTable) Set(val *MatchActionTable) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchActionTable) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchActionTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchActionTable(val *MatchActionTable) *NullableMatchActionTable {
	return &NullableMatchActionTable{value: val, isSet: true}
}

func (v NullableMatchActionTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchActionTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


