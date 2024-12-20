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

// SortSortField the model 'SortSortField'
type SortSortField string

// List of SortSortField
const (
	SORTSORTFIELD_UNKNOWN SortSortField = "UNKNOWN"
	SORTSORTFIELD_DATE    SortSortField = "DATE"
)

// All allowed values of SortSortField enum
var AllowedSortSortFieldEnumValues = []SortSortField{
	"UNKNOWN",
	"DATE",
}

func (v *SortSortField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortSortField(value)
	for _, existing := range AllowedSortSortFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortSortField", value)
}

// NewSortSortFieldFromValue returns a pointer to a valid SortSortField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortSortFieldFromValue(v string) (*SortSortField, error) {
	ev := SortSortField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortSortField: valid values are %v", v, AllowedSortSortFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortSortField) IsValid() bool {
	for _, existing := range AllowedSortSortFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortSortField value
func (v SortSortField) Ptr() *SortSortField {
	return &v
}

type NullableSortSortField struct {
	value *SortSortField
	isSet bool
}

func (v NullableSortSortField) Get() *SortSortField {
	return v.value
}

func (v *NullableSortSortField) Set(val *SortSortField) {
	v.value = val
	v.isSet = true
}

func (v NullableSortSortField) IsSet() bool {
	return v.isSet
}

func (v *NullableSortSortField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortSortField(val *SortSortField) *NullableSortSortField {
	return &NullableSortSortField{value: val, isSet: true}
}

func (v NullableSortSortField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortSortField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
