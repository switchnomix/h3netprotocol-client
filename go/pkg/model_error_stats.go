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

// checks if the ErrorStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorStats{}

// ErrorStats Detailed error statistics
type ErrorStats struct {
	Crc *float32 `json:"crc,omitempty"`
	Fragment *float32 `json:"fragment,omitempty"`
	Jabber *float32 `json:"jabber,omitempty"`
	Alignment *float32 `json:"alignment,omitempty"`
	Symbol *float32 `json:"symbol,omitempty"`
	Unknown *float32 `json:"unknown,omitempty"`
}

// NewErrorStats instantiates a new ErrorStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorStats() *ErrorStats {
	this := ErrorStats{}
	return &this
}

// NewErrorStatsWithDefaults instantiates a new ErrorStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorStatsWithDefaults() *ErrorStats {
	this := ErrorStats{}
	return &this
}

// GetCrc returns the Crc field value if set, zero value otherwise.
func (o *ErrorStats) GetCrc() float32 {
	if o == nil || IsNil(o.Crc) {
		var ret float32
		return ret
	}
	return *o.Crc
}

// GetCrcOk returns a tuple with the Crc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetCrcOk() (*float32, bool) {
	if o == nil || IsNil(o.Crc) {
		return nil, false
	}
	return o.Crc, true
}

// HasCrc returns a boolean if a field has been set.
func (o *ErrorStats) HasCrc() bool {
	if o != nil && !IsNil(o.Crc) {
		return true
	}

	return false
}

// SetCrc gets a reference to the given float32 and assigns it to the Crc field.
func (o *ErrorStats) SetCrc(v float32) {
	o.Crc = &v
}

// GetFragment returns the Fragment field value if set, zero value otherwise.
func (o *ErrorStats) GetFragment() float32 {
	if o == nil || IsNil(o.Fragment) {
		var ret float32
		return ret
	}
	return *o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetFragmentOk() (*float32, bool) {
	if o == nil || IsNil(o.Fragment) {
		return nil, false
	}
	return o.Fragment, true
}

// HasFragment returns a boolean if a field has been set.
func (o *ErrorStats) HasFragment() bool {
	if o != nil && !IsNil(o.Fragment) {
		return true
	}

	return false
}

// SetFragment gets a reference to the given float32 and assigns it to the Fragment field.
func (o *ErrorStats) SetFragment(v float32) {
	o.Fragment = &v
}

// GetJabber returns the Jabber field value if set, zero value otherwise.
func (o *ErrorStats) GetJabber() float32 {
	if o == nil || IsNil(o.Jabber) {
		var ret float32
		return ret
	}
	return *o.Jabber
}

// GetJabberOk returns a tuple with the Jabber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetJabberOk() (*float32, bool) {
	if o == nil || IsNil(o.Jabber) {
		return nil, false
	}
	return o.Jabber, true
}

// HasJabber returns a boolean if a field has been set.
func (o *ErrorStats) HasJabber() bool {
	if o != nil && !IsNil(o.Jabber) {
		return true
	}

	return false
}

// SetJabber gets a reference to the given float32 and assigns it to the Jabber field.
func (o *ErrorStats) SetJabber(v float32) {
	o.Jabber = &v
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *ErrorStats) GetAlignment() float32 {
	if o == nil || IsNil(o.Alignment) {
		var ret float32
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetAlignmentOk() (*float32, bool) {
	if o == nil || IsNil(o.Alignment) {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *ErrorStats) HasAlignment() bool {
	if o != nil && !IsNil(o.Alignment) {
		return true
	}

	return false
}

// SetAlignment gets a reference to the given float32 and assigns it to the Alignment field.
func (o *ErrorStats) SetAlignment(v float32) {
	o.Alignment = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ErrorStats) GetSymbol() float32 {
	if o == nil || IsNil(o.Symbol) {
		var ret float32
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetSymbolOk() (*float32, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ErrorStats) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given float32 and assigns it to the Symbol field.
func (o *ErrorStats) SetSymbol(v float32) {
	o.Symbol = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *ErrorStats) GetUnknown() float32 {
	if o == nil || IsNil(o.Unknown) {
		var ret float32
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorStats) GetUnknownOk() (*float32, bool) {
	if o == nil || IsNil(o.Unknown) {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *ErrorStats) HasUnknown() bool {
	if o != nil && !IsNil(o.Unknown) {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given float32 and assigns it to the Unknown field.
func (o *ErrorStats) SetUnknown(v float32) {
	o.Unknown = &v
}

func (o ErrorStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Crc) {
		toSerialize["crc"] = o.Crc
	}
	if !IsNil(o.Fragment) {
		toSerialize["fragment"] = o.Fragment
	}
	if !IsNil(o.Jabber) {
		toSerialize["jabber"] = o.Jabber
	}
	if !IsNil(o.Alignment) {
		toSerialize["alignment"] = o.Alignment
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Unknown) {
		toSerialize["unknown"] = o.Unknown
	}
	return toSerialize, nil
}

type NullableErrorStats struct {
	value *ErrorStats
	isSet bool
}

func (v NullableErrorStats) Get() *ErrorStats {
	return v.value
}

func (v *NullableErrorStats) Set(val *ErrorStats) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorStats) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorStats(val *ErrorStats) *NullableErrorStats {
	return &NullableErrorStats{value: val, isSet: true}
}

func (v NullableErrorStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


