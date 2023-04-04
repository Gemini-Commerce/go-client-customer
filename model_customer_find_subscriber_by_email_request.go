/*
customer/customer.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomerFindSubscriberByEmailRequest struct for CustomerFindSubscriberByEmailRequest
type CustomerFindSubscriberByEmailRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	SubscriberEmail *string `json:"subscriberEmail,omitempty"`
}

// NewCustomerFindSubscriberByEmailRequest instantiates a new CustomerFindSubscriberByEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFindSubscriberByEmailRequest() *CustomerFindSubscriberByEmailRequest {
	this := CustomerFindSubscriberByEmailRequest{}
	return &this
}

// NewCustomerFindSubscriberByEmailRequestWithDefaults instantiates a new CustomerFindSubscriberByEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFindSubscriberByEmailRequestWithDefaults() *CustomerFindSubscriberByEmailRequest {
	this := CustomerFindSubscriberByEmailRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerFindSubscriberByEmailRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindSubscriberByEmailRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerFindSubscriberByEmailRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerFindSubscriberByEmailRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSubscriberEmail returns the SubscriberEmail field value if set, zero value otherwise.
func (o *CustomerFindSubscriberByEmailRequest) GetSubscriberEmail() string {
	if o == nil || isNil(o.SubscriberEmail) {
		var ret string
		return ret
	}
	return *o.SubscriberEmail
}

// GetSubscriberEmailOk returns a tuple with the SubscriberEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindSubscriberByEmailRequest) GetSubscriberEmailOk() (*string, bool) {
	if o == nil || isNil(o.SubscriberEmail) {
    return nil, false
	}
	return o.SubscriberEmail, true
}

// HasSubscriberEmail returns a boolean if a field has been set.
func (o *CustomerFindSubscriberByEmailRequest) HasSubscriberEmail() bool {
	if o != nil && !isNil(o.SubscriberEmail) {
		return true
	}

	return false
}

// SetSubscriberEmail gets a reference to the given string and assigns it to the SubscriberEmail field.
func (o *CustomerFindSubscriberByEmailRequest) SetSubscriberEmail(v string) {
	o.SubscriberEmail = &v
}

func (o CustomerFindSubscriberByEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.SubscriberEmail) {
		toSerialize["subscriberEmail"] = o.SubscriberEmail
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFindSubscriberByEmailRequest struct {
	value *CustomerFindSubscriberByEmailRequest
	isSet bool
}

func (v NullableCustomerFindSubscriberByEmailRequest) Get() *CustomerFindSubscriberByEmailRequest {
	return v.value
}

func (v *NullableCustomerFindSubscriberByEmailRequest) Set(val *CustomerFindSubscriberByEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFindSubscriberByEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFindSubscriberByEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFindSubscriberByEmailRequest(val *CustomerFindSubscriberByEmailRequest) *NullableCustomerFindSubscriberByEmailRequest {
	return &NullableCustomerFindSubscriberByEmailRequest{value: val, isSet: true}
}

func (v NullableCustomerFindSubscriberByEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFindSubscriberByEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


