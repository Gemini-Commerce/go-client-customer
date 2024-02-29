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

// checks if the CustomerCustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCustomerResponse{}

// CustomerCustomerResponse struct for CustomerCustomerResponse
type CustomerCustomerResponse struct {
	Em *CustomerEMFields `json:"em,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Email *string `json:"email,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Gender *string `json:"gender,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Source *string `json:"source,omitempty"`
	Addresses []CustomerAddressCustomerResponse `json:"addresses,omitempty"`
	DefaultBillingAddressId *string `json:"defaultBillingAddressId,omitempty"`
	DefaultShippingAddressId *string `json:"defaultShippingAddressId,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	PreferredLocale *string `json:"preferredLocale,omitempty"`
	Groups []string `json:"groups,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Newsletters []CustomerNewsletterResponse `json:"newsletters,omitempty"`
	MigratedPassword *CustomerPassword `json:"migratedPassword,omitempty"`
	Attributes *map[string]ProtobufAny `json:"attributes,omitempty"`
	TaxCode *string `json:"taxCode,omitempty"`
	CertifiedEmail *string `json:"certifiedEmail,omitempty"`
	SdiCode *string `json:"sdiCode,omitempty"`
	FiscalCode *string `json:"fiscalCode,omitempty"`
	CompanyName *string `json:"companyName,omitempty"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo,omitempty"`
	Market *string `json:"market,omitempty"`
	Permissions []CustomerPermission `json:"permissions,omitempty"`
	ExternalIds *map[string]string `json:"externalIds,omitempty"`
	AgentGrn *string `json:"agentGrn,omitempty"`
	AggregationId *string `json:"aggregationId,omitempty"`
}

// NewCustomerCustomerResponse instantiates a new CustomerCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCustomerResponse() *CustomerCustomerResponse {
	this := CustomerCustomerResponse{}
	return &this
}

// NewCustomerCustomerResponseWithDefaults instantiates a new CustomerCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCustomerResponseWithDefaults() *CustomerCustomerResponse {
	this := CustomerCustomerResponse{}
	return &this
}

// GetEm returns the Em field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetEm() CustomerEMFields {
	if o == nil || IsNil(o.Em) {
		var ret CustomerEMFields
		return ret
	}
	return *o.Em
}

// GetEmOk returns a tuple with the Em field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetEmOk() (*CustomerEMFields, bool) {
	if o == nil || IsNil(o.Em) {
		return nil, false
	}
	return o.Em, true
}

// HasEm returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasEm() bool {
	if o != nil && !IsNil(o.Em) {
		return true
	}

	return false
}

// SetEm gets a reference to the given CustomerEMFields and assigns it to the Em field.
func (o *CustomerCustomerResponse) SetEm(v CustomerEMFields) {
	o.Em = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *CustomerCustomerResponse) SetGrn(v string) {
	o.Grn = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerCustomerResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerCustomerResponse) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CustomerCustomerResponse) SetSurname(v string) {
	o.Surname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CustomerCustomerResponse) SetEmail(v string) {
	o.Email = &v
}

// GetBirthdate returns the Birthdate field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetBirthdate() time.Time {
	if o == nil || IsNil(o.Birthdate) {
		var ret time.Time
		return ret
	}
	return *o.Birthdate
}

// GetBirthdateOk returns a tuple with the Birthdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetBirthdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Birthdate) {
		return nil, false
	}
	return o.Birthdate, true
}

// HasBirthdate returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasBirthdate() bool {
	if o != nil && !IsNil(o.Birthdate) {
		return true
	}

	return false
}

