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
	"time"
)

// checks if the ListCustomersRequestFilterDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomersRequestFilterDate{}

// ListCustomersRequestFilterDate struct for ListCustomersRequestFilterDate
type ListCustomersRequestFilterDate struct {
	From                 *time.Time `json:"from,omitempty"`
	To                   *time.Time `json:"to,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCustomersRequestFilterDate ListCustomersRequestFilterDate

// NewListCustomersRequestFilterDate instantiates a new ListCustomersRequestFilterDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomersRequestFilterDate() *ListCustomersRequestFilterDate {
	this := ListCustomersRequestFilterDate{}
	return &this
}

// NewListCustomersRequestFilterDateWithDefaults instantiates a new ListCustomersRequestFilterDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomersRequestFilterDateWithDefaults() *ListCustomersRequestFilterDate {
	this := ListCustomersRequestFilterDate{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ListCustomersRequestFilterDate) GetFrom() time.Time {
	if o == nil || IsNil(o.From) {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomersRequestFilterDate) GetFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ListCustomersRequestFilterDate) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *ListCustomersRequestFilterDate) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ListCustomersRequestFilterDate) GetTo() time.Time {
	if o == nil || IsNil(o.To) {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomersRequestFilterDate) GetToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ListCustomersRequestFilterDate) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *ListCustomersRequestFilterDate) SetTo(v time.Time) {
	o.To = &v
}

func (o ListCustomersRequestFilterDate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomersRequestFilterDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCustomersRequestFilterDate) UnmarshalJSON(data []byte) (err error) {
	varListCustomersRequestFilterDate := _ListCustomersRequestFilterDate{}

	err = json.Unmarshal(data, &varListCustomersRequestFilterDate)

	if err != nil {
		return err
	}

	*o = ListCustomersRequestFilterDate(varListCustomersRequestFilterDate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ListCustomersRequestFilterDate) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ListCustomersRequestFilterDate) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableListCustomersRequestFilterDate struct {
	value *ListCustomersRequestFilterDate
	isSet bool
}

func (v NullableListCustomersRequestFilterDate) Get() *ListCustomersRequestFilterDate {
	return v.value
}

func (v *NullableListCustomersRequestFilterDate) Set(val *ListCustomersRequestFilterDate) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomersRequestFilterDate) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomersRequestFilterDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomersRequestFilterDate(val *ListCustomersRequestFilterDate) *NullableListCustomersRequestFilterDate {
	return &NullableListCustomersRequestFilterDate{value: val, isSet: true}
}

func (v NullableListCustomersRequestFilterDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomersRequestFilterDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
