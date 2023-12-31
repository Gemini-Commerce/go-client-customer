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

// CustomerSearchRequestFilter struct for CustomerSearchRequestFilter
type CustomerSearchRequestFilter struct {
	Newsletter *bool `json:"newsletter,omitempty"`
}

// NewCustomerSearchRequestFilter instantiates a new CustomerSearchRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSearchRequestFilter() *CustomerSearchRequestFilter {
	this := CustomerSearchRequestFilter{}
	return &this
}

// NewCustomerSearchRequestFilterWithDefaults instantiates a new CustomerSearchRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSearchRequestFilterWithDefaults() *CustomerSearchRequestFilter {
	this := CustomerSearchRequestFilter{}
	return &this
}

// GetNewsletter returns the Newsletter field value if set, zero value otherwise.
func (o *CustomerSearchRequestFilter) GetNewsletter() bool {
	if o == nil || isNil(o.Newsletter) {
		var ret bool
		return ret
	}
	return *o.Newsletter
}

// GetNewsletterOk returns a tuple with the Newsletter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequestFilter) GetNewsletterOk() (*bool, bool) {
	if o == nil || isNil(o.Newsletter) {
    return nil, false
	}
	return o.Newsletter, true
}

// HasNewsletter returns a boolean if a field has been set.
func (o *CustomerSearchRequestFilter) HasNewsletter() bool {
	if o != nil && !isNil(o.Newsletter) {
		return true
	}

	return false
}

// SetNewsletter gets a reference to the given bool and assigns it to the Newsletter field.
func (o *CustomerSearchRequestFilter) SetNewsletter(v bool) {
	o.Newsletter = &v
}

func (o CustomerSearchRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Newsletter) {
		toSerialize["newsletter"] = o.Newsletter
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSearchRequestFilter struct {
	value *CustomerSearchRequestFilter
	isSet bool
}

func (v NullableCustomerSearchRequestFilter) Get() *CustomerSearchRequestFilter {
	return v.value
}

func (v *NullableCustomerSearchRequestFilter) Set(val *CustomerSearchRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSearchRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSearchRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSearchRequestFilter(val *CustomerSearchRequestFilter) *NullableCustomerSearchRequestFilter {
	return &NullableCustomerSearchRequestFilter{value: val, isSet: true}
}

func (v NullableCustomerSearchRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSearchRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


