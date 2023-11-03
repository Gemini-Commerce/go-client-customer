/*
CDP Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomerPassword struct for CustomerPassword
type CustomerPassword struct {
	Data *map[string]string `json:"data,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Type *PasswordPasswordType `json:"type,omitempty"`
}

// NewCustomerPassword instantiates a new CustomerPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPassword() *CustomerPassword {
	this := CustomerPassword{}
	var type_ PasswordPasswordType = PASSWORDPASSWORDTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewCustomerPasswordWithDefaults instantiates a new CustomerPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordWithDefaults() *CustomerPassword {
	this := CustomerPassword{}
	var type_ PasswordPasswordType = PASSWORDPASSWORDTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerPassword) GetData() map[string]string {
	if o == nil || isNil(o.Data) {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPassword) GetDataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerPassword) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *CustomerPassword) SetData(v map[string]string) {
	o.Data = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomerPassword) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPassword) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomerPassword) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomerPassword) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerPassword) GetType() PasswordPasswordType {
	if o == nil || isNil(o.Type) {
		var ret PasswordPasswordType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPassword) GetTypeOk() (*PasswordPasswordType, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerPassword) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PasswordPasswordType and assigns it to the Type field.
func (o *CustomerPassword) SetType(v PasswordPasswordType) {
	o.Type = &v
}

func (o CustomerPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPassword struct {
	value *CustomerPassword
	isSet bool
}

func (v NullableCustomerPassword) Get() *CustomerPassword {
	return v.value
}

func (v *NullableCustomerPassword) Set(val *CustomerPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPassword(val *CustomerPassword) *NullableCustomerPassword {
	return &NullableCustomerPassword{value: val, isSet: true}
}

func (v NullableCustomerPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


