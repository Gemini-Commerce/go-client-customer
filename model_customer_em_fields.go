/*
customer/customer.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomerEMFields struct for CustomerEMFields
type CustomerEMFields struct {
	TenantId *string `json:"tenantId,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
	EntityCode *string `json:"entityCode,omitempty"`
}

// NewCustomerEMFields instantiates a new CustomerEMFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerEMFields() *CustomerEMFields {
	this := CustomerEMFields{}
	return &this
}

// NewCustomerEMFieldsWithDefaults instantiates a new CustomerEMFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerEMFieldsWithDefaults() *CustomerEMFields {
	this := CustomerEMFields{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerEMFields) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerEMFields) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerEMFields) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CustomerEMFields) GetEntityType() string {
	if o == nil || isNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetEntityTypeOk() (*string, bool) {
	if o == nil || isNil(o.EntityType) {
    return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CustomerEMFields) HasEntityType() bool {
	if o != nil && !isNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CustomerEMFields) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityCode returns the EntityCode field value if set, zero value otherwise.
func (o *CustomerEMFields) GetEntityCode() string {
	if o == nil || isNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetEntityCodeOk() (*string, bool) {
	if o == nil || isNil(o.EntityCode) {
    return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *CustomerEMFields) HasEntityCode() bool {
	if o != nil && !isNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *CustomerEMFields) SetEntityCode(v string) {
	o.EntityCode = &v
}

func (o CustomerEMFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !isNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerEMFields struct {
	value *CustomerEMFields
	isSet bool
}

func (v NullableCustomerEMFields) Get() *CustomerEMFields {
	return v.value
}

func (v *NullableCustomerEMFields) Set(val *CustomerEMFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerEMFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerEMFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerEMFields(val *CustomerEMFields) *NullableCustomerEMFields {
	return &NullableCustomerEMFields{value: val, isSet: true}
}

func (v NullableCustomerEMFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerEMFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


