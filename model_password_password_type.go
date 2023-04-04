/*
customer/customer.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer

import (
	"encoding/json"
	"fmt"
)

// PasswordPasswordType the model 'PasswordPasswordType'
type PasswordPasswordType string

// List of PasswordPasswordType
const (
	UNKNOWN PasswordPasswordType = "PASSWORD_TYPE_UNKNOWN"
	MAGENTO1 PasswordPasswordType = "PASSWORD_TYPE_MAGENTO1"
	SYLIUS PasswordPasswordType = "PASSWORD_TYPE_SYLIUS"
	MAGENTO2 PasswordPasswordType = "PASSWORD_TYPE_MAGENTO2"
)

// All allowed values of PasswordPasswordType enum
var AllowedPasswordPasswordTypeEnumValues = []PasswordPasswordType{
	"PASSWORD_TYPE_UNKNOWN",
	"PASSWORD_TYPE_MAGENTO1",
	"PASSWORD_TYPE_SYLIUS",
	"PASSWORD_TYPE_MAGENTO2",
}

func (v *PasswordPasswordType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PasswordPasswordType(value)
	for _, existing := range AllowedPasswordPasswordTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PasswordPasswordType", value)
}

// NewPasswordPasswordTypeFromValue returns a pointer to a valid PasswordPasswordType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPasswordPasswordTypeFromValue(v string) (*PasswordPasswordType, error) {
	ev := PasswordPasswordType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PasswordPasswordType: valid values are %v", v, AllowedPasswordPasswordTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PasswordPasswordType) IsValid() bool {
	for _, existing := range AllowedPasswordPasswordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PasswordPasswordType value
func (v PasswordPasswordType) Ptr() *PasswordPasswordType {
	return &v
}

type NullablePasswordPasswordType struct {
	value *PasswordPasswordType
	isSet bool
}

func (v NullablePasswordPasswordType) Get() *PasswordPasswordType {
	return v.value
}

func (v *NullablePasswordPasswordType) Set(val *PasswordPasswordType) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPasswordType) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPasswordType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPasswordType(val *PasswordPasswordType) *NullablePasswordPasswordType {
	return &NullablePasswordPasswordType{value: val, isSet: true}
}

func (v NullablePasswordPasswordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPasswordType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

