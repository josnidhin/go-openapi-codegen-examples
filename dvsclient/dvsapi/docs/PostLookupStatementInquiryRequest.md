# PostLookupStatementInquiryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **int32** |  | 
**AccountNumber** | **string** |  | 
**AccountQualifier** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** | Page number | [optional] [default to 1]
**PerPage** | Pointer to **int32** | Number of records per page | [optional] [default to 50]

## Methods

### NewPostLookupStatementInquiryRequest

`func NewPostLookupStatementInquiryRequest(productId int32, accountNumber string, ) *PostLookupStatementInquiryRequest`

NewPostLookupStatementInquiryRequest instantiates a new PostLookupStatementInquiryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLookupStatementInquiryRequestWithDefaults

`func NewPostLookupStatementInquiryRequestWithDefaults() *PostLookupStatementInquiryRequest`

NewPostLookupStatementInquiryRequestWithDefaults instantiates a new PostLookupStatementInquiryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *PostLookupStatementInquiryRequest) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PostLookupStatementInquiryRequest) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PostLookupStatementInquiryRequest) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetAccountNumber

`func (o *PostLookupStatementInquiryRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PostLookupStatementInquiryRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PostLookupStatementInquiryRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountQualifier

`func (o *PostLookupStatementInquiryRequest) GetAccountQualifier() string`

GetAccountQualifier returns the AccountQualifier field if non-nil, zero value otherwise.

### GetAccountQualifierOk

`func (o *PostLookupStatementInquiryRequest) GetAccountQualifierOk() (*string, bool)`

GetAccountQualifierOk returns a tuple with the AccountQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQualifier

`func (o *PostLookupStatementInquiryRequest) SetAccountQualifier(v string)`

SetAccountQualifier sets AccountQualifier field to given value.

### HasAccountQualifier

`func (o *PostLookupStatementInquiryRequest) HasAccountQualifier() bool`

HasAccountQualifier returns a boolean if a field has been set.

### GetPage

`func (o *PostLookupStatementInquiryRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PostLookupStatementInquiryRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PostLookupStatementInquiryRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PostLookupStatementInquiryRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *PostLookupStatementInquiryRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *PostLookupStatementInquiryRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *PostLookupStatementInquiryRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *PostLookupStatementInquiryRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


