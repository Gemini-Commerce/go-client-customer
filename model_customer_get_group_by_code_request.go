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

// CustomerGetGroupByCodeRequest struct for CustomerGetGroupByCodeRequest
type CustomerGetGroupByCodeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewCustomerGetGroupByCodeRequest instantiates a new CustomerGetGroupByCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGetGroupByCodeRequest() *CustomerGetGroupByCodeRequest {
	this := CustomerGetGroupByCodeRequest{}
	return &this
}

// NewCustomerGetGroupByCodeRequestWithDefaults instantiates a new CustomerGetGroupByCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGetGroupByCodeRequestWithDefaults() *CustomerGetGroupByCodeRequest {
	this := CustomerGetGroupByCodeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerGetGroupByCodeRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGetGroupByCodeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerGetGroupByCodeRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerGetGroupByCodeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CustomerGetGroupByCodeRequest) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGetGroupByCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomerGetGroupByCodeRequest) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomerGetGroupByCodeRequest) SetCode(v string) {
	o.Code = &v
}

func (o CustomerGetGroupByCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerGetGroupByCodeRequest struct {
	value *CustomerGetGroupByCodeRequest
	isSet bool
}

func (v NullableCustomerGetGroupByCodeRequest) Get() *CustomerGetGroupByCodeRequest {
	return v.value
}

func (v *NullableCustomerGetGroupByCodeRequest) Set(val *CustomerGetGroupByCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGetGroupByCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGetGroupByCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGetGroupByCodeRequest(val *CustomerGetGroupByCodeRequest) *NullableCustomerGetGroupByCodeRequest {
	return &NullableCustomerGetGroupByCodeRequest{value: val, isSet: true}
}

func (v NullableCustomerGetGroupByCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGetGroupByCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


