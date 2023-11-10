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

// CustomerUnsubscribeRequest struct for CustomerUnsubscribeRequest
type CustomerUnsubscribeRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Email *string `json:"email,omitempty"`
	NewsletterGrn []string `json:"newsletterGrn,omitempty"`
}

// NewCustomerUnsubscribeRequest instantiates a new CustomerUnsubscribeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUnsubscribeRequest() *CustomerUnsubscribeRequest {
	this := CustomerUnsubscribeRequest{}
	return &this
}

// NewCustomerUnsubscribeRequestWithDefaults instantiates a new CustomerUnsubscribeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUnsubscribeRequestWithDefaults() *CustomerUnsubscribeRequest {
	this := CustomerUnsubscribeRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerUnsubscribeRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUnsubscribeRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerUnsubscribeRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerUnsubscribeRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerUnsubscribeRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUnsubscribeRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerUnsubscribeRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerUnsubscribeRequest) SetEmail(v string) {
	o.Email = &v
}

// GetNewsletterGrn returns the NewsletterGrn field value if set, zero value otherwise.
func (o *CustomerUnsubscribeRequest) GetNewsletterGrn() []string {
	if o == nil || isNil(o.NewsletterGrn) {
		var ret []string
		return ret
	}
	return o.NewsletterGrn
}

// GetNewsletterGrnOk returns a tuple with the NewsletterGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUnsubscribeRequest) GetNewsletterGrnOk() ([]string, bool) {
	if o == nil || isNil(o.NewsletterGrn) {
    return nil, false
	}
	return o.NewsletterGrn, true
}

// HasNewsletterGrn returns a boolean if a field has been set.
func (o *CustomerUnsubscribeRequest) HasNewsletterGrn() bool {
	if o != nil && !isNil(o.NewsletterGrn) {
		return true
	}

	return false
}

// SetNewsletterGrn gets a reference to the given []string and assigns it to the NewsletterGrn field.
func (o *CustomerUnsubscribeRequest) SetNewsletterGrn(v []string) {
	o.NewsletterGrn = v
}

func (o CustomerUnsubscribeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.NewsletterGrn) {
		toSerialize["newsletterGrn"] = o.NewsletterGrn
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerUnsubscribeRequest struct {
	value *CustomerUnsubscribeRequest
	isSet bool
}

func (v NullableCustomerUnsubscribeRequest) Get() *CustomerUnsubscribeRequest {
	return v.value
}

func (v *NullableCustomerUnsubscribeRequest) Set(val *CustomerUnsubscribeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUnsubscribeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUnsubscribeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUnsubscribeRequest(val *CustomerUnsubscribeRequest) *NullableCustomerUnsubscribeRequest {
	return &NullableCustomerUnsubscribeRequest{value: val, isSet: true}
}

func (v NullableCustomerUnsubscribeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUnsubscribeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


