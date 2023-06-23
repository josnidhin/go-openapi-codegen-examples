# Transaction

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
**AdditionalIdentifier** | Pointer to [**TransactionAdditionalIdentifier**](TransactionAdditionalIdentifier.md) |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction(externalId string, productId int32, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *Transaction) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Transaction) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Transaction) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCreationDate

`func (o *Transaction) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Transaction) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Transaction) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Transaction) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetConfirmationExpirationDate

`func (o *Transaction) GetConfirmationExpirationDate() time.Time`

GetConfirmationExpirationDate returns the ConfirmationExpirationDate field if non-nil, zero value otherwise.

### GetConfirmationExpirationDateOk

`func (o *Transaction) GetConfirmationExpirationDateOk() (*time.Time, bool)`

GetConfirmationExpirationDateOk returns a tuple with the ConfirmationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationExpirationDate

`func (o *Transaction) SetConfirmationExpirationDate(v time.Time)`

SetConfirmationExpirationDate sets ConfirmationExpirationDate field to given value.

### HasConfirmationExpirationDate

`func (o *Transaction) HasConfirmationExpirationDate() bool`

HasConfirmationExpirationDate returns a boolean if a field has been set.

### GetConfirmationDate

`func (o *Transaction) GetConfirmationDate() time.Time`

GetConfirmationDate returns the ConfirmationDate field if non-nil, zero value otherwise.

### GetConfirmationDateOk

`func (o *Transaction) GetConfirmationDateOk() (*time.Time, bool)`

GetConfirmationDateOk returns a tuple with the ConfirmationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationDate

`func (o *Transaction) SetConfirmationDate(v time.Time)`

SetConfirmationDate sets ConfirmationDate field to given value.

### HasConfirmationDate

`func (o *Transaction) HasConfirmationDate() bool`

HasConfirmationDate returns a boolean if a field has been set.

### GetProductId

`func (o *Transaction) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Transaction) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Transaction) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetCalculationMode

`func (o *Transaction) GetCalculationMode() CalculationModes`

GetCalculationMode returns the CalculationMode field if non-nil, zero value otherwise.

### GetCalculationModeOk

`func (o *Transaction) GetCalculationModeOk() (*CalculationModes, bool)`

GetCalculationModeOk returns a tuple with the CalculationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMode

`func (o *Transaction) SetCalculationMode(v CalculationModes)`

SetCalculationMode sets CalculationMode field to given value.

### HasCalculationMode

`func (o *Transaction) HasCalculationMode() bool`

HasCalculationMode returns a boolean if a field has been set.

### GetSource

`func (o *Transaction) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Transaction) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Transaction) SetSource(v TransactionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Transaction) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *Transaction) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *Transaction) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *Transaction) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Transaction) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Transaction) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Transaction) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *Transaction) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *Transaction) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetAutoConfirm

`func (o *Transaction) GetAutoConfirm() bool`

GetAutoConfirm returns the AutoConfirm field if non-nil, zero value otherwise.

### GetAutoConfirmOk

`func (o *Transaction) GetAutoConfirmOk() (*bool, bool)`

GetAutoConfirmOk returns a tuple with the AutoConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirm

`func (o *Transaction) SetAutoConfirm(v bool)`

SetAutoConfirm sets AutoConfirm field to given value.

### HasAutoConfirm

`func (o *Transaction) HasAutoConfirm() bool`

HasAutoConfirm returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOperatorReference

`func (o *Transaction) GetOperatorReference() string`

GetOperatorReference returns the OperatorReference field if non-nil, zero value otherwise.

### GetOperatorReferenceOk

`func (o *Transaction) GetOperatorReferenceOk() (*string, bool)`

GetOperatorReferenceOk returns a tuple with the OperatorReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorReference

`func (o *Transaction) SetOperatorReference(v string)`

SetOperatorReference sets OperatorReference field to given value.

### HasOperatorReference

`func (o *Transaction) HasOperatorReference() bool`

HasOperatorReference returns a boolean if a field has been set.

### SetOperatorReferenceNil

`func (o *Transaction) SetOperatorReferenceNil(b bool)`

 SetOperatorReferenceNil sets the value for OperatorReference to be an explicit nil

### UnsetOperatorReference
`func (o *Transaction) UnsetOperatorReference()`

UnsetOperatorReference ensures that no value is present for OperatorReference, not even an explicit nil
### GetPin

`func (o *Transaction) GetPin() TransactionPin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *Transaction) GetPinOk() (*TransactionPin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *Transaction) SetPin(v TransactionPin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *Transaction) HasPin() bool`

HasPin returns a boolean if a field has been set.

### SetPinNil

`func (o *Transaction) SetPinNil(b bool)`

 SetPinNil sets the value for Pin to be an explicit nil

### UnsetPin
`func (o *Transaction) UnsetPin()`

UnsetPin ensures that no value is present for Pin, not even an explicit nil
### GetProduct

`func (o *Transaction) GetProduct() TransactionProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Transaction) GetProductOk() (*TransactionProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Transaction) SetProduct(v TransactionProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Transaction) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPrices

