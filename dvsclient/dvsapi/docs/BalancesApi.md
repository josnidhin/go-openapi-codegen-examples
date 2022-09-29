# \BalancesApi

All URIs are relative to *https://preprod-dvs-api.dtone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalances**](BalancesApi.md#GetBalances) | **Get** /balances | Retrieve balances



## GetBalances

> []Balance GetBalances(ctx).UnitType(unitType).Unit(unit).Page(page).PerPage(perPage).Execute()

Retrieve balances

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
    unitType := openapiclient.UnitTypes("CURRENCY") // UnitTypes |  (optional)
    unit := "unit_example" // string |  (optional)
    page := int32(56) // int32 | Page number (optional) (default to 1)
    perPage := int32(56) // int32 | Number of records per page (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancesApi.GetBalances(context.Background()).UnitType(unitType).Unit(unit).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancesApi.GetBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalances`: []Balance
    fmt.Fprintf(os.Stdout, "Response from `BalancesApi.GetBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unitType** | [**UnitTypes**](UnitTypes.md) |  | 
 **unit** | **string** |  | 
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]

### Return type

[**[]Balance**](Balance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

