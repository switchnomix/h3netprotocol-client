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

// checks if the BaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseResponse{}

// BaseResponse struct for BaseResponse
type BaseResponse struct {
	// Message type identifier
	Type string `json:"type"`
	// ISO 8601 formatted timestamp with timezone
	Timestamp time.Time `json:"timestamp"`
	// Schema version
	SchemaVersion string `json:"schemaVersion"`
	// Response status
	Status string `json:"status"`
	// Original request ID
	RequestId *string `json:"requestId,omitempty"`
}

type _BaseResponse BaseResponse

// NewBaseResponse instantiates a new BaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseResponse(type_ string, timestamp time.Time, schemaVersion string, status string) *BaseResponse {
	this := BaseResponse{}
	this.Type = type_
	this.Timestamp = timestamp
	this.SchemaVersion = schemaVersion
	this.Status = status
	return &this
}

// NewBaseResponseWithDefaults instantiates a new BaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseResponseWithDefaults() *BaseResponse {
	this := BaseResponse{}
	return &this
}

// GetType returns the Type field value
func (o *BaseResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BaseResponse) SetType(v string) {
	o.Type = v
}

// GetTimestamp returns the Timestamp field value
func (o *BaseResponse) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BaseResponse) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *BaseResponse) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *BaseResponse) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetStatus returns the Status field value
func (o *BaseResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BaseResponse) SetStatus(v string) {
	o.Status = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *BaseResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *BaseResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *BaseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o BaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["schemaVersion"] = o.SchemaVersion
	toSerialize["status"] = o.Status
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

func (o *BaseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"timestamp",
		"schemaVersion",
		"status",
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

	varBaseResponse := _BaseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseResponse)

	if err != nil {
		return err
	}

	*o = BaseResponse(varBaseResponse)

	return err
}

type NullableBaseResponse struct {
	value *BaseResponse
	isSet bool
}

func (v NullableBaseResponse) Get() *BaseResponse {
	return v.value
}

func (v *NullableBaseResponse) Set(val *BaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseResponse(val *BaseResponse) *NullableBaseResponse {
	return &NullableBaseResponse{value: val, isSet: true}
}

func (v NullableBaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


