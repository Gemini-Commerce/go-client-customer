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

// BulkUpdateRequestAction the model 'BulkUpdateRequestAction'
type BulkUpdateRequestAction string

// List of BulkUpdateRequestAction
const (
	ENABLE BulkUpdateRequestAction = "ENABLE"
	DISABLE BulkUpdateRequestAction = "DISABLE"
	DELETE BulkUpdateRequestAction = "DELETE"
)

// All allowed values of BulkUpdateRequestAction enum
var AllowedBulkUpdateRequestActionEnumValues = []BulkUpdateRequestAction{
	"ENABLE",
	"DISABLE",
	"DELETE",
}

func (v *BulkUpdateRequestAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkUpdateRequestAction(value)
	for _, existing := range AllowedBulkUpdateRequestActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkUpdateRequestAction", value)
}

// NewBulkUpdateRequestActionFromValue returns a pointer to a valid BulkUpdateRequestAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkUpdateRequestActionFromValue(v string) (*BulkUpdateRequestAction, error) {
	ev := BulkUpdateRequestAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkUpdateRequestAction: valid values are %v", v, AllowedBulkUpdateRequestActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkUpdateRequestAction) IsValid() bool {
	for _, existing := range AllowedBulkUpdateRequestActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BulkUpdateRequestAction value
func (v BulkUpdateRequestAction) Ptr() *BulkUpdateRequestAction {
	return &v
}

type NullableBulkUpdateRequestAction struct {
	value *BulkUpdateRequestAction
	isSet bool
}

func (v NullableBulkUpdateRequestAction) Get() *BulkUpdateRequestAction {
	return v.value
}

func (v *NullableBulkUpdateRequestAction) Set(val *BulkUpdateRequestAction) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateRequestAction) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateRequestAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateRequestAction(val *BulkUpdateRequestAction) *NullableBulkUpdateRequestAction {
	return &NullableBulkUpdateRequestAction{value: val, isSet: true}
}

func (v NullableBulkUpdateRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateRequestAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

