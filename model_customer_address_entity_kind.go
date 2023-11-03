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

// CustomerAddressEntityKind the model 'CustomerAddressEntityKind'
type CustomerAddressEntityKind string

// List of customerAddressEntityKind
const (
	CUSTOMERADDRESSENTITYKIND_SHIPPING CustomerAddressEntityKind = "SHIPPING"
	CUSTOMERADDRESSENTITYKIND_BILLING CustomerAddressEntityKind = "BILLING"
	CUSTOMERADDRESSENTITYKIND_BOTH CustomerAddressEntityKind = "BOTH"
)

// All allowed values of CustomerAddressEntityKind enum
var AllowedCustomerAddressEntityKindEnumValues = []CustomerAddressEntityKind{
	"SHIPPING",
	"BILLING",
	"BOTH",
}

func (v *CustomerAddressEntityKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerAddressEntityKind(value)
	for _, existing := range AllowedCustomerAddressEntityKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerAddressEntityKind", value)
}

// NewCustomerAddressEntityKindFromValue returns a pointer to a valid CustomerAddressEntityKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerAddressEntityKindFromValue(v string) (*CustomerAddressEntityKind, error) {
	ev := CustomerAddressEntityKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerAddressEntityKind: valid values are %v", v, AllowedCustomerAddressEntityKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerAddressEntityKind) IsValid() bool {
	for _, existing := range AllowedCustomerAddressEntityKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customerAddressEntityKind value
func (v CustomerAddressEntityKind) Ptr() *CustomerAddressEntityKind {
	return &v
}

type NullableCustomerAddressEntityKind struct {
	value *CustomerAddressEntityKind
	isSet bool
}

func (v NullableCustomerAddressEntityKind) Get() *CustomerAddressEntityKind {
	return v.value
}

func (v *NullableCustomerAddressEntityKind) Set(val *CustomerAddressEntityKind) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressEntityKind) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressEntityKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressEntityKind(val *CustomerAddressEntityKind) *NullableCustomerAddressEntityKind {
	return &NullableCustomerAddressEntityKind{value: val, isSet: true}
}

func (v NullableCustomerAddressEntityKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressEntityKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