`func (o *Transaction) GetPrices() TransactionPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *Transaction) GetPricesOk() (*TransactionPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *Transaction) SetPrices(v TransactionPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *Transaction) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetRates

`func (o *Transaction) GetRates() TransactionRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *Transaction) GetRatesOk() (*TransactionRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *Transaction) SetRates(v TransactionRates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *Transaction) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetBenefits

`func (o *Transaction) GetBenefits() []TransactionBenefitsInner`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *Transaction) GetBenefitsOk() (*[]TransactionBenefitsInner, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *Transaction) SetBenefits(v []TransactionBenefitsInner)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *Transaction) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *Transaction) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *Transaction) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *Transaction) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *Transaction) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *Transaction) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *Transaction) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### SetPromotionsNil

`func (o *Transaction) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *Transaction) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil
### GetRequestedValues

`func (o *Transaction) GetRequestedValues() TransactionRequestedValues`

GetRequestedValues returns the RequestedValues field if non-nil, zero value otherwise.

### GetRequestedValuesOk

`func (o *Transaction) GetRequestedValuesOk() (*TransactionRequestedValues, bool)`

GetRequestedValuesOk returns a tuple with the RequestedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedValues

`func (o *Transaction) SetRequestedValues(v TransactionRequestedValues)`

SetRequestedValues sets RequestedValues field to given value.

### HasRequestedValues

`func (o *Transaction) HasRequestedValues() bool`

HasRequestedValues returns a boolean if a field has been set.

### SetRequestedValuesNil

`func (o *Transaction) SetRequestedValuesNil(b bool)`

 SetRequestedValuesNil sets the value for RequestedValues to be an explicit nil

### UnsetRequestedValues
`func (o *Transaction) UnsetRequestedValues()`

UnsetRequestedValues ensures that no value is present for RequestedValues, not even an explicit nil
### GetAdjustedValues

`func (o *Transaction) GetAdjustedValues() TransactionRequestedValues`

GetAdjustedValues returns the AdjustedValues field if non-nil, zero value otherwise.

### GetAdjustedValuesOk

`func (o *Transaction) GetAdjustedValuesOk() (*TransactionRequestedValues, bool)`

GetAdjustedValuesOk returns a tuple with the AdjustedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedValues

`func (o *Transaction) SetAdjustedValues(v TransactionRequestedValues)`

SetAdjustedValues sets AdjustedValues field to given value.

### HasAdjustedValues

`func (o *Transaction) HasAdjustedValues() bool`

HasAdjustedValues returns a boolean if a field has been set.

### SetAdjustedValuesNil

`func (o *Transaction) SetAdjustedValuesNil(b bool)`

 SetAdjustedValuesNil sets the value for AdjustedValues to be an explicit nil

### UnsetAdjustedValues
`func (o *Transaction) UnsetAdjustedValues()`

UnsetAdjustedValues ensures that no value is present for AdjustedValues, not even an explicit nil
### GetSender

`func (o *Transaction) GetSender() TransactionSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Transaction) GetSenderOk() (*TransactionSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Transaction) SetSender(v TransactionSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Transaction) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *Transaction) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *Transaction) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetBeneficiary

