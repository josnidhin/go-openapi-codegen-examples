# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | **string** |  | 
**Dates** | [**StatementDates**](StatementDates.md) |  | 
**Balance** | [**StatementBalance**](StatementBalance.md) |  | 

## Methods

### NewStatement

`func NewStatement(reference string, dates StatementDates, balance StatementBalance, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *Statement) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Statement) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Statement) SetReference(v string)`

SetReference sets Reference field to given value.


### GetDates

`func (o *Statement) GetDates() StatementDates`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *Statement) GetDatesOk() (*StatementDates, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *Statement) SetDates(v StatementDates)`

SetDates sets Dates field to given value.


### GetBalance

`func (o *Statement) GetBalance() StatementBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Statement) GetBalanceOk() (*StatementBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Statement) SetBalance(v StatementBalance)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


