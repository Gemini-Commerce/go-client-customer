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

// checks if the CustomerSetPermissionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSetPermissionsRequest{}

// CustomerSetPermissionsRequest struct for CustomerSetPermissionsRequest
type CustomerSetPermissionsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CustomerId *string `json:"customerId,omitempty"`
	Permissions []CustomerPermission `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerSetPermissionsRequest CustomerSetPermissionsRequest

// NewCustomerSetPermissionsRequest instantiates a new CustomerSetPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSetPermissionsRequest() *CustomerSetPermissionsRequest {
	this := CustomerSetPermissionsRequest{}
	return &this
}

// NewCustomerSetPermissionsRequestWithDefaults instantiates a new CustomerSetPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSetPermissionsRequestWithDefaults() *CustomerSetPermissionsRequest {
	this := CustomerSetPermissionsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerSetPermissionsRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSetPermissionsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerSetPermissionsRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerSetPermissionsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerSetPermissionsRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSetPermissionsRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerSetPermissionsRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerSetPermissionsRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CustomerSetPermissionsRequest) GetPermissions() []CustomerPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []CustomerPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSetPermissionsRequest) GetPermissionsOk() ([]CustomerPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CustomerSetPermissionsRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CustomerPermission and assigns it to the Permissions field.
func (o *CustomerSetPermissionsRequest) SetPermissions(v []CustomerPermission) {
	o.Permissions = v
}

func (o CustomerSetPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSetPermissionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerSetPermissionsRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomerSetPermissionsRequest := _CustomerSetPermissionsRequest{}

	err = json.Unmarshal(data, &varCustomerSetPermissionsRequest)

	if err != nil {
		return err
	}

	*o = CustomerSetPermissionsRequest(varCustomerSetPermissionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerSetPermissionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *CustomerSetPermissionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableCustomerSetPermissionsRequest struct {
	value *CustomerSetPermissionsRequest
	isSet bool
}

func (v NullableCustomerSetPermissionsRequest) Get() *CustomerSetPermissionsRequest {
	return v.value
}

func (v *NullableCustomerSetPermissionsRequest) Set(val *CustomerSetPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSetPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSetPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSetPermissionsRequest(val *CustomerSetPermissionsRequest) *NullableCustomerSetPermissionsRequest {
	return &NullableCustomerSetPermissionsRequest{value: val, isSet: true}
}

func (v NullableCustomerSetPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSetPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


