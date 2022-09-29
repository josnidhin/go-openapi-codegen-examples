# ProductRangedValuePinPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**Service** | [**Service**](Service.md) |  | 
**Operator** | [**GetOperators200ResponseInner**](GetOperators200ResponseInner.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 
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

### NewProductRangedValuePinPurchase

`func NewProductRangedValuePinPurchase(id int32, name string, description string, tags []string, service Service, operator GetOperators200ResponseInner, regions []ServiceRegion, type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, acceptedCalculationModes []CalculationModes, availabilityZones []AvailabilityZones, source ProductRangedValueRechargeAllOfSource, destination ProductRangedValueRechargeAllOfDestination, prices ProductRangedValueRechargeAllOfPrices, rates Rates, pin PIN, benefits []RangedBenefit, promotions []ProductPromotion, ) *ProductRangedValuePinPurchase`

NewProductRangedValuePinPurchase instantiates a new ProductRangedValuePinPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRangedValuePinPurchaseWithDefaults

`func NewProductRangedValuePinPurchaseWithDefaults() *ProductRangedValuePinPurchase`

NewProductRangedValuePinPurchaseWithDefaults instantiates a new ProductRangedValuePinPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductRangedValuePinPurchase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductRangedValuePinPurchase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductRangedValuePinPurchase) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductRangedValuePinPurchase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductRangedValuePinPurchase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductRangedValuePinPurchase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductRangedValuePinPurchase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductRangedValuePinPurchase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductRangedValuePinPurchase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ProductRangedValuePinPurchase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductRangedValuePinPurchase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductRangedValuePinPurchase) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ProductRangedValuePinPurchase) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProductRangedValuePinPurchase) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetService

`func (o *ProductRangedValuePinPurchase) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProductRangedValuePinPurchase) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProductRangedValuePinPurchase) SetService(v Service)`

SetService sets Service field to given value.


### GetOperator

`func (o *ProductRangedValuePinPurchase) GetOperator() GetOperators200ResponseInner`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProductRangedValuePinPurchase) GetOperatorOk() (*GetOperators200ResponseInner, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProductRangedValuePinPurchase) SetOperator(v GetOperators200ResponseInner)`

SetOperator sets Operator field to given value.


### GetRegions

`func (o *ProductRangedValuePinPurchase) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProductRangedValuePinPurchase) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProductRangedValuePinPurchase) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *ProductRangedValuePinPurchase) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ProductRangedValuePinPurchase) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetType

`func (o *ProductRangedValuePinPurchase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductRangedValuePinPurchase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductRangedValuePinPurchase) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *ProductRangedValuePinPurchase) GetValidity() ProductFixedValueRechargeAllOfValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ProductRangedValuePinPurchase) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ProductRangedValuePinPurchase) SetValidity(v ProductFixedValueRechargeAllOfValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *ProductRangedValuePinPurchase) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ProductRangedValuePinPurchase) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValuePinPurchase) GetRequiredDebitPartyIdentifierFields() [][]string`

GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredDebitPartyIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchase) GetRequiredDebitPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValuePinPurchase) SetRequiredDebitPartyIdentifierFields(v [][]string)`

SetRequiredDebitPartyIdentifierFields sets RequiredDebitPartyIdentifierFields field to given value.


### SetRequiredDebitPartyIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchase) SetRequiredDebitPartyIdentifierFieldsNil(b bool)`

 SetRequiredDebitPartyIdentifierFieldsNil sets the value for RequiredDebitPartyIdentifierFields to be an explicit nil

### UnsetRequiredDebitPartyIdentifierFields
`func (o *ProductRangedValuePinPurchase) UnsetRequiredDebitPartyIdentifierFields()`

UnsetRequiredDebitPartyIdentifierFields ensures that no value is present for RequiredDebitPartyIdentifierFields, not even an explicit nil
### GetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValuePinPurchase) GetRequiredCreditPartyIdentifierFields() [][]string`

GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredCreditPartyIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchase) GetRequiredCreditPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValuePinPurchase) SetRequiredCreditPartyIdentifierFields(v [][]string)`

SetRequiredCreditPartyIdentifierFields sets RequiredCreditPartyIdentifierFields field to given value.


### SetRequiredCreditPartyIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchase) SetRequiredCreditPartyIdentifierFieldsNil(b bool)`

 SetRequiredCreditPartyIdentifierFieldsNil sets the value for RequiredCreditPartyIdentifierFields to be an explicit nil

### UnsetRequiredCreditPartyIdentifierFields
`func (o *ProductRangedValuePinPurchase) UnsetRequiredCreditPartyIdentifierFields()`

UnsetRequiredCreditPartyIdentifierFields ensures that no value is present for RequiredCreditPartyIdentifierFields, not even an explicit nil
### GetRequiredSenderFields

`func (o *ProductRangedValuePinPurchase) GetRequiredSenderFields() [][]string`

GetRequiredSenderFields returns the RequiredSenderFields field if non-nil, zero value otherwise.

### GetRequiredSenderFieldsOk

`func (o *ProductRangedValuePinPurchase) GetRequiredSenderFieldsOk() (*[][]string, bool)`

GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSenderFields

`func (o *ProductRangedValuePinPurchase) SetRequiredSenderFields(v [][]string)`

SetRequiredSenderFields sets RequiredSenderFields field to given value.


### SetRequiredSenderFieldsNil

`func (o *ProductRangedValuePinPurchase) SetRequiredSenderFieldsNil(b bool)`

 SetRequiredSenderFieldsNil sets the value for RequiredSenderFields to be an explicit nil

### UnsetRequiredSenderFields
`func (o *ProductRangedValuePinPurchase) UnsetRequiredSenderFields()`

