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

// CustomerRevokePermissionsRequest struct for CustomerRevokePermissionsRequest
type CustomerRevokePermissionsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CustomerId *string `json:"customerId,omitempty"`
	Permissions []CustomerPermission `json:"permissions,omitempty"`
}

// NewCustomerRevokePermissionsRequest instantiates a new CustomerRevokePermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerRevokePermissionsRequest() *CustomerRevokePermissionsRequest {
	this := CustomerRevokePermissionsRequest{}
	return &this
}

// NewCustomerRevokePermissionsRequestWithDefaults instantiates a new CustomerRevokePermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerRevokePermissionsRequestWithDefaults() *CustomerRevokePermissionsRequest {
	this := CustomerRevokePermissionsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerRevokePermissionsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRevokePermissionsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerRevokePermissionsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerRevokePermissionsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerRevokePermissionsRequest) GetCustomerId() string {
	if o == nil || isNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRevokePermissionsRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomerId) {
    return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerRevokePermissionsRequest) HasCustomerId() bool {
	if o != nil && !isNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerRevokePermissionsRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CustomerRevokePermissionsRequest) GetPermissions() []CustomerPermission {
	if o == nil || isNil(o.Permissions) {
		var ret []CustomerPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRevokePermissionsRequest) GetPermissionsOk() ([]CustomerPermission, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CustomerRevokePermissionsRequest) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CustomerPermission and assigns it to the Permissions field.
func (o *CustomerRevokePermissionsRequest) SetPermissions(v []CustomerPermission) {
	o.Permissions = v
}

func (o CustomerRevokePermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerRevokePermissionsRequest struct {
	value *CustomerRevokePermissionsRequest
	isSet bool
}

func (v NullableCustomerRevokePermissionsRequest) Get() *CustomerRevokePermissionsRequest {
	return v.value
}

func (v *NullableCustomerRevokePermissionsRequest) Set(val *CustomerRevokePermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerRevokePermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerRevokePermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerRevokePermissionsRequest(val *CustomerRevokePermissionsRequest) *NullableCustomerRevokePermissionsRequest {
	return &NullableCustomerRevokePermissionsRequest{value: val, isSet: true}
}

func (v NullableCustomerRevokePermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerRevokePermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


