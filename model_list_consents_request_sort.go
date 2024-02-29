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

// checks if the ListConsentsRequestSort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConsentsRequestSort{}

// ListConsentsRequestSort struct for ListConsentsRequestSort
type ListConsentsRequestSort struct {
	Field *SortSortField `json:"field,omitempty"`
	Order *SortSortOrder `json:"order,omitempty"`
}

// NewListConsentsRequestSort instantiates a new ListConsentsRequestSort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConsentsRequestSort() *ListConsentsRequestSort {
	this := ListConsentsRequestSort{}
	var field SortSortField = UNKNOWN
	this.Field = &field
	var order SortSortOrder = DESC
	this.Order = &order
	return &this
}

// NewListConsentsRequestSortWithDefaults instantiates a new ListConsentsRequestSort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConsentsRequestSortWithDefaults() *ListConsentsRequestSort {
	this := ListConsentsRequestSort{}
	var field SortSortField = UNKNOWN
	this.Field = &field
	var order SortSortOrder = DESC
	this.Order = &order
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ListConsentsRequestSort) GetField() SortSortField {
	if o == nil || IsNil(o.Field) {
		var ret SortSortField
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConsentsRequestSort) GetFieldOk() (*SortSortField, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ListConsentsRequestSort) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given SortSortField and assigns it to the Field field.
func (o *ListConsentsRequestSort) SetField(v SortSortField) {
	o.Field = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ListConsentsRequestSort) GetOrder() SortSortOrder {
	if o == nil || IsNil(o.Order) {
		var ret SortSortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConsentsRequestSort) GetOrderOk() (*SortSortOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ListConsentsRequestSort) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given SortSortOrder and assigns it to the Order field.
func (o *ListConsentsRequestSort) SetOrder(v SortSortOrder) {
	o.Order = &v
}

func (o ListConsentsRequestSort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConsentsRequestSort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableListConsentsRequestSort struct {
	value *ListConsentsRequestSort
	isSet bool
}

func (v NullableListConsentsRequestSort) Get() *ListConsentsRequestSort {
	return v.value
}

func (v *NullableListConsentsRequestSort) Set(val *ListConsentsRequestSort) {
	v.value = val
	v.isSet = true
}

func (v NullableListConsentsRequestSort) IsSet() bool {
	return v.isSet
}

func (v *NullableListConsentsRequestSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConsentsRequestSort(val *ListConsentsRequestSort) *NullableListConsentsRequestSort {
	return &NullableListConsentsRequestSort{value: val, isSet: true}
}

func (v NullableListConsentsRequestSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConsentsRequestSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


