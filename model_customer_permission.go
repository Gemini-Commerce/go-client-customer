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
	"fmt"
)

// CustomerPermission the model 'CustomerPermission'
type CustomerPermission string

// List of customerPermission
const (
	CUSTOMERPERMISSION_UNKNOWN CustomerPermission = "PERMISSION_UNKNOWN"
	CUSTOMERPERMISSION_LOGIN CustomerPermission = "PERMISSION_LOGIN"
	CUSTOMERPERMISSION_BUY CustomerPermission = "PERMISSION_BUY"
)

// All allowed values of CustomerPermission enum
var AllowedCustomerPermissionEnumValues = []CustomerPermission{
	"PERMISSION_UNKNOWN",
	"PERMISSION_LOGIN",
	"PERMISSION_BUY",
}

func (v *CustomerPermission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerPermission(value)
	for _, existing := range AllowedCustomerPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerPermission", value)
}

// NewCustomerPermissionFromValue returns a pointer to a valid CustomerPermission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerPermissionFromValue(v string) (*CustomerPermission, error) {
	ev := CustomerPermission(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerPermission: valid values are %v", v, AllowedCustomerPermissionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerPermission) IsValid() bool {
	for _, existing := range AllowedCustomerPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customerPermission value
func (v CustomerPermission) Ptr() *CustomerPermission {
	return &v
}

type NullableCustomerPermission struct {
	value *CustomerPermission
	isSet bool
}

func (v NullableCustomerPermission) Get() *CustomerPermission {
	return v.value
}

func (v *NullableCustomerPermission) Set(val *CustomerPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPermission(val *CustomerPermission) *NullableCustomerPermission {
	return &NullableCustomerPermission{value: val, isSet: true}
}

func (v NullableCustomerPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

