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
	"fmt"
)

// CustomerAddressCustomerResponseKind the model 'CustomerAddressCustomerResponseKind'
type CustomerAddressCustomerResponseKind string

// List of customerAddressCustomerResponseKind
const (
	CUSTOMERADDRESSCUSTOMERRESPONSEKIND_SHIPPING CustomerAddressCustomerResponseKind = "SHIPPING"
	CUSTOMERADDRESSCUSTOMERRESPONSEKIND_BILLING  CustomerAddressCustomerResponseKind = "BILLING"
	CUSTOMERADDRESSCUSTOMERRESPONSEKIND_BOTH     CustomerAddressCustomerResponseKind = "BOTH"
)

// All allowed values of CustomerAddressCustomerResponseKind enum
var AllowedCustomerAddressCustomerResponseKindEnumValues = []CustomerAddressCustomerResponseKind{
	"SHIPPING",
	"BILLING",
	"BOTH",
}

func (v *CustomerAddressCustomerResponseKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomerAddressCustomerResponseKind(value)
	for _, existing := range AllowedCustomerAddressCustomerResponseKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomerAddressCustomerResponseKind", value)
}

// NewCustomerAddressCustomerResponseKindFromValue returns a pointer to a valid CustomerAddressCustomerResponseKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerAddressCustomerResponseKindFromValue(v string) (*CustomerAddressCustomerResponseKind, error) {
	ev := CustomerAddressCustomerResponseKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerAddressCustomerResponseKind: valid values are %v", v, AllowedCustomerAddressCustomerResponseKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerAddressCustomerResponseKind) IsValid() bool {
	for _, existing := range AllowedCustomerAddressCustomerResponseKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customerAddressCustomerResponseKind value
func (v CustomerAddressCustomerResponseKind) Ptr() *CustomerAddressCustomerResponseKind {
	return &v
}

type NullableCustomerAddressCustomerResponseKind struct {
	value *CustomerAddressCustomerResponseKind
	isSet bool
}

func (v NullableCustomerAddressCustomerResponseKind) Get() *CustomerAddressCustomerResponseKind {
	return v.value
}

func (v *NullableCustomerAddressCustomerResponseKind) Set(val *CustomerAddressCustomerResponseKind) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCustomerResponseKind) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCustomerResponseKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCustomerResponseKind(val *CustomerAddressCustomerResponseKind) *NullableCustomerAddressCustomerResponseKind {
	return &NullableCustomerAddressCustomerResponseKind{value: val, isSet: true}
}

func (v NullableCustomerAddressCustomerResponseKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCustomerResponseKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
