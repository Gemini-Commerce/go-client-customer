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

// CustomerDeleteGroupRequest struct for CustomerDeleteGroupRequest
type CustomerDeleteGroupRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
}

// NewCustomerDeleteGroupRequest instantiates a new CustomerDeleteGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDeleteGroupRequest() *CustomerDeleteGroupRequest {
	this := CustomerDeleteGroupRequest{}
	return &this
}

// NewCustomerDeleteGroupRequestWithDefaults instantiates a new CustomerDeleteGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDeleteGroupRequestWithDefaults() *CustomerDeleteGroupRequest {
	this := CustomerDeleteGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerDeleteGroupRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDeleteGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerDeleteGroupRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerDeleteGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomerDeleteGroupRequest) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerDeleteGroupRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomerDeleteGroupRequest) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomerDeleteGroupRequest) SetGroupId(v string) {
	o.GroupId = &v
}

func (o CustomerDeleteGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDeleteGroupRequest struct {
	value *CustomerDeleteGroupRequest
	isSet bool
}

func (v NullableCustomerDeleteGroupRequest) Get() *CustomerDeleteGroupRequest {
	return v.value
}

func (v *NullableCustomerDeleteGroupRequest) Set(val *CustomerDeleteGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDeleteGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDeleteGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDeleteGroupRequest(val *CustomerDeleteGroupRequest) *NullableCustomerDeleteGroupRequest {
	return &NullableCustomerDeleteGroupRequest{value: val, isSet: true}
}

func (v NullableCustomerDeleteGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDeleteGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


