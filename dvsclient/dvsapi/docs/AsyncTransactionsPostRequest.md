# AsyncTransactionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | **string** |  | 
**CreationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**ConfirmationExpirationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**ConfirmationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**ProductId** | **int32** |  | 
**CalculationMode** | Pointer to [**CalculationModes**](CalculationModes.md) |  | [optional] 
**Source** | Pointer to [**NullableTransactionSource**](TransactionSource.md) |  | [optional] 
**Destination** | Pointer to [**NullableTransactionDestination**](TransactionDestination.md) |  | [optional] 
**AutoConfirm** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to [**TransactionStatus**](TransactionStatus.md) |  | [optional] 
**OperatorReference** | Pointer to **NullableString** |  | [optional] [readonly] 
**Pin** | Pointer to [**NullableTransactionPin**](TransactionPin.md) |  | [optional] 
**Product** | Pointer to [**TransactionProduct**](TransactionProduct.md) |  | [optional] 
**Prices** | Pointer to [**TransactionPrices**](TransactionPrices.md) |  | [optional] 
**Rates** | Pointer to [**TransactionRates**](TransactionRates.md) |  | [optional] 
**Benefits** | Pointer to [**[]TransactionBenefitsInner**](TransactionBenefitsInner.md) |  | [optional] [readonly] 
**Promotions** | Pointer to [**[]ProductPromotion**](ProductPromotion.md) |  | [optional] [readonly] 
**RequestedValues** | Pointer to [**NullableTransactionRequestedValues**](TransactionRequestedValues.md) |  | [optional] 
**AdjustedValues** | Pointer to [**NullableTransactionRequestedValues**](TransactionRequestedValues.md) |  | [optional] 
**Sender** | Pointer to [**NullableTransactionSender**](TransactionSender.md) |  | [optional] 
**Beneficiary** | Pointer to [**NullableTransactionSender**](TransactionSender.md) |  | [optional] 
**DebitPartyIdentifier** | Pointer to [**NullableTransactionDebitPartyIdentifier**](TransactionDebitPartyIdentifier.md) |  | [optional] 
**CreditPartyIdentifier** | Pointer to [**TransactionCreditPartyIdentifier**](TransactionCreditPartyIdentifier.md) |  | [optional] 
**StatementIdentifier** | Pointer to [**TransactionStatementIdentifier**](TransactionStatementIdentifier.md) |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewAsyncTransactionsPostRequest

`func NewAsyncTransactionsPostRequest(externalId string, productId int32, ) *AsyncTransactionsPostRequest`

NewAsyncTransactionsPostRequest instantiates a new AsyncTransactionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncTransactionsPostRequestWithDefaults

`func NewAsyncTransactionsPostRequestWithDefaults() *AsyncTransactionsPostRequest`

NewAsyncTransactionsPostRequestWithDefaults instantiates a new AsyncTransactionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AsyncTransactionsPostRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AsyncTransactionsPostRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AsyncTransactionsPostRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AsyncTransactionsPostRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *AsyncTransactionsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AsyncTransactionsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AsyncTransactionsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCreationDate

`func (o *AsyncTransactionsPostRequest) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AsyncTransactionsPostRequest) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AsyncTransactionsPostRequest) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AsyncTransactionsPostRequest) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetConfirmationExpirationDate

`func (o *AsyncTransactionsPostRequest) GetConfirmationExpirationDate() time.Time`

GetConfirmationExpirationDate returns the ConfirmationExpirationDate field if non-nil, zero value otherwise.

### GetConfirmationExpirationDateOk

`func (o *AsyncTransactionsPostRequest) GetConfirmationExpirationDateOk() (*time.Time, bool)`

GetConfirmationExpirationDateOk returns a tuple with the ConfirmationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationExpirationDate

`func (o *AsyncTransactionsPostRequest) SetConfirmationExpirationDate(v time.Time)`

SetConfirmationExpirationDate sets ConfirmationExpirationDate field to given value.

### HasConfirmationExpirationDate

`func (o *AsyncTransactionsPostRequest) HasConfirmationExpirationDate() bool`

HasConfirmationExpirationDate returns a boolean if a field has been set.

### GetConfirmationDate

`func (o *AsyncTransactionsPostRequest) GetConfirmationDate() time.Time`

GetConfirmationDate returns the ConfirmationDate field if non-nil, zero value otherwise.

### GetConfirmationDateOk

`func (o *AsyncTransactionsPostRequest) GetConfirmationDateOk() (*time.Time, bool)`

GetConfirmationDateOk returns a tuple with the ConfirmationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationDate

`func (o *AsyncTransactionsPostRequest) SetConfirmationDate(v time.Time)`

SetConfirmationDate sets ConfirmationDate field to given value.

### HasConfirmationDate

`func (o *AsyncTransactionsPostRequest) HasConfirmationDate() bool`

HasConfirmationDate returns a boolean if a field has been set.

### GetProductId

`func (o *AsyncTransactionsPostRequest) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AsyncTransactionsPostRequest) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AsyncTransactionsPostRequest) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetCalculationMode

