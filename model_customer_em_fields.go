/*
CDP Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer

import (
	"encoding/json"
)

// checks if the CustomerEMFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerEMFields{}

// CustomerEMFields struct for CustomerEMFields
type CustomerEMFields struct {
	TenantId             *string `json:"tenantId,omitempty"`
	EntityType           *string `json:"entityType,omitempty"`
	EntityCode           *string `json:"entityCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerEMFields CustomerEMFields

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
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerEMFields) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
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
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CustomerEMFields) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
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
	if o == nil || IsNil(o.EntityCode) {
		var ret string
		return ret
	}
	return *o.EntityCode
}

// GetEntityCodeOk returns a tuple with the EntityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerEMFields) GetEntityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityCode) {
		return nil, false
	}
	return o.EntityCode, true
}

// HasEntityCode returns a boolean if a field has been set.
func (o *CustomerEMFields) HasEntityCode() bool {
	if o != nil && !IsNil(o.EntityCode) {
		return true
	}

	return false
}

// SetEntityCode gets a reference to the given string and assigns it to the EntityCode field.
func (o *CustomerEMFields) SetEntityCode(v string) {
	o.EntityCode = &v
}

func (o CustomerEMFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerEMFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.EntityCode) {
		toSerialize["entityCode"] = o.EntityCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerEMFields) UnmarshalJSON(data []byte) (err error) {
	varCustomerEMFields := _CustomerEMFields{}

	err = json.Unmarshal(data, &varCustomerEMFields)

	if err != nil {
		return err
	}

	*o = CustomerEMFields(varCustomerEMFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "entityType")
		delete(additionalProperties, "entityCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerEMFields) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CustomerEMFields) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
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
