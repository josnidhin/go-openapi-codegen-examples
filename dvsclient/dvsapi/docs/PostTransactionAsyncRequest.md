# PostTransactionAsyncRequest

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
**CallbackUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewPostTransactionAsyncRequest

`func NewPostTransactionAsyncRequest(externalId string, productId int32, ) *PostTransactionAsyncRequest`

NewPostTransactionAsyncRequest instantiates a new PostTransactionAsyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTransactionAsyncRequestWithDefaults

`func NewPostTransactionAsyncRequestWithDefaults() *PostTransactionAsyncRequest`

NewPostTransactionAsyncRequestWithDefaults instantiates a new PostTransactionAsyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostTransactionAsyncRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTransactionAsyncRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTransactionAsyncRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostTransactionAsyncRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *PostTransactionAsyncRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PostTransactionAsyncRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PostTransactionAsyncRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetCreationDate

`func (o *PostTransactionAsyncRequest) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *PostTransactionAsyncRequest) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *PostTransactionAsyncRequest) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *PostTransactionAsyncRequest) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetConfirmationExpirationDate

`func (o *PostTransactionAsyncRequest) GetConfirmationExpirationDate() time.Time`

GetConfirmationExpirationDate returns the ConfirmationExpirationDate field if non-nil, zero value otherwise.

### GetConfirmationExpirationDateOk

`func (o *PostTransactionAsyncRequest) GetConfirmationExpirationDateOk() (*time.Time, bool)`

GetConfirmationExpirationDateOk returns a tuple with the ConfirmationExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationExpirationDate

`func (o *PostTransactionAsyncRequest) SetConfirmationExpirationDate(v time.Time)`

SetConfirmationExpirationDate sets ConfirmationExpirationDate field to given value.

### HasConfirmationExpirationDate

`func (o *PostTransactionAsyncRequest) HasConfirmationExpirationDate() bool`

HasConfirmationExpirationDate returns a boolean if a field has been set.

### GetConfirmationDate

`func (o *PostTransactionAsyncRequest) GetConfirmationDate() time.Time`

GetConfirmationDate returns the ConfirmationDate field if non-nil, zero value otherwise.

### GetConfirmationDateOk

`func (o *PostTransactionAsyncRequest) GetConfirmationDateOk() (*time.Time, bool)`

GetConfirmationDateOk returns a tuple with the ConfirmationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationDate

`func (o *PostTransactionAsyncRequest) SetConfirmationDate(v time.Time)`

SetConfirmationDate sets ConfirmationDate field to given value.

### HasConfirmationDate

`func (o *PostTransactionAsyncRequest) HasConfirmationDate() bool`

HasConfirmationDate returns a boolean if a field has been set.

### GetProductId

`func (o *PostTransactionAsyncRequest) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PostTransactionAsyncRequest) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PostTransactionAsyncRequest) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetCalculationMode

`func (o *PostTransactionAsyncRequest) GetCalculationMode() CalculationModes`

GetCalculationMode returns the CalculationMode field if non-nil, zero value otherwise.

### GetCalculationModeOk

`func (o *PostTransactionAsyncRequest) GetCalculationModeOk() (*CalculationModes, bool)`

GetCalculationModeOk returns a tuple with the CalculationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationMode

`func (o *PostTransactionAsyncRequest) SetCalculationMode(v CalculationModes)`

SetCalculationMode sets CalculationMode field to given value.

### HasCalculationMode

`func (o *PostTransactionAsyncRequest) HasCalculationMode() bool`

HasCalculationMode returns a boolean if a field has been set.

### GetSource

`func (o *PostTransactionAsyncRequest) GetSource() TransactionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostTransactionAsyncRequest) GetSourceOk() (*TransactionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostTransactionAsyncRequest) SetSource(v TransactionSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *PostTransactionAsyncRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *PostTransactionAsyncRequest) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *PostTransactionAsyncRequest) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetDestination

`func (o *PostTransactionAsyncRequest) GetDestination() TransactionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PostTransactionAsyncRequest) GetDestinationOk() (*TransactionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PostTransactionAsyncRequest) SetDestination(v TransactionDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *PostTransactionAsyncRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *PostTransactionAsyncRequest) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *PostTransactionAsyncRequest) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetAutoConfirm

