# \StatementInquiryApi

All URIs are relative to *https://preprod-dvs-api.dtone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLookupStatementInquiry**](StatementInquiryApi.md#PostLookupStatementInquiry) | **Post** /lookup/statement-inquiry | Inquire statements for a given account number



## PostLookupStatementInquiry

> []Statement PostLookupStatementInquiry(ctx).PostLookupStatementInquiryRequest(postLookupStatementInquiryRequest).Execute()

Inquire statements for a given account number

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
    postLookupStatementInquiryRequest := *openapiclient.NewPostLookupStatementInquiryRequest(int32(123), "AccountNumber_example") // PostLookupStatementInquiryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatementInquiryApi.PostLookupStatementInquiry(context.Background()).PostLookupStatementInquiryRequest(postLookupStatementInquiryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatementInquiryApi.PostLookupStatementInquiry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLookupStatementInquiry`: []Statement
    fmt.Fprintf(os.Stdout, "Response from `StatementInquiryApi.PostLookupStatementInquiry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLookupStatementInquiryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postLookupStatementInquiryRequest** | [**PostLookupStatementInquiryRequest**](PostLookupStatementInquiryRequest.md) |  | 

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

