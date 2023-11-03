# # CustomerCreateRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   | [optional]
**Em**| [**CustomerEMFields**](CustomerEMFields.md) |   | [optional]
**Name**| **string** |   | [optional]
**Surname**| **string** |   | [optional]
**Email**| **string** |   | [optional]
**Birthdate**| [**time.Time**](time.Time.md) |   | [optional]
**Gender**| **string** |   | [optional]
**Enabled**| **bool** |   | [optional]
**Source**| **string** |   | [optional]
**Addresses**| [**[]CustomerAddressEntity**](CustomerAddressEntity.md) |   | [optional]
**PhoneNumber**| **string** |   | [optional]
**Nationality**| **string** |   | [optional]
**Groups**| **[]string** |   | [optional]
**Deleted**| **bool** |   | [optional]
**Newsletters**| [**[]CustomerNewsletterRequest**](CustomerNewsletterRequest.md) |   | [optional]
**DoNotNotify**| **bool** |   | [optional]
**Attributes**| [**map[string]ProtobufAny**](ProtobufAny.md) |   | [optional]
**MigratedPassword**| [**CustomerPassword**](CustomerPassword.md) |   | [optional]
**Market**| **string** |   | [optional]
**PreferredLocale**| **string** |   | [optional]
**TaxCode**| **string** |   | [optional]
**CertifiedEmail**| **string** |   | [optional]
**ExternalIds**| **map[string]string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

