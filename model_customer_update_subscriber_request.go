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

// CustomerUpdateSubscriberRequest struct for CustomerUpdateSubscriberRequest
type CustomerUpdateSubscriberRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Subscriber *CustomerSubscriberResponseWithNewsletterRequest `json:"subscriber,omitempty"`
	FieldMask []string `json:"fieldMask,omitempty"`
}

// NewCustomerUpdateSubscriberRequest instantiates a new CustomerUpdateSubscriberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateSubscriberRequest() *CustomerUpdateSubscriberRequest {
	this := CustomerUpdateSubscriberRequest{}
	return &this
}

// NewCustomerUpdateSubscriberRequestWithDefaults instantiates a new CustomerUpdateSubscriberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateSubscriberRequestWithDefaults() *CustomerUpdateSubscriberRequest {
	this := CustomerUpdateSubscriberRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerUpdateSubscriberRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateSubscriberRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerUpdateSubscriberRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerUpdateSubscriberRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSubscriber returns the Subscriber field value if set, zero value otherwise.
func (o *CustomerUpdateSubscriberRequest) GetSubscriber() CustomerSubscriberResponseWithNewsletterRequest {
	if o == nil || isNil(o.Subscriber) {
		var ret CustomerSubscriberResponseWithNewsletterRequest
		return ret
	}
	return *o.Subscriber
}

// GetSubscriberOk returns a tuple with the Subscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateSubscriberRequest) GetSubscriberOk() (*CustomerSubscriberResponseWithNewsletterRequest, bool) {
	if o == nil || isNil(o.Subscriber) {
    return nil, false
	}
	return o.Subscriber, true
}

// HasSubscriber returns a boolean if a field has been set.
func (o *CustomerUpdateSubscriberRequest) HasSubscriber() bool {
	if o != nil && !isNil(o.Subscriber) {
		return true
	}

	return false
}

// SetSubscriber gets a reference to the given CustomerSubscriberResponseWithNewsletterRequest and assigns it to the Subscriber field.
func (o *CustomerUpdateSubscriberRequest) SetSubscriber(v CustomerSubscriberResponseWithNewsletterRequest) {
	o.Subscriber = &v
}

// GetFieldMask returns the FieldMask field value if set, zero value otherwise.
func (o *CustomerUpdateSubscriberRequest) GetFieldMask() []string {
	if o == nil || isNil(o.FieldMask) {
		var ret []string
		return ret
	}
	return o.FieldMask
}

// GetFieldMaskOk returns a tuple with the FieldMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateSubscriberRequest) GetFieldMaskOk() ([]string, bool) {
	if o == nil || isNil(o.FieldMask) {
    return nil, false
	}
	return o.FieldMask, true
}

// HasFieldMask returns a boolean if a field has been set.
func (o *CustomerUpdateSubscriberRequest) HasFieldMask() bool {
	if o != nil && !isNil(o.FieldMask) {
		return true
	}

	return false
}

// SetFieldMask gets a reference to the given []string and assigns it to the FieldMask field.
func (o *CustomerUpdateSubscriberRequest) SetFieldMask(v []string) {
	o.FieldMask = v
}

func (o CustomerUpdateSubscriberRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Subscriber) {
		toSerialize["subscriber"] = o.Subscriber
	}
	if !isNil(o.FieldMask) {
		toSerialize["fieldMask"] = o.FieldMask
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerUpdateSubscriberRequest struct {
	value *CustomerUpdateSubscriberRequest
	isSet bool
}

func (v NullableCustomerUpdateSubscriberRequest) Get() *CustomerUpdateSubscriberRequest {
	return v.value
}

func (v *NullableCustomerUpdateSubscriberRequest) Set(val *CustomerUpdateSubscriberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateSubscriberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateSubscriberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateSubscriberRequest(val *CustomerUpdateSubscriberRequest) *NullableCustomerUpdateSubscriberRequest {
	return &NullableCustomerUpdateSubscriberRequest{value: val, isSet: true}
}

func (v NullableCustomerUpdateSubscriberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateSubscriberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


