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

// checks if the MetricThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricThresholds{}

// MetricThresholds struct for MetricThresholds
type MetricThresholds struct {
	Cpu *UpdateMetricThresholdsRequestCpu `json:"cpu,omitempty"`
	Memory *UpdateMetricThresholdsRequestCpu `json:"memory,omitempty"`
	Temperature *UpdateMetricThresholdsRequestCpu `json:"temperature,omitempty"`
}

// NewMetricThresholds instantiates a new MetricThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricThresholds() *MetricThresholds {
	this := MetricThresholds{}
	return &this
}

// NewMetricThresholdsWithDefaults instantiates a new MetricThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricThresholdsWithDefaults() *MetricThresholds {
	this := MetricThresholds{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *MetricThresholds) GetCpu() UpdateMetricThresholdsRequestCpu {
	if o == nil || IsNil(o.Cpu) {
		var ret UpdateMetricThresholdsRequestCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricThresholds) GetCpuOk() (*UpdateMetricThresholdsRequestCpu, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *MetricThresholds) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given UpdateMetricThresholdsRequestCpu and assigns it to the Cpu field.
func (o *MetricThresholds) SetCpu(v UpdateMetricThresholdsRequestCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *MetricThresholds) GetMemory() UpdateMetricThresholdsRequestCpu {
	if o == nil || IsNil(o.Memory) {
		var ret UpdateMetricThresholdsRequestCpu
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricThresholds) GetMemoryOk() (*UpdateMetricThresholdsRequestCpu, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *MetricThresholds) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given UpdateMetricThresholdsRequestCpu and assigns it to the Memory field.
func (o *MetricThresholds) SetMemory(v UpdateMetricThresholdsRequestCpu) {
	o.Memory = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *MetricThresholds) GetTemperature() UpdateMetricThresholdsRequestCpu {
	if o == nil || IsNil(o.Temperature) {
		var ret UpdateMetricThresholdsRequestCpu
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricThresholds) GetTemperatureOk() (*UpdateMetricThresholdsRequestCpu, bool) {
	if o == nil || IsNil(o.Temperature) {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *MetricThresholds) HasTemperature() bool {
	if o != nil && !IsNil(o.Temperature) {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given UpdateMetricThresholdsRequestCpu and assigns it to the Temperature field.
func (o *MetricThresholds) SetTemperature(v UpdateMetricThresholdsRequestCpu) {
	o.Temperature = &v
}

func (o MetricThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Temperature) {
		toSerialize["temperature"] = o.Temperature
	}
	return toSerialize, nil
}

type NullableMetricThresholds struct {
	value *MetricThresholds
	isSet bool
}

func (v NullableMetricThresholds) Get() *MetricThresholds {
	return v.value
}

func (v *NullableMetricThresholds) Set(val *MetricThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricThresholds(val *MetricThresholds) *NullableMetricThresholds {
	return &NullableMetricThresholds{value: val, isSet: true}
}

func (v NullableMetricThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