`func (o *AsyncTransactionsPostRequest) GetCalculationMode() CalculationModes`

GetCalculationMode returns the CalculationMode field if non-nil, zero value otherwise.

### GetCalculationModeOk

`func (o *AsyncTransactionsPostRequest) GetCalculationModeOk() (*CalculationModes, bool)`

GetCalculationModeOk returns a tuple with the CalculationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMode

`func (o *AsyncTransactionsPostRequest) SetCalculationMode(v CalculationModes)`

SetCalculationMode sets CalculationMode field to given value.

### HasCalculationMode

`func (o *AsyncTransactionsPostRequest) HasCalculationMode() bool`

HasCalculationMode returns a boolean if a field has been set.

### GetSource

`func (o *AsyncTransactionsPostRequest) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AsyncTransactionsPostRequest) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AsyncTransactionsPostRequest) SetSource(v TransactionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *AsyncTransactionsPostRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AsyncTransactionsPostRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AsyncTransactionsPostRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *AsyncTransactionsPostRequest) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AsyncTransactionsPostRequest) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AsyncTransactionsPostRequest) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AsyncTransactionsPostRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *AsyncTransactionsPostRequest) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *AsyncTransactionsPostRequest) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetAutoConfirm

`func (o *AsyncTransactionsPostRequest) GetAutoConfirm() bool`

GetAutoConfirm returns the AutoConfirm field if non-nil, zero value otherwise.

### GetAutoConfirmOk

`func (o *AsyncTransactionsPostRequest) GetAutoConfirmOk() (*bool, bool)`

GetAutoConfirmOk returns a tuple with the AutoConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirm

`func (o *AsyncTransactionsPostRequest) SetAutoConfirm(v bool)`

SetAutoConfirm sets AutoConfirm field to given value.

### HasAutoConfirm

`func (o *AsyncTransactionsPostRequest) HasAutoConfirm() bool`

HasAutoConfirm returns a boolean if a field has been set.

### GetStatus

`func (o *AsyncTransactionsPostRequest) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncTransactionsPostRequest) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncTransactionsPostRequest) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AsyncTransactionsPostRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOperatorReference

`func (o *AsyncTransactionsPostRequest) GetOperatorReference() string`

GetOperatorReference returns the OperatorReference field if non-nil, zero value otherwise.

### GetOperatorReferenceOk

`func (o *AsyncTransactionsPostRequest) GetOperatorReferenceOk() (*string, bool)`

GetOperatorReferenceOk returns a tuple with the OperatorReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorReference

`func (o *AsyncTransactionsPostRequest) SetOperatorReference(v string)`

SetOperatorReference sets OperatorReference field to given value.

### HasOperatorReference

`func (o *AsyncTransactionsPostRequest) HasOperatorReference() bool`

HasOperatorReference returns a boolean if a field has been set.

### SetOperatorReferenceNil

`func (o *AsyncTransactionsPostRequest) SetOperatorReferenceNil(b bool)`

 SetOperatorReferenceNil sets the value for OperatorReference to be an explicit nil

### UnsetOperatorReference
`func (o *AsyncTransactionsPostRequest) UnsetOperatorReference()`

UnsetOperatorReference ensures that no value is present for OperatorReference, not even an explicit nil
### GetPin

