# \CampaignsApi

All URIs are relative to *http://127.0.0.1:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignById**](CampaignsApi.md#GetCampaignById) | **Get** /campaigns/{campaign_id} | Retrieve campaign by ID
[**GetCampaigns**](CampaignsApi.md#GetCampaigns) | **Get** /campaigns | Retrieve list of active campaigns



## GetCampaignById

> Campaign GetCampaignById(ctx, campaignId).Execute()

Retrieve campaign by ID

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
    campaignId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CampaignsApi.GetCampaignById(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaignById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignById`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> []Campaign GetCampaigns(ctx).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).OperatorId(operatorId).ProductId(productId).Execute()

Retrieve list of active campaigns

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
    resp, r, err := apiClient.CampaignsApi.GetCampaigns(context.Background()).Page(page).PerPage(perPage).CountryIsoCode(countryIsoCode).OperatorId(operatorId).ProductId(productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsApi.GetCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaigns`: []Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignsApi.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Number of records per page | [default to 50]
 **countryIsoCode** | **string** |  | 
 **operatorId** | **int32** |  | 
 **productId** | **int32** |  | 

### Return type

[**[]Campaign**](Campaign.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

