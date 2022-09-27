# \PromotionsApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromotionsGet**](PromotionsApi.md#PromotionsGet) | **Get** /promotions | Retrieve list of promotions
[**PromotionsPromotionIdGet**](PromotionsApi.md#PromotionsPromotionIdGet) | **Get** /promotions/{promotion_id} | Retrieve promotion by ID



## PromotionsGet

> []Promotion PromotionsGet(ctx).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).OperatorId(operatorId).ProductId(productId).Execute()

Retrieve list of promotions

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
    operatorId := int32(56) // int32 |  (optional)
    productId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionsApi.PromotionsGet(context.Background()).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).OperatorId(operatorId).ProductId(productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionsApi.PromotionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromotionsGet`: []Promotion
    fmt.Fprintf(os.Stdout, "Response from `PromotionsApi.PromotionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromotionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]
 **countryIsoCode** | **string** |  | 
 **operatorId** | **int32** |  | 
 **productId** | **int32** |  | 

### Return type

[**[]Promotion**](Promotion.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromotionsPromotionIdGet

> Promotion PromotionsPromotionIdGet(ctx, promotionId).Execute()

Retrieve promotion by ID

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
    promotionId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionsApi.PromotionsPromotionIdGet(context.Background(), promotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionsApi.PromotionsPromotionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromotionsPromotionIdGet`: Promotion
    fmt.Fprintf(os.Stdout, "Response from `PromotionsApi.PromotionsPromotionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromotionsPromotionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Promotion**](Promotion.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

