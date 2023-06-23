# ProductFixedValuePinPurchaseAllOf

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

### NewProductFixedValuePinPurchaseAllOf

`func NewProductFixedValuePinPurchaseAllOf(type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, requiredAdditionalIdentifierFields [][]string, availabilityZones []AvailabilityZones, source ProductFixedValueRechargeAllOfSource, destination ProductFixedValueRechargeAllOfSource, prices ProductFixedValueRechargeAllOfPrices, rates Rates, pin PIN, benefits []FixedBenefit, promotions []ProductPromotion, ) *ProductFixedValuePinPurchaseAllOf`

NewProductFixedValuePinPurchaseAllOf instantiates a new ProductFixedValuePinPurchaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductFixedValuePinPurchaseAllOfWithDefaults

`func NewProductFixedValuePinPurchaseAllOfWithDefaults() *ProductFixedValuePinPurchaseAllOf`

NewProductFixedValuePinPurchaseAllOfWithDefaults instantiates a new ProductFixedValuePinPurchaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProductFixedValuePinPurchaseAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductFixedValuePinPurchaseAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetValidity

`func (o *ProductFixedValuePinPurchaseAllOf) GetValidity() ProductFixedValueRechargeAllOfValidity`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ProductFixedValuePinPurchaseAllOf) SetValidity(v ProductFixedValueRechargeAllOfValidity)`

SetValidity sets Validity field to given value.


### SetValidityNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetRequiredDebitPartyIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredDebitPartyIdentifierFields() [][]string`

GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredDebitPartyIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredDebitPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDebitPartyIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredDebitPartyIdentifierFields(v [][]string)`

SetRequiredDebitPartyIdentifierFields sets RequiredDebitPartyIdentifierFields field to given value.


### SetRequiredDebitPartyIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredDebitPartyIdentifierFieldsNil(b bool)`

 SetRequiredDebitPartyIdentifierFieldsNil sets the value for RequiredDebitPartyIdentifierFields to be an explicit nil

### UnsetRequiredDebitPartyIdentifierFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredDebitPartyIdentifierFields()`

UnsetRequiredDebitPartyIdentifierFields ensures that no value is present for RequiredDebitPartyIdentifierFields, not even an explicit nil
### GetRequiredCreditPartyIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredCreditPartyIdentifierFields() [][]string`

GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredCreditPartyIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredCreditPartyIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCreditPartyIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredCreditPartyIdentifierFields(v [][]string)`

SetRequiredCreditPartyIdentifierFields sets RequiredCreditPartyIdentifierFields field to given value.


### SetRequiredCreditPartyIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredCreditPartyIdentifierFieldsNil(b bool)`

 SetRequiredCreditPartyIdentifierFieldsNil sets the value for RequiredCreditPartyIdentifierFields to be an explicit nil

### UnsetRequiredCreditPartyIdentifierFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredCreditPartyIdentifierFields()`

UnsetRequiredCreditPartyIdentifierFields ensures that no value is present for RequiredCreditPartyIdentifierFields, not even an explicit nil
### GetRequiredSenderFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredSenderFields() [][]string`

GetRequiredSenderFields returns the RequiredSenderFields field if non-nil, zero value otherwise.

### GetRequiredSenderFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredSenderFieldsOk() (*[][]string, bool)`

GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSenderFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredSenderFields(v [][]string)`

SetRequiredSenderFields sets RequiredSenderFields field to given value.


### SetRequiredSenderFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredSenderFieldsNil(b bool)`

 SetRequiredSenderFieldsNil sets the value for RequiredSenderFields to be an explicit nil

### UnsetRequiredSenderFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredSenderFields()`

UnsetRequiredSenderFields ensures that no value is present for RequiredSenderFields, not even an explicit nil
### GetRequiredBeneficiaryFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredBeneficiaryFields() [][]string`

GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field if non-nil, zero value otherwise.

### GetRequiredBeneficiaryFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredBeneficiaryFieldsOk() (*[][]string, bool)`

GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredBeneficiaryFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredBeneficiaryFields(v [][]string)`

SetRequiredBeneficiaryFields sets RequiredBeneficiaryFields field to given value.


### SetRequiredBeneficiaryFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredBeneficiaryFieldsNil(b bool)`

 SetRequiredBeneficiaryFieldsNil sets the value for RequiredBeneficiaryFields to be an explicit nil

### UnsetRequiredBeneficiaryFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredBeneficiaryFields()`

UnsetRequiredBeneficiaryFields ensures that no value is present for RequiredBeneficiaryFields, not even an explicit nil
### GetRequiredStatementIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredStatementIdentifierFields() [][]string`

GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredStatementIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredStatementIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatementIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredStatementIdentifierFields(v [][]string)`

SetRequiredStatementIdentifierFields sets RequiredStatementIdentifierFields field to given value.


### SetRequiredStatementIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredStatementIdentifierFieldsNil(b bool)`

 SetRequiredStatementIdentifierFieldsNil sets the value for RequiredStatementIdentifierFields to be an explicit nil

### UnsetRequiredStatementIdentifierFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredStatementIdentifierFields()`

UnsetRequiredStatementIdentifierFields ensures that no value is present for RequiredStatementIdentifierFields, not even an explicit nil
### GetRequiredAdditionalIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredAdditionalIdentifierFields() [][]string`

GetRequiredAdditionalIdentifierFields returns the RequiredAdditionalIdentifierFields field if non-nil, zero value otherwise.

### GetRequiredAdditionalIdentifierFieldsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRequiredAdditionalIdentifierFieldsOk() (*[][]string, bool)`

GetRequiredAdditionalIdentifierFieldsOk returns a tuple with the RequiredAdditionalIdentifierFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAdditionalIdentifierFields

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredAdditionalIdentifierFields(v [][]string)`

SetRequiredAdditionalIdentifierFields sets RequiredAdditionalIdentifierFields field to given value.


### SetRequiredAdditionalIdentifierFieldsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetRequiredAdditionalIdentifierFieldsNil(b bool)`

 SetRequiredAdditionalIdentifierFieldsNil sets the value for RequiredAdditionalIdentifierFields to be an explicit nil

### UnsetRequiredAdditionalIdentifierFields
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetRequiredAdditionalIdentifierFields()`

UnsetRequiredAdditionalIdentifierFields ensures that no value is present for RequiredAdditionalIdentifierFields, not even an explicit nil
### GetAvailabilityZones

`func (o *ProductFixedValuePinPurchaseAllOf) GetAvailabilityZones() []AvailabilityZones`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetAvailabilityZonesOk() (*[]AvailabilityZones, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProductFixedValuePinPurchaseAllOf) SetAvailabilityZones(v []AvailabilityZones)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetSource

`func (o *ProductFixedValuePinPurchaseAllOf) GetSource() ProductFixedValueRechargeAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetSourceOk() (*ProductFixedValueRechargeAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductFixedValuePinPurchaseAllOf) SetSource(v ProductFixedValueRechargeAllOfSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *ProductFixedValuePinPurchaseAllOf) GetDestination() ProductFixedValueRechargeAllOfSource`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetDestinationOk() (*ProductFixedValueRechargeAllOfSource, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ProductFixedValuePinPurchaseAllOf) SetDestination(v ProductFixedValueRechargeAllOfSource)`

SetDestination sets Destination field to given value.


### GetPrices

`func (o *ProductFixedValuePinPurchaseAllOf) GetPrices() ProductFixedValueRechargeAllOfPrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetPricesOk() (*ProductFixedValueRechargeAllOfPrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductFixedValuePinPurchaseAllOf) SetPrices(v ProductFixedValueRechargeAllOfPrices)`

SetPrices sets Prices field to given value.


### GetRates

`func (o *ProductFixedValuePinPurchaseAllOf) GetRates() Rates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetRatesOk() (*Rates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *ProductFixedValuePinPurchaseAllOf) SetRates(v Rates)`

SetRates sets Rates field to given value.


### GetPin

`func (o *ProductFixedValuePinPurchaseAllOf) GetPin() PIN`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetPinOk() (*PIN, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ProductFixedValuePinPurchaseAllOf) SetPin(v PIN)`

SetPin sets Pin field to given value.


### GetBenefits

`func (o *ProductFixedValuePinPurchaseAllOf) GetBenefits() []FixedBenefit`

GetBenefits returns the Benefits field if non-nil, zero value otherwise.

### GetBenefitsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetBenefitsOk() (*[]FixedBenefit, bool)`

GetBenefitsOk returns a tuple with the Benefits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefits

`func (o *ProductFixedValuePinPurchaseAllOf) SetBenefits(v []FixedBenefit)`

SetBenefits sets Benefits field to given value.


### SetBenefitsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetBenefitsNil(b bool)`

 SetBenefitsNil sets the value for Benefits to be an explicit nil

### UnsetBenefits
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetBenefits()`

UnsetBenefits ensures that no value is present for Benefits, not even an explicit nil
### GetPromotions

`func (o *ProductFixedValuePinPurchaseAllOf) GetPromotions() []ProductPromotion`

GetPromotions returns the Promotions field if non-nil, zero value otherwise.

### GetPromotionsOk

`func (o *ProductFixedValuePinPurchaseAllOf) GetPromotionsOk() (*[]ProductPromotion, bool)`

GetPromotionsOk returns a tuple with the Promotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotions

`func (o *ProductFixedValuePinPurchaseAllOf) SetPromotions(v []ProductPromotion)`

SetPromotions sets Promotions field to given value.


### SetPromotionsNil

`func (o *ProductFixedValuePinPurchaseAllOf) SetPromotionsNil(b bool)`

 SetPromotionsNil sets the value for Promotions to be an explicit nil

### UnsetPromotions
`func (o *ProductFixedValuePinPurchaseAllOf) UnsetPromotions()`

UnsetPromotions ensures that no value is present for Promotions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


