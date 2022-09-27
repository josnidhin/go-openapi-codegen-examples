# ProductRangedValuePinPurchaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Validity** | [**NullableProductFixedValueRechargeAllOfValidity**](ProductFixedValueRechargeAllOfValidity.md) |  | 
**RequiredDebitPartyIdentifierFields** | **[][]string** |  | 
**RequiredCreditPartyIdentifierFields** | **[][]string** |  | 
**RequiredSenderFields** | **[][]string** |  | 
**RequiredBeneficiaryFields** | **[][]string** |  | 
**RequiredStatementIdentifierFields** | **[][]string** |  | 
**AcceptedCalculationModes** | [**[]CalculationModes**](CalculationModes.md) |  | 
**AvailabilityZones** | [**[]AvailabilityZones**](AvailabilityZones.md) |  | 
**Source** | [**ProductRangedValueRechargeAllOfSource**](ProductRangedValueRechargeAllOfSource.md) |  | 
**Destination** | [**ProductRangedValueRechargeAllOfDestination**](ProductRangedValueRechargeAllOfDestination.md) |  | 
**Prices** | [**ProductRangedValueRechargeAllOfPrices**](ProductRangedValueRechargeAllOfPrices.md) |  | 
**Rates** | [**Rates**](Rates.md) |  | 
**Pin** | [**PIN**](PIN.md) |  | 
**Benefits** | [**[]RangedBenefit**](RangedBenefit.md) |  | 
**Promotions** | [**[]ProductPromotion**](ProductPromotion.md) |  | 

## Methods

### NewProductRangedValuePinPurchaseAllOf

`func NewProductRangedValuePinPurchaseAllOf(type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, acceptedCalculationModes []CalculationModes, availabilityZones []AvailabilityZones, source ProductRangedValueRechargeAllOfSource, destination ProductRangedValueRechargeAllOfDestination, prices ProductRangedValueRechargeAllOfPrices, rates Rates, pin PIN, benefits []RangedBenefit, promotions []ProductPromotion, ) *ProductRangedValuePinPurchaseAllOf`

NewProductRangedValuePinPurchaseAllOf instantiates a new ProductRangedValuePinPurchaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRangedValuePinPurchaseAllOfWithDefaults

`func NewProductRangedValuePinPurchaseAllOfWithDefaults() *ProductRangedValuePinPurchaseAllOf`