// SetBirthdate gets a reference to the given time.Time and assigns it to the Birthdate field.
func (o *CustomerCustomerResponse) SetBirthdate(v time.Time) {
	o.Birthdate = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *CustomerCustomerResponse) SetGender(v string) {
	o.Gender = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomerCustomerResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CustomerCustomerResponse) SetSource(v string) {
	o.Source = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetAddresses() []CustomerAddressCustomerResponse {
	if o == nil || IsNil(o.Addresses) {
		var ret []CustomerAddressCustomerResponse
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetAddressesOk() ([]CustomerAddressCustomerResponse, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []CustomerAddressCustomerResponse and assigns it to the Addresses field.
func (o *CustomerCustomerResponse) SetAddresses(v []CustomerAddressCustomerResponse) {
	o.Addresses = v
}

// GetDefaultBillingAddressId returns the DefaultBillingAddressId field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetDefaultBillingAddressId() string {
	if o == nil || IsNil(o.DefaultBillingAddressId) {
		var ret string
		return ret
	}
	return *o.DefaultBillingAddressId
}

// GetDefaultBillingAddressIdOk returns a tuple with the DefaultBillingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetDefaultBillingAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBillingAddressId) {
		return nil, false
	}
	return o.DefaultBillingAddressId, true
}

// HasDefaultBillingAddressId returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasDefaultBillingAddressId() bool {
	if o != nil && !IsNil(o.DefaultBillingAddressId) {
		return true
	}

	return false
}

// SetDefaultBillingAddressId gets a reference to the given string and assigns it to the DefaultBillingAddressId field.
func (o *CustomerCustomerResponse) SetDefaultBillingAddressId(v string) {
	o.DefaultBillingAddressId = &v
}

// GetDefaultShippingAddressId returns the DefaultShippingAddressId field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetDefaultShippingAddressId() string {
	if o == nil || IsNil(o.DefaultShippingAddressId) {
		var ret string
		return ret
	}
	return *o.DefaultShippingAddressId
}

// GetDefaultShippingAddressIdOk returns a tuple with the DefaultShippingAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetDefaultShippingAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultShippingAddressId) {
		return nil, false
	}
	return o.DefaultShippingAddressId, true
}

// HasDefaultShippingAddressId returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasDefaultShippingAddressId() bool {
	if o != nil && !IsNil(o.DefaultShippingAddressId) {
		return true
	}

	return false
}

// SetDefaultShippingAddressId gets a reference to the given string and assigns it to the DefaultShippingAddressId field.
func (o *CustomerCustomerResponse) SetDefaultShippingAddressId(v string) {
	o.DefaultShippingAddressId = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomerCustomerResponse) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *CustomerCustomerResponse) SetNationality(v string) {
	o.Nationality = &v
}

// GetPreferredLocale returns the PreferredLocale field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetPreferredLocale() string {
	if o == nil || IsNil(o.PreferredLocale) {
		var ret string
		return ret
	}
	return *o.PreferredLocale
}

// GetPreferredLocaleOk returns a tuple with the PreferredLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetPreferredLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredLocale) {
		return nil, false
	}
	return o.PreferredLocale, true
}

// HasPreferredLocale returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasPreferredLocale() bool {
	if o != nil && !IsNil(o.PreferredLocale) {
		return true
	}

	return false
}

// SetPreferredLocale gets a reference to the given string and assigns it to the PreferredLocale field.
func (o *CustomerCustomerResponse) SetPreferredLocale(v string) {
	o.PreferredLocale = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *CustomerCustomerResponse) SetGroups(v []string) {
	o.Groups = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomerCustomerResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CustomerCustomerResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CustomerCustomerResponse) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *CustomerCustomerResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetNewsletters returns the Newsletters field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetNewsletters() []CustomerNewsletterResponse {
	if o == nil || IsNil(o.Newsletters) {
		var ret []CustomerNewsletterResponse
		return ret
	}
	return o.Newsletters
}

// GetNewslettersOk returns a tuple with the Newsletters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetNewslettersOk() ([]CustomerNewsletterResponse, bool) {
	if o == nil || IsNil(o.Newsletters) {
		return nil, false
	}
	return o.Newsletters, true
}

// HasNewsletters returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasNewsletters() bool {
	if o != nil && !IsNil(o.Newsletters) {
		return true
	}

	return false
}

// SetNewsletters gets a reference to the given []CustomerNewsletterResponse and assigns it to the Newsletters field.
func (o *CustomerCustomerResponse) SetNewsletters(v []CustomerNewsletterResponse) {
	o.Newsletters = v
}

// GetMigratedPassword returns the MigratedPassword field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetMigratedPassword() CustomerPassword {
	if o == nil || IsNil(o.MigratedPassword) {
		var ret CustomerPassword
		return ret
	}
	return *o.MigratedPassword
}

