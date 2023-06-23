# \TransactionsApi

All URIs are relative to *https://preprod-dvs-api.dtone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTransactionById**](TransactionsApi.md#GetTransactionById) | **Get** /transactions/{transaction_id} | Query a transaction by ID
[**GetTransactions**](TransactionsApi.md#GetTransactions) | **Get** /transactions | Query list of transactions
[**PostTransactionAsync**](TransactionsApi.md#PostTransactionAsync) | **Post** /async/transactions | Create a transaction asynchronously
[**PostTransactionCancel**](TransactionsApi.md#PostTransactionCancel) | **Post** /transactions/{transaction_id}/cancel | Cancel a transaction
[**PostTransactionConfirmAsync**](TransactionsApi.md#PostTransactionConfirmAsync) | **Post** /async/transactions/{transaction_id}/confirm | Confirm a transaction asynchronously
[**PostTransactionConfirmSync**](TransactionsApi.md#PostTransactionConfirmSync) | **Post** /sync/transactions/{transaction_id}/confirm | Confirm a transaction synchronously
[**PostTransactionSync**](TransactionsApi.md#PostTransactionSync) | **Post** /sync/transactions | Create a transaction synchronously



## GetTransactionById

> PostTransactionAsyncRequest GetTransactionById(ctx, transactionId).Execute()

Query a transaction by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.GetTransactionById(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionById`: PostTransactionAsyncRequest
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostTransactionAsyncRequest**](PostTransactionAsyncRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> []PostTransactionAsyncRequest GetTransactions(ctx).ExternalId(externalId).ProductType(productType).ServiceId(serviceId).SubserviceId(subserviceId).CountryIsoCode(countryIsoCode).OperatorId(operatorId).StatusId(statusId).CreditPartyMobileNumber(creditPartyMobileNumber).CreditPartyAccountNumber(creditPartyAccountNumber).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Execute()

Query list of transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    externalId := "externalId_example" // string |  (optional)
    productType := openapiclient.ProductTypes("FIXED_VALUE_RECHARGE") // ProductTypes |  (optional)
    serviceId := int32(56) // int32 |  (optional)
    subserviceId := int32(56) // int32 |  (optional)
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
    resp, r, err := apiClient.TransactionsApi.GetTransactions(context.Background()).ExternalId(externalId).ProductType(productType).ServiceId(serviceId).SubserviceId(subserviceId).CountryIsoCode(countryIsoCode).OperatorId(operatorId).StatusId(statusId).CreditPartyMobileNumber(creditPartyMobileNumber).CreditPartyAccountNumber(creditPartyAccountNumber).FromDate(fromDate).ToDate(toDate).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactions`: []PostTransactionAsyncRequest
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalId** | **string** |  | 
 **productType** | [**ProductTypes**](ProductTypes.md) |  | 
 **serviceId** | **int32** |  | 
 **subserviceId** | **int32** |  | 
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

[**[]PostTransactionAsyncRequest**](PostTransactionAsyncRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransactionAsync

> Transaction PostTransactionAsync(ctx).PostTransactionAsyncRequest(postTransactionAsyncRequest).Execute()

Create a transaction asynchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    postTransactionAsyncRequest := *openapiclient.NewPostTransactionAsyncRequest(string(123), int32(123)) // PostTransactionAsyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionAsync(context.Background()).PostTransactionAsyncRequest(postTransactionAsyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionAsync`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTransactionAsyncRequest** | [**PostTransactionAsyncRequest**](PostTransactionAsyncRequest.md) |  | 

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


## PostTransactionCancel

> Transaction PostTransactionCancel(ctx, transactionId).Execute()

Cancel a transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionCancel(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionCancel`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionCancelRequest struct via the builder pattern


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


## PostTransactionConfirmAsync

> Transaction PostTransactionConfirmAsync(ctx, transactionId).Execute()

Confirm a transaction asynchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionConfirmAsync(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionConfirmAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionConfirmAsync`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionConfirmAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionConfirmAsyncRequest struct via the builder pattern


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


## PostTransactionConfirmSync

> Transaction PostTransactionConfirmSync(ctx, transactionId).Execute()

Confirm a transaction synchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    transactionId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionConfirmSync(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionConfirmSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionConfirmSync`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionConfirmSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionConfirmSyncRequest struct via the builder pattern


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


## PostTransactionSync

> Transaction PostTransactionSync(ctx).Transaction(transaction).Execute()

Create a transaction synchronously



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
)

func main() {
    transaction := *openapiclient.NewTransaction(string(123), int32(123)) // Transaction |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsApi.PostTransactionSync(context.Background()).Transaction(transaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.PostTransactionSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTransactionSync`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.PostTransactionSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionSyncRequest struct via the builder pattern


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

