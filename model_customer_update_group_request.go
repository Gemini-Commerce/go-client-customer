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

// CustomerUpdateGroupRequest struct for CustomerUpdateGroupRequest
type CustomerUpdateGroupRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	GroupId *string `json:"groupId,omitempty"`
	Payload *CustomerUpdateGroupRequestPayload `json:"payload,omitempty"`
	FieldMask []string `json:"fieldMask,omitempty"`
}

// NewCustomerUpdateGroupRequest instantiates a new CustomerUpdateGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateGroupRequest() *CustomerUpdateGroupRequest {
	this := CustomerUpdateGroupRequest{}
	return &this
}

// NewCustomerUpdateGroupRequestWithDefaults instantiates a new CustomerUpdateGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateGroupRequestWithDefaults() *CustomerUpdateGroupRequest {
	this := CustomerUpdateGroupRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerUpdateGroupRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateGroupRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerUpdateGroupRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerUpdateGroupRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomerUpdateGroupRequest) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateGroupRequest) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomerUpdateGroupRequest) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomerUpdateGroupRequest) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CustomerUpdateGroupRequest) GetPayload() CustomerUpdateGroupRequestPayload {
	if o == nil || isNil(o.Payload) {
		var ret CustomerUpdateGroupRequestPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateGroupRequest) GetPayloadOk() (*CustomerUpdateGroupRequestPayload, bool) {
	if o == nil || isNil(o.Payload) {
    return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CustomerUpdateGroupRequest) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CustomerUpdateGroupRequestPayload and assigns it to the Payload field.
func (o *CustomerUpdateGroupRequest) SetPayload(v CustomerUpdateGroupRequestPayload) {
	o.Payload = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *CustomerUpdateGroupRequest) GetFieldMask() []string {
	if o == nil || isNil(o.FieldMask) {
		var ret []string
		return ret
	}
	return o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateGroupRequest) GetFieldMaskOk() ([]string, bool) {
	if o == nil || isNil(o.FieldMask) {
    return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *CustomerUpdateGroupRequest) HasFieldMask() bool {
	if o != nil && !isNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given []string and assigns it to the FieldMask field.
func (o *CustomerUpdateGroupRequest) SetFieldMask(v []string) {
	o.FieldMask = v
}

func (o CustomerUpdateGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerUpdateGroupRequest struct {
	value *CustomerUpdateGroupRequest
	isSet bool
}

func (v NullableCustomerUpdateGroupRequest) Get() *CustomerUpdateGroupRequest {
	return v.value
}

func (v *NullableCustomerUpdateGroupRequest) Set(val *CustomerUpdateGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateGroupRequest(val *CustomerUpdateGroupRequest) *NullableCustomerUpdateGroupRequest {
	return &NullableCustomerUpdateGroupRequest{value: val, isSet: true}
}

func (v NullableCustomerUpdateGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


