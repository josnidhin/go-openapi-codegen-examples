# ProductFixedValuePinPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Tags** | **[]string** |  | 
**Service** | [**ProductService**](ProductService.md) |  | 
**Operator** | [**GetOperators200ResponseInner**](GetOperators200ResponseInner.md) |  | 
**Regions** | [**[]ServiceRegion**](ServiceRegion.md) |  | 
**Type** | **string** |  | 
**Validity** | [**NullableProductFixedValueRechargeAllOfValidity**](ProductFixedValueRechargeAllOfValidity.md) |  | 
**RequiredDebitPartyIdentifierFields** | **[][]string** |  | 
**RequiredCreditPartyIdentifierFields** | **[][]string** |  | 
**RequiredSenderFields** | **[][]string** |  | 
**RequiredBeneficiaryFields** | **[][]string** |  | 
**RequiredStatementIdentifierFields** | **[][]string** |  | 
**RequiredAdditionalIdentifierFields** | **[][]string** |  | 
**AvailabilityZones** | [**[]AvailabilityZones**](AvailabilityZones.md) |  | 
**Source** | [**ProductFixedValueRechargeAllOfSource**](ProductFixedValueRechargeAllOfSource.md) |  | 
**Destination** | [**ProductFixedValueRechargeAllOfSource**](ProductFixedValueRechargeAllOfSource.md) |  | 
**Prices** | [**ProductFixedValueRechargeAllOfPrices**](ProductFixedValueRechargeAllOfPrices.md) |  | 
**Rates** | [**Rates**](Rates.md) |  | 
**Pin** | [**PIN**](PIN.md) |  | 
**Benefits** | [**[]FixedBenefit**](FixedBenefit.md) |  | 
**Promotions** | [**[]ProductPromotion**](ProductPromotion.md) |  | 

## Methods

### NewProductFixedValuePinPurchase

`func NewProductFixedValuePinPurchase(id int32, name string, description string, tags []string, service ProductService, operator GetOperators200ResponseInner, regions []ServiceRegion, type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, requiredAdditionalIdentifierFields [][]string, availabilityZones []AvailabilityZones, source ProductFixedValueRechargeAllOfSource, destination ProductFixedValueRechargeAllOfSource, prices ProductFixedValueRechargeAllOfPrices, rates Rates, pin PIN, benefits []FixedBenefit, promotions []ProductPromotion, ) *ProductFixedValuePinPurchase`

NewProductFixedValuePinPurchase instantiates a new ProductFixedValuePinPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductFixedValuePinPurchaseWithDefaults

`func NewProductFixedValuePinPurchaseWithDefaults() *ProductFixedValuePinPurchase`

NewProductFixedValuePinPurchaseWithDefaults instantiates a new ProductFixedValuePinPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductFixedValuePinPurchase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductFixedValuePinPurchase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductFixedValuePinPurchase) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductFixedValuePinPurchase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductFixedValuePinPurchase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductFixedValuePinPurchase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductFixedValuePinPurchase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductFixedValuePinPurchase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductFixedValuePinPurchase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *ProductFixedValuePinPurchase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductFixedValuePinPurchase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductFixedValuePinPurchase) SetTags(v []string)`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *ProductFixedValuePinPurchase) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ProductFixedValuePinPurchase) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetService

`func (o *ProductFixedValuePinPurchase) GetService() ProductService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ProductFixedValuePinPurchase) GetServiceOk() (*ProductService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ProductFixedValuePinPurchase) SetService(v ProductService)`

SetService sets Service field to given value.


### GetOperator

`func (o *ProductFixedValuePinPurchase) GetOperator() GetOperators200ResponseInner`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ProductFixedValuePinPurchase) GetOperatorOk() (*GetOperators200ResponseInner, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ProductFixedValuePinPurchase) SetOperator(v GetOperators200ResponseInner)`

SetOperator sets Operator field to given value.


### GetRegions

`func (o *ProductFixedValuePinPurchase) GetRegions() []ServiceRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ProductFixedValuePinPurchase) GetRegionsOk() (*[]ServiceRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ProductFixedValuePinPurchase) SetRegions(v []ServiceRegion)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *ProductFixedValuePinPurchase) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *ProductFixedValuePinPurchase) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetType

`func (o *ProductFixedValuePinPurchase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductFixedValuePinPurchase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductFixedValuePinPurchase) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *ProductFixedValuePinPurchase) GetValidity() ProductFixedValueRechargeAllOfValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ProductFixedValuePinPurchase) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ProductFixedValuePinPurchase) SetValidity(v ProductFixedValueRechargeAllOfValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *ProductFixedValuePinPurchase) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ProductFixedValuePinPurchase) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetRequiredDebitPartyIdentifierFields

`func (o *ProductFixedValuePinPurchase) GetRequiredDebitPartyIdentifierFields() [][]string`

GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredDebitPartyIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredDebitPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDebitPartyIdentifierFields

`func (o *ProductFixedValuePinPurchase) SetRequiredDebitPartyIdentifierFields(v [][]string)`

SetRequiredDebitPartyIdentifierFields sets RequiredDebitPartyIdentifierFields field to given value.


### SetRequiredDebitPartyIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredDebitPartyIdentifierFieldsNil(b bool)`

 SetRequiredDebitPartyIdentifierFieldsNil sets the value for RequiredDebitPartyIdentifierFields to be an explicit nil

### UnsetRequiredDebitPartyIdentifierFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredDebitPartyIdentifierFields()`

UnsetRequiredDebitPartyIdentifierFields ensures that no value is present for RequiredDebitPartyIdentifierFields, not even an explicit nil
### GetRequiredCreditPartyIdentifierFields

`func (o *ProductFixedValuePinPurchase) GetRequiredCreditPartyIdentifierFields() [][]string`

GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredCreditPartyIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredCreditPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCreditPartyIdentifierFields

`func (o *ProductFixedValuePinPurchase) SetRequiredCreditPartyIdentifierFields(v [][]string)`

SetRequiredCreditPartyIdentifierFields sets RequiredCreditPartyIdentifierFields field to given value.


### SetRequiredCreditPartyIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredCreditPartyIdentifierFieldsNil(b bool)`

 SetRequiredCreditPartyIdentifierFieldsNil sets the value for RequiredCreditPartyIdentifierFields to be an explicit nil

### UnsetRequiredCreditPartyIdentifierFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredCreditPartyIdentifierFields()`

UnsetRequiredCreditPartyIdentifierFields ensures that no value is present for RequiredCreditPartyIdentifierFields, not even an explicit nil
### GetRequiredSenderFields

`func (o *ProductFixedValuePinPurchase) GetRequiredSenderFields() [][]string`

GetRequiredSenderFields returns the RequiredSenderFields field if non-nil, zero value otherwise.

### GetRequiredSenderFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredSenderFieldsOk() (*[][]string, bool)`

GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSenderFields

`func (o *ProductFixedValuePinPurchase) SetRequiredSenderFields(v [][]string)`

SetRequiredSenderFields sets RequiredSenderFields field to given value.


### SetRequiredSenderFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredSenderFieldsNil(b bool)`

 SetRequiredSenderFieldsNil sets the value for RequiredSenderFields to be an explicit nil

### UnsetRequiredSenderFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredSenderFields()`

UnsetRequiredSenderFields ensures that no value is present for RequiredSenderFields, not even an explicit nil
### GetRequiredBeneficiaryFields

`func (o *ProductFixedValuePinPurchase) GetRequiredBeneficiaryFields() [][]string`

GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field if non-nil, zero value otherwise.

### GetRequiredBeneficiaryFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredBeneficiaryFieldsOk() (*[][]string, bool)`

GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeneficiaryFields

`func (o *ProductFixedValuePinPurchase) SetRequiredBeneficiaryFields(v [][]string)`

SetRequiredBeneficiaryFields sets RequiredBeneficiaryFields field to given value.


### SetRequiredBeneficiaryFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredBeneficiaryFieldsNil(b bool)`

 SetRequiredBeneficiaryFieldsNil sets the value for RequiredBeneficiaryFields to be an explicit nil

### UnsetRequiredBeneficiaryFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredBeneficiaryFields()`

UnsetRequiredBeneficiaryFields ensures that no value is present for RequiredBeneficiaryFields, not even an explicit nil
### GetRequiredStatementIdentifierFields

`func (o *ProductFixedValuePinPurchase) GetRequiredStatementIdentifierFields() [][]string`

GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredStatementIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredStatementIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatementIdentifierFields

`func (o *ProductFixedValuePinPurchase) SetRequiredStatementIdentifierFields(v [][]string)`

