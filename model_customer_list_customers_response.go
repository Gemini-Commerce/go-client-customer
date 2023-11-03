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
)

// CustomerListCustomersResponse struct for CustomerListCustomersResponse
type CustomerListCustomersResponse struct {
	Customers []CustomerCustomerResponse `json:"customers,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewCustomerListCustomersResponse instantiates a new CustomerListCustomersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerListCustomersResponse() *CustomerListCustomersResponse {
	this := CustomerListCustomersResponse{}
	return &this
}

// NewCustomerListCustomersResponseWithDefaults instantiates a new CustomerListCustomersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerListCustomersResponseWithDefaults() *CustomerListCustomersResponse {
	this := CustomerListCustomersResponse{}
	return &this
}

// GetCustomers returns the Customers field value if set, zero value otherwise.
func (o *CustomerListCustomersResponse) GetCustomers() []CustomerCustomerResponse {
	if o == nil || isNil(o.Customers) {
		var ret []CustomerCustomerResponse
		return ret
	}
	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListCustomersResponse) GetCustomersOk() ([]CustomerCustomerResponse, bool) {
	if o == nil || isNil(o.Customers) {
    return nil, false
	}
	return o.Customers, true
}

// HasCustomers returns a boolean if a field has been set.
func (o *CustomerListCustomersResponse) HasCustomers() bool {
	if o != nil && !isNil(o.Customers) {
		return true
	}

	return false
}

// SetCustomers gets a reference to the given []CustomerCustomerResponse and assigns it to the Customers field.
func (o *CustomerListCustomersResponse) SetCustomers(v []CustomerCustomerResponse) {
	o.Customers = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerListCustomersResponse) GetNextPageToken() string {
	if o == nil || isNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListCustomersResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.NextPageToken) {
    return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerListCustomersResponse) HasNextPageToken() bool {
	if o != nil && !isNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerListCustomersResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CustomerListCustomersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Customers) {
		toSerialize["customers"] = o.Customers
	}
	if !isNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerListCustomersResponse struct {
	value *CustomerListCustomersResponse
	isSet bool
}

func (v NullableCustomerListCustomersResponse) Get() *CustomerListCustomersResponse {
	return v.value
}

func (v *NullableCustomerListCustomersResponse) Set(val *CustomerListCustomersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerListCustomersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerListCustomersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerListCustomersResponse(val *CustomerListCustomersResponse) *NullableCustomerListCustomersResponse {
	return &NullableCustomerListCustomersResponse{value: val, isSet: true}
}

func (v NullableCustomerListCustomersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerListCustomersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