`func (o *PostTransactionAsyncRequest) GetAutoConfirm() bool`

GetAutoConfirm returns the AutoConfirm field if non-nil, zero value otherwise.

### GetAutoConfirmOk

`func (o *PostTransactionAsyncRequest) GetAutoConfirmOk() (*bool, bool)`

GetAutoConfirmOk returns a tuple with the AutoConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConfirm

`func (o *PostTransactionAsyncRequest) SetAutoConfirm(v bool)`

SetAutoConfirm sets AutoConfirm field to given value.

### HasAutoConfirm

`func (o *PostTransactionAsyncRequest) HasAutoConfirm() bool`

HasAutoConfirm returns a boolean if a field has been set.

### GetStatus

`func (o *PostTransactionAsyncRequest) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostTransactionAsyncRequest) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostTransactionAsyncRequest) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostTransactionAsyncRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOperatorReference

`func (o *PostTransactionAsyncRequest) GetOperatorReference() string`

GetOperatorReference returns the OperatorReference field if non-nil, zero value otherwise.

### GetOperatorReferenceOk

`func (o *PostTransactionAsyncRequest) GetOperatorReferenceOk() (*string, bool)`

GetOperatorReferenceOk returns a tuple with the OperatorReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorReference

`func (o *PostTransactionAsyncRequest) SetOperatorReference(v string)`

SetOperatorReference sets OperatorReference field to given value.

### HasOperatorReference

`func (o *PostTransactionAsyncRequest) HasOperatorReference() bool`

HasOperatorReference returns a boolean if a field has been set.

### SetOperatorReferenceNil

`func (o *PostTransactionAsyncRequest) SetOperatorReferenceNil(b bool)`

 SetOperatorReferenceNil sets the value for OperatorReference to be an explicit nil

### UnsetOperatorReference
`func (o *PostTransactionAsyncRequest) UnsetOperatorReference()`

UnsetOperatorReference ensures that no value is present for OperatorReference, not even an explicit nil
### GetPin

`func (o *PostTransactionAsyncRequest) GetPin() TransactionPin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *PostTransactionAsyncRequest) GetPinOk() (*TransactionPin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *PostTransactionAsyncRequest) SetPin(v TransactionPin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *PostTransactionAsyncRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.

### SetPinNil

`func (o *PostTransactionAsyncRequest) SetPinNil(b bool)`

 SetPinNil sets the value for Pin to be an explicit nil

### UnsetPin
`func (o *PostTransactionAsyncRequest) UnsetPin()`

UnsetPin ensures that no value is present for Pin, not even an explicit nil
### GetProduct

`func (o *PostTransactionAsyncRequest) GetProduct() TransactionProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PostTransactionAsyncRequest) GetProductOk() (*TransactionProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PostTransactionAsyncRequest) SetProduct(v TransactionProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PostTransactionAsyncRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPrices

`func (o *PostTransactionAsyncRequest) GetPrices() TransactionPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *PostTransactionAsyncRequest) GetPricesOk() (*TransactionPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *PostTransactionAsyncRequest) SetPrices(v TransactionPrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *PostTransactionAsyncRequest) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetRates

`func (o *PostTransactionAsyncRequest) GetRates() TransactionRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *PostTransactionAsyncRequest) GetRatesOk() (*TransactionRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *PostTransactionAsyncRequest) SetRates(v TransactionRates)`

SetRates sets Rates field to given value.

### HasRates

`func (o *PostTransactionAsyncRequest) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetBenefits

`func (o *PostTransactionAsyncRequest) GetBenefits() []TransactionBenefitsInner`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *PostTransactionAsyncRequest) GetBenefitsOk() (*[]TransactionBenefitsInner, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *PostTransactionAsyncRequest) SetBenefits(v []TransactionBenefitsInner)`

SetBenefits sets Benefits field to given value.

### HasBenefits

`func (o *PostTransactionAsyncRequest) HasBenefits() bool`

HasBenefits returns a boolean if a field has been set.

### SetBenefitsNil

