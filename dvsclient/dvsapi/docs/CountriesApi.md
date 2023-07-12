# \CountriesApi

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](CountriesApi.md#GetCountries) | **Get** /countries | Retrieve list of countries
[**GetCountryByIsoCode**](CountriesApi.md#GetCountryByIsoCode) | **Get** /countries/{country_iso_code} | Retrieve country by ISO code



## GetCountries

> []Country GetCountries(ctx).Page(page).PerPage(perPage).Execute()

Retrieve list of countries

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
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountriesApi.GetCountries(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountriesApi.GetCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountries`: []Country
    fmt.Fprintf(os.Stdout, "Response from `CountriesApi.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]Country**](Country.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByIsoCode

> Country GetCountryByIsoCode(ctx, countryIsoCode).Execute()

Retrieve country by ISO code

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
    countryIsoCode := "countryIsoCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountriesApi.GetCountryByIsoCode(context.Background(), countryIsoCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountriesApi.GetCountryByIsoCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountryByIsoCode`: Country
    fmt.Fprintf(os.Stdout, "Response from `CountriesApi.GetCountryByIsoCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryIsoCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByIsoCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Country**](Country.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

