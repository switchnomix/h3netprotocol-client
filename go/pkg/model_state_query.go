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

// checks if the StateQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateQuery{}

// StateQuery struct for StateQuery
type StateQuery struct {
	QueryType string `json:"queryType"`
}

type _StateQuery StateQuery

// NewStateQuery instantiates a new StateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateQuery(queryType string) *StateQuery {
	this := StateQuery{}
	this.QueryType = queryType
	return &this
}

// NewStateQueryWithDefaults instantiates a new StateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateQueryWithDefaults() *StateQuery {
	this := StateQuery{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *StateQuery) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *StateQuery) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *StateQuery) SetQueryType(v string) {
	o.QueryType = v
}

func (o StateQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryType"] = o.QueryType
	return toSerialize, nil
}

func (o *StateQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"queryType",
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

	varStateQuery := _StateQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStateQuery)

	if err != nil {
		return err
	}

	*o = StateQuery(varStateQuery)

	return err
}

type NullableStateQuery struct {
	value *StateQuery
	isSet bool
}

func (v NullableStateQuery) Get() *StateQuery {
	return v.value
}

func (v *NullableStateQuery) Set(val *StateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableStateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableStateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateQuery(val *StateQuery) *NullableStateQuery {
	return &NullableStateQuery{value: val, isSet: true}
}

func (v NullableStateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


