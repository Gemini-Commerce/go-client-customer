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

// checks if the CustomerAddressDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddressDeleteResponse{}

// CustomerAddressDeleteResponse struct for CustomerAddressDeleteResponse
type CustomerAddressDeleteResponse struct {
	Customer *CustomerCustomerResponse `json:"customer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomerAddressDeleteResponse CustomerAddressDeleteResponse

// NewCustomerAddressDeleteResponse instantiates a new CustomerAddressDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressDeleteResponse() *CustomerAddressDeleteResponse {
	this := CustomerAddressDeleteResponse{}
	return &this
}

// NewCustomerAddressDeleteResponseWithDefaults instantiates a new CustomerAddressDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressDeleteResponseWithDefaults() *CustomerAddressDeleteResponse {
	this := CustomerAddressDeleteResponse{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomerAddressDeleteResponse) GetCustomer() CustomerCustomerResponse {
	if o == nil || IsNil(o.Customer) {
		var ret CustomerCustomerResponse
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressDeleteResponse) GetCustomerOk() (*CustomerCustomerResponse, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// &#39;Has&#39;Customer returns a boolean if a field has been set.
func (o *CustomerAddressDeleteResponse) &#39;Has&#39;Customer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerCustomerResponse and assigns it to the Customer field.
func (o *CustomerAddressDeleteResponse) SetCustomer(v CustomerCustomerResponse) {
	o.Customer = &v
}

func (o CustomerAddressDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddressDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	varCustomerAddressDeleteResponse := _CustomerAddressDeleteResponse{}

	err = json.Unmarshal(data, &varCustomerAddressDeleteResponse)

	if err != nil {
		return err
	}

	*o = CustomerAddressDeleteResponse(varCustomerAddressDeleteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *CustomerAddressDeleteResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *CustomerAddressDeleteResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableCustomerAddressDeleteResponse struct {
	value *CustomerAddressDeleteResponse
	isSet bool
}

func (v NullableCustomerAddressDeleteResponse) Get() *CustomerAddressDeleteResponse {
	return v.value
}

func (v *NullableCustomerAddressDeleteResponse) Set(val *CustomerAddressDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressDeleteResponse(val *CustomerAddressDeleteResponse) *NullableCustomerAddressDeleteResponse {
	return &NullableCustomerAddressDeleteResponse{value: val, isSet: true}
}

func (v NullableCustomerAddressDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


