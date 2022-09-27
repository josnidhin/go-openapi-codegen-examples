# ProductRangedValueRecharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**Service** | [**Service**](Service.md) |  | 
**Operator** | [**OperatorsGet200ResponseInner**](OperatorsGet200ResponseInner.md) |  | 
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
**Benefits** | [**[]RangedBenefit**](RangedBenefit.md) |  | 
**Promotions** | [**[]ProductPromotion**](ProductPromotion.md) |  | 

## Methods

### NewProductRangedValueRecharge

`func NewProductRangedValueRecharge(id int32, name string, description string, tags []string, service Service, operator OperatorsGet200ResponseInner, regions []ServiceRegion, type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, acceptedCalculationModes []CalculationModes, availabilityZones []AvailabilityZones, source ProductRangedValueRechargeAllOfSource, destination ProductRangedValueRechargeAllOfDestination, prices ProductRangedValueRechargeAllOfPrices, rates Rates, benefits []RangedBenefit, promotions []ProductPromotion, ) *ProductRangedValueRecharge`

NewProductRangedValueRecharge instantiates a new ProductRangedValueRecharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRangedValueRechargeWithDefaults

`func NewProductRangedValueRechargeWithDefaults() *ProductRangedValueRecharge`

NewProductRangedValueRechargeWithDefaults instantiates a new ProductRangedValueRecharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductRangedValueRecharge) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductRangedValueRecharge) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductRangedValueRecharge) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductRangedValueRecharge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductRangedValueRecharge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductRangedValueRecharge) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductRangedValueRecharge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductRangedValueRecharge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductRangedValueRecharge) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ProductRangedValueRecharge) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductRangedValueRecharge) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductRangedValueRecharge) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ProductRangedValueRecharge) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProductRangedValueRecharge) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetService

`func (o *ProductRangedValueRecharge) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProductRangedValueRecharge) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProductRangedValueRecharge) SetService(v Service)`

SetService sets Service field to given value.


### GetOperator

`func (o *ProductRangedValueRecharge) GetOperator() OperatorsGet200ResponseInner`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProductRangedValueRecharge) GetOperatorOk() (*OperatorsGet200ResponseInner, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProductRangedValueRecharge) SetOperator(v OperatorsGet200ResponseInner)`

SetOperator sets Operator field to given value.


### GetRegions

`func (o *ProductRangedValueRecharge) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProductRangedValueRecharge) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProductRangedValueRecharge) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *ProductRangedValueRecharge) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ProductRangedValueRecharge) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetType

`func (o *ProductRangedValueRecharge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductRangedValueRecharge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductRangedValueRecharge) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *ProductRangedValueRecharge) GetValidity() ProductFixedValueRechargeAllOfValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ProductRangedValueRecharge) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ProductRangedValueRecharge) SetValidity(v ProductFixedValueRechargeAllOfValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *ProductRangedValueRecharge) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ProductRangedValueRecharge) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValueRecharge) GetRequiredDebitPartyIdentifierFields() [][]string`

GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredDebitPartyIdentifierFieldsOk

`func (o *ProductRangedValueRecharge) GetRequiredDebitPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDebitPartyIdentifierFields

`func (o *ProductRangedValueRecharge) SetRequiredDebitPartyIdentifierFields(v [][]string)`

SetRequiredDebitPartyIdentifierFields sets RequiredDebitPartyIdentifierFields field to given value.


### SetRequiredDebitPartyIdentifierFieldsNil

`func (o *ProductRangedValueRecharge) SetRequiredDebitPartyIdentifierFieldsNil(b bool)`

 SetRequiredDebitPartyIdentifierFieldsNil sets the value for RequiredDebitPartyIdentifierFields to be an explicit nil

### UnsetRequiredDebitPartyIdentifierFields
`func (o *ProductRangedValueRecharge) UnsetRequiredDebitPartyIdentifierFields()`

UnsetRequiredDebitPartyIdentifierFields ensures that no value is present for RequiredDebitPartyIdentifierFields, not even an explicit nil
### GetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValueRecharge) GetRequiredCreditPartyIdentifierFields() [][]string`

GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredCreditPartyIdentifierFieldsOk

