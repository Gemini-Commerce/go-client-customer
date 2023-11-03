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

// CustomerFindByIdRequest struct for CustomerFindByIdRequest
type CustomerFindByIdRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCustomerFindByIdRequest instantiates a new CustomerFindByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerFindByIdRequest() *CustomerFindByIdRequest {
	this := CustomerFindByIdRequest{}
	return &this
}

// NewCustomerFindByIdRequestWithDefaults instantiates a new CustomerFindByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerFindByIdRequestWithDefaults() *CustomerFindByIdRequest {
	this := CustomerFindByIdRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerFindByIdRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindByIdRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerFindByIdRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerFindByIdRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerFindByIdRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerFindByIdRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerFindByIdRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerFindByIdRequest) SetId(v string) {
	o.Id = &v
}

func (o CustomerFindByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerFindByIdRequest struct {
	value *CustomerFindByIdRequest
	isSet bool
}

func (v NullableCustomerFindByIdRequest) Get() *CustomerFindByIdRequest {
	return v.value
}

func (v *NullableCustomerFindByIdRequest) Set(val *CustomerFindByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerFindByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerFindByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerFindByIdRequest(val *CustomerFindByIdRequest) *NullableCustomerFindByIdRequest {
	return &NullableCustomerFindByIdRequest{value: val, isSet: true}
}

func (v NullableCustomerFindByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerFindByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