// GetMigratedPasswordOk returns a tuple with the MigratedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetMigratedPasswordOk() (*CustomerPassword, bool) {
	if o == nil || IsNil(o.MigratedPassword) {
		return nil, false
	}
	return o.MigratedPassword, true
}

// HasMigratedPassword returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasMigratedPassword() bool {
	if o != nil && !IsNil(o.MigratedPassword) {
		return true
	}

	return false
}

// SetMigratedPassword gets a reference to the given CustomerPassword and assigns it to the MigratedPassword field.
func (o *CustomerCustomerResponse) SetMigratedPassword(v CustomerPassword) {
	o.MigratedPassword = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetAttributes() map[string]ProtobufAny {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]ProtobufAny
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetAttributesOk() (*map[string]ProtobufAny, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]ProtobufAny and assigns it to the Attributes field.
func (o *CustomerCustomerResponse) SetAttributes(v map[string]ProtobufAny) {
	o.Attributes = &v
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode) {
		var ret string
		return ret
	}
	return *o.TaxCode
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetTaxCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCode) {
		return nil, false
	}
	return o.TaxCode, true
}

// HasTaxCode returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasTaxCode() bool {
	if o != nil && !IsNil(o.TaxCode) {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given string and assigns it to the TaxCode field.
func (o *CustomerCustomerResponse) SetTaxCode(v string) {
	o.TaxCode = &v
}

// GetCertifiedEmail returns the CertifiedEmail field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetCertifiedEmail() string {
	if o == nil || IsNil(o.CertifiedEmail) {
		var ret string
		return ret
	}
	return *o.CertifiedEmail
}

// GetCertifiedEmailOk returns a tuple with the CertifiedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetCertifiedEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CertifiedEmail) {
		return nil, false
	}
	return o.CertifiedEmail, true
}

// HasCertifiedEmail returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasCertifiedEmail() bool {
	if o != nil && !IsNil(o.CertifiedEmail) {
		return true
	}

	return false
}

// SetCertifiedEmail gets a reference to the given string and assigns it to the CertifiedEmail field.
func (o *CustomerCustomerResponse) SetCertifiedEmail(v string) {
	o.CertifiedEmail = &v
}

// GetSdiCode returns the SdiCode field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetSdiCode() string {
	if o == nil || IsNil(o.SdiCode) {
		var ret string
		return ret
	}
	return *o.SdiCode
}

// GetSdiCodeOk returns a tuple with the SdiCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetSdiCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SdiCode) {
		return nil, false
	}
	return o.SdiCode, true
}

// HasSdiCode returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasSdiCode() bool {
	if o != nil && !IsNil(o.SdiCode) {
		return true
	}

	return false
}

// SetSdiCode gets a reference to the given string and assigns it to the SdiCode field.
func (o *CustomerCustomerResponse) SetSdiCode(v string) {
	o.SdiCode = &v
}

// GetFiscalCode returns the FiscalCode field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetFiscalCode() string {
	if o == nil || IsNil(o.FiscalCode) {
		var ret string
		return ret
	}
	return *o.FiscalCode
}

// GetFiscalCodeOk returns a tuple with the FiscalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetFiscalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalCode) {
		return nil, false
	}
	return o.FiscalCode, true
}

// HasFiscalCode returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasFiscalCode() bool {
	if o != nil && !IsNil(o.FiscalCode) {
		return true
	}

	return false
}

// SetFiscalCode gets a reference to the given string and assigns it to the FiscalCode field.
func (o *CustomerCustomerResponse) SetFiscalCode(v string) {
	o.FiscalCode = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *CustomerCustomerResponse) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetAdditionalInfo() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetAdditionalInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given map[string]interface{} and assigns it to the AdditionalInfo field.
