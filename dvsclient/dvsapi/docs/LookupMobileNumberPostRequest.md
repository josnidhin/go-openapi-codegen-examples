# LookupMobileNumberPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileNumber** | **string** |  | 
**Page** | Pointer to **int32** | Page number | [optional] [default to 1]
**PerPage** | Pointer to **int32** | Number of records per page | [optional] [default to 50]

## Methods

### NewLookupMobileNumberPostRequest

`func NewLookupMobileNumberPostRequest(mobileNumber string, ) *LookupMobileNumberPostRequest`

NewLookupMobileNumberPostRequest instantiates a new LookupMobileNumberPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupMobileNumberPostRequestWithDefaults

`func NewLookupMobileNumberPostRequestWithDefaults() *LookupMobileNumberPostRequest`

NewLookupMobileNumberPostRequestWithDefaults instantiates a new LookupMobileNumberPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileNumber

`func (o *LookupMobileNumberPostRequest) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *LookupMobileNumberPostRequest) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *LookupMobileNumberPostRequest) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.


### GetPage

`func (o *LookupMobileNumberPostRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LookupMobileNumberPostRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LookupMobileNumberPostRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *LookupMobileNumberPostRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *LookupMobileNumberPostRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *LookupMobileNumberPostRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *LookupMobileNumberPostRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *LookupMobileNumberPostRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


