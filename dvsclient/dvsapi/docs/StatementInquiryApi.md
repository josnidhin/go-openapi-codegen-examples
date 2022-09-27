# \StatementInquiryApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookupStatementInquiryPost**](StatementInquiryApi.md#LookupStatementInquiryPost) | **Post** /lookup/statement-inquiry | Inquire statements for a given account number



## LookupStatementInquiryPost

> []Statement LookupStatementInquiryPost(ctx).LookupStatementInquiryPostRequest(lookupStatementInquiryPostRequest).Execute()

Inquire statements for a given account number

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
    lookupStatementInquiryPostRequest := *openapiclient.NewLookupStatementInquiryPostRequest(int32(123), "AccountNumber_example") // LookupStatementInquiryPostRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementInquiryApi.LookupStatementInquiryPost(context.Background()).LookupStatementInquiryPostRequest(lookupStatementInquiryPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementInquiryApi.LookupStatementInquiryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupStatementInquiryPost`: []Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementInquiryApi.LookupStatementInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupStatementInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupStatementInquiryPostRequest** | [**LookupStatementInquiryPostRequest**](LookupStatementInquiryPostRequest.md) |  | 

### Return type

[**[]Statement**](Statement.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