`func (o *PostTransactionAsyncRequest) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *PostTransactionAsyncRequest) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *PostTransactionAsyncRequest) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *PostTransactionAsyncRequest) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *PostTransactionAsyncRequest) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.

### HasPromotions

`func (o *PostTransactionAsyncRequest) HasPromotions() bool`

HasPromotions returns a boolean if a field has been set.

### SetPromotionsNil

`func (o *PostTransactionAsyncRequest) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *PostTransactionAsyncRequest) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil
### GetRequestedValues

`func (o *PostTransactionAsyncRequest) GetRequestedValues() TransactionRequestedValues`

GetRequestedValues returns the RequestedValues field if non-nil, zero value otherwise.

### GetRequestedValuesOk

`func (o *PostTransactionAsyncRequest) GetRequestedValuesOk() (*TransactionRequestedValues, bool)`

GetRequestedValuesOk returns a tuple with the RequestedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedValues

`func (o *PostTransactionAsyncRequest) SetRequestedValues(v TransactionRequestedValues)`

SetRequestedValues sets RequestedValues field to given value.

### HasRequestedValues

`func (o *PostTransactionAsyncRequest) HasRequestedValues() bool`

HasRequestedValues returns a boolean if a field has been set.

### SetRequestedValuesNil

`func (o *PostTransactionAsyncRequest) SetRequestedValuesNil(b bool)`

 SetRequestedValuesNil sets the value for RequestedValues to be an explicit nil

### UnsetRequestedValues
`func (o *PostTransactionAsyncRequest) UnsetRequestedValues()`

UnsetRequestedValues ensures that no value is present for RequestedValues, not even an explicit nil
### GetAdjustedValues

`func (o *PostTransactionAsyncRequest) GetAdjustedValues() TransactionRequestedValues`

GetAdjustedValues returns the AdjustedValues field if non-nil, zero value otherwise.

### GetAdjustedValuesOk

`func (o *PostTransactionAsyncRequest) GetAdjustedValuesOk() (*TransactionRequestedValues, bool)`

GetAdjustedValuesOk returns a tuple with the AdjustedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedValues

`func (o *PostTransactionAsyncRequest) SetAdjustedValues(v TransactionRequestedValues)`

SetAdjustedValues sets AdjustedValues field to given value.

### HasAdjustedValues

`func (o *PostTransactionAsyncRequest) HasAdjustedValues() bool`

HasAdjustedValues returns a boolean if a field has been set.

### SetAdjustedValuesNil

`func (o *PostTransactionAsyncRequest) SetAdjustedValuesNil(b bool)`

 SetAdjustedValuesNil sets the value for AdjustedValues to be an explicit nil

### UnsetAdjustedValues
`func (o *PostTransactionAsyncRequest) UnsetAdjustedValues()`

UnsetAdjustedValues ensures that no value is present for AdjustedValues, not even an explicit nil
### GetSender

`func (o *PostTransactionAsyncRequest) GetSender() TransactionSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *PostTransactionAsyncRequest) GetSenderOk() (*TransactionSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *PostTransactionAsyncRequest) SetSender(v TransactionSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *PostTransactionAsyncRequest) HasSender() bool`

HasSender returns a boolean if a field has been set.

### SetSenderNil

`func (o *PostTransactionAsyncRequest) SetSenderNil(b bool)`

 SetSenderNil sets the value for Sender to be an explicit nil

### UnsetSender
`func (o *PostTransactionAsyncRequest) UnsetSender()`

UnsetSender ensures that no value is present for Sender, not even an explicit nil
### GetBeneficiary