NewProductRangedValuePinPurchaseAllOfWithDefaults instantiates a new ProductRangedValuePinPurchaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProductRangedValuePinPurchaseAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductRangedValuePinPurchaseAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *ProductRangedValuePinPurchaseAllOf) GetValidity() ProductFixedValueRechargeAllOfValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ProductRangedValuePinPurchaseAllOf) SetValidity(v ProductFixedValueRechargeAllOfValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredDebitPartyIdentifierFields() [][]string`

GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredDebitPartyIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredDebitPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredDebitPartyIdentifierFields(v [][]string)`

SetRequiredDebitPartyIdentifierFields sets RequiredDebitPartyIdentifierFields field to given value.


### SetRequiredDebitPartyIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredDebitPartyIdentifierFieldsNil(b bool)`

 SetRequiredDebitPartyIdentifierFieldsNil sets the value for RequiredDebitPartyIdentifierFields to be an explicit nil

### UnsetRequiredDebitPartyIdentifierFields
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetRequiredDebitPartyIdentifierFields()`

UnsetRequiredDebitPartyIdentifierFields ensures that no value is present for RequiredDebitPartyIdentifierFields, not even an explicit nil
### GetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredCreditPartyIdentifierFields() [][]string`

GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredCreditPartyIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredCreditPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredCreditPartyIdentifierFields(v [][]string)`

SetRequiredCreditPartyIdentifierFields sets RequiredCreditPartyIdentifierFields field to given value.


### SetRequiredCreditPartyIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredCreditPartyIdentifierFieldsNil(b bool)`

 SetRequiredCreditPartyIdentifierFieldsNil sets the value for RequiredCreditPartyIdentifierFields to be an explicit nil

### UnsetRequiredCreditPartyIdentifierFields
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetRequiredCreditPartyIdentifierFields()`

UnsetRequiredCreditPartyIdentifierFields ensures that no value is present for RequiredCreditPartyIdentifierFields, not even an explicit nil
### GetRequiredSenderFields

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredSenderFields() [][]string`

GetRequiredSenderFields returns the RequiredSenderFields field if non-nil, zero value otherwise.

### GetRequiredSenderFieldsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredSenderFieldsOk() (*[][]string, bool)`

GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSenderFields

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredSenderFields(v [][]string)`

SetRequiredSenderFields sets RequiredSenderFields field to given value.


### SetRequiredSenderFieldsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredSenderFieldsNil(b bool)`

 SetRequiredSenderFieldsNil sets the value for RequiredSenderFields to be an explicit nil

### UnsetRequiredSenderFields
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetRequiredSenderFields()`

UnsetRequiredSenderFields ensures that no value is present for RequiredSenderFields, not even an explicit nil
### GetRequiredBeneficiaryFields

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredBeneficiaryFields() [][]string`

GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field if non-nil, zero value otherwise.

### GetRequiredBeneficiaryFieldsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredBeneficiaryFieldsOk() (*[][]string, bool)`

GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeneficiaryFields

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredBeneficiaryFields(v [][]string)`

SetRequiredBeneficiaryFields sets RequiredBeneficiaryFields field to given value.


### SetRequiredBeneficiaryFieldsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredBeneficiaryFieldsNil(b bool)`

 SetRequiredBeneficiaryFieldsNil sets the value for RequiredBeneficiaryFields to be an explicit nil

### UnsetRequiredBeneficiaryFields
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetRequiredBeneficiaryFields()`

UnsetRequiredBeneficiaryFields ensures that no value is present for RequiredBeneficiaryFields, not even an explicit nil
### GetRequiredStatementIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredStatementIdentifierFields() [][]string`

GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredStatementIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRequiredStatementIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatementIdentifierFields

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredStatementIdentifierFields(v [][]string)`

SetRequiredStatementIdentifierFields sets RequiredStatementIdentifierFields field to given value.


### SetRequiredStatementIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetRequiredStatementIdentifierFieldsNil(b bool)`

 SetRequiredStatementIdentifierFieldsNil sets the value for RequiredStatementIdentifierFields to be an explicit nil

### UnsetRequiredStatementIdentifierFields
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetRequiredStatementIdentifierFields()`

UnsetRequiredStatementIdentifierFields ensures that no value is present for RequiredStatementIdentifierFields, not even an explicit nil
### GetAcceptedCalculationModes

`func (o *ProductRangedValuePinPurchaseAllOf) GetAcceptedCalculationModes() []CalculationModes`

GetAcceptedCalculationModes returns the AcceptedCalculationModes field if non-nil, zero value otherwise.

### GetAcceptedCalculationModesOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetAcceptedCalculationModesOk() (*[]CalculationModes, bool)`

GetAcceptedCalculationModesOk returns a tuple with the AcceptedCalculationModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCalculationModes

`func (o *ProductRangedValuePinPurchaseAllOf) SetAcceptedCalculationModes(v []CalculationModes)`

SetAcceptedCalculationModes sets AcceptedCalculationModes field to given value.


### GetAvailabilityZones

`func (o *ProductRangedValuePinPurchaseAllOf) GetAvailabilityZones() []AvailabilityZones`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetAvailabilityZonesOk() (*[]AvailabilityZones, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProductRangedValuePinPurchaseAllOf) SetAvailabilityZones(v []AvailabilityZones)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSource

`func (o *ProductRangedValuePinPurchaseAllOf) GetSource() ProductRangedValueRechargeAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetSourceOk() (*ProductRangedValueRechargeAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductRangedValuePinPurchaseAllOf) SetSource(v ProductRangedValueRechargeAllOfSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ProductRangedValuePinPurchaseAllOf) GetDestination() ProductRangedValueRechargeAllOfDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetDestinationOk() (*ProductRangedValueRechargeAllOfDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProductRangedValuePinPurchaseAllOf) SetDestination(v ProductRangedValueRechargeAllOfDestination)`

SetDestination sets Destination field to given value.


### GetPrices

`func (o *ProductRangedValuePinPurchaseAllOf) GetPrices() ProductRangedValueRechargeAllOfPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetPricesOk() (*ProductRangedValueRechargeAllOfPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductRangedValuePinPurchaseAllOf) SetPrices(v ProductRangedValueRechargeAllOfPrices)`

SetPrices sets Prices field to given value.


### GetRates

`func (o *ProductRangedValuePinPurchaseAllOf) GetRates() Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetRatesOk() (*Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ProductRangedValuePinPurchaseAllOf) SetRates(v Rates)`

SetRates sets Rates field to given value.


### GetPin

`func (o *ProductRangedValuePinPurchaseAllOf) GetPin() PIN`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetPinOk() (*PIN, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ProductRangedValuePinPurchaseAllOf) SetPin(v PIN)`

SetPin sets Pin field to given value.


### GetBenefits

`func (o *ProductRangedValuePinPurchaseAllOf) GetBenefits() []RangedBenefit`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetBenefitsOk() (*[]RangedBenefit, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *ProductRangedValuePinPurchaseAllOf) SetBenefits(v []RangedBenefit)`

SetBenefits sets Benefits field to given value.


### SetBenefitsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *ProductRangedValuePinPurchaseAllOf) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductRangedValuePinPurchaseAllOf) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductRangedValuePinPurchaseAllOf) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.


### SetPromotionsNil

`func (o *ProductRangedValuePinPurchaseAllOf) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *ProductRangedValuePinPurchaseAllOf) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


