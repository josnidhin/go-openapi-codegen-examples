# \TransactionsApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncTransactionsPost**](TransactionsApi.md#AsyncTransactionsPost) | **Post** /async/transactions | Create a transaction asynchronously
[**AsyncTransactionsTransactionIdConfirmPost**](TransactionsApi.md#AsyncTransactionsTransactionIdConfirmPost) | **Post** /async/transactions/{transaction_id}/confirm | Confirm a transaction asynchronously
[**SyncTransactionsPost**](TransactionsApi.md#SyncTransactionsPost) | **Post** /sync/transactions | Create a transaction synchronously
[**SyncTransactionsTransactionIdConfirmPost**](TransactionsApi.md#SyncTransactionsTransactionIdConfirmPost) | **Post** /sync/transactions/{transaction_id}/confirm | Confirm a transaction synchronously
[**TransactionsGet**](TransactionsApi.md#TransactionsGet) | **Get** /transactions | Query list of transactions
[**TransactionsTransactionIdCancelPost**](TransactionsApi.md#TransactionsTransactionIdCancelPost) | **Post** /transactions/{transaction_id}/cancel | Cancel a transaction
[**TransactionsTransactionIdGet**](TransactionsApi.md#TransactionsTransactionIdGet) | **Get** /transactions/{transaction_id} | Query a transaction by ID



## AsyncTransactionsPost

> Transaction AsyncTransactionsPost(ctx).AsyncTransactionsPostRequest(asyncTransactionsPostRequest).Execute()

Create a transaction asynchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    asyncTransactionsPostRequest := *openapiclient.NewAsyncTransactionsPostRequest(string(123), int32(123)) // AsyncTransactionsPostRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.AsyncTransactionsPost(context.Background()).AsyncTransactionsPostRequest(asyncTransactionsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AsyncTransactionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AsyncTransactionsPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.AsyncTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAsyncTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asyncTransactionsPostRequest** | [**AsyncTransactionsPostRequest**](AsyncTransactionsPostRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AsyncTransactionsTransactionIdConfirmPost

> Transaction AsyncTransactionsTransactionIdConfirmPost(ctx, transactionId).Execute()

Confirm a transaction asynchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.AsyncTransactionsTransactionIdConfirmPost(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.AsyncTransactionsTransactionIdConfirmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AsyncTransactionsTransactionIdConfirmPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.AsyncTransactionsTransactionIdConfirmPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncTransactionsTransactionIdConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncTransactionsPost

> Transaction SyncTransactionsPost(ctx).Transaction(transaction).Execute()

Create a transaction synchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transaction := *openapiclient.NewTransaction(string(123), int32(123)) // Transaction |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.SyncTransactionsPost(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.SyncTransactionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncTransactionsPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.SyncTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | [**Transaction**](Transaction.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncTransactionsTransactionIdConfirmPost

> Transaction SyncTransactionsTransactionIdConfirmPost(ctx, transactionId).Execute()

Confirm a transaction synchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.SyncTransactionsTransactionIdConfirmPost(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.SyncTransactionsTransactionIdConfirmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncTransactionsTransactionIdConfirmPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.SyncTransactionsTransactionIdConfirmPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncTransactionsTransactionIdConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> []AsyncTransactionsPostRequest TransactionsGet(ctx).ExternalId(externalId).ProductType(productType).ServiceId(serviceId).CountryIsoCode(countryIsoCode).OperatorId(operatorId).StatusId(statusId).CreditPartyMobileNumber(creditPartyMobileNumber).CreditPartyAccountNumber(creditPartyAccountNumber).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Execute()

Query list of transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    externalId := "externalId_example" // string |  (optional)
    productType := openapiclient.ProductTypes("FIXED_VALUE_RECHARGE") // ProductTypes |  (optional)
    serviceId := int32(56) // int32 |  (optional)
    countryIsoCode := "countryIsoCode_example" // string |  (optional)
    operatorId := int32(56) // int32 |  (optional)
    statusId := int32(56) // int32 |  (optional)
    creditPartyMobileNumber := "creditPartyMobileNumber_example" // string |  (optional)
    creditPartyAccountNumber := "creditPartyAccountNumber_example" // string |  (optional)
    fromDate := time.Now() // time.Time | Starting date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between `from_date` and `to_date` can not exceed 24 hours. (optional)
    toDate := time.Now() // time.Time | Ending date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between `from_date` and `to_date` can not exceed 24 hours. (optional)
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.TransactionsGet(context.Background()).ExternalId(externalId).ProductType(productType).ServiceId(serviceId).CountryIsoCode(countryIsoCode).OperatorId(operatorId).StatusId(statusId).CreditPartyMobileNumber(creditPartyMobileNumber).CreditPartyAccountNumber(creditPartyAccountNumber).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsGet`: []AsyncTransactionsPostRequest
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** |  | 
 **productType** | [**ProductTypes**](ProductTypes.md) |  | 
 **serviceId** | **int32** |  | 
 **countryIsoCode** | **string** |  | 
 **operatorId** | **int32** |  | 
 **statusId** | **int32** |  | 
 **creditPartyMobileNumber** | **string** |  | 
 **creditPartyAccountNumber** | **string** |  | 
 **fromDate** | **time.Time** | Starting date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between &#x60;from_date&#x60; and &#x60;to_date&#x60; can not exceed 24 hours. | 
 **toDate** | **time.Time** | Ending date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between &#x60;from_date&#x60; and &#x60;to_date&#x60; can not exceed 24 hours. | 
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]AsyncTransactionsPostRequest**](AsyncTransactionsPostRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdCancelPost

> Transaction TransactionsTransactionIdCancelPost(ctx, transactionId).Execute()

Cancel a transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.TransactionsTransactionIdCancelPost(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TransactionsTransactionIdCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsTransactionIdCancelPost`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.TransactionsTransactionIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdGet

> AsyncTransactionsPostRequest TransactionsTransactionIdGet(ctx, transactionId).Execute()

Query a transaction by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.TransactionsTransactionIdGet(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TransactionsTransactionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsTransactionIdGet`: AsyncTransactionsPostRequest
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.TransactionsTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncTransactionsPostRequest**](AsyncTransactionsPostRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