SetRequiredStatementIdentifierFields sets RequiredStatementIdentifierFields field to given value.


### SetRequiredStatementIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredStatementIdentifierFieldsNil(b bool)`

 SetRequiredStatementIdentifierFieldsNil sets the value for RequiredStatementIdentifierFields to be an explicit nil

### UnsetRequiredStatementIdentifierFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredStatementIdentifierFields()`

UnsetRequiredStatementIdentifierFields ensures that no value is present for RequiredStatementIdentifierFields, not even an explicit nil
### GetRequiredAdditionalIdentifierFields

`func (o *ProductFixedValuePinPurchase) GetRequiredAdditionalIdentifierFields() [][]string`

GetRequiredAdditionalIdentifierFields returns the RequiredAdditionalIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredAdditionalIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchase) GetRequiredAdditionalIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredAdditionalIdentifierFieldsOk returns a tuple with the RequiredAdditionalIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAdditionalIdentifierFields

`func (o *ProductFixedValuePinPurchase) SetRequiredAdditionalIdentifierFields(v [][]string)`

SetRequiredAdditionalIdentifierFields sets RequiredAdditionalIdentifierFields field to given value.


### SetRequiredAdditionalIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchase) SetRequiredAdditionalIdentifierFieldsNil(b bool)`

 SetRequiredAdditionalIdentifierFieldsNil sets the value for RequiredAdditionalIdentifierFields to be an explicit nil

### UnsetRequiredAdditionalIdentifierFields
`func (o *ProductFixedValuePinPurchase) UnsetRequiredAdditionalIdentifierFields()`

UnsetRequiredAdditionalIdentifierFields ensures that no value is present for RequiredAdditionalIdentifierFields, not even an explicit nil
### GetAvailabilityZones

`func (o *ProductFixedValuePinPurchase) GetAvailabilityZones() []AvailabilityZones`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProductFixedValuePinPurchase) GetAvailabilityZonesOk() (*[]AvailabilityZones, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProductFixedValuePinPurchase) SetAvailabilityZones(v []AvailabilityZones)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSource

`func (o *ProductFixedValuePinPurchase) GetSource() ProductFixedValueRechargeAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductFixedValuePinPurchase) GetSourceOk() (*ProductFixedValueRechargeAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductFixedValuePinPurchase) SetSource(v ProductFixedValueRechargeAllOfSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ProductFixedValuePinPurchase) GetDestination() ProductFixedValueRechargeAllOfSource`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProductFixedValuePinPurchase) GetDestinationOk() (*ProductFixedValueRechargeAllOfSource, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProductFixedValuePinPurchase) SetDestination(v ProductFixedValueRechargeAllOfSource)`

SetDestination sets Destination field to given value.


### GetPrices

`func (o *ProductFixedValuePinPurchase) GetPrices() ProductFixedValueRechargeAllOfPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductFixedValuePinPurchase) GetPricesOk() (*ProductFixedValueRechargeAllOfPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductFixedValuePinPurchase) SetPrices(v ProductFixedValueRechargeAllOfPrices)`

SetPrices sets Prices field to given value.


### GetRates

`func (o *ProductFixedValuePinPurchase) GetRates() Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ProductFixedValuePinPurchase) GetRatesOk() (*Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ProductFixedValuePinPurchase) SetRates(v Rates)`

SetRates sets Rates field to given value.


### GetPin

`func (o *ProductFixedValuePinPurchase) GetPin() PIN`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ProductFixedValuePinPurchase) GetPinOk() (*PIN, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ProductFixedValuePinPurchase) SetPin(v PIN)`

SetPin sets Pin field to given value.


### GetBenefits

`func (o *ProductFixedValuePinPurchase) GetBenefits() []FixedBenefit`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *ProductFixedValuePinPurchase) GetBenefitsOk() (*[]FixedBenefit, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *ProductFixedValuePinPurchase) SetBenefits(v []FixedBenefit)`

SetBenefits sets Benefits field to given value.


### SetBenefitsNil

`func (o *ProductFixedValuePinPurchase) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *ProductFixedValuePinPurchase) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *ProductFixedValuePinPurchase) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductFixedValuePinPurchase) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductFixedValuePinPurchase) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.


### SetPromotionsNil

`func (o *ProductFixedValuePinPurchase) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *ProductFixedValuePinPurchase) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


