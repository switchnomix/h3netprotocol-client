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

// checks if the ACLActionRedirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACLActionRedirect{}

// ACLActionRedirect struct for ACLActionRedirect
type ACLActionRedirect struct {
	// Interface to redirect traffic to
	Interface *string `json:"interface,omitempty"`
	// Next hop IP address
	NextHop *string `json:"nextHop,omitempty"`
}

// NewACLActionRedirect instantiates a new ACLActionRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLActionRedirect() *ACLActionRedirect {
	this := ACLActionRedirect{}
	return &this
}

// NewACLActionRedirectWithDefaults instantiates a new ACLActionRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLActionRedirectWithDefaults() *ACLActionRedirect {
	this := ACLActionRedirect{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *ACLActionRedirect) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLActionRedirect) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *ACLActionRedirect) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *ACLActionRedirect) SetInterface(v string) {
	o.Interface = &v
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *ACLActionRedirect) GetNextHop() string {
	if o == nil || IsNil(o.NextHop) {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLActionRedirect) GetNextHopOk() (*string, bool) {
	if o == nil || IsNil(o.NextHop) {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *ACLActionRedirect) HasNextHop() bool {
	if o != nil && !IsNil(o.NextHop) {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *ACLActionRedirect) SetNextHop(v string) {
	o.NextHop = &v
}

func (o ACLActionRedirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACLActionRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.NextHop) {
		toSerialize["nextHop"] = o.NextHop
	}
	return toSerialize, nil
}

type NullableACLActionRedirect struct {
	value *ACLActionRedirect
	isSet bool
}

func (v NullableACLActionRedirect) Get() *ACLActionRedirect {
	return v.value
}

func (v *NullableACLActionRedirect) Set(val *ACLActionRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableACLActionRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableACLActionRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLActionRedirect(val *ACLActionRedirect) *NullableACLActionRedirect {
	return &NullableACLActionRedirect{value: val, isSet: true}
}

func (v NullableACLActionRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLActionRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