`func (o *AsyncTransactionsPostRequest) GetPin() TransactionPin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *AsyncTransactionsPostRequest) GetPinOk() (*TransactionPin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *AsyncTransactionsPostRequest) SetPin(v TransactionPin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *AsyncTransactionsPostRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.

### SetPinNil

`func (o *AsyncTransactionsPostRequest) SetPinNil(b bool)`

 SetPinNil sets the value for Pin to be an explicit nil

### UnsetPin
`func (o *AsyncTransactionsPostRequest) UnsetPin()`

UnsetPin ensures that no value is present for Pin, not even an explicit nil
### GetProduct

`func (o *AsyncTransactionsPostRequest) GetProduct() TransactionProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AsyncTransactionsPostRequest) GetProductOk() (*TransactionProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AsyncTransactionsPostRequest) SetProduct(v TransactionProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AsyncTransactionsPostRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPrices

`func (o *AsyncTransactionsPostRequest) GetPrices() TransactionPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *AsyncTransactionsPostRequest) GetPricesOk() (*TransactionPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *AsyncTransactionsPostRequest) SetPrices(v TransactionPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *AsyncTransactionsPostRequest) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetRates

`func (o *AsyncTransactionsPostRequest) GetRates() TransactionRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *AsyncTransactionsPostRequest) GetRatesOk() (*TransactionRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *AsyncTransactionsPostRequest) SetRates(v TransactionRates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *AsyncTransactionsPostRequest) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetBenefits

`func (o *AsyncTransactionsPostRequest) GetBenefits() []TransactionBenefitsInner`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *AsyncTransactionsPostRequest) GetBenefitsOk() (*[]TransactionBenefitsInner, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *AsyncTransactionsPostRequest) SetBenefits(v []TransactionBenefitsInner)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *AsyncTransactionsPostRequest) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *AsyncTransactionsPostRequest) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *AsyncTransactionsPostRequest) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *AsyncTransactionsPostRequest) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *AsyncTransactionsPostRequest) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *AsyncTransactionsPostRequest) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *AsyncTransactionsPostRequest) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### SetPromotionsNil

`func (o *AsyncTransactionsPostRequest) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *AsyncTransactionsPostRequest) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil
### GetRequestedValues

`func (o *AsyncTransactionsPostRequest) GetRequestedValues() TransactionRequestedValues`

GetRequestedValues returns the RequestedValues field if non-nil, zero value otherwise.

### GetRequestedValuesOk

`func (o *AsyncTransactionsPostRequest) GetRequestedValuesOk() (*TransactionRequestedValues, bool)`

GetRequestedValuesOk returns a tuple with the RequestedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedValues

`func (o *AsyncTransactionsPostRequest) SetRequestedValues(v TransactionRequestedValues)`

SetRequestedValues sets RequestedValues field to given value.

### HasRequestedValues

`func (o *AsyncTransactionsPostRequest) HasRequestedValues() bool`

HasRequestedValues returns a boolean if a field has been set.

### SetRequestedValuesNil

`func (o *AsyncTransactionsPostRequest) SetRequestedValuesNil(b bool)`

 SetRequestedValuesNil sets the value for RequestedValues to be an explicit nil

### UnsetRequestedValues
`func (o *AsyncTransactionsPostRequest) UnsetRequestedValues()`

UnsetRequestedValues ensures that no value is present for RequestedValues, not even an explicit nil
### GetAdjustedValues

`func (o *AsyncTransactionsPostRequest) GetAdjustedValues() TransactionRequestedValues`

GetAdjustedValues returns the AdjustedValues field if non-nil, zero value otherwise.

### GetAdjustedValuesOk

`func (o *AsyncTransactionsPostRequest) GetAdjustedValuesOk() (*TransactionRequestedValues, bool)`

GetAdjustedValuesOk returns a tuple with the AdjustedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedValues

`func (o *AsyncTransactionsPostRequest) SetAdjustedValues(v TransactionRequestedValues)`

SetAdjustedValues sets AdjustedValues field to given value.

### HasAdjustedValues

`func (o *AsyncTransactionsPostRequest) HasAdjustedValues() bool`

HasAdjustedValues returns a boolean if a field has been set.

### SetAdjustedValuesNil

`func (o *AsyncTransactionsPostRequest) SetAdjustedValuesNil(b bool)`

 SetAdjustedValuesNil sets the value for AdjustedValues to be an explicit nil

### UnsetAdjustedValues
`func (o *AsyncTransactionsPostRequest) UnsetAdjustedValues()`

UnsetAdjustedValues ensures that no value is present for AdjustedValues, not even an explicit nil
### GetSender

`func (o *AsyncTransactionsPostRequest) GetSender() TransactionSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AsyncTransactionsPostRequest) GetSenderOk() (*TransactionSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AsyncTransactionsPostRequest) SetSender(v TransactionSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AsyncTransactionsPostRequest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *AsyncTransactionsPostRequest) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *AsyncTransactionsPostRequest) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetBeneficiary

