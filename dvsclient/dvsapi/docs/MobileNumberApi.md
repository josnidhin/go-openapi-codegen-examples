# \MobileNumberApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupMobileNumberMobileNumberGet**](MobileNumberApi.md#LookupMobileNumberMobileNumberGet) | **Get** /lookup/mobile-number/{mobile_number} | Look up operators for a given mobile number
[**LookupMobileNumberPost**](MobileNumberApi.md#LookupMobileNumberPost) | **Post** /lookup/mobile-number | Look up operators for a given mobile number



## LookupMobileNumberMobileNumberGet

> []LookupMobileNumberMobileNumberGet200ResponseInner LookupMobileNumberMobileNumberGet(ctx, mobileNumber).Page(page).PerPage(perPage).Execute()

Look up operators for a given mobile number

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
    mobileNumber := "mobileNumber_example" // string | 
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileNumberApi.LookupMobileNumberMobileNumberGet(context.Background(), mobileNumber).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileNumberApi.LookupMobileNumberMobileNumberGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupMobileNumberMobileNumberGet`: []LookupMobileNumberMobileNumberGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MobileNumberApi.LookupMobileNumberMobileNumberGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobileNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupMobileNumberMobileNumberGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]LookupMobileNumberMobileNumberGet200ResponseInner**](LookupMobileNumberMobileNumberGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupMobileNumberPost

> []LookupMobileNumberMobileNumberGet200ResponseInner LookupMobileNumberPost(ctx).LookupMobileNumberPostRequest(lookupMobileNumberPostRequest).Execute()

Look up operators for a given mobile number

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
    lookupMobileNumberPostRequest := *openapiclient.NewLookupMobileNumberPostRequest("MobileNumber_example") // LookupMobileNumberPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MobileNumberApi.LookupMobileNumberPost(context.Background()).LookupMobileNumberPostRequest(lookupMobileNumberPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileNumberApi.LookupMobileNumberPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupMobileNumberPost`: []LookupMobileNumberMobileNumberGet200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MobileNumberApi.LookupMobileNumberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupMobileNumberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupMobileNumberPostRequest** | [**LookupMobileNumberPostRequest**](LookupMobileNumberPostRequest.md) |  | 

### Return type

[**[]LookupMobileNumberMobileNumberGet200ResponseInner**](LookupMobileNumberMobileNumberGet200ResponseInner.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

