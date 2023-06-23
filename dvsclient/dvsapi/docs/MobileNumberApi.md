# \MobileNumberApi

All URIs are relative to *https://preprod-dvs-api.dtone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLookupMobileNumber**](MobileNumberApi.md#GetLookupMobileNumber) | **Get** /lookup/mobile-number/{mobile_number} | Look up operators for a given mobile number
[**PostLookupMobileNumber**](MobileNumberApi.md#PostLookupMobileNumber) | **Post** /lookup/mobile-number | Look up operators for a given mobile number



## GetLookupMobileNumber

> []GetLookupMobileNumber200ResponseInner GetLookupMobileNumber(ctx, mobileNumber).Page(page).PerPage(perPage).Execute()

Look up operators for a given mobile number

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
    mobileNumber := "mobileNumber_example" // string | 
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileNumberApi.GetLookupMobileNumber(context.Background(), mobileNumber).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileNumberApi.GetLookupMobileNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLookupMobileNumber`: []GetLookupMobileNumber200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MobileNumberApi.GetLookupMobileNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLookupMobileNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]GetLookupMobileNumber200ResponseInner**](GetLookupMobileNumber200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLookupMobileNumber

> []GetLookupMobileNumber200ResponseInner PostLookupMobileNumber(ctx).PostLookupMobileNumberRequest(postLookupMobileNumberRequest).Execute()

Look up operators for a given mobile number

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
    postLookupMobileNumberRequest := *openapiclient.NewPostLookupMobileNumberRequest("MobileNumber_example") // PostLookupMobileNumberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileNumberApi.PostLookupMobileNumber(context.Background()).PostLookupMobileNumberRequest(postLookupMobileNumberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileNumberApi.PostLookupMobileNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLookupMobileNumber`: []GetLookupMobileNumber200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MobileNumberApi.PostLookupMobileNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLookupMobileNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLookupMobileNumberRequest** | [**PostLookupMobileNumberRequest**](PostLookupMobileNumberRequest.md) |  | 

### Return type

[**[]GetLookupMobileNumber200ResponseInner**](GetLookupMobileNumber200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

