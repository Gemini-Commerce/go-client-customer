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

// checks if the CustomerGetGroupByCodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerGetGroupByCodeRequest{}

// CustomerGetGroupByCodeRequest struct for CustomerGetGroupByCodeRequest
type CustomerGetGroupByCodeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Code *string `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerGetGroupByCodeRequest CustomerGetGroupByCodeRequest

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
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGetGroupByCodeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerGetGroupByCodeRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
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
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGetGroupByCodeRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomerGetGroupByCodeRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomerGetGroupByCodeRequest) SetCode(v string) {
	o.Code = &v
}

func (o CustomerGetGroupByCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerGetGroupByCodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerGetGroupByCodeRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomerGetGroupByCodeRequest := _CustomerGetGroupByCodeRequest{}

	err = json.Unmarshal(data, &varCustomerGetGroupByCodeRequest)

	if err != nil {
		return err
	}

	*o = CustomerGetGroupByCodeRequest(varCustomerGetGroupByCodeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerGetGroupByCodeRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *CustomerGetGroupByCodeRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
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


