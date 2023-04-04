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

// CustomerAddressCreateRequestKind the model 'CustomerAddressCreateRequestKind'
type CustomerAddressCreateRequestKind string

// List of customerAddressCreateRequestKind
const (
	SHIPPING CustomerAddressCreateRequestKind = "SHIPPING"
	BILLING CustomerAddressCreateRequestKind = "BILLING"
	BOTH CustomerAddressCreateRequestKind = "BOTH"
)

// All allowed values of CustomerAddressCreateRequestKind enum
var AllowedCustomerAddressCreateRequestKindEnumValues = []CustomerAddressCreateRequestKind{
	"SHIPPING",
	"BILLING",
	"BOTH",
}

func (v *CustomerAddressCreateRequestKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerAddressCreateRequestKind(value)
	for _, existing := range AllowedCustomerAddressCreateRequestKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerAddressCreateRequestKind", value)
}

// NewCustomerAddressCreateRequestKindFromValue returns a pointer to a valid CustomerAddressCreateRequestKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerAddressCreateRequestKindFromValue(v string) (*CustomerAddressCreateRequestKind, error) {
	ev := CustomerAddressCreateRequestKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerAddressCreateRequestKind: valid values are %v", v, AllowedCustomerAddressCreateRequestKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerAddressCreateRequestKind) IsValid() bool {
	for _, existing := range AllowedCustomerAddressCreateRequestKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customerAddressCreateRequestKind value
func (v CustomerAddressCreateRequestKind) Ptr() *CustomerAddressCreateRequestKind {
	return &v
}

type NullableCustomerAddressCreateRequestKind struct {
	value *CustomerAddressCreateRequestKind
	isSet bool
}

func (v NullableCustomerAddressCreateRequestKind) Get() *CustomerAddressCreateRequestKind {
	return v.value
}

func (v *NullableCustomerAddressCreateRequestKind) Set(val *CustomerAddressCreateRequestKind) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreateRequestKind) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreateRequestKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreateRequestKind(val *CustomerAddressCreateRequestKind) *NullableCustomerAddressCreateRequestKind {
	return &NullableCustomerAddressCreateRequestKind{value: val, isSet: true}
}

func (v NullableCustomerAddressCreateRequestKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreateRequestKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

