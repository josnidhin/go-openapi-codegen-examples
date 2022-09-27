# LookupStatementInquiryPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int32** |  | 
**AccountNumber** | **string** |  | 
**AccountQualifier** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** | Page number | [optional] [default to 1]
**PerPage** | Pointer to **int32** | Number of records per page | [optional] [default to 50]

## Methods

### NewLookupStatementInquiryPostRequest

`func NewLookupStatementInquiryPostRequest(productId int32, accountNumber string, ) *LookupStatementInquiryPostRequest`

NewLookupStatementInquiryPostRequest instantiates a new LookupStatementInquiryPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupStatementInquiryPostRequestWithDefaults

`func NewLookupStatementInquiryPostRequestWithDefaults() *LookupStatementInquiryPostRequest`

NewLookupStatementInquiryPostRequestWithDefaults instantiates a new LookupStatementInquiryPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *LookupStatementInquiryPostRequest) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *LookupStatementInquiryPostRequest) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *LookupStatementInquiryPostRequest) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetAccountNumber

`func (o *LookupStatementInquiryPostRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *LookupStatementInquiryPostRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *LookupStatementInquiryPostRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountQualifier

`func (o *LookupStatementInquiryPostRequest) GetAccountQualifier() string`

GetAccountQualifier returns the AccountQualifier field if non-nil, zero value otherwise.

### GetAccountQualifierOk

`func (o *LookupStatementInquiryPostRequest) GetAccountQualifierOk() (*string, bool)`

GetAccountQualifierOk returns a tuple with the AccountQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQualifier

`func (o *LookupStatementInquiryPostRequest) SetAccountQualifier(v string)`

SetAccountQualifier sets AccountQualifier field to given value.

### HasAccountQualifier

`func (o *LookupStatementInquiryPostRequest) HasAccountQualifier() bool`

HasAccountQualifier returns a boolean if a field has been set.

### GetPage

`func (o *LookupStatementInquiryPostRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LookupStatementInquiryPostRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LookupStatementInquiryPostRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *LookupStatementInquiryPostRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *LookupStatementInquiryPostRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *LookupStatementInquiryPostRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *LookupStatementInquiryPostRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *LookupStatementInquiryPostRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


