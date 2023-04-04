# # CustomerUpdateRequestPayload


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Em**| [**CustomerEMFields**](CustomerEMFields.md) |   | [optional]
**Name**| **string** |   | [optional]
**Surname**| **string** |   | [optional]
**Email**| **string** |   | [optional]
**Birthdate**| [**time.Time**](time.Time.md) |   | [optional]
**Gender**| **string** |   | [optional]
**Enabled**| **bool** |   | [optional]
**Source**| **string** |   | [optional]
**Addresses**| [**[]CustomerAddressEntity**](CustomerAddressEntity.md) |   | [optional]
**DefaultBillingAddressId**| **string** |   | [optional]
**DefaultShippingAddressId**| **string** |   | [optional]
**PhoneNumber**| **string** |   | [optional]
**Nationality**| **string** |   | [optional]
**Groups**| **[]string** |   | [optional]
**Deleted**| **bool** |   | [optional]
**Newsletters**| [**[]CustomerNewsletterRequest**](CustomerNewsletterRequest.md) |   | [optional]
**Attributes**| [**map[string]ProtobufAny**](ProtobufAny.md) |   | [optional]
**MigratedPassword**| [**CustomerPassword**](CustomerPassword.md) |   | [optional]
**PreferredLocale**| **string** |   | [optional]
**TaxCode**| **string** |   | [optional]
**CertifiedEmail**| **string** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

