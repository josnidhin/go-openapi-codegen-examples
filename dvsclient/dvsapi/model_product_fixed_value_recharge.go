/*
Digital Value Services API

# Overview  Welcome to the Digital Value Services (DVS) API reference.  This API serves as the primary gateway to facilitate digital value transfers through [DT One](https://www.dtone.com), a leading global network covering more than 160 countries and 550 mobile operators.  The Digital Value Services API is organized according to [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, using [JSON](https://en.wikipedia.org/wiki/JSON) as format for data interchange, and provides the following services:   - [Discovery Services](#tag/Services)   - [Transaction Services](#tag/Transactions)   - [Account Services](#tag/Balances)   - [Look-Up Services](#tag/Mobile-Number)  If you have any questions and/or suggestions related to our API, please do not hesitate to send an email to the [DVS API support team](mailto:dvs-api-support@dtone.com).  ## Integration Libraries  Officially supported [SDK](https://en.wikipedia.org/wiki/Software_development_kit)s are available for the following languages:  * [Java](https://github.com/dtone/dtone-dvs-api-java-client) * [Node.js](https://www.npmjs.com/package/@dtone/dvs)  These SDKs offer an accelerated path to developing your applications as an alternative to accessing the REST API directly.  Separately, we would love to hear from you! If you have any questions and/or suggestions related to our SDKs, please do not hesitate to create corresponding [GitHub issues](https://guides.github.com/features/issues/) or send an email to the [DVS Open Source team](mailto:opensource@dtone.com).  ## Sandbox  A sandbox environment is available for testing integrations with the DVS API. It is available at [https://preprod-dvs-api.dtone.com/v1/](https://preprod-dvs-api.dtone.com/v1/).  You can generate sandbox API keys from your [DT Shop](https://dtshop.dtone.com/account?tab=developer) account, under the **Pre-Production API Keys** section.  All transactions on the sandbox environment are simulated: no real transaction goes through. To simulate different responses, the last three digits of the recipient mobile number (i.e. `credit_party_identifier.mobile_number`) will have to be replaced with one of the following suffixes:  | Suffix              | Transaction Status                        | Example       | | ---                 | ---                                       | ---           | | `100`, `200`, `300` | `COMPLETED` (PIN-less)                    | `+6595123100` | | `101`, `201`, `301` | `COMPLETED` (PIN-based)                   | `+6595123201` | | `102`, `202`, `302` | `DECLINED-INVALID-CREDIT-PARTY`           | `+6595123102` | | `103`, `203`, `303` | `DECLINED-BARRED-CREDIT-PARTY`            | `+6595123103` | | `104`, `204`, `304` | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE` | `+6595123205` | | `105`, `205`, `305` | `DECLINED-DUPLICATED-TRANSACTION`         | `+6595123105` | | `106`, `206`, `306` | `DECLINED`                                | `+6595123206` | | `107`, `207`, `307` | `DECLINED-EXCEPTION`                      | `+6595123107` |  The different suffixes for a given transaction status can be used to simulate delays, as follows:   * `10X` suffix will take at least 3 seconds to finish   * `20X` suffix will take at least 20 seconds to finish   * `30X` suffix will take at least 5 minutes to finish  Please note that there are products that do not require any credit party identifier such as Gift Cards. For these products, the simulated transaction will always be in `COMPLETED` status.  ## Versioning  Endpoints of the API are prefixed with a corresponding version number.  This is done to provide complete isolation between implementations and to ensure that subsequent major changes to the API will never affect existing integrations.  When a new version of the API is available and you are keen to upgrade, testing in the sandbox environment to ensure that everything works well with your implementation before switching to the production environment comes highly recommended.  Feel free to [contact us](mailto:dvs-api-support@dtone.com) should you have any questions and/or are in need of assistance during your tests.  ## Transactions  The main purpose of this API is to deliver value (e.g. mobile airtime top-up, data bundles, etc.) to a beneficiary. This is what we call a \"transaction\".  During the course of a transfer, a transaction undergoes various status changes (or transitions) as illustrated below.  ![transaction states](/images/transaction_states.png)  As changes in transaction status occur, updates are sent in real-time when a callback URL is provided. In conjunction, transaction status can be queried through one of two means: via the returned `id` or a provided `external_id`.  The latter serves as your unique reference and provides a utility to retrieve transaction details when exceptions occur, such as when the supposed API response was not received, as an example.  ## Balances  Transactions can be created through the platform as long as there is enough balance available in your account. A given balance is composed of the following:  | Balance   | Description                                              | | ---       | ---                                                      | | Available | Balance amount available for use                         | | Holding   | Amount being held while transactions are being processed |  As a given transaction goes through various changes in status as outlined [here](#section/Overview/Transactions), corresponding balance movements will be made. The following table illustrates the relationship between transaction status and balance movements:  | Transaction Status                  | Balance Movement | Description                                                | | ---                                 | ---              | ---                                                        | | `CREATED`                           | Authorize        | Transfer wholesale price and fee from available to holding | | `CANCELLED`, `REJECTED`, `DECLINED` | Void             | Amount authorized in holding moves back to available       | | `COMPLETED`                         | Capture          | Amount authorized in holding is captured, i.e. debited     | | `REVERSED`                          | Reverse          | Debited amount is reversed back into available             |  ## Flow  Once a product has been selected through one of the [discovery methods](/#tag/Services) provided by the API, the actual transfer (i.e. transaction) can be performed in either one of the following modes:   - Asynchronous (recommended)   - Synchronous  Each mode is accessible via a specific endpoint.  As soon as a transaction is confirmed, the transfer order will be sent to the operator for immediate processing. During this time, the transaction will remain in a `CONFIRMED` status until the final status is received from the operator.  ### Asynchronous Mode  When a transaction is created and confirmed in an asynchronous fashion, the HTTP connection won't have to be kept open. This preserves system resources on your applications. As such, performing transactions **asynchronously** is **recommended**.  ### Synchronous Mode  When a transaction is created and confirmed in a synchronous fashion, the HTTP connection will be kept open in an attempt to capture the final status from the receiving operator so it can be returned in the API response. The processing time usually takes just a few seconds. However, with some receiving operators, it may take longer.  Our system will keep the HTTP connection open for up to 180 seconds (this value can be configured upon request) and will return a status before closing this connection. This status can be in one of the final status (e.g. `COMPLETED`, `DECLINED`) or not (e.g. `SUBMITTED`). In the latter case, this denotes the transaction is still being processed by the receiving operator.  **Note:** your application does not have to wait for the connection to close, it can listen for a shorter period of time and query the final status later on (refer to the \"Final Status\" section below for more details).  ### Final Status  Regardless of the processing mode, the application should be designed to capture the final status of a transaction. This can be done through one of the following means:   - Checking the status of a specific transaction via the corresponding API method (\"pull\" mechanism)   - Configuring a callback URL passed in the request when creating a transaction (\"push\" mechanism)  ## Callbacks  As a transaction is being processed, changes in status will be notified in real-time if a callback URL was provided. Even though one callback per transaction is expected (i.e. change to either `COMPLETED` or `DECLINED`), a manual reversal from the [DT One](https://dtone.com/) team, which may happen in very rare occasions, will also trigger a callback to inform your application of a change in transaction status to `REVERSED`.  This callback endpoint must be implemented by the sending partner, which should expect an HTTP `POST` request containing a transaction object represented in [JSON](https://en.wikipedia.org/wiki/JSON). As callbacks will be sent from the [DT One](https://dtone.com/) servers, these endpoints will have to be publicly-accessible in most cases. During development, a service such as [ngrok](https://ngrok.com/) can be used to expose local servers to the internet.  Upon successful receipt of data, the callback endpoint should respond with an HTTP `2XX` status. In the event that the platform did not receive a successful status, callback notifications will be retried several times, beyond which, the transaction status will have to be queried through the API.  ## Status and Errors  ### HTTP Status Codes  [DT One](https://dtone.com/) uses standard HTTP response codes to indicate whether an API request was successful or not.  | Status | Description                                        | | ---    | ---                                                | | `200`  | OK                                                 | | `201`  | Created: Resource created                          | | `202`  | Accepted: Request has been accepted for processing | | `400`  | Bad Request: Request was malformed                 | | `401`  | Unauthorized: Credentials missing or invalid       | | `404`  | Not Found: Resource doesn't exist                  | | `429`  | Too Many Requests                                  | | `500`  | Server Error: Error occurred on DT One             |  ### API Error Codes  | Code      | Description                                       | | ---       | ---                                               | | `1000400` | Bad Request                                       | | `1000401` | Unauthorized                                      | | `1000404` | Resource not found                                | | `1000429` | Too many requests                                 | | `1003001` | Product is not available in your account          | | `1003002` | Requested product amount is out of range          | | `1003003` | Requested product unit is invalid                 | | `1003101` | Benefits not defined for available products       | | `1003201` | Promotion not found                               | | `1003301` | Campaign not found                                | | `1005003` | Credit party mobile number is invalid             | | `1005004` | Service not found                                 | | `1005005` | Country not found                                 | | `1005006` | Operator not found                                | | `1005503` | Sender mobile number is invalid                   | | `1006001` | Insufficient balance                              | | `1006003` | Debit party mobile number is invalid              | | `1006009` | Account balance not found                         | | `1006503` | Beneficiary mobile number is invalid              | | `1007001` | Transaction external ID has already been used     | | `1007002` | Transaction has already been confirmed            | | `1007004` | Transaction can no longer be confirmed            | | `1007005` | Transaction has already been cancelled            | | `1007007` | Transaction can no longer be cancelled            | | `1007500` | Method not supported by operator                  | | `1008004` | Transaction not found                             | | `1009001` | Unexpected error, please contact our support team |  ### Transaction Status  | Class       | Status                                          | Description                                                            | | ---         | ---                                             | ---                                                                    | | `CREATED`   | `CREATED`                                       | Created                                                                | | `CONFIRMED` | `CONFIRMED`                                     | Confirmed                                                              | | `REJECTED`  | `REJECTED`                                      | Rejected                                                               | | `REJECTED`  | `REJECTED-INVALID-CREDIT-PARTY`                 | Rejected - Credit party is invalid                                     | | `REJECTED`  | `REJECTED-BARRED-CREDIT-PARTY`                  | Rejected - Credit party is barred                                      | | `REJECTED`  | `REJECTED-INELIGIBLE-CREDIT-PARTY`              | Rejected - Credit party is ineligible for chosen product               | | `REJECTED`  | `REJECTED-INVALID-DEBIT-PARTY`                  | Rejected - Debit party is invalid                                      | | `REJECTED`  | `REJECTED-BARRED-DEBIT-PARTY`                   | Rejected - Debit party is barred                                       | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Rejected - Limitations on credit party cumulative transaction amount   | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Rejected - Limitations on credit party cumulative transaction quantity | | `REJECTED`  | `REJECTED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Rejected - Operator currently unavailable                              | | `REJECTED`  | `REJECTED-INSUFFICIENT-BALANCE`                 | Rejected - Insufficient balance                                        | | `CANCELLED` | `CANCELLED`                                     | Cancelled                                                              | | `SUBMITTED` | `SUBMITTED`                                     | Submitted                                                              | | `COMPLETED` | `COMPLETED`                                     | Completed                                                              | | `REVERSED`  | `REVERSED`                                      | Reversed                                                               | | `DECLINED`  | `DECLINED`                                      | Declined (no additional information available)                         | | `DECLINED`  | `DECLINED-INVALID-CREDIT-PARTY`                 | Declined - Credit party is invalid                                     | | `DECLINED`  | `DECLINED-BARRED-CREDIT-PARTY`                  | Declined - Credit party is barred                                      | | `DECLINED`  | `DECLINED-INELIGIBLE-CREDIT-PARTY`              | Declined - Credit party is ineligible for chosen product               | | `DECLINED`  | `DECLINED-INVALID-DEBIT-PARTY`                  | Declined - Debit party is invalid                                      | | `DECLINED`  | `DECLINED-BARRED-DEBIT-PARTY`                   | Declined - Debit party is barred                                       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-AMOUNT`       | Declined - Limitations on operator cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Declined - Limitations on credit party cumulative transaction amount   | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-AMOUNT`       | Declined - Limitations on customer cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-QUANTITY`     | Declined - Limitations on operator cumulative transaction quantity     | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Declined - Limitations on credit party cumulative transaction quantity | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-QUANTITY`     | Declined - Limitations on customer cumulative transaction quantity     | | `DECLINED`  | `DECLINED-DUPLICATED-TRANSACTION`               | Declined - Duplicated transaction                                      | | `DECLINED`  | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Declined - Operator currently unavailable                              |  `REJECTED` and `DECLINED` status classes both denote unsuccessful transactions. The primary distinction between these two relates to the party that determined the failure:   * `REJECTED` are issued by the DVS platform based on various business rules (e.g. insufficient balance, limitations, etc)   * `DECLINED` are issued by the operators  Separately, it is recommended to define application logic based on **classes**, while additional distinction and/or insight are reflected on the actual **status**.  ## Pagination  API resources supporting bulk fetches via \"list\" API methods will be returned in a paginated fashion.  ### Input Parameters  | Field      | Required | Type    | Description                                      | | ---        | ---      | ---     | ---                                              | | `page`     | No       | Integer | Page number                                      | | `per_page` | No       | Integer | Number of results per page (default 50, max 100) |  ### Output Headers  | Field           | Description                   | | ---             | ---                           | | `X-Total`       | Total number of records       | | `X-Total-Pages` | Total number of pages         | | `X-Per-Page`    | Number of records per page    | | `X-Page`        | Current page number           | | `X-Next-Page`   | Next page number (if any)     | | `X-Prev-Page`   | Previous page number (if any) |  ## Rate Limiting  The API endpoints have rate limiting in place to protect our service from excessive number of requests.  If the limit is reached, an [HTTP error 429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) will be returned by the server.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dvsapi

import (
	"encoding/json"
)

// ProductFixedValueRecharge struct for ProductFixedValueRecharge
type ProductFixedValueRecharge struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Service Service `json:"service"`
	Operator OperatorsGet200ResponseInner `json:"operator"`
	Regions []ServiceRegion `json:"regions"`
	Type string `json:"type"`
	Validity NullableProductFixedValueRechargeAllOfValidity `json:"validity"`
	RequiredDebitPartyIdentifierFields [][]string `json:"required_debit_party_identifier_fields"`
	RequiredCreditPartyIdentifierFields [][]string `json:"required_credit_party_identifier_fields"`
	RequiredSenderFields [][]string `json:"required_sender_fields"`
	RequiredBeneficiaryFields [][]string `json:"required_beneficiary_fields"`
	RequiredStatementIdentifierFields [][]string `json:"required_statement_identifier_fields"`
	AvailabilityZones []AvailabilityZones `json:"availability_zones"`
	Source ProductFixedValueRechargeAllOfSource `json:"source"`
	Destination ProductFixedValueRechargeAllOfSource `json:"destination"`
	Prices ProductFixedValueRechargeAllOfPrices `json:"prices"`
	Rates Rates `json:"rates"`
	Benefits []FixedBenefit `json:"benefits"`
	Promotions []ProductPromotion `json:"promotions"`
}

// NewProductFixedValueRecharge instantiates a new ProductFixedValueRecharge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductFixedValueRecharge(id int32, name string, description string, tags []string, service Service, operator OperatorsGet200ResponseInner, regions []ServiceRegion, type_ string, validity NullableProductFixedValueRechargeAllOfValidity, requiredDebitPartyIdentifierFields [][]string, requiredCreditPartyIdentifierFields [][]string, requiredSenderFields [][]string, requiredBeneficiaryFields [][]string, requiredStatementIdentifierFields [][]string, availabilityZones []AvailabilityZones, source ProductFixedValueRechargeAllOfSource, destination ProductFixedValueRechargeAllOfSource, prices ProductFixedValueRechargeAllOfPrices, rates Rates, benefits []FixedBenefit, promotions []ProductPromotion) *ProductFixedValueRecharge {
	this := ProductFixedValueRecharge{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Tags = tags
	this.Service = service
	this.Operator = operator
	this.Regions = regions
	this.Type = type_
	this.Validity = validity
	this.RequiredDebitPartyIdentifierFields = requiredDebitPartyIdentifierFields
	this.RequiredCreditPartyIdentifierFields = requiredCreditPartyIdentifierFields
	this.RequiredSenderFields = requiredSenderFields
	this.RequiredBeneficiaryFields = requiredBeneficiaryFields
	this.RequiredStatementIdentifierFields = requiredStatementIdentifierFields
	this.AvailabilityZones = availabilityZones
	this.Source = source
	this.Destination = destination
	this.Prices = prices
	this.Rates = rates
	this.Benefits = benefits
	this.Promotions = promotions
	return &this
}

// NewProductFixedValueRechargeWithDefaults instantiates a new ProductFixedValueRecharge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductFixedValueRechargeWithDefaults() *ProductFixedValueRecharge {
	this := ProductFixedValueRecharge{}
	return &this
}

// GetId returns the Id field value
func (o *ProductFixedValueRecharge) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductFixedValueRecharge) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProductFixedValueRecharge) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductFixedValueRecharge) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ProductFixedValueRecharge) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProductFixedValueRecharge) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ProductFixedValueRecharge) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ProductFixedValueRecharge) SetTags(v []string) {
	o.Tags = v
}

// GetService returns the Service field value
func (o *ProductFixedValueRecharge) GetService() Service {
	if o == nil {
		var ret Service
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetServiceOk() (*Service, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ProductFixedValueRecharge) SetService(v Service) {
	o.Service = v
}

// GetOperator returns the Operator field value
func (o *ProductFixedValueRecharge) GetOperator() OperatorsGet200ResponseInner {
	if o == nil {
		var ret OperatorsGet200ResponseInner
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetOperatorOk() (*OperatorsGet200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ProductFixedValueRecharge) SetOperator(v OperatorsGet200ResponseInner) {
	o.Operator = v
}

// GetRegions returns the Regions field value
// If the value is explicit nil, the zero value for []ServiceRegion will be returned
func (o *ProductFixedValueRecharge) GetRegions() []ServiceRegion {
	if o == nil {
		var ret []ServiceRegion
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetRegionsOk() ([]ServiceRegion, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// SetRegions sets field value
func (o *ProductFixedValueRecharge) SetRegions(v []ServiceRegion) {
	o.Regions = v
}

// GetType returns the Type field value
func (o *ProductFixedValueRecharge) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProductFixedValueRecharge) SetType(v string) {
	o.Type = v
}

// GetValidity returns the Validity field value
// If the value is explicit nil, the zero value for ProductFixedValueRechargeAllOfValidity will be returned
func (o *ProductFixedValueRecharge) GetValidity() ProductFixedValueRechargeAllOfValidity {
	if o == nil || o.Validity.Get() == nil {
		var ret ProductFixedValueRechargeAllOfValidity
		return ret
	}

	return *o.Validity.Get()
}

// GetValidityOk returns a tuple with the Validity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetValidityOk() (*ProductFixedValueRechargeAllOfValidity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Validity.Get(), o.Validity.IsSet()
}

// SetValidity sets field value
func (o *ProductFixedValueRecharge) SetValidity(v ProductFixedValueRechargeAllOfValidity) {
	o.Validity.Set(&v)
}

// GetRequiredDebitPartyIdentifierFields returns the RequiredDebitPartyIdentifierFields field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *ProductFixedValueRecharge) GetRequiredDebitPartyIdentifierFields() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.RequiredDebitPartyIdentifierFields
}

// GetRequiredDebitPartyIdentifierFieldsOk returns a tuple with the RequiredDebitPartyIdentifierFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetRequiredDebitPartyIdentifierFieldsOk() ([][]string, bool) {
	if o == nil || o.RequiredDebitPartyIdentifierFields == nil {
		return nil, false
	}
	return o.RequiredDebitPartyIdentifierFields, true
}

// SetRequiredDebitPartyIdentifierFields sets field value
func (o *ProductFixedValueRecharge) SetRequiredDebitPartyIdentifierFields(v [][]string) {
	o.RequiredDebitPartyIdentifierFields = v
}

// GetRequiredCreditPartyIdentifierFields returns the RequiredCreditPartyIdentifierFields field value
func (o *ProductFixedValueRecharge) GetRequiredCreditPartyIdentifierFields() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.RequiredCreditPartyIdentifierFields
}

// GetRequiredCreditPartyIdentifierFieldsOk returns a tuple with the RequiredCreditPartyIdentifierFields field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetRequiredCreditPartyIdentifierFieldsOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredCreditPartyIdentifierFields, true
}

// SetRequiredCreditPartyIdentifierFields sets field value
func (o *ProductFixedValueRecharge) SetRequiredCreditPartyIdentifierFields(v [][]string) {
	o.RequiredCreditPartyIdentifierFields = v
}

// GetRequiredSenderFields returns the RequiredSenderFields field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *ProductFixedValueRecharge) GetRequiredSenderFields() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.RequiredSenderFields
}

// GetRequiredSenderFieldsOk returns a tuple with the RequiredSenderFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetRequiredSenderFieldsOk() ([][]string, bool) {
	if o == nil || o.RequiredSenderFields == nil {
		return nil, false
	}
	return o.RequiredSenderFields, true
}

// SetRequiredSenderFields sets field value
func (o *ProductFixedValueRecharge) SetRequiredSenderFields(v [][]string) {
	o.RequiredSenderFields = v
}

// GetRequiredBeneficiaryFields returns the RequiredBeneficiaryFields field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *ProductFixedValueRecharge) GetRequiredBeneficiaryFields() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.RequiredBeneficiaryFields
}

// GetRequiredBeneficiaryFieldsOk returns a tuple with the RequiredBeneficiaryFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetRequiredBeneficiaryFieldsOk() ([][]string, bool) {
	if o == nil || o.RequiredBeneficiaryFields == nil {
		return nil, false
	}
	return o.RequiredBeneficiaryFields, true
}

// SetRequiredBeneficiaryFields sets field value
func (o *ProductFixedValueRecharge) SetRequiredBeneficiaryFields(v [][]string) {
	o.RequiredBeneficiaryFields = v
}

// GetRequiredStatementIdentifierFields returns the RequiredStatementIdentifierFields field value
// If the value is explicit nil, the zero value for [][]string will be returned
func (o *ProductFixedValueRecharge) GetRequiredStatementIdentifierFields() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.RequiredStatementIdentifierFields
}

// GetRequiredStatementIdentifierFieldsOk returns a tuple with the RequiredStatementIdentifierFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetRequiredStatementIdentifierFieldsOk() ([][]string, bool) {
	if o == nil || o.RequiredStatementIdentifierFields == nil {
		return nil, false
	}
	return o.RequiredStatementIdentifierFields, true
}

// SetRequiredStatementIdentifierFields sets field value
func (o *ProductFixedValueRecharge) SetRequiredStatementIdentifierFields(v [][]string) {
	o.RequiredStatementIdentifierFields = v
}

// GetAvailabilityZones returns the AvailabilityZones field value
func (o *ProductFixedValueRecharge) GetAvailabilityZones() []AvailabilityZones {
	if o == nil {
		var ret []AvailabilityZones
		return ret
	}

	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetAvailabilityZonesOk() ([]AvailabilityZones, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZones, true
}

// SetAvailabilityZones sets field value
func (o *ProductFixedValueRecharge) SetAvailabilityZones(v []AvailabilityZones) {
	o.AvailabilityZones = v
}

// GetSource returns the Source field value
func (o *ProductFixedValueRecharge) GetSource() ProductFixedValueRechargeAllOfSource {
	if o == nil {
		var ret ProductFixedValueRechargeAllOfSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetSourceOk() (*ProductFixedValueRechargeAllOfSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ProductFixedValueRecharge) SetSource(v ProductFixedValueRechargeAllOfSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *ProductFixedValueRecharge) GetDestination() ProductFixedValueRechargeAllOfSource {
	if o == nil {
		var ret ProductFixedValueRechargeAllOfSource
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetDestinationOk() (*ProductFixedValueRechargeAllOfSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *ProductFixedValueRecharge) SetDestination(v ProductFixedValueRechargeAllOfSource) {
	o.Destination = v
}

// GetPrices returns the Prices field value
func (o *ProductFixedValueRecharge) GetPrices() ProductFixedValueRechargeAllOfPrices {
	if o == nil {
		var ret ProductFixedValueRechargeAllOfPrices
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetPricesOk() (*ProductFixedValueRechargeAllOfPrices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prices, true
}

// SetPrices sets field value
func (o *ProductFixedValueRecharge) SetPrices(v ProductFixedValueRechargeAllOfPrices) {
	o.Prices = v
}

// GetRates returns the Rates field value
func (o *ProductFixedValueRecharge) GetRates() Rates {
	if o == nil {
		var ret Rates
		return ret
	}

	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *ProductFixedValueRecharge) GetRatesOk() (*Rates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rates, true
}

// SetRates sets field value
func (o *ProductFixedValueRecharge) SetRates(v Rates) {
	o.Rates = v
}

// GetBenefits returns the Benefits field value
// If the value is explicit nil, the zero value for []FixedBenefit will be returned
func (o *ProductFixedValueRecharge) GetBenefits() []FixedBenefit {
	if o == nil {
		var ret []FixedBenefit
		return ret
	}

	return o.Benefits
}

// GetBenefitsOk returns a tuple with the Benefits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetBenefitsOk() ([]FixedBenefit, bool) {
	if o == nil || o.Benefits == nil {
		return nil, false
	}
	return o.Benefits, true
}

// SetBenefits sets field value
func (o *ProductFixedValueRecharge) SetBenefits(v []FixedBenefit) {
	o.Benefits = v
}

// GetPromotions returns the Promotions field value
// If the value is explicit nil, the zero value for []ProductPromotion will be returned
func (o *ProductFixedValueRecharge) GetPromotions() []ProductPromotion {
	if o == nil {
		var ret []ProductPromotion
		return ret
	}

	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductFixedValueRecharge) GetPromotionsOk() ([]ProductPromotion, bool) {
	if o == nil || o.Promotions == nil {
		return nil, false
	}
	return o.Promotions, true
}

// SetPromotions sets field value
func (o *ProductFixedValueRecharge) SetPromotions(v []ProductPromotion) {
	o.Promotions = v
}

func (o ProductFixedValueRecharge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["validity"] = o.Validity.Get()
	}
	if o.RequiredDebitPartyIdentifierFields != nil {
		toSerialize["required_debit_party_identifier_fields"] = o.RequiredDebitPartyIdentifierFields
	}
	if true {
		toSerialize["required_credit_party_identifier_fields"] = o.RequiredCreditPartyIdentifierFields
	}
	if o.RequiredSenderFields != nil {
		toSerialize["required_sender_fields"] = o.RequiredSenderFields
	}
	if o.RequiredBeneficiaryFields != nil {
		toSerialize["required_beneficiary_fields"] = o.RequiredBeneficiaryFields
	}
	if o.RequiredStatementIdentifierFields != nil {
		toSerialize["required_statement_identifier_fields"] = o.RequiredStatementIdentifierFields
	}
	if true {
		toSerialize["availability_zones"] = o.AvailabilityZones
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["prices"] = o.Prices
	}
	if true {
		toSerialize["rates"] = o.Rates
	}
	if o.Benefits != nil {
		toSerialize["benefits"] = o.Benefits
	}
	if o.Promotions != nil {
		toSerialize["promotions"] = o.Promotions
	}
	return json.Marshal(toSerialize)
}

type NullableProductFixedValueRecharge struct {
	value *ProductFixedValueRecharge
	isSet bool
}

func (v NullableProductFixedValueRecharge) Get() *ProductFixedValueRecharge {
	return v.value
}

func (v *NullableProductFixedValueRecharge) Set(val *ProductFixedValueRecharge) {
	v.value = val
	v.isSet = true
}

func (v NullableProductFixedValueRecharge) IsSet() bool {
	return v.isSet
}

func (v *NullableProductFixedValueRecharge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductFixedValueRecharge(val *ProductFixedValueRecharge) *NullableProductFixedValueRecharge {
	return &NullableProductFixedValueRecharge{value: val, isSet: true}
}

func (v NullableProductFixedValueRecharge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductFixedValueRecharge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


