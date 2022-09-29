# \OperatorsApi

All URIs are relative to *http://127.0.0.1:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOperatorById**](OperatorsApi.md#GetOperatorById) | **Get** /operators/{operator_id} | Retrieve operator by ID
[**GetOperators**](OperatorsApi.md#GetOperators) | **Get** /operators | Retrieve list of operators



## GetOperatorById

> Operator GetOperatorById(ctx, operatorId).Execute()

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
    resp, r, err := apiClient.OperatorsApi.GetOperatorById(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorsApi.GetOperatorById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperatorById`: Operator
    fmt.Fprintf(os.Stdout, "Response from `OperatorsApi.GetOperatorById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Operator**](Operator.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperators

> []Operator GetOperators(ctx).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).Execute()

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
    resp, r, err := apiClient.OperatorsApi.GetOperators(context.Background()).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorsApi.GetOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperators`: []Operator
    fmt.Fprintf(os.Stdout, "Response from `OperatorsApi.GetOperators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]
 **countryIsoCode** | **string** |  | 

### Return type

[**[]Operator**](Operator.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

