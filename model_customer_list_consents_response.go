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

// checks if the CustomerListConsentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerListConsentsResponse{}

// CustomerListConsentsResponse struct for CustomerListConsentsResponse
type CustomerListConsentsResponse struct {
	Consents []CustomerConsent `json:"consents,omitempty"`
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewCustomerListConsentsResponse instantiates a new CustomerListConsentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerListConsentsResponse() *CustomerListConsentsResponse {
	this := CustomerListConsentsResponse{}
	return &this
}

// NewCustomerListConsentsResponseWithDefaults instantiates a new CustomerListConsentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerListConsentsResponseWithDefaults() *CustomerListConsentsResponse {
	this := CustomerListConsentsResponse{}
	return &this
}

// GetConsents returns the Consents field value if set, zero value otherwise.
func (o *CustomerListConsentsResponse) GetConsents() []CustomerConsent {
	if o == nil || IsNil(o.Consents) {
		var ret []CustomerConsent
		return ret
	}
	return o.Consents
}

// GetConsentsOk returns a tuple with the Consents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListConsentsResponse) GetConsentsOk() ([]CustomerConsent, bool) {
	if o == nil || IsNil(o.Consents) {
		return nil, false
	}
	return o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *CustomerListConsentsResponse) HasConsents() bool {
	if o != nil && !IsNil(o.Consents) {
		return true
	}

	return false
}

// SetConsents gets a reference to the given []CustomerConsent and assigns it to the Consents field.
func (o *CustomerListConsentsResponse) SetConsents(v []CustomerConsent) {
	o.Consents = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *CustomerListConsentsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerListConsentsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *CustomerListConsentsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *CustomerListConsentsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o CustomerListConsentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerListConsentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consents) {
		toSerialize["consents"] = o.Consents
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableCustomerListConsentsResponse struct {
	value *CustomerListConsentsResponse
	isSet bool
}

func (v NullableCustomerListConsentsResponse) Get() *CustomerListConsentsResponse {
	return v.value
}

func (v *NullableCustomerListConsentsResponse) Set(val *CustomerListConsentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerListConsentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerListConsentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerListConsentsResponse(val *CustomerListConsentsResponse) *NullableCustomerListConsentsResponse {
	return &NullableCustomerListConsentsResponse{value: val, isSet: true}
}

func (v NullableCustomerListConsentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerListConsentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

