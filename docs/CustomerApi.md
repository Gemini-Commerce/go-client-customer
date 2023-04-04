# \CustomerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerAcquireSubscriber**](CustomerApi.md#CustomerAcquireSubscriber) | **Post** /customer.Customer/AcquireSubscriber | 
[**CustomerAcquireUnsubscriber**](CustomerApi.md#CustomerAcquireUnsubscriber) | **Post** /customer.Customer/AcquireUnsubscriber | 
[**CustomerAddCustomerToGroup**](CustomerApi.md#CustomerAddCustomerToGroup) | **Post** /customer.Customer/AddCustomerToGroup | 
[**CustomerBulkUpdate**](CustomerApi.md#CustomerBulkUpdate) | **Post** /customer.Customer/BulkUpdate | 
[**CustomerCreate**](CustomerApi.md#CustomerCreate) | **Post** /customer.Customer/Create | 
[**CustomerCreateAddress**](CustomerApi.md#CustomerCreateAddress) | **Post** /customer.Customer/CreateAddress | 
[**CustomerCreateGroup**](CustomerApi.md#CustomerCreateGroup) | **Post** /customer.Customer/CreateGroup | 
[**CustomerCreateSubscriber**](CustomerApi.md#CustomerCreateSubscriber) | **Post** /customer.Customer/CreateSubscriber | 
[**CustomerDeleteAddress**](CustomerApi.md#CustomerDeleteAddress) | **Post** /customer.Customer/DeleteAddress | 
[**CustomerDeleteGroup**](CustomerApi.md#CustomerDeleteGroup) | **Post** /customer.Customer/DeleteGroup | 
[**CustomerFind**](CustomerApi.md#CustomerFind) | **Post** /customer.Customer/Find | 
[**CustomerFindByEmail**](CustomerApi.md#CustomerFindByEmail) | **Post** /customer.Customer/FindByEmail | 
[**CustomerFindById**](CustomerApi.md#CustomerFindById) | **Post** /customer.Customer/FindById | 
[**CustomerFindSubscriberByEmail**](CustomerApi.md#CustomerFindSubscriberByEmail) | **Post** /customer.Customer/FindSubscriberByEmail | 
[**CustomerFindSubscriberById**](CustomerApi.md#CustomerFindSubscriberById) | **Post** /customer.Customer/FindSubscriberById | 
[**CustomerGetGroupByCode**](CustomerApi.md#CustomerGetGroupByCode) | **Post** /customer.Customer/GetGroupByCode | 
[**CustomerGetGroupById**](CustomerApi.md#CustomerGetGroupById) | **Post** /customer.Customer/GetGroupById | 
[**CustomerListGroups**](CustomerApi.md#CustomerListGroups) | **Post** /customer.Customer/ListGroups | 
[**CustomerRemoveCustomerFromGroup**](CustomerApi.md#CustomerRemoveCustomerFromGroup) | **Post** /customer.Customer/RemoveCustomerFromGroup | 
[**CustomerRemoveDefaultAddress**](CustomerApi.md#CustomerRemoveDefaultAddress) | **Post** /customer.Customer/RemoveDefaultAddress | 
[**CustomerSearch**](CustomerApi.md#CustomerSearch) | **Post** /customer.Customer/Search | 
[**CustomerSetDefaultAddress**](CustomerApi.md#CustomerSetDefaultAddress) | **Post** /customer.Customer/SetDefaultAddress | 
[**CustomerUnsubscribe**](CustomerApi.md#CustomerUnsubscribe) | **Post** /customer.Customer/Unsubscribe | 
[**CustomerUpdate**](CustomerApi.md#CustomerUpdate) | **Post** /customer.Customer/Update | 
[**CustomerUpdateAddress**](CustomerApi.md#CustomerUpdateAddress) | **Post** /customer.Customer/UpdateAddress | 
[**CustomerUpdateGroup**](CustomerApi.md#CustomerUpdateGroup) | **Post** /customer.Customer/UpdateGroup | 
[**CustomerUpdateSubscriber**](CustomerApi.md#CustomerUpdateSubscriber) | **Post** /customer.Customer/UpdateSubscriber | 



## CustomerAcquireSubscriber

> CustomerSubscriberResponse CustomerAcquireSubscriber(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerCreateSubscriberRequest() // CustomerCreateSubscriberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerAcquireSubscriber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerAcquireSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerAcquireSubscriber`: CustomerSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerAcquireSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAcquireSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerCreateSubscriberRequest**](CustomerCreateSubscriberRequest.md) |  | 

### Return type

[**CustomerSubscriberResponse**](CustomerSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAcquireUnsubscriber

> CustomerUnsubscribeResponse CustomerAcquireUnsubscriber(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerUnsubscribeRequest() // CustomerUnsubscribeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerAcquireUnsubscriber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerAcquireUnsubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerAcquireUnsubscriber`: CustomerUnsubscribeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerAcquireUnsubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAcquireUnsubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerUnsubscribeRequest**](CustomerUnsubscribeRequest.md) |  | 

### Return type

[**CustomerUnsubscribeResponse**](CustomerUnsubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerAddCustomerToGroup

> CustomerGroupResponse CustomerAddCustomerToGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerAddCustomerToGroupRequest() // CustomerAddCustomerToGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerAddCustomerToGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerAddCustomerToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerAddCustomerToGroup`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerAddCustomerToGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerAddCustomerToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerAddCustomerToGroupRequest**](CustomerAddCustomerToGroupRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerBulkUpdate

> CustomerBulkUpdateResponse CustomerBulkUpdate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerBulkUpdateRequest() // CustomerBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerBulkUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerBulkUpdate`: CustomerBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerBulkUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerBulkUpdateRequest**](CustomerBulkUpdateRequest.md) |  | 

### Return type

[**CustomerBulkUpdateResponse**](CustomerBulkUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCreate

> CustomerCustomerResponse CustomerCreate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerCreateRequest() // CustomerCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCreate`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerCreateRequest**](CustomerCreateRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCreateAddress

> CustomerAddressCustomerResponse CustomerCreateAddress(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerAddressCreateRequest() // CustomerAddressCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerCreateAddress(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerCreateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCreateAddress`: CustomerAddressCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerCreateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerAddressCreateRequest**](CustomerAddressCreateRequest.md) |  | 

### Return type

[**CustomerAddressCustomerResponse**](CustomerAddressCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCreateGroup

> CustomerGroupResponse CustomerCreateGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerCreateGroupRequest() // CustomerCreateGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerCreateGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerCreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCreateGroup`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerCreateGroupRequest**](CustomerCreateGroupRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCreateSubscriber

> CustomerSubscriberResponse CustomerCreateSubscriber(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerCreateSubscriberRequest() // CustomerCreateSubscriberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerCreateSubscriber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerCreateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCreateSubscriber`: CustomerSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerCreateSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerCreateSubscriberRequest**](CustomerCreateSubscriberRequest.md) |  | 

### Return type

[**CustomerSubscriberResponse**](CustomerSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerDeleteAddress

> CustomerAddressDeleteResponse CustomerDeleteAddress(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerAddressDeleteRequest() // CustomerAddressDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerDeleteAddress(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerDeleteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerDeleteAddress`: CustomerAddressDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerDeleteAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerDeleteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerAddressDeleteRequest**](CustomerAddressDeleteRequest.md) |  | 

### Return type

[**CustomerAddressDeleteResponse**](CustomerAddressDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerDeleteGroup

> CustomerDeleteGroupResponse CustomerDeleteGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerDeleteGroupRequest() // CustomerDeleteGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerDeleteGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerDeleteGroup`: CustomerDeleteGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerDeleteGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerDeleteGroupRequest**](CustomerDeleteGroupRequest.md) |  | 

### Return type

[**CustomerDeleteGroupResponse**](CustomerDeleteGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFind

> CustomerFindManyResponse CustomerFind(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerFindManyRequest() // CustomerFindManyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerFind(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerFind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerFind`: CustomerFindManyResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerFind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerFindManyRequest**](CustomerFindManyRequest.md) |  | 

### Return type

[**CustomerFindManyResponse**](CustomerFindManyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFindByEmail

> CustomerCustomerResponse CustomerFindByEmail(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerFindByEmailRequest() // CustomerFindByEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerFindByEmail(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerFindByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerFindByEmail`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerFindByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerFindByEmailRequest**](CustomerFindByEmailRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFindById

> CustomerCustomerResponse CustomerFindById(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerFindByIdRequest() // CustomerFindByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerFindById(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerFindById`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerFindById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerFindByIdRequest**](CustomerFindByIdRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFindSubscriberByEmail

> CustomerSubscriberResponse CustomerFindSubscriberByEmail(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerFindSubscriberByEmailRequest() // CustomerFindSubscriberByEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerFindSubscriberByEmail(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerFindSubscriberByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerFindSubscriberByEmail`: CustomerSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerFindSubscriberByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindSubscriberByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerFindSubscriberByEmailRequest**](CustomerFindSubscriberByEmailRequest.md) |  | 

### Return type

[**CustomerSubscriberResponse**](CustomerSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerFindSubscriberById

> CustomerSubscriberResponse CustomerFindSubscriberById(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerFindSubscriberByIdRequest() // CustomerFindSubscriberByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerFindSubscriberById(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerFindSubscriberById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerFindSubscriberById`: CustomerSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerFindSubscriberById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFindSubscriberByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerFindSubscriberByIdRequest**](CustomerFindSubscriberByIdRequest.md) |  | 

### Return type

[**CustomerSubscriberResponse**](CustomerSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetGroupByCode

> CustomerGroupResponse CustomerGetGroupByCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerGetGroupByCodeRequest() // CustomerGetGroupByCodeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerGetGroupByCode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerGetGroupByCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerGetGroupByCode`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerGetGroupByCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetGroupByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerGetGroupByCodeRequest**](CustomerGetGroupByCodeRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerGetGroupById

> CustomerGroupResponse CustomerGetGroupById(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerGetGroupByIdRequest() // CustomerGetGroupByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerGetGroupById(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerGetGroupById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerGetGroupById`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerGetGroupById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerGetGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerGetGroupByIdRequest**](CustomerGetGroupByIdRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerListGroups

> CustomerListGroupsResponse CustomerListGroups(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerListGroupsRequest() // CustomerListGroupsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerListGroups(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerListGroups`: CustomerListGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerListGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerListGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerListGroupsRequest**](CustomerListGroupsRequest.md) |  | 

### Return type

[**CustomerListGroupsResponse**](CustomerListGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerRemoveCustomerFromGroup

> CustomerGroupResponse CustomerRemoveCustomerFromGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerRemoveCustomerFromGroupRequest() // CustomerRemoveCustomerFromGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerRemoveCustomerFromGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerRemoveCustomerFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerRemoveCustomerFromGroup`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerRemoveCustomerFromGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerRemoveCustomerFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerRemoveCustomerFromGroupRequest**](CustomerRemoveCustomerFromGroupRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerRemoveDefaultAddress

> CustomerCustomerResponse CustomerRemoveDefaultAddress(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerRemoveDefaultAddressRequest() // CustomerRemoveDefaultAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerRemoveDefaultAddress(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerRemoveDefaultAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerRemoveDefaultAddress`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerRemoveDefaultAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerRemoveDefaultAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerRemoveDefaultAddressRequest**](CustomerRemoveDefaultAddressRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerSearch

> CustomerSearchResponse CustomerSearch(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerSearchRequest() // CustomerSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerSearch(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerSearch`: CustomerSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerSearchRequest**](CustomerSearchRequest.md) |  | 

### Return type

[**CustomerSearchResponse**](CustomerSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerSetDefaultAddress

> CustomerCustomerResponse CustomerSetDefaultAddress(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerSetDefaultAddressRequest() // CustomerSetDefaultAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerSetDefaultAddress(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerSetDefaultAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerSetDefaultAddress`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerSetDefaultAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerSetDefaultAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerSetDefaultAddressRequest**](CustomerSetDefaultAddressRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUnsubscribe

> CustomerUnsubscribeResponse CustomerUnsubscribe(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerUnsubscribeRequest() // CustomerUnsubscribeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerUnsubscribe(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerUnsubscribe`: CustomerUnsubscribeResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerUnsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerUnsubscribeRequest**](CustomerUnsubscribeRequest.md) |  | 

### Return type

[**CustomerUnsubscribeResponse**](CustomerUnsubscribeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUpdate

> CustomerCustomerResponse CustomerUpdate(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerUpdateRequest() // CustomerUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerUpdate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerUpdate`: CustomerCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerUpdateRequest**](CustomerUpdateRequest.md) |  | 

### Return type

[**CustomerCustomerResponse**](CustomerCustomerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUpdateAddress

> CustomerAddressUpdateResponse CustomerUpdateAddress(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerAddressUpdateRequest() // CustomerAddressUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerUpdateAddress(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerUpdateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerUpdateAddress`: CustomerAddressUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerUpdateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerAddressUpdateRequest**](CustomerAddressUpdateRequest.md) |  | 

### Return type

[**CustomerAddressUpdateResponse**](CustomerAddressUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUpdateGroup

> CustomerGroupResponse CustomerUpdateGroup(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerUpdateGroupRequest() // CustomerUpdateGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerUpdateGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerUpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerUpdateGroup`: CustomerGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerUpdateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerUpdateGroupRequest**](CustomerUpdateGroupRequest.md) |  | 

### Return type

[**CustomerGroupResponse**](CustomerGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerUpdateSubscriber

> CustomerSubscriberResponse CustomerUpdateSubscriber(ctx).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := *openapiclient.NewCustomerUpdateSubscriberRequest() // CustomerUpdateSubscriberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerApi.CustomerUpdateSubscriber(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CustomerUpdateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerUpdateSubscriber`: CustomerSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CustomerUpdateSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerUpdateSubscriberRequest**](CustomerUpdateSubscriberRequest.md) |  | 

### Return type

[**CustomerSubscriberResponse**](CustomerSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

