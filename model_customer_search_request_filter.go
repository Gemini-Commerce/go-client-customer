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

// checks if the CustomerSearchRequestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSearchRequestFilter{}

// CustomerSearchRequestFilter struct for CustomerSearchRequestFilter
type CustomerSearchRequestFilter struct {
	Newsletter           *bool   `json:"newsletter,omitempty"`
	AgentGrn             *string `json:"agentGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerSearchRequestFilter CustomerSearchRequestFilter

// NewCustomerSearchRequestFilter instantiates a new CustomerSearchRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSearchRequestFilter() *CustomerSearchRequestFilter {
	this := CustomerSearchRequestFilter{}
	return &this
}

// NewCustomerSearchRequestFilterWithDefaults instantiates a new CustomerSearchRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSearchRequestFilterWithDefaults() *CustomerSearchRequestFilter {
	this := CustomerSearchRequestFilter{}
	return &this
}

// GetNewsletter returns the Newsletter field value if set, zero value otherwise.
func (o *CustomerSearchRequestFilter) GetNewsletter() bool {
	if o == nil || IsNil(o.Newsletter) {
		var ret bool
		return ret
	}
	return *o.Newsletter
}

// GetNewsletterOk returns a tuple with the Newsletter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequestFilter) GetNewsletterOk() (*bool, bool) {
	if o == nil || IsNil(o.Newsletter) {
		return nil, false
	}
	return o.Newsletter, true
}

// HasNewsletter returns a boolean if a field has been set.
func (o *CustomerSearchRequestFilter) HasNewsletter() bool {
	if o != nil && !IsNil(o.Newsletter) {
		return true
	}

	return false
}

// SetNewsletter gets a reference to the given bool and assigns it to the Newsletter field.
func (o *CustomerSearchRequestFilter) SetNewsletter(v bool) {
	o.Newsletter = &v
}

// GetAgentGrn returns the AgentGrn field value if set, zero value otherwise.
func (o *CustomerSearchRequestFilter) GetAgentGrn() string {
	if o == nil || IsNil(o.AgentGrn) {
		var ret string
		return ret
	}
	return *o.AgentGrn
}

// GetAgentGrnOk returns a tuple with the AgentGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSearchRequestFilter) GetAgentGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AgentGrn) {
		return nil, false
	}
	return o.AgentGrn, true
}

// HasAgentGrn returns a boolean if a field has been set.
func (o *CustomerSearchRequestFilter) HasAgentGrn() bool {
	if o != nil && !IsNil(o.AgentGrn) {
		return true
	}

	return false
}

// SetAgentGrn gets a reference to the given string and assigns it to the AgentGrn field.
func (o *CustomerSearchRequestFilter) SetAgentGrn(v string) {
	o.AgentGrn = &v
}

func (o CustomerSearchRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSearchRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Newsletter) {
		toSerialize["newsletter"] = o.Newsletter
	}
	if !IsNil(o.AgentGrn) {
		toSerialize["agentGrn"] = o.AgentGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerSearchRequestFilter) UnmarshalJSON(data []byte) (err error) {
	varCustomerSearchRequestFilter := _CustomerSearchRequestFilter{}

	err = json.Unmarshal(data, &varCustomerSearchRequestFilter)

	if err != nil {
		return err
	}

	*o = CustomerSearchRequestFilter(varCustomerSearchRequestFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "newsletter")
		delete(additionalProperties, "agentGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerSearchRequestFilter) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *CustomerSearchRequestFilter) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableCustomerSearchRequestFilter struct {
	value *CustomerSearchRequestFilter
	isSet bool
}

func (v NullableCustomerSearchRequestFilter) Get() *CustomerSearchRequestFilter {
	return v.value
}

func (v *NullableCustomerSearchRequestFilter) Set(val *CustomerSearchRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSearchRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSearchRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSearchRequestFilter(val *CustomerSearchRequestFilter) *NullableCustomerSearchRequestFilter {
	return &NullableCustomerSearchRequestFilter{value: val, isSet: true}
}

func (v NullableCustomerSearchRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSearchRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