func (o *CustomerCustomerResponse) SetAdditionalInfo(v map[string]interface{}) {
	o.AdditionalInfo = v
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetMarket() string {
	if o == nil || IsNil(o.Market) {
		var ret string
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetMarketOk() (*string, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given string and assigns it to the Market field.
func (o *CustomerCustomerResponse) SetMarket(v string) {
	o.Market = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetPermissions() []CustomerPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []CustomerPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetPermissionsOk() ([]CustomerPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CustomerPermission and assigns it to the Permissions field.
func (o *CustomerCustomerResponse) SetPermissions(v []CustomerPermission) {
	o.Permissions = v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetExternalIds() map[string]string {
	if o == nil || IsNil(o.ExternalIds) {
		var ret map[string]string
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetExternalIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given map[string]string and assigns it to the ExternalIds field.
func (o *CustomerCustomerResponse) SetExternalIds(v map[string]string) {
	o.ExternalIds = &v
}

// GetAgentGrn returns the AgentGrn field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetAgentGrn() string {
	if o == nil || IsNil(o.AgentGrn) {
		var ret string
		return ret
	}
	return *o.AgentGrn
}

// GetAgentGrnOk returns a tuple with the AgentGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetAgentGrnOk() (*string, bool) {
	if o == nil || IsNil(o.AgentGrn) {
		return nil, false
	}
	return o.AgentGrn, true
}

// HasAgentGrn returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasAgentGrn() bool {
	if o != nil && !IsNil(o.AgentGrn) {
		return true
	}

	return false
}

// SetAgentGrn gets a reference to the given string and assigns it to the AgentGrn field.
func (o *CustomerCustomerResponse) SetAgentGrn(v string) {
	o.AgentGrn = &v
}

// GetAggregationId returns the AggregationId field value if set, zero value otherwise.
func (o *CustomerCustomerResponse) GetAggregationId() string {
	if o == nil || IsNil(o.AggregationId) {
		var ret string
		return ret
	}
	return *o.AggregationId
}

// GetAggregationIdOk returns a tuple with the AggregationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCustomerResponse) GetAggregationIdOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationId) {
		return nil, false
	}
	return o.AggregationId, true
}

// HasAggregationId returns a boolean if a field has been set.
func (o *CustomerCustomerResponse) HasAggregationId() bool {
	if o != nil && !IsNil(o.AggregationId) {
		return true
	}

	return false
}

// SetAggregationId gets a reference to the given string and assigns it to the AggregationId field.
func (o *CustomerCustomerResponse) SetAggregationId(v string) {
	o.AggregationId = &v
}

func (o CustomerCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCustomerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Em) {
		toSerialize["em"] = o.Em
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Birthdate) {
		toSerialize["birthdate"] = o.Birthdate
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.DefaultBillingAddressId) {
		toSerialize["defaultBillingAddressId"] = o.DefaultBillingAddressId
	}
	if !IsNil(o.DefaultShippingAddressId) {
		toSerialize["defaultShippingAddressId"] = o.DefaultShippingAddressId
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.PreferredLocale) {
		toSerialize["preferredLocale"] = o.PreferredLocale
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if !IsNil(o.Newsletters) {
		toSerialize["newsletters"] = o.Newsletters
	}
	if !IsNil(o.MigratedPassword) {
		toSerialize["migratedPassword"] = o.MigratedPassword
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.TaxCode) {
		toSerialize["taxCode"] = o.TaxCode
	}
	if !IsNil(o.CertifiedEmail) {
		toSerialize["certifiedEmail"] = o.CertifiedEmail
	}
	if !IsNil(o.SdiCode) {
		toSerialize["sdiCode"] = o.SdiCode
	}
	if !IsNil(o.FiscalCode) {
		toSerialize["fiscalCode"] = o.FiscalCode
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.AgentGrn) {
		toSerialize["agentGrn"] = o.AgentGrn
	}
	if !IsNil(o.AggregationId) {
		toSerialize["aggregationId"] = o.AggregationId
	}
	return toSerialize, nil
}

type NullableCustomerCustomerResponse struct {
	value *CustomerCustomerResponse
	isSet bool
}

func (v NullableCustomerCustomerResponse) Get() *CustomerCustomerResponse {
	return v.value
}

func (v *NullableCustomerCustomerResponse) Set(val *CustomerCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCustomerResponse(val *CustomerCustomerResponse) *NullableCustomerCustomerResponse {
	return &NullableCustomerCustomerResponse{value: val, isSet: true}
}

func (v NullableCustomerCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


