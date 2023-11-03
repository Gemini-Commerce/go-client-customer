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

// CustomerAddressCreateRequest struct for CustomerAddressCreateRequest
type CustomerAddressCreateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	CustomerId *string `json:"customerId,omitempty"`
	Em *CustomerEMFields `json:"em,omitempty"`
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Street *string `json:"street,omitempty"`
	Number *string `json:"number,omitempty"`
	Zip *string `json:"zip,omitempty"`
	City *string `json:"city,omitempty"`
	Province *string `json:"province,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	FiscalCode *string `json:"fiscalCode,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
	Kind *CustomerAddressCreateRequestKind `json:"kind,omitempty"`
	Default *bool `json:"default,omitempty"`
	Country *string `json:"country,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
}

// NewCustomerAddressCreateRequest instantiates a new CustomerAddressCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCreateRequest() *CustomerAddressCreateRequest {
	this := CustomerAddressCreateRequest{}
	var kind CustomerAddressCreateRequestKind = CUSTOMERADDRESSCREATEREQUESTKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// NewCustomerAddressCreateRequestWithDefaults instantiates a new CustomerAddressCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCreateRequestWithDefaults() *CustomerAddressCreateRequest {
	this := CustomerAddressCreateRequest{}
	var kind CustomerAddressCreateRequestKind = CUSTOMERADDRESSCREATEREQUESTKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CustomerAddressCreateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetCustomerId() string {
	if o == nil || isNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomerId) {
    return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasCustomerId() bool {
	if o != nil && !isNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CustomerAddressCreateRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerAddressCreateRequest) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerAddressCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetSurname() string {
	if o == nil || isNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetSurnameOk() (*string, bool) {
	if o == nil || isNil(o.Surname) {
    return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasSurname() bool {
	if o != nil && !isNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerAddressCreateRequest) SetSurname(v string) {
	o.Surname = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetStreet() string {
	if o == nil || isNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetStreetOk() (*string, bool) {
	if o == nil || isNil(o.Street) {
    return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasStreet() bool {
	if o != nil && !isNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *CustomerAddressCreateRequest) SetStreet(v string) {
	o.Street = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetNumber() string {
	if o == nil || isNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetNumberOk() (*string, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CustomerAddressCreateRequest) SetNumber(v string) {
	o.Number = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetZip() string {
	if o == nil || isNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetZipOk() (*string, bool) {
	if o == nil || isNil(o.Zip) {
    return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasZip() bool {
	if o != nil && !isNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CustomerAddressCreateRequest) SetZip(v string) {
	o.Zip = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerAddressCreateRequest) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetProvince() string {
	if o == nil || isNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetProvinceOk() (*string, bool) {
	if o == nil || isNil(o.Province) {
    return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasProvince() bool {
	if o != nil && !isNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *CustomerAddressCreateRequest) SetProvince(v string) {
	o.Province = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetPhoneNumber() string {
	if o == nil || isNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumber) {
    return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasPhoneNumber() bool {
	if o != nil && !isNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerAddressCreateRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetFiscalCode returns the FiscalCode field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetFiscalCode() string {
	if o == nil || isNil(o.FiscalCode) {
		var ret string
		return ret
	}
	return *o.FiscalCode
}

// GetFiscalCodeOk returns a tuple with the FiscalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetFiscalCodeOk() (*string, bool) {
	if o == nil || isNil(o.FiscalCode) {
    return nil, false
	}
	return o.FiscalCode, true
}

// HasFiscalCode returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasFiscalCode() bool {
	if o != nil && !isNil(o.FiscalCode) {
		return true
	}

	return false
}

// SetFiscalCode gets a reference to the given string and assigns it to the FiscalCode field.
func (o *CustomerAddressCreateRequest) SetFiscalCode(v string) {
	o.FiscalCode = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetVatNumber() string {
	if o == nil || isNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetVatNumberOk() (*string, bool) {
	if o == nil || isNil(o.VatNumber) {
    return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasVatNumber() bool {
	if o != nil && !isNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CustomerAddressCreateRequest) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetKind() CustomerAddressCreateRequestKind {
	if o == nil || isNil(o.Kind) {
		var ret CustomerAddressCreateRequestKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetKindOk() (*CustomerAddressCreateRequestKind, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CustomerAddressCreateRequestKind and assigns it to the Kind field.
func (o *CustomerAddressCreateRequest) SetKind(v CustomerAddressCreateRequestKind) {
	o.Kind = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CustomerAddressCreateRequest) SetDefault(v bool) {
	o.Default = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerAddressCreateRequest) SetCountry(v string) {
	o.Country = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerAddressCreateRequest) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateRequest) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerAddressCreateRequest) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *CustomerAddressCreateRequest) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o CustomerAddressCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !isNil(o.Em) {
		toSerialize["em"] = o.Em
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !isNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.Province) {
		toSerialize["province"] = o.Province
	}
	if !isNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !isNil(o.FiscalCode) {
		toSerialize["fiscalCode"] = o.FiscalCode
	}
	if !isNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressCreateRequest struct {
	value *CustomerAddressCreateRequest
	isSet bool
}

func (v NullableCustomerAddressCreateRequest) Get() *CustomerAddressCreateRequest {
	return v.value
}

func (v *NullableCustomerAddressCreateRequest) Set(val *CustomerAddressCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreateRequest(val *CustomerAddressCreateRequest) *NullableCustomerAddressCreateRequest {
	return &NullableCustomerAddressCreateRequest{value: val, isSet: true}
}

func (v NullableCustomerAddressCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


