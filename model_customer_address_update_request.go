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

// checks if the CustomerAddressUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddressUpdateRequest{}

// CustomerAddressUpdateRequest struct for CustomerAddressUpdateRequest
type CustomerAddressUpdateRequest struct {
	TenantId             *string                `json:"tenantId,omitempty"`
	CustomerId           *string                `json:"customerId,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Address              *CustomerAddressEntity `json:"address,omitempty"`
	FieldMask            *string                `json:"fieldMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerAddressUpdateRequest CustomerAddressUpdateRequest

// NewCustomerAddressUpdateRequest instantiates a new CustomerAddressUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressUpdateRequest() *CustomerAddressUpdateRequest {
	this := CustomerAddressUpdateRequest{}
	return &this
}

// NewCustomerAddressUpdateRequestWithDefaults instantiates a new CustomerAddressUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressUpdateRequestWithDefaults() *CustomerAddressUpdateRequest {
	this := CustomerAddressUpdateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerAddressUpdateRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerAddressUpdateRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerAddressUpdateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerAddressUpdateRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdateRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerAddressUpdateRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerAddressUpdateRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerAddressUpdateRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerAddressUpdateRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerAddressUpdateRequest) SetId(v string) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CustomerAddressUpdateRequest) GetAddress() CustomerAddressEntity {
	if o == nil || IsNil(o.Address) {
		var ret CustomerAddressEntity
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdateRequest) GetAddressOk() (*CustomerAddressEntity, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CustomerAddressUpdateRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerAddressEntity and assigns it to the Address field.
func (o *CustomerAddressUpdateRequest) SetAddress(v CustomerAddressEntity) {
	o.Address = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *CustomerAddressUpdateRequest) GetFieldMask() string {
	if o == nil || IsNil(o.FieldMask) {
		var ret string
		return ret
	}
	return *o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressUpdateRequest) GetFieldMaskOk() (*string, bool) {
	if o == nil || IsNil(o.FieldMask) {
		return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *CustomerAddressUpdateRequest) HasFieldMask() bool {
	if o != nil && !IsNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given string and assigns it to the FieldMask field.
func (o *CustomerAddressUpdateRequest) SetFieldMask(v string) {
	o.FieldMask = &v
}

func (o CustomerAddressUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddressUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerAddressUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomerAddressUpdateRequest := _CustomerAddressUpdateRequest{}

	err = json.Unmarshal(data, &varCustomerAddressUpdateRequest)

	if err != nil {
		return err
	}

	*o = CustomerAddressUpdateRequest(varCustomerAddressUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "address")
		delete(additionalProperties, "fieldMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerAddressUpdateRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CustomerAddressUpdateRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCustomerAddressUpdateRequest struct {
	value *CustomerAddressUpdateRequest
	isSet bool
}

func (v NullableCustomerAddressUpdateRequest) Get() *CustomerAddressUpdateRequest {
	return v.value
}

func (v *NullableCustomerAddressUpdateRequest) Set(val *CustomerAddressUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressUpdateRequest(val *CustomerAddressUpdateRequest) *NullableCustomerAddressUpdateRequest {
	return &NullableCustomerAddressUpdateRequest{value: val, isSet: true}
}

func (v NullableCustomerAddressUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