`func (o *AsyncTransactionsPostRequest) GetBeneficiary() TransactionSender`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *AsyncTransactionsPostRequest) GetBeneficiaryOk() (*TransactionSender, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *AsyncTransactionsPostRequest) SetBeneficiary(v TransactionSender)`

SetBeneficiary sets Beneficiary field to given value.

### HasBeneficiary

`func (o *AsyncTransactionsPostRequest) HasBeneficiary() bool`

HasBeneficiary returns a boolean if a field has been set.

### SetBeneficiaryNil

`func (o *AsyncTransactionsPostRequest) SetBeneficiaryNil(b bool)`

 SetBeneficiaryNil sets the value for Beneficiary to be an explicit nil

### UnsetBeneficiary
`func (o *AsyncTransactionsPostRequest) UnsetBeneficiary()`

UnsetBeneficiary ensures that no value is present for Beneficiary, not even an explicit nil
### GetDebitPartyIdentifier

`func (o *AsyncTransactionsPostRequest) GetDebitPartyIdentifier() TransactionDebitPartyIdentifier`

GetDebitPartyIdentifier returns the DebitPartyIdentifier field if non-nil, zero value otherwise.

### GetDebitPartyIdentifierOk

`func (o *AsyncTransactionsPostRequest) GetDebitPartyIdentifierOk() (*TransactionDebitPartyIdentifier, bool)`

GetDebitPartyIdentifierOk returns a tuple with the DebitPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitPartyIdentifier

`func (o *AsyncTransactionsPostRequest) SetDebitPartyIdentifier(v TransactionDebitPartyIdentifier)`

SetDebitPartyIdentifier sets DebitPartyIdentifier field to given value.

### HasDebitPartyIdentifier

`func (o *AsyncTransactionsPostRequest) HasDebitPartyIdentifier() bool`

HasDebitPartyIdentifier returns a boolean if a field has been set.

### SetDebitPartyIdentifierNil

`func (o *AsyncTransactionsPostRequest) SetDebitPartyIdentifierNil(b bool)`

 SetDebitPartyIdentifierNil sets the value for DebitPartyIdentifier to be an explicit nil

### UnsetDebitPartyIdentifier
`func (o *AsyncTransactionsPostRequest) UnsetDebitPartyIdentifier()`

UnsetDebitPartyIdentifier ensures that no value is present for DebitPartyIdentifier, not even an explicit nil
### GetCreditPartyIdentifier

`func (o *AsyncTransactionsPostRequest) GetCreditPartyIdentifier() TransactionCreditPartyIdentifier`

GetCreditPartyIdentifier returns the CreditPartyIdentifier field if non-nil, zero value otherwise.

### GetCreditPartyIdentifierOk

`func (o *AsyncTransactionsPostRequest) GetCreditPartyIdentifierOk() (*TransactionCreditPartyIdentifier, bool)`

GetCreditPartyIdentifierOk returns a tuple with the CreditPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPartyIdentifier

`func (o *AsyncTransactionsPostRequest) SetCreditPartyIdentifier(v TransactionCreditPartyIdentifier)`

SetCreditPartyIdentifier sets CreditPartyIdentifier field to given value.

### HasCreditPartyIdentifier

`func (o *AsyncTransactionsPostRequest) HasCreditPartyIdentifier() bool`

HasCreditPartyIdentifier returns a boolean if a field has been set.

### GetStatementIdentifier

`func (o *AsyncTransactionsPostRequest) GetStatementIdentifier() TransactionStatementIdentifier`

GetStatementIdentifier returns the StatementIdentifier field if non-nil, zero value otherwise.

### GetStatementIdentifierOk

`func (o *AsyncTransactionsPostRequest) GetStatementIdentifierOk() (*TransactionStatementIdentifier, bool)`

GetStatementIdentifierOk returns a tuple with the StatementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIdentifier

`func (o *AsyncTransactionsPostRequest) SetStatementIdentifier(v TransactionStatementIdentifier)`

SetStatementIdentifier sets StatementIdentifier field to given value.

### HasStatementIdentifier

`func (o *AsyncTransactionsPostRequest) HasStatementIdentifier() bool`

HasStatementIdentifier returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *AsyncTransactionsPostRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *AsyncTransactionsPostRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *AsyncTransactionsPostRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *AsyncTransactionsPostRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