`func (o *Transaction) GetBeneficiary() TransactionSender`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *Transaction) GetBeneficiaryOk() (*TransactionSender, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *Transaction) SetBeneficiary(v TransactionSender)`

SetBeneficiary sets Beneficiary field to given value.

### HasBeneficiary

`func (o *Transaction) HasBeneficiary() bool`

HasBeneficiary returns a boolean if a field has been set.

### SetBeneficiaryNil

`func (o *Transaction) SetBeneficiaryNil(b bool)`

 SetBeneficiaryNil sets the value for Beneficiary to be an explicit nil

### UnsetBeneficiary
`func (o *Transaction) UnsetBeneficiary()`

UnsetBeneficiary ensures that no value is present for Beneficiary, not even an explicit nil
### GetDebitPartyIdentifier

`func (o *Transaction) GetDebitPartyIdentifier() TransactionDebitPartyIdentifier`

GetDebitPartyIdentifier returns the DebitPartyIdentifier field if non-nil, zero value otherwise.

### GetDebitPartyIdentifierOk

`func (o *Transaction) GetDebitPartyIdentifierOk() (*TransactionDebitPartyIdentifier, bool)`

GetDebitPartyIdentifierOk returns a tuple with the DebitPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitPartyIdentifier

`func (o *Transaction) SetDebitPartyIdentifier(v TransactionDebitPartyIdentifier)`

SetDebitPartyIdentifier sets DebitPartyIdentifier field to given value.

### HasDebitPartyIdentifier

`func (o *Transaction) HasDebitPartyIdentifier() bool`

HasDebitPartyIdentifier returns a boolean if a field has been set.

### SetDebitPartyIdentifierNil

`func (o *Transaction) SetDebitPartyIdentifierNil(b bool)`

 SetDebitPartyIdentifierNil sets the value for DebitPartyIdentifier to be an explicit nil

### UnsetDebitPartyIdentifier
`func (o *Transaction) UnsetDebitPartyIdentifier()`

UnsetDebitPartyIdentifier ensures that no value is present for DebitPartyIdentifier, not even an explicit nil
### GetCreditPartyIdentifier

`func (o *Transaction) GetCreditPartyIdentifier() TransactionCreditPartyIdentifier`

GetCreditPartyIdentifier returns the CreditPartyIdentifier field if non-nil, zero value otherwise.

### GetCreditPartyIdentifierOk

`func (o *Transaction) GetCreditPartyIdentifierOk() (*TransactionCreditPartyIdentifier, bool)`

GetCreditPartyIdentifierOk returns a tuple with the CreditPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPartyIdentifier

`func (o *Transaction) SetCreditPartyIdentifier(v TransactionCreditPartyIdentifier)`

SetCreditPartyIdentifier sets CreditPartyIdentifier field to given value.

### HasCreditPartyIdentifier

`func (o *Transaction) HasCreditPartyIdentifier() bool`

HasCreditPartyIdentifier returns a boolean if a field has been set.

### GetStatementIdentifier

`func (o *Transaction) GetStatementIdentifier() TransactionStatementIdentifier`

GetStatementIdentifier returns the StatementIdentifier field if non-nil, zero value otherwise.

### GetStatementIdentifierOk

`func (o *Transaction) GetStatementIdentifierOk() (*TransactionStatementIdentifier, bool)`

GetStatementIdentifierOk returns a tuple with the StatementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIdentifier

`func (o *Transaction) SetStatementIdentifier(v TransactionStatementIdentifier)`

SetStatementIdentifier sets StatementIdentifier field to given value.

### HasStatementIdentifier

`func (o *Transaction) HasStatementIdentifier() bool`

HasStatementIdentifier returns a boolean if a field has been set.

### GetAdditionalIdentifier

`func (o *Transaction) GetAdditionalIdentifier() TransactionAdditionalIdentifier`

GetAdditionalIdentifier returns the AdditionalIdentifier field if non-nil, zero value otherwise.

### GetAdditionalIdentifierOk

`func (o *Transaction) GetAdditionalIdentifierOk() (*TransactionAdditionalIdentifier, bool)`

GetAdditionalIdentifierOk returns a tuple with the AdditionalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIdentifier

`func (o *Transaction) SetAdditionalIdentifier(v TransactionAdditionalIdentifier)`

SetAdditionalIdentifier sets AdditionalIdentifier field to given value.

### HasAdditionalIdentifier

`func (o *Transaction) HasAdditionalIdentifier() bool`

HasAdditionalIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


