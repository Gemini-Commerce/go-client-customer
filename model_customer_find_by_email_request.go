/*
customer/customer.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer

import (
	"encoding/json"
)

// CustomerFindByEmailRequest struct for CustomerFindByEmailRequest
type CustomerFindByEmailRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewCustomerFindByEmailRequest instantiates a new CustomerFindByEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFindByEmailRequest() *CustomerFindByEmailRequest {
	this := CustomerFindByEmailRequest{}
	return &this
}

// NewCustomerFindByEmailRequestWithDefaults instantiates a new CustomerFindByEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFindByEmailRequestWithDefaults() *CustomerFindByEmailRequest {
	this := CustomerFindByEmailRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerFindByEmailRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindByEmailRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerFindByEmailRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerFindByEmailRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerFindByEmailRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindByEmailRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerFindByEmailRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerFindByEmailRequest) SetEmail(v string) {
	o.Email = &v
}

func (o CustomerFindByEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFindByEmailRequest struct {
	value *CustomerFindByEmailRequest
	isSet bool
}

func (v NullableCustomerFindByEmailRequest) Get() *CustomerFindByEmailRequest {
	return v.value
}

func (v *NullableCustomerFindByEmailRequest) Set(val *CustomerFindByEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFindByEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFindByEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFindByEmailRequest(val *CustomerFindByEmailRequest) *NullableCustomerFindByEmailRequest {
	return &NullableCustomerFindByEmailRequest{value: val, isSet: true}
}

func (v NullableCustomerFindByEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFindByEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


