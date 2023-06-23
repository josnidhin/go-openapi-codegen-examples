# TransactionAdditionalIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaserId** | Pointer to **string** | Anonymized ID related to the purchase for compliance reasons (e.g. terminal, POS, agent, end user, etc.) | [optional] 

## Methods

### NewTransactionAdditionalIdentifier

`func NewTransactionAdditionalIdentifier() *TransactionAdditionalIdentifier`

NewTransactionAdditionalIdentifier instantiates a new TransactionAdditionalIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAdditionalIdentifierWithDefaults

`func NewTransactionAdditionalIdentifierWithDefaults() *TransactionAdditionalIdentifier`

NewTransactionAdditionalIdentifierWithDefaults instantiates a new TransactionAdditionalIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchaserId

`func (o *TransactionAdditionalIdentifier) GetPurchaserId() string`

GetPurchaserId returns the PurchaserId field if non-nil, zero value otherwise.

### GetPurchaserIdOk

`func (o *TransactionAdditionalIdentifier) GetPurchaserIdOk() (*string, bool)`

GetPurchaserIdOk returns a tuple with the PurchaserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaserId

`func (o *TransactionAdditionalIdentifier) SetPurchaserId(v string)`

SetPurchaserId sets PurchaserId field to given value.

### HasPurchaserId

`func (o *TransactionAdditionalIdentifier) HasPurchaserId() bool`

HasPurchaserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


