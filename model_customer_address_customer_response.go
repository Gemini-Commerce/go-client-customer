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
	"time"
)

// CustomerAddressCustomerResponse struct for CustomerAddressCustomerResponse
type CustomerAddressCustomerResponse struct {
	Em *CustomerEMFields `json:"em,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Id *string `json:"id,omitempty"`
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
	Kind *CustomerAddressCustomerResponseKind `json:"kind,omitempty"`
	Default *bool `json:"default,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
}

// NewCustomerAddressCustomerResponse instantiates a new CustomerAddressCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCustomerResponse() *CustomerAddressCustomerResponse {
	this := CustomerAddressCustomerResponse{}
	var kind CustomerAddressCustomerResponseKind = CUSTOMERADDRESSCUSTOMERRESPONSEKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// NewCustomerAddressCustomerResponseWithDefaults instantiates a new CustomerAddressCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCustomerResponseWithDefaults() *CustomerAddressCustomerResponse {
	this := CustomerAddressCustomerResponse{}
	var kind CustomerAddressCustomerResponseKind = CUSTOMERADDRESSCUSTOMERRESPONSEKIND_SHIPPING
	this.Kind = &kind
	return &this
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerAddressCustomerResponse) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetGrn() string {
	if o == nil || isNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetGrnOk() (*string, bool) {
	if o == nil || isNil(o.Grn) {
    return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasGrn() bool {
	if o != nil && !isNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *CustomerAddressCustomerResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerAddressCustomerResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerAddressCustomerResponse) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetSurname() string {
	if o == nil || isNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetSurnameOk() (*string, bool) {
	if o == nil || isNil(o.Surname) {
    return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasSurname() bool {
	if o != nil && !isNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerAddressCustomerResponse) SetSurname(v string) {
	o.Surname = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetStreet() string {
	if o == nil || isNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetStreetOk() (*string, bool) {
	if o == nil || isNil(o.Street) {
    return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasStreet() bool {
	if o != nil && !isNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *CustomerAddressCustomerResponse) SetStreet(v string) {
	o.Street = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetNumber() string {
	if o == nil || isNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetNumberOk() (*string, bool) {
	if o == nil || isNil(o.Number) {
    return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *CustomerAddressCustomerResponse) SetNumber(v string) {
	o.Number = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetZip() string {
	if o == nil || isNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetZipOk() (*string, bool) {
	if o == nil || isNil(o.Zip) {
    return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasZip() bool {
	if o != nil && !isNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CustomerAddressCustomerResponse) SetZip(v string) {
	o.Zip = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CustomerAddressCustomerResponse) SetCity(v string) {
	o.City = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetProvince() string {
	if o == nil || isNil(o.Province) {
		var ret string
		return ret
	}
	return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetProvinceOk() (*string, bool) {
	if o == nil || isNil(o.Province) {
    return nil, false
	}
	return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasProvince() bool {
	if o != nil && !isNil(o.Province) {
		return true
	}

	return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *CustomerAddressCustomerResponse) SetProvince(v string) {
	o.Province = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetPhoneNumber() string {
	if o == nil || isNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetPhoneNumberOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumber) {
    return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasPhoneNumber() bool {
	if o != nil && !isNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerAddressCustomerResponse) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetFiscalCode returns the FiscalCode field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetFiscalCode() string {
	if o == nil || isNil(o.FiscalCode) {
		var ret string
		return ret
	}
	return *o.FiscalCode
}

// GetFiscalCodeOk returns a tuple with the FiscalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetFiscalCodeOk() (*string, bool) {
	if o == nil || isNil(o.FiscalCode) {
    return nil, false
	}
	return o.FiscalCode, true
}

// HasFiscalCode returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasFiscalCode() bool {
	if o != nil && !isNil(o.FiscalCode) {
		return true
	}

	return false
}

// SetFiscalCode gets a reference to the given string and assigns it to the FiscalCode field.
func (o *CustomerAddressCustomerResponse) SetFiscalCode(v string) {
	o.FiscalCode = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetVatNumber() string {
	if o == nil || isNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetVatNumberOk() (*string, bool) {
	if o == nil || isNil(o.VatNumber) {
    return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasVatNumber() bool {
	if o != nil && !isNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *CustomerAddressCustomerResponse) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetKind() CustomerAddressCustomerResponseKind {
	if o == nil || isNil(o.Kind) {
		var ret CustomerAddressCustomerResponseKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetKindOk() (*CustomerAddressCustomerResponseKind, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CustomerAddressCustomerResponseKind and assigns it to the Kind field.
func (o *CustomerAddressCustomerResponse) SetKind(v CustomerAddressCustomerResponseKind) {
	o.Kind = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
    return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CustomerAddressCustomerResponse) SetDefault(v bool) {
	o.Default = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CustomerAddressCustomerResponse) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerAddressCustomerResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerAddressCustomerResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerAddressCustomerResponse) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerAddressCustomerResponse) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerAddressCustomerResponse) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *CustomerAddressCustomerResponse) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

func (o CustomerAddressCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Em) {
		toSerialize["em"] = o.Em
	}
	if !isNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerAddressCustomerResponse struct {
	value *CustomerAddressCustomerResponse
	isSet bool
}

func (v NullableCustomerAddressCustomerResponse) Get() *CustomerAddressCustomerResponse {
	return v.value
}

func (v *NullableCustomerAddressCustomerResponse) Set(val *CustomerAddressCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCustomerResponse(val *CustomerAddressCustomerResponse) *NullableCustomerAddressCustomerResponse {
	return &NullableCustomerAddressCustomerResponse{value: val, isSet: true}
}

func (v NullableCustomerAddressCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


