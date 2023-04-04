# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: version not set
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CustomerApi* | [**CustomerAcquireSubscriber**](docs/CustomerApi.md#customeracquiresubscriber) | **Post** /customer.Customer/AcquireSubscriber | 
*CustomerApi* | [**CustomerAcquireUnsubscriber**](docs/CustomerApi.md#customeracquireunsubscriber) | **Post** /customer.Customer/AcquireUnsubscriber | 
*CustomerApi* | [**CustomerAddCustomerToGroup**](docs/CustomerApi.md#customeraddcustomertogroup) | **Post** /customer.Customer/AddCustomerToGroup | 
*CustomerApi* | [**CustomerBulkUpdate**](docs/CustomerApi.md#customerbulkupdate) | **Post** /customer.Customer/BulkUpdate | 
*CustomerApi* | [**CustomerCreate**](docs/CustomerApi.md#customercreate) | **Post** /customer.Customer/Create | 
*CustomerApi* | [**CustomerCreateAddress**](docs/CustomerApi.md#customercreateaddress) | **Post** /customer.Customer/CreateAddress | 
*CustomerApi* | [**CustomerCreateGroup**](docs/CustomerApi.md#customercreategroup) | **Post** /customer.Customer/CreateGroup | 
*CustomerApi* | [**CustomerCreateSubscriber**](docs/CustomerApi.md#customercreatesubscriber) | **Post** /customer.Customer/CreateSubscriber | 
*CustomerApi* | [**CustomerDeleteAddress**](docs/CustomerApi.md#customerdeleteaddress) | **Post** /customer.Customer/DeleteAddress | 
*CustomerApi* | [**CustomerDeleteGroup**](docs/CustomerApi.md#customerdeletegroup) | **Post** /customer.Customer/DeleteGroup | 
*CustomerApi* | [**CustomerFind**](docs/CustomerApi.md#customerfind) | **Post** /customer.Customer/Find | 
*CustomerApi* | [**CustomerFindByEmail**](docs/CustomerApi.md#customerfindbyemail) | **Post** /customer.Customer/FindByEmail | 
*CustomerApi* | [**CustomerFindById**](docs/CustomerApi.md#customerfindbyid) | **Post** /customer.Customer/FindById | 
*CustomerApi* | [**CustomerFindSubscriberByEmail**](docs/CustomerApi.md#customerfindsubscriberbyemail) | **Post** /customer.Customer/FindSubscriberByEmail | 
*CustomerApi* | [**CustomerFindSubscriberById**](docs/CustomerApi.md#customerfindsubscriberbyid) | **Post** /customer.Customer/FindSubscriberById | 
*CustomerApi* | [**CustomerGetGroupByCode**](docs/CustomerApi.md#customergetgroupbycode) | **Post** /customer.Customer/GetGroupByCode | 
*CustomerApi* | [**CustomerGetGroupById**](docs/CustomerApi.md#customergetgroupbyid) | **Post** /customer.Customer/GetGroupById | 
*CustomerApi* | [**CustomerListGroups**](docs/CustomerApi.md#customerlistgroups) | **Post** /customer.Customer/ListGroups | 
*CustomerApi* | [**CustomerRemoveCustomerFromGroup**](docs/CustomerApi.md#customerremovecustomerfromgroup) | **Post** /customer.Customer/RemoveCustomerFromGroup | 
*CustomerApi* | [**CustomerRemoveDefaultAddress**](docs/CustomerApi.md#customerremovedefaultaddress) | **Post** /customer.Customer/RemoveDefaultAddress | 
*CustomerApi* | [**CustomerSearch**](docs/CustomerApi.md#customersearch) | **Post** /customer.Customer/Search | 
*CustomerApi* | [**CustomerSetDefaultAddress**](docs/CustomerApi.md#customersetdefaultaddress) | **Post** /customer.Customer/SetDefaultAddress | 
*CustomerApi* | [**CustomerUnsubscribe**](docs/CustomerApi.md#customerunsubscribe) | **Post** /customer.Customer/Unsubscribe | 
*CustomerApi* | [**CustomerUpdate**](docs/CustomerApi.md#customerupdate) | **Post** /customer.Customer/Update | 
*CustomerApi* | [**CustomerUpdateAddress**](docs/CustomerApi.md#customerupdateaddress) | **Post** /customer.Customer/UpdateAddress | 
*CustomerApi* | [**CustomerUpdateGroup**](docs/CustomerApi.md#customerupdategroup) | **Post** /customer.Customer/UpdateGroup | 
*CustomerApi* | [**CustomerUpdateSubscriber**](docs/CustomerApi.md#customerupdatesubscriber) | **Post** /customer.Customer/UpdateSubscriber | 


## Documentation For Models

 - [BulkUpdateRequestAction](docs/BulkUpdateRequestAction.md)
 - [CustomerAddCustomerToGroupRequest](docs/CustomerAddCustomerToGroupRequest.md)
 - [CustomerAddressCreateRequest](docs/CustomerAddressCreateRequest.md)
 - [CustomerAddressCreateRequestKind](docs/CustomerAddressCreateRequestKind.md)
 - [CustomerAddressCustomerResponse](docs/CustomerAddressCustomerResponse.md)
 - [CustomerAddressCustomerResponseKind](docs/CustomerAddressCustomerResponseKind.md)
 - [CustomerAddressDeleteRequest](docs/CustomerAddressDeleteRequest.md)
 - [CustomerAddressDeleteResponse](docs/CustomerAddressDeleteResponse.md)
 - [CustomerAddressEntity](docs/CustomerAddressEntity.md)
 - [CustomerAddressEntityKind](docs/CustomerAddressEntityKind.md)
 - [CustomerAddressUpdateRequest](docs/CustomerAddressUpdateRequest.md)
 - [CustomerAddressUpdateResponse](docs/CustomerAddressUpdateResponse.md)
 - [CustomerBulkUpdateRequest](docs/CustomerBulkUpdateRequest.md)
 - [CustomerBulkUpdateResponse](docs/CustomerBulkUpdateResponse.md)
 - [CustomerCreateGroupRequest](docs/CustomerCreateGroupRequest.md)
 - [CustomerCreateRequest](docs/CustomerCreateRequest.md)
 - [CustomerCreateSubscriberRequest](docs/CustomerCreateSubscriberRequest.md)
 - [CustomerCustomerResponse](docs/CustomerCustomerResponse.md)
 - [CustomerDeleteGroupRequest](docs/CustomerDeleteGroupRequest.md)
 - [CustomerDeleteGroupResponse](docs/CustomerDeleteGroupResponse.md)
 - [CustomerEMFields](docs/CustomerEMFields.md)
 - [CustomerFindByEmailRequest](docs/CustomerFindByEmailRequest.md)
 - [CustomerFindByIdRequest](docs/CustomerFindByIdRequest.md)
 - [CustomerFindManyRequest](docs/CustomerFindManyRequest.md)
 - [CustomerFindManyRequestFilter](docs/CustomerFindManyRequestFilter.md)
 - [CustomerFindManyResponse](docs/CustomerFindManyResponse.md)
 - [CustomerFindSubscriberByEmailRequest](docs/CustomerFindSubscriberByEmailRequest.md)
 - [CustomerFindSubscriberByIdRequest](docs/CustomerFindSubscriberByIdRequest.md)
 - [CustomerGetGroupByCodeRequest](docs/CustomerGetGroupByCodeRequest.md)
 - [CustomerGetGroupByIdRequest](docs/CustomerGetGroupByIdRequest.md)
 - [CustomerGroupResponse](docs/CustomerGroupResponse.md)
 - [CustomerListGroupsRequest](docs/CustomerListGroupsRequest.md)
 - [CustomerListGroupsResponse](docs/CustomerListGroupsResponse.md)
 - [CustomerNewsletterRequest](docs/CustomerNewsletterRequest.md)
 - [CustomerNewsletterResponse](docs/CustomerNewsletterResponse.md)
 - [CustomerPassword](docs/CustomerPassword.md)
 - [CustomerRemoveCustomerFromGroupRequest](docs/CustomerRemoveCustomerFromGroupRequest.md)
 - [CustomerRemoveDefaultAddressRequest](docs/CustomerRemoveDefaultAddressRequest.md)
 - [CustomerSearchRequest](docs/CustomerSearchRequest.md)
 - [CustomerSearchRequestFilter](docs/CustomerSearchRequestFilter.md)
 - [CustomerSearchResponse](docs/CustomerSearchResponse.md)
 - [CustomerSetDefaultAddressRequest](docs/CustomerSetDefaultAddressRequest.md)
 - [CustomerSubscriberRequest](docs/CustomerSubscriberRequest.md)
 - [CustomerSubscriberResponse](docs/CustomerSubscriberResponse.md)
 - [CustomerSubscriberResponseWithNewsletterRequest](docs/CustomerSubscriberResponseWithNewsletterRequest.md)
 - [CustomerUnsubscribeRequest](docs/CustomerUnsubscribeRequest.md)
 - [CustomerUnsubscribeResponse](docs/CustomerUnsubscribeResponse.md)
 - [CustomerUpdateGroupRequest](docs/CustomerUpdateGroupRequest.md)
 - [CustomerUpdateGroupRequestPayload](docs/CustomerUpdateGroupRequestPayload.md)
 - [CustomerUpdateRequest](docs/CustomerUpdateRequest.md)
 - [CustomerUpdateRequestPayload](docs/CustomerUpdateRequestPayload.md)
 - [CustomerUpdateSubscriberRequest](docs/CustomerUpdateSubscriberRequest.md)
 - [GooglerpcStatus](docs/GooglerpcStatus.md)
 - [PasswordPasswordType](docs/PasswordPasswordType.md)
 - [ProtobufAny](docs/ProtobufAny.md)


## Documentation For Authorization



### geminiAuthorization

- **Type**: API key
- **API key parameter name**: X-Gem-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Gem-Token and passed in as the auth context for each request.


### standardAuthorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



