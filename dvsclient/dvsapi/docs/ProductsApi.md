# \ProductsApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsGet**](ProductsApi.md#ProductsGet) | **Get** /products | Retrieve list of products
[**ProductsProductIdGet**](ProductsApi.md#ProductsProductIdGet) | **Get** /products/{product_id} | Retrieve product by ID



## ProductsGet

> []Product ProductsGet(ctx).Type_(type_).ServiceId(serviceId).Tags(tags).CountryIsoCode(countryIsoCode).OperatorId(operatorId).Region(region).BenefitTypes(benefitTypes).Page(page).PerPage(perPage).Execute()

Retrieve list of products

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
    type_ := openapiclient.ProductTypes("FIXED_VALUE_RECHARGE") // ProductTypes |  (optional)
    serviceId := int32(56) // int32 |  (optional)
    tags := []string{"Inner_example"} // []string |  (optional)
    countryIsoCode := "countryIsoCode_example" // string |  (optional)
    operatorId := int32(56) // int32 |  (optional)
    region := "region_example" // string |  (optional)
    benefitTypes := []openapiclient.BenefitTypes{openapiclient.BenefitTypes("TALKTIME")} // []BenefitTypes |  (optional)
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.ProductsGet(context.Background()).Type_(type_).ServiceId(serviceId).Tags(tags).CountryIsoCode(countryIsoCode).OperatorId(operatorId).Region(region).BenefitTypes(benefitTypes).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsGet`: []Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ProductTypes**](ProductTypes.md) |  | 
 **serviceId** | **int32** |  | 
 **tags** | **[]string** |  | 
 **countryIsoCode** | **string** |  | 
 **operatorId** | **int32** |  | 
 **region** | **string** |  | 
 **benefitTypes** | [**[]BenefitTypes**](BenefitTypes.md) |  | 
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]Product**](Product.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsProductIdGet

> Product ProductsProductIdGet(ctx, productId).Execute()

Retrieve product by ID

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
    productId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.ProductsProductIdGet(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductsProductIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsProductIdGet`: Product
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductsProductIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsProductIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Product**](Product.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

