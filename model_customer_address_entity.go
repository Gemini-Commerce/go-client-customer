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

// CustomerAddressEntity struct for CustomerAddressEntity
type CustomerAddressEntity struct {
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
	Kind *CustomerAddressEntityKind `json:"kind,omitempty"`
	Default *bool `json:"default,omitempty"`
	Country *string `json:"country,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
}

// NewCustomerAddressEntity instantiates a new CustomerAddressEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressEntity() *CustomerAddressEntity {
	this := CustomerAddressEntity{}
	var kind CustomerAddressEntityKind = CUSTOMERADDRESSENTITYKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// NewCustomerAddressEntityWithDefaults instantiates a new CustomerAddressEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressEntityWithDefaults() *CustomerAddressEntity {
	this := CustomerAddressEntity{}
	var kind CustomerAddressEntityKind = CUSTOMERADDRESSENTITYKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerAddressEntity) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerAddressEntity) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetSurname() string {
	if o == nil || isNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetSurnameOk() (*string, bool) {
	if o == nil || isNil(o.Surname) {
    return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasSurname() bool {
	if o != nil && !isNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerAddressEntity) SetSurname(v string) {
	o.Surname = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetStreet() string {
	if o == nil || isNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetStreetOk() (*string, bool) {
	if o == nil || isNil(o.Street) {
    return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasStreet() bool {
	if o != nil && !isNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *CustomerAddressEntity) SetStreet(v string) {
	o.Street = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetNumber() string {
	if o == nil || isNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetNumberOk() (*string, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CustomerAddressEntity) SetNumber(v string) {
	o.Number = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetZip() string {
	if o == nil || isNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetZipOk() (*string, bool) {
	if o == nil || isNil(o.Zip) {
    return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasZip() bool {
	if o != nil && !isNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CustomerAddressEntity) SetZip(v string) {
	o.Zip = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerAddressEntity) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetProvince() string {
	if o == nil || isNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetProvinceOk() (*string, bool) {
	if o == nil || isNil(o.Province) {
    return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasProvince() bool {
	if o != nil && !isNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *CustomerAddressEntity) SetProvince(v string) {
	o.Province = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetPhoneNumber() string {
	if o == nil || isNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetPhoneNumberOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumber) {
    return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasPhoneNumber() bool {
	if o != nil && !isNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerAddressEntity) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetFiscalCode returns the FiscalCode field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetFiscalCode() string {
	if o == nil || isNil(o.FiscalCode) {
		var ret string
		return ret
	}
	return *o.FiscalCode
}

// GetFiscalCodeOk returns a tuple with the FiscalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetFiscalCodeOk() (*string, bool) {
	if o == nil || isNil(o.FiscalCode) {
    return nil, false
	}
	return o.FiscalCode, true
}

// HasFiscalCode returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasFiscalCode() bool {
	if o != nil && !isNil(o.FiscalCode) {
		return true
	}

	return false
}

// SetFiscalCode gets a reference to the given string and assigns it to the FiscalCode field.
func (o *CustomerAddressEntity) SetFiscalCode(v string) {
	o.FiscalCode = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetVatNumber() string {
	if o == nil || isNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetVatNumberOk() (*string, bool) {
	if o == nil || isNil(o.VatNumber) {
    return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasVatNumber() bool {
	if o != nil && !isNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CustomerAddressEntity) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetKind() CustomerAddressEntityKind {
	if o == nil || isNil(o.Kind) {
		var ret CustomerAddressEntityKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetKindOk() (*CustomerAddressEntityKind, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CustomerAddressEntityKind and assigns it to the Kind field.
func (o *CustomerAddressEntity) SetKind(v CustomerAddressEntityKind) {
	o.Kind = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CustomerAddressEntity) SetDefault(v bool) {
	o.Default = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerAddressEntity) SetCountry(v string) {
	o.Country = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerAddressEntity) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressEntity) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerAddressEntity) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *CustomerAddressEntity) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o CustomerAddressEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCustomerAddressEntity struct {
	value *CustomerAddressEntity
	isSet bool
}

func (v NullableCustomerAddressEntity) Get() *CustomerAddressEntity {
	return v.value
}

func (v *NullableCustomerAddressEntity) Set(val *CustomerAddressEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressEntity(val *CustomerAddressEntity) *NullableCustomerAddressEntity {
	return &NullableCustomerAddressEntity{value: val, isSet: true}
}

func (v NullableCustomerAddressEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


