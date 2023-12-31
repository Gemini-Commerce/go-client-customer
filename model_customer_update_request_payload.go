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

// CustomerUpdateRequestPayload struct for CustomerUpdateRequestPayload
type CustomerUpdateRequestPayload struct {
	Em *CustomerEMFields `json:"em,omitempty"`
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Email *string `json:"email,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Source *string `json:"source,omitempty"`
	Addresses []CustomerAddressEntity `json:"addresses,omitempty"`
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	DefaultShippingAddressId *string `json:"defaultShippingAddressId,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Newsletters []CustomerNewsletterRequest `json:"newsletters,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	MigratedPassword *CustomerPassword `json:"migratedPassword,omitempty"`
	PreferredLocale *string `json:"preferredLocale,omitempty"`
	TaxCode *string `json:"taxCode,omitempty"`
	CertifiedEmail *string `json:"certifiedEmail,omitempty"`
	Market *string `json:"market,omitempty"`
	ExternalIds *map[string]string `json:"externalIds,omitempty"`
}

// NewCustomerUpdateRequestPayload instantiates a new CustomerUpdateRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerUpdateRequestPayload() *CustomerUpdateRequestPayload {
	this := CustomerUpdateRequestPayload{}
	return &this
}

// NewCustomerUpdateRequestPayloadWithDefaults instantiates a new CustomerUpdateRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerUpdateRequestPayloadWithDefaults() *CustomerUpdateRequestPayload {
	this := CustomerUpdateRequestPayload{}
	return &this
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetEm() CustomerEMFields {
	if o == nil || isNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || isNil(o.Em) {
    return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasEm() bool {
	if o != nil && !isNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerUpdateRequestPayload) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerUpdateRequestPayload) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetSurname() string {
	if o == nil || isNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetSurnameOk() (*string, bool) {
	if o == nil || isNil(o.Surname) {
    return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasSurname() bool {
	if o != nil && !isNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerUpdateRequestPayload) SetSurname(v string) {
	o.Surname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerUpdateRequestPayload) SetEmail(v string) {
	o.Email = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetBirthdate() time.Time {
	if o == nil || isNil(o.Birthdate) {
		var ret time.Time
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetBirthdateOk() (*time.Time, bool) {
	if o == nil || isNil(o.Birthdate) {
    return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasBirthdate() bool {
	if o != nil && !isNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given time.Time and assigns it to the Birthdate field.
func (o *CustomerUpdateRequestPayload) SetBirthdate(v time.Time) {
	o.Birthdate = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CustomerUpdateRequestPayload) SetGender(v string) {
	o.Gender = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomerUpdateRequestPayload) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CustomerUpdateRequestPayload) SetSource(v string) {
	o.Source = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetAddresses() []CustomerAddressEntity {
	if o == nil || isNil(o.Addresses) {
		var ret []CustomerAddressEntity
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetAddressesOk() ([]CustomerAddressEntity, bool) {
	if o == nil || isNil(o.Addresses) {
    return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasAddresses() bool {
	if o != nil && !isNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []CustomerAddressEntity and assigns it to the Addresses field.
func (o *CustomerUpdateRequestPayload) SetAddresses(v []CustomerAddressEntity) {
	o.Addresses = v
}

// GetDefaultBillingAddressId returns the DefaultBillingAddressId field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetDefaultBillingAddressId() string {
	if o == nil || isNil(o.DefaultBillingAddressId) {
		var ret string
		return ret
	}
	return *o.DefaultBillingAddressId
}

// GetDefaultBillingAddressIdOk returns a tuple with the DefaultBillingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetDefaultBillingAddressIdOk() (*string, bool) {
	if o == nil || isNil(o.DefaultBillingAddressId) {
    return nil, false
	}
	return o.DefaultBillingAddressId, true
}

// HasDefaultBillingAddressId returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasDefaultBillingAddressId() bool {
	if o != nil && !isNil(o.DefaultBillingAddressId) {
		return true
	}

	return false
}

// SetDefaultBillingAddressId gets a reference to the given string and assigns it to the DefaultBillingAddressId field.
func (o *CustomerUpdateRequestPayload) SetDefaultBillingAddressId(v string) {
	o.DefaultBillingAddressId = &v
}

// GetDefaultShippingAddressId returns the DefaultShippingAddressId field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetDefaultShippingAddressId() string {
	if o == nil || isNil(o.DefaultShippingAddressId) {
		var ret string
		return ret
	}
	return *o.DefaultShippingAddressId
}

// GetDefaultShippingAddressIdOk returns a tuple with the DefaultShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetDefaultShippingAddressIdOk() (*string, bool) {
	if o == nil || isNil(o.DefaultShippingAddressId) {
    return nil, false
	}
	return o.DefaultShippingAddressId, true
}

// HasDefaultShippingAddressId returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasDefaultShippingAddressId() bool {
	if o != nil && !isNil(o.DefaultShippingAddressId) {
		return true
	}

	return false
}

// SetDefaultShippingAddressId gets a reference to the given string and assigns it to the DefaultShippingAddressId field.
func (o *CustomerUpdateRequestPayload) SetDefaultShippingAddressId(v string) {
	o.DefaultShippingAddressId = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetPhoneNumber() string {
	if o == nil || isNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetPhoneNumberOk() (*string, bool) {
	if o == nil || isNil(o.PhoneNumber) {
    return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasPhoneNumber() bool {
	if o != nil && !isNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerUpdateRequestPayload) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetNationality() string {
	if o == nil || isNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetNationalityOk() (*string, bool) {
	if o == nil || isNil(o.Nationality) {
    return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasNationality() bool {
	if o != nil && !isNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *CustomerUpdateRequestPayload) SetNationality(v string) {
	o.Nationality = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *CustomerUpdateRequestPayload) SetGroups(v []string) {
	o.Groups = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetDeleted() bool {
	if o == nil || isNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetDeletedOk() (*bool, bool) {
	if o == nil || isNil(o.Deleted) {
    return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasDeleted() bool {
	if o != nil && !isNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CustomerUpdateRequestPayload) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetNewsletters() []CustomerNewsletterRequest {
	if o == nil || isNil(o.Newsletters) {
		var ret []CustomerNewsletterRequest
		return ret
	}
	return o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetNewslettersOk() ([]CustomerNewsletterRequest, bool) {
	if o == nil || isNil(o.Newsletters) {
    return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasNewsletters() bool {
	if o != nil && !isNil(o.Newsletters) {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given []CustomerNewsletterRequest and assigns it to the Newsletters field.
func (o *CustomerUpdateRequestPayload) SetNewsletters(v []CustomerNewsletterRequest) {
	o.Newsletters = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetAttributes() map[string]ProtobufAny {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *CustomerUpdateRequestPayload) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetMigratedPassword returns the MigratedPassword field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetMigratedPassword() CustomerPassword {
	if o == nil || isNil(o.MigratedPassword) {
		var ret CustomerPassword
		return ret
	}
	return *o.MigratedPassword
}

// GetMigratedPasswordOk returns a tuple with the MigratedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetMigratedPasswordOk() (*CustomerPassword, bool) {
	if o == nil || isNil(o.MigratedPassword) {
    return nil, false
	}
	return o.MigratedPassword, true
}

// HasMigratedPassword returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasMigratedPassword() bool {
	if o != nil && !isNil(o.MigratedPassword) {
		return true
	}

	return false
}

// SetMigratedPassword gets a reference to the given CustomerPassword and assigns it to the MigratedPassword field.
func (o *CustomerUpdateRequestPayload) SetMigratedPassword(v CustomerPassword) {
	o.MigratedPassword = &v
}

// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetPreferredLocale() string {
	if o == nil || isNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || isNil(o.PreferredLocale) {
    return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasPreferredLocale() bool {
	if o != nil && !isNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *CustomerUpdateRequestPayload) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetTaxCode() string {
	if o == nil || isNil(o.TaxCode) {
		var ret string
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetTaxCodeOk() (*string, bool) {
	if o == nil || isNil(o.TaxCode) {
    return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasTaxCode() bool {
	if o != nil && !isNil(o.TaxCode) {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given string and assigns it to the TaxCode field.
func (o *CustomerUpdateRequestPayload) SetTaxCode(v string) {
	o.TaxCode = &v
}

// GetCertifiedEmail returns the CertifiedEmail field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetCertifiedEmail() string {
	if o == nil || isNil(o.CertifiedEmail) {
		var ret string
		return ret
	}
	return *o.CertifiedEmail
}

// GetCertifiedEmailOk returns a tuple with the CertifiedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetCertifiedEmailOk() (*string, bool) {
	if o == nil || isNil(o.CertifiedEmail) {
    return nil, false
	}
	return o.CertifiedEmail, true
}

// HasCertifiedEmail returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasCertifiedEmail() bool {
	if o != nil && !isNil(o.CertifiedEmail) {
		return true
	}

	return false
}

// SetCertifiedEmail gets a reference to the given string and assigns it to the CertifiedEmail field.
func (o *CustomerUpdateRequestPayload) SetCertifiedEmail(v string) {
	o.CertifiedEmail = &v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetMarket() string {
	if o == nil || isNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetMarketOk() (*string, bool) {
	if o == nil || isNil(o.Market) {
    return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasMarket() bool {
	if o != nil && !isNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CustomerUpdateRequestPayload) SetMarket(v string) {
	o.Market = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *CustomerUpdateRequestPayload) GetExternalIds() map[string]string {
	if o == nil || isNil(o.ExternalIds) {
		var ret map[string]string
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerUpdateRequestPayload) GetExternalIdsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ExternalIds) {
    return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *CustomerUpdateRequestPayload) HasExternalIds() bool {
	if o != nil && !isNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given map[string]string and assigns it to the ExternalIds field.
func (o *CustomerUpdateRequestPayload) SetExternalIds(v map[string]string) {
	o.ExternalIds = &v
}

func (o CustomerUpdateRequestPayload) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !isNil(o.DefaultBillingAddressId) {
		toSerialize["defaultBillingAddressId"] = o.DefaultBillingAddressId
	}
	if !isNil(o.DefaultShippingAddressId) {
		toSerialize["defaultShippingAddressId"] = o.DefaultShippingAddressId
	}
	if !isNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !isNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !isNil(o.Newsletters) {
		toSerialize["newsletters"] = o.Newsletters
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !isNil(o.MigratedPassword) {
		toSerialize["migratedPassword"] = o.MigratedPassword
	}
	if !isNil(o.PreferredLocale) {
		toSerialize["preferredLocale"] = o.PreferredLocale
	}
	if !isNil(o.TaxCode) {
		toSerialize["taxCode"] = o.TaxCode
	}
	if !isNil(o.CertifiedEmail) {
		toSerialize["certifiedEmail"] = o.CertifiedEmail
	}
	if !isNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !isNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerUpdateRequestPayload struct {
	value *CustomerUpdateRequestPayload
	isSet bool
}

func (v NullableCustomerUpdateRequestPayload) Get() *CustomerUpdateRequestPayload {
	return v.value
}

func (v *NullableCustomerUpdateRequestPayload) Set(val *CustomerUpdateRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerUpdateRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerUpdateRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerUpdateRequestPayload(val *CustomerUpdateRequestPayload) *NullableCustomerUpdateRequestPayload {
	return &NullableCustomerUpdateRequestPayload{value: val, isSet: true}
}

func (v NullableCustomerUpdateRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerUpdateRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