UnsetRequiredSenderFields ensures that no value is present for RequiredSenderFields, not even an explicit nil
### GetRequiredBeneficiaryFields

`func (o *ProductRangedValuePinPurchase) GetRequiredBeneficiaryFields() [][]string`

GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field if non-nil, zero value otherwise.

### GetRequiredBeneficiaryFieldsOk

`func (o *ProductRangedValuePinPurchase) GetRequiredBeneficiaryFieldsOk() (*[][]string, bool)`

GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeneficiaryFields

`func (o *ProductRangedValuePinPurchase) SetRequiredBeneficiaryFields(v [][]string)`

SetRequiredBeneficiaryFields sets RequiredBeneficiaryFields field to given value.


### SetRequiredBeneficiaryFieldsNil

`func (o *ProductRangedValuePinPurchase) SetRequiredBeneficiaryFieldsNil(b bool)`

 SetRequiredBeneficiaryFieldsNil sets the value for RequiredBeneficiaryFields to be an explicit nil

### UnsetRequiredBeneficiaryFields
`func (o *ProductRangedValuePinPurchase) UnsetRequiredBeneficiaryFields()`

UnsetRequiredBeneficiaryFields ensures that no value is present for RequiredBeneficiaryFields, not even an explicit nil
### GetRequiredStatementIdentifierFields

`func (o *ProductRangedValuePinPurchase) GetRequiredStatementIdentifierFields() [][]string`

GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredStatementIdentifierFieldsOk

`func (o *ProductRangedValuePinPurchase) GetRequiredStatementIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatementIdentifierFields

`func (o *ProductRangedValuePinPurchase) SetRequiredStatementIdentifierFields(v [][]string)`

SetRequiredStatementIdentifierFields sets RequiredStatementIdentifierFields field to given value.


### SetRequiredStatementIdentifierFieldsNil

`func (o *ProductRangedValuePinPurchase) SetRequiredStatementIdentifierFieldsNil(b bool)`

 SetRequiredStatementIdentifierFieldsNil sets the value for RequiredStatementIdentifierFields to be an explicit nil

### UnsetRequiredStatementIdentifierFields
`func (o *ProductRangedValuePinPurchase) UnsetRequiredStatementIdentifierFields()`

UnsetRequiredStatementIdentifierFields ensures that no value is present for RequiredStatementIdentifierFields, not even an explicit nil
### GetAcceptedCalculationModes

`func (o *ProductRangedValuePinPurchase) GetAcceptedCalculationModes() []CalculationModes`

GetAcceptedCalculationModes returns the AcceptedCalculationModes field if non-nil, zero value otherwise.

### GetAcceptedCalculationModesOk

`func (o *ProductRangedValuePinPurchase) GetAcceptedCalculationModesOk() (*[]CalculationModes, bool)`

GetAcceptedCalculationModesOk returns a tuple with the AcceptedCalculationModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCalculationModes

`func (o *ProductRangedValuePinPurchase) SetAcceptedCalculationModes(v []CalculationModes)`

SetAcceptedCalculationModes sets AcceptedCalculationModes field to given value.


### GetAvailabilityZones

`func (o *ProductRangedValuePinPurchase) GetAvailabilityZones() []AvailabilityZones`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProductRangedValuePinPurchase) GetAvailabilityZonesOk() (*[]AvailabilityZones, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProductRangedValuePinPurchase) SetAvailabilityZones(v []AvailabilityZones)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSource

`func (o *ProductRangedValuePinPurchase) GetSource() ProductRangedValueRechargeAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductRangedValuePinPurchase) GetSourceOk() (*ProductRangedValueRechargeAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductRangedValuePinPurchase) SetSource(v ProductRangedValueRechargeAllOfSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ProductRangedValuePinPurchase) GetDestination() ProductRangedValueRechargeAllOfDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProductRangedValuePinPurchase) GetDestinationOk() (*ProductRangedValueRechargeAllOfDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProductRangedValuePinPurchase) SetDestination(v ProductRangedValueRechargeAllOfDestination)`

SetDestination sets Destination field to given value.


### GetPrices

`func (o *ProductRangedValuePinPurchase) GetPrices() ProductRangedValueRechargeAllOfPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductRangedValuePinPurchase) GetPricesOk() (*ProductRangedValueRechargeAllOfPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductRangedValuePinPurchase) SetPrices(v ProductRangedValueRechargeAllOfPrices)`

SetPrices sets Prices field to given value.


### GetRates

`func (o *ProductRangedValuePinPurchase) GetRates() Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ProductRangedValuePinPurchase) GetRatesOk() (*Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ProductRangedValuePinPurchase) SetRates(v Rates)`

SetRates sets Rates field to given value.


### GetPin

`func (o *ProductRangedValuePinPurchase) GetPin() PIN`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ProductRangedValuePinPurchase) GetPinOk() (*PIN, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ProductRangedValuePinPurchase) SetPin(v PIN)`

SetPin sets Pin field to given value.


### GetBenefits

`func (o *ProductRangedValuePinPurchase) GetBenefits() []RangedBenefit`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *ProductRangedValuePinPurchase) GetBenefitsOk() (*[]RangedBenefit, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *ProductRangedValuePinPurchase) SetBenefits(v []RangedBenefit)`

SetBenefits sets Benefits field to given value.


### SetBenefitsNil

`func (o *ProductRangedValuePinPurchase) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *ProductRangedValuePinPurchase) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *ProductRangedValuePinPurchase) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductRangedValuePinPurchase) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductRangedValuePinPurchase) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.


### SetPromotionsNil

`func (o *ProductRangedValuePinPurchase) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *ProductRangedValuePinPurchase) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