`func (o *PostTransactionAsyncRequest) GetBeneficiary() TransactionSender`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *PostTransactionAsyncRequest) GetBeneficiaryOk() (*TransactionSender, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *PostTransactionAsyncRequest) SetBeneficiary(v TransactionSender)`

SetBeneficiary sets Beneficiary field to given value.

### HasBeneficiary

`func (o *PostTransactionAsyncRequest) HasBeneficiary() bool`

HasBeneficiary returns a boolean if a field has been set.

### SetBeneficiaryNil

`func (o *PostTransactionAsyncRequest) SetBeneficiaryNil(b bool)`

 SetBeneficiaryNil sets the value for Beneficiary to be an explicit nil

### UnsetBeneficiary
`func (o *PostTransactionAsyncRequest) UnsetBeneficiary()`

UnsetBeneficiary ensures that no value is present for Beneficiary, not even an explicit nil
### GetDebitPartyIdentifier

`func (o *PostTransactionAsyncRequest) GetDebitPartyIdentifier() TransactionDebitPartyIdentifier`

GetDebitPartyIdentifier returns the DebitPartyIdentifier field if non-nil, zero value otherwise.

### GetDebitPartyIdentifierOk

`func (o *PostTransactionAsyncRequest) GetDebitPartyIdentifierOk() (*TransactionDebitPartyIdentifier, bool)`

GetDebitPartyIdentifierOk returns a tuple with the DebitPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitPartyIdentifier

`func (o *PostTransactionAsyncRequest) SetDebitPartyIdentifier(v TransactionDebitPartyIdentifier)`

SetDebitPartyIdentifier sets DebitPartyIdentifier field to given value.

### HasDebitPartyIdentifier

`func (o *PostTransactionAsyncRequest) HasDebitPartyIdentifier() bool`

HasDebitPartyIdentifier returns a boolean if a field has been set.

### SetDebitPartyIdentifierNil

`func (o *PostTransactionAsyncRequest) SetDebitPartyIdentifierNil(b bool)`

 SetDebitPartyIdentifierNil sets the value for DebitPartyIdentifier to be an explicit nil

### UnsetDebitPartyIdentifier
`func (o *PostTransactionAsyncRequest) UnsetDebitPartyIdentifier()`

UnsetDebitPartyIdentifier ensures that no value is present for DebitPartyIdentifier, not even an explicit nil
### GetCreditPartyIdentifier

`func (o *PostTransactionAsyncRequest) GetCreditPartyIdentifier() TransactionCreditPartyIdentifier`

GetCreditPartyIdentifier returns the CreditPartyIdentifier field if non-nil, zero value otherwise.

### GetCreditPartyIdentifierOk

`func (o *PostTransactionAsyncRequest) GetCreditPartyIdentifierOk() (*TransactionCreditPartyIdentifier, bool)`

GetCreditPartyIdentifierOk returns a tuple with the CreditPartyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPartyIdentifier

`func (o *PostTransactionAsyncRequest) SetCreditPartyIdentifier(v TransactionCreditPartyIdentifier)`

SetCreditPartyIdentifier sets CreditPartyIdentifier field to given value.

### HasCreditPartyIdentifier

`func (o *PostTransactionAsyncRequest) HasCreditPartyIdentifier() bool`

HasCreditPartyIdentifier returns a boolean if a field has been set.

### GetStatementIdentifier

`func (o *PostTransactionAsyncRequest) GetStatementIdentifier() TransactionStatementIdentifier`

GetStatementIdentifier returns the StatementIdentifier field if non-nil, zero value otherwise.

### GetStatementIdentifierOk

`func (o *PostTransactionAsyncRequest) GetStatementIdentifierOk() (*TransactionStatementIdentifier, bool)`

GetStatementIdentifierOk returns a tuple with the StatementIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementIdentifier

`func (o *PostTransactionAsyncRequest) SetStatementIdentifier(v TransactionStatementIdentifier)`

SetStatementIdentifier sets StatementIdentifier field to given value.

### HasStatementIdentifier

`func (o *PostTransactionAsyncRequest) HasStatementIdentifier() bool`

HasStatementIdentifier returns a boolean if a field has been set.

### GetAdditionalIdentifier

`func (o *PostTransactionAsyncRequest) GetAdditionalIdentifier() TransactionAdditionalIdentifier`

GetAdditionalIdentifier returns the AdditionalIdentifier field if non-nil, zero value otherwise.

### GetAdditionalIdentifierOk

`func (o *PostTransactionAsyncRequest) GetAdditionalIdentifierOk() (*TransactionAdditionalIdentifier, bool)`

GetAdditionalIdentifierOk returns a tuple with the AdditionalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIdentifier

`func (o *PostTransactionAsyncRequest) SetAdditionalIdentifier(v TransactionAdditionalIdentifier)`

SetAdditionalIdentifier sets AdditionalIdentifier field to given value.

### HasAdditionalIdentifier

`func (o *PostTransactionAsyncRequest) HasAdditionalIdentifier() bool`

HasAdditionalIdentifier returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *PostTransactionAsyncRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PostTransactionAsyncRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PostTransactionAsyncRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PostTransactionAsyncRequest) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