`func (o *ProductRangedValueRecharge) GetRequiredCreditPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCreditPartyIdentifierFields

`func (o *ProductRangedValueRecharge) SetRequiredCreditPartyIdentifierFields(v [][]string)`

SetRequiredCreditPartyIdentifierFields sets RequiredCreditPartyIdentifierFields field to given value.


### GetRequiredSenderFields

`func (o *ProductRangedValueRecharge) GetRequiredSenderFields() [][]string`

GetRequiredSenderFields returns the RequiredSenderFields field if non-nil, zero value otherwise.

### GetRequiredSenderFieldsOk

`func (o *ProductRangedValueRecharge) GetRequiredSenderFieldsOk() (*[][]string, bool)`

GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSenderFields

`func (o *ProductRangedValueRecharge) SetRequiredSenderFields(v [][]string)`

SetRequiredSenderFields sets RequiredSenderFields field to given value.


### SetRequiredSenderFieldsNil

`func (o *ProductRangedValueRecharge) SetRequiredSenderFieldsNil(b bool)`

 SetRequiredSenderFieldsNil sets the value for RequiredSenderFields to be an explicit nil

### UnsetRequiredSenderFields
`func (o *ProductRangedValueRecharge) UnsetRequiredSenderFields()`

UnsetRequiredSenderFields ensures that no value is present for RequiredSenderFields, not even an explicit nil
### GetRequiredBeneficiaryFields

`func (o *ProductRangedValueRecharge) GetRequiredBeneficiaryFields() [][]string`

GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field if non-nil, zero value otherwise.

### GetRequiredBeneficiaryFieldsOk

`func (o *ProductRangedValueRecharge) GetRequiredBeneficiaryFieldsOk() (*[][]string, bool)`

GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeneficiaryFields

`func (o *ProductRangedValueRecharge) SetRequiredBeneficiaryFields(v [][]string)`

SetRequiredBeneficiaryFields sets RequiredBeneficiaryFields field to given value.


### SetRequiredBeneficiaryFieldsNil

`func (o *ProductRangedValueRecharge) SetRequiredBeneficiaryFieldsNil(b bool)`

 SetRequiredBeneficiaryFieldsNil sets the value for RequiredBeneficiaryFields to be an explicit nil

### UnsetRequiredBeneficiaryFields
`func (o *ProductRangedValueRecharge) UnsetRequiredBeneficiaryFields()`

UnsetRequiredBeneficiaryFields ensures that no value is present for RequiredBeneficiaryFields, not even an explicit nil
### GetRequiredStatementIdentifierFields

`func (o *ProductRangedValueRecharge) GetRequiredStatementIdentifierFields() [][]string`

GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredStatementIdentifierFieldsOk

`func (o *ProductRangedValueRecharge) GetRequiredStatementIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatementIdentifierFields

`func (o *ProductRangedValueRecharge) SetRequiredStatementIdentifierFields(v [][]string)`

SetRequiredStatementIdentifierFields sets RequiredStatementIdentifierFields field to given value.


### SetRequiredStatementIdentifierFieldsNil

`func (o *ProductRangedValueRecharge) SetRequiredStatementIdentifierFieldsNil(b bool)`

 SetRequiredStatementIdentifierFieldsNil sets the value for RequiredStatementIdentifierFields to be an explicit nil

### UnsetRequiredStatementIdentifierFields
`func (o *ProductRangedValueRecharge) UnsetRequiredStatementIdentifierFields()`

UnsetRequiredStatementIdentifierFields ensures that no value is present for RequiredStatementIdentifierFields, not even an explicit nil
### GetAcceptedCalculationModes

`func (o *ProductRangedValueRecharge) GetAcceptedCalculationModes() []CalculationModes`

GetAcceptedCalculationModes returns the AcceptedCalculationModes field if non-nil, zero value otherwise.

### GetAcceptedCalculationModesOk

`func (o *ProductRangedValueRecharge) GetAcceptedCalculationModesOk() (*[]CalculationModes, bool)`

GetAcceptedCalculationModesOk returns a tuple with the AcceptedCalculationModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCalculationModes

`func (o *ProductRangedValueRecharge) SetAcceptedCalculationModes(v []CalculationModes)`

SetAcceptedCalculationModes sets AcceptedCalculationModes field to given value.


### GetAvailabilityZones

`func (o *ProductRangedValueRecharge) GetAvailabilityZones() []AvailabilityZones`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProductRangedValueRecharge) GetAvailabilityZonesOk() (*[]AvailabilityZones, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProductRangedValueRecharge) SetAvailabilityZones(v []AvailabilityZones)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSource

`func (o *ProductRangedValueRecharge) GetSource() ProductRangedValueRechargeAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductRangedValueRecharge) GetSourceOk() (*ProductRangedValueRechargeAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductRangedValueRecharge) SetSource(v ProductRangedValueRechargeAllOfSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ProductRangedValueRecharge) GetDestination() ProductRangedValueRechargeAllOfDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProductRangedValueRecharge) GetDestinationOk() (*ProductRangedValueRechargeAllOfDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProductRangedValueRecharge) SetDestination(v ProductRangedValueRechargeAllOfDestination)`

SetDestination sets Destination field to given value.


### GetPrices

`func (o *ProductRangedValueRecharge) GetPrices() ProductRangedValueRechargeAllOfPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductRangedValueRecharge) GetPricesOk() (*ProductRangedValueRechargeAllOfPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductRangedValueRecharge) SetPrices(v ProductRangedValueRechargeAllOfPrices)`

SetPrices sets Prices field to given value.


### GetRates

`func (o *ProductRangedValueRecharge) GetRates() Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ProductRangedValueRecharge) GetRatesOk() (*Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ProductRangedValueRecharge) SetRates(v Rates)`

SetRates sets Rates field to given value.


### GetBenefits

`func (o *ProductRangedValueRecharge) GetBenefits() []RangedBenefit`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *ProductRangedValueRecharge) GetBenefitsOk() (*[]RangedBenefit, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *ProductRangedValueRecharge) SetBenefits(v []RangedBenefit)`

SetBenefits sets Benefits field to given value.


### SetBenefitsNil

`func (o *ProductRangedValueRecharge) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *ProductRangedValueRecharge) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *ProductRangedValueRecharge) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductRangedValueRecharge) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductRangedValueRecharge) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.


### SetPromotionsNil

`func (o *ProductRangedValueRecharge) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *ProductRangedValueRecharge) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


