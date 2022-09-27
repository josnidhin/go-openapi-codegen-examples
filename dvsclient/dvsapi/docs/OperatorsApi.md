# \OperatorsApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperatorsGet**](OperatorsApi.md#OperatorsGet) | **Get** /operators | Retrieve list of operators
[**OperatorsOperatorIdGet**](OperatorsApi.md#OperatorsOperatorIdGet) | **Get** /operators/{operator_id} | Retrieve operator by ID



## OperatorsGet

> []OperatorsGet200ResponseInner OperatorsGet(ctx).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).Execute()

Retrieve list of operators

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
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)
    countryIsoCode := "countryIsoCode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorsApi.OperatorsGet(context.Background()).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorsApi.OperatorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorsGet`: []OperatorsGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OperatorsApi.OperatorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperatorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]
 **countryIsoCode** | **string** |  | 

### Return type

[**[]OperatorsGet200ResponseInner**](OperatorsGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperatorsOperatorIdGet

> OperatorsGet200ResponseInner OperatorsOperatorIdGet(ctx, operatorId).Execute()

Retrieve operator by ID

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
    operatorId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorsApi.OperatorsOperatorIdGet(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorsApi.OperatorsOperatorIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OperatorsOperatorIdGet`: OperatorsGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OperatorsApi.OperatorsOperatorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorsOperatorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperatorsGet200ResponseInner**](OperatorsGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

