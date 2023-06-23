/*
Digital Value Services API

# Overview  Welcome to the Digital Value Services (DVS) API reference.  This API serves as the primary gateway to facilitate digital value transfers through [DT One](https://www.dtone.com), a leading global network covering more than 160 countries and 550 mobile operators.  The Digital Value Services API is organized according to [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, using [JSON](https://en.wikipedia.org/wiki/JSON) as format for data interchange, and provides the following services:   - [Discovery Services](#tag/Services)   - [Transaction Services](#tag/Transactions)   - [Account Services](#tag/Balances)   - [Look-Up Services](#tag/Mobile-Number)  If you have any questions and/or suggestions related to our API, please do not hesitate to send an email to the [DVS API support team](mailto:dvs-api-support@dtone.com).  ## Integration Libraries  Officially supported [SDK](https://en.wikipedia.org/wiki/Software_development_kit)s are available for the following languages:  * [Java](https://github.com/dtone/dtone-dvs-api-java-client) * [Node.js](https://www.npmjs.com/package/@dtone/dvs)  These SDKs offer an accelerated path to developing your applications as an alternative to accessing the REST API directly.  Separately, we would love to hear from you! If you have any questions and/or suggestions related to our SDKs, please do not hesitate to create corresponding [GitHub issues](https://guides.github.com/features/issues/) or send an email to the [DVS Open Source team](mailto:opensource@dtone.com).  ## Sandbox  A sandbox environment is available for testing integrations with the DVS API. It is available at [https://preprod-dvs-api.dtone.com/v1/](https://preprod-dvs-api.dtone.com/v1/).  You can generate sandbox API keys from your [DT Shop](https://dtshop.dtone.com/account?tab=developer) account, under the **Pre-Production API Keys** section.  All transactions on the sandbox environment are simulated: no real transaction goes through. To simulate different responses, the last three digits of the `credit_party_identifier` (i.e. `mobile_number` or `account_number`, depending on the `required_credit_party_identifier_fields` of a given [Product](/#tag/Products)) will have to be replaced with one of the following suffixes:  | Suffix              | Transaction Status                        | Example       | | ---                 | ---                                       | ---           | | `100`, `200`, `300` | `COMPLETED` (PIN-less)                    | `+6595123100` | | `101`, `201`, `301` | `COMPLETED` (PIN-based)                   | `+6595123201` | | `102`, `202`, `302` | `DECLINED-INVALID-CREDIT-PARTY`           | `+6595123102` | | `103`, `203`, `303` | `DECLINED-BARRED-CREDIT-PARTY`            | `+6595123103` | | `104`, `204`, `304` | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE` | `+6595123204` | | `105`, `205`, `305` | `DECLINED-DUPLICATED-TRANSACTION`         | `+6595123105` | | `106`, `206`, `306` | `DECLINED`                                | `+6595123206` | | `107`, `207`, `307` | `DECLINED-EXCEPTION`                      | `+6595123107` |  The different suffixes for a given transaction status can be used to simulate delays, as follows:   * `10X` suffix will take at least 3 seconds to finish   * `20X` suffix will take at least 20 seconds to finish   * `30X` suffix will take at least 5 minutes to finish  Please note that there are products that do not require any credit party identifier such as Gift Cards. For these products, the simulated transaction will always be in `COMPLETED` status.  ## Versioning  Endpoints of the API are prefixed with a corresponding version number.  This is done to provide complete isolation between implementations and to ensure that subsequent major changes to the API will never affect existing integrations.  When a new version of the API is available and you are keen to upgrade, testing in the sandbox environment to ensure that everything works well with your implementation before switching to the production environment comes highly recommended.  Feel free to [contact us](mailto:dvs-api-support@dtone.com) should you have any questions and/or are in need of assistance during your tests.  ## Transactions  The main purpose of this API is to deliver value (e.g. mobile airtime top-up, data bundles, etc.) to a beneficiary. This is what we call a \"transaction\".  During the course of a transfer, a transaction undergoes various status changes (or transitions) as illustrated below.  ![transaction states](/images/transaction_states.png)  As changes in transaction status occur, updates are sent in real-time when a callback URL is provided. In conjunction, transaction status can be queried through one of two means: via the returned `id` or a provided `external_id`.  The latter serves as your unique reference and provides a utility to retrieve transaction details when exceptions occur, such as when the supposed API response was not received, as an example.  ## Balances  Transactions can be created through the platform as long as there is enough balance available in your account. A given balance is composed of the following:  | Balance   | Description                                              | | ---       | ---                                                      | | Available | Balance amount available for use                         | | Holding   | Amount being held while transactions are being processed |  As a given transaction goes through various changes in status as outlined [here](#section/Overview/Transactions), corresponding balance movements will be made. The following table illustrates the relationship between transaction status and balance movements:  | Transaction Status                  | Balance Movement | Description                                                | | ---                                 | ---              | ---                                                        | | `CREATED`                           | Authorize        | Transfer wholesale price and fee from available to holding | | `CANCELLED`, `REJECTED`, `DECLINED` | Void             | Amount authorized in holding moves back to available       | | `COMPLETED`                         | Capture          | Amount authorized in holding is captured, i.e. debited     | | `REVERSED`                          | Reverse          | Debited amount is reversed back into available             |  ## Flow  Once a product has been selected through one of the [discovery methods](/#tag/Services) provided by the API, the actual transfer (i.e. transaction) can be performed in either one of the following modes:   - Asynchronous (recommended)   - Synchronous  Each mode is accessible via a specific endpoint.  As soon as a transaction is confirmed, the transfer order will be sent to the operator for immediate processing. During this time, the transaction will remain in a `CONFIRMED` status until the final status is received from the operator.  ### Asynchronous Mode  When a transaction is created and confirmed in an asynchronous fashion, the HTTP connection won't have to be kept open. This preserves system resources on your applications. As such, performing transactions **asynchronously** is **recommended**.  ### Synchronous Mode  When a transaction is created and confirmed in a synchronous fashion, the HTTP connection will be kept open in an attempt to capture the final status from the receiving operator so it can be returned in the API response. The processing time usually takes just a few seconds. However, with some receiving operators, it may take longer.  Our system will keep the HTTP connection open for up to 180 seconds (this value can be configured upon request) and will return a status before closing this connection. This status can be in one of the final status (e.g. `COMPLETED`, `DECLINED`) or not (e.g. `SUBMITTED`). In the latter case, this denotes the transaction is still being processed by the receiving operator.  **Note:** your application does not have to wait for the connection to close, it can listen for a shorter period of time and query the final status later on (refer to the \"Final Status\" section below for more details).  ### Final Status  Regardless of the processing mode, the application should be designed to capture the final status of a transaction. This can be done through one of the following means:   - Checking the status of a specific transaction via the corresponding API method (\"pull\" mechanism)   - Configuring a callback URL passed in the request when creating a transaction (\"push\" mechanism)  ## Callbacks  As a transaction is being processed, changes in status will be notified in real-time if a callback URL was provided. Even though one callback per transaction is expected (i.e. change to either `COMPLETED` or `DECLINED`), a manual reversal from the [DT One](https://dtone.com/) team, which may happen in very rare occasions, will also trigger a callback to inform your application of a change in transaction status to `REVERSED`.  This callback endpoint must be implemented by the sending partner, which should expect an HTTP `POST` request containing a transaction object represented in [JSON](https://en.wikipedia.org/wiki/JSON). As callbacks will be sent from the [DT One](https://dtone.com/) servers, these endpoints will have to be publicly-accessible in most cases. During development, a service such as [ngrok](https://ngrok.com/) can be used to expose local servers to the internet.  Upon successful receipt of data, the callback endpoint should respond with an HTTP `2XX` status. In the event that the platform did not receive a successful status, callback notifications will be retried several times, beyond which, the transaction status will have to be queried through the API.  ## Status and Errors  ### HTTP Status Codes  [DT One](https://dtone.com/) uses standard HTTP response codes to indicate whether an API request was successful or not.  | Status | Description                                        | | ---    | ---                                                | | `200`  | OK                                                 | | `201`  | Created: Resource created                          | | `202`  | Accepted: Request has been accepted for processing | | `400`  | Bad Request: Request was malformed                 | | `401`  | Unauthorized: Credentials missing or invalid       | | `404`  | Not Found: Resource doesn't exist                  | | `429`  | Too Many Requests                                  | | `500`  | Server Error: Error occurred on DT One             | | `503`  | Service Unavailable                                |  ### API Error Codes  | Code      | Description                                       | HTTP Status | | ---       | ---                                               | ---         | | `1000400` | Bad Request                                       | `400`       | | `1000401` | Unauthorized                                      | `401`       | | `1000404` | Resource not found                                | `404`       | | `1000429` | Too many requests                                 | `429`       | | `1003001` | Product is not available in your account          | `404`       | | `1003002` | Requested product amount is out of range          | `400`       | | `1003003` | Requested product unit is invalid                 | `400`       | | `1003101` | Benefits not defined for available products       | `404`       | | `1003201` | Promotion not found                               | `404`       | | `1003301` | Campaign not found                                | `404`       | | `1005003` | Credit party mobile number is invalid             | `400`       | | `1005004` | Service not found                                 | `404`       | | `1005005` | Country not found                                 | `404`       | | `1005006` | Operator not found                                | `404`       | | `1005503` | Sender mobile number is invalid                   | `400`       | | `1006001` | Insufficient balance                              | `400`       | | `1006003` | Debit party mobile number is invalid              | `400`       | | `1006009` | Account balance not found                         | `404`       | | `1006503` | Beneficiary mobile number is invalid              | `400`       | | `1007001` | Transaction external ID has already been used     | `400`       | | `1007002` | Transaction has already been confirmed            | `400`       | | `1007004` | Transaction can no longer be confirmed            | `400`       | | `1007005` | Transaction has already been cancelled            | `400`       | | `1007007` | Transaction can no longer be cancelled            | `400`       | | `1007500` | Method not supported by operator                  | `400`       | | `1008004` | Transaction not found                             | `404`       | | `1009001` | Unexpected error, please contact our support team | `500`       | | `1009503` | Service unavailable, please retry later           | `503`       |  ### Transaction Status  | Class       | Status                                          | Description                                                            | | ---         | ---                                             | ---                                                                    | | `CREATED`   | `CREATED`                                       | Created                                                                | | `CONFIRMED` | `CONFIRMED`                                     | Confirmed                                                              | | `REJECTED`  | `REJECTED`                                      | Rejected                                                               | | `REJECTED`  | `REJECTED-INVALID-CREDIT-PARTY`                 | Rejected - Credit party is invalid                                     | | `REJECTED`  | `REJECTED-BARRED-CREDIT-PARTY`                  | Rejected - Credit party is barred                                      | | `REJECTED`  | `REJECTED-INELIGIBLE-CREDIT-PARTY`              | Rejected - Credit party is ineligible for chosen product               | | `REJECTED`  | `REJECTED-INVALID-DEBIT-PARTY`                  | Rejected - Debit party is invalid                                      | | `REJECTED`  | `REJECTED-BARRED-DEBIT-PARTY`                   | Rejected - Debit party is barred                                       | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Rejected - Limitations on credit party cumulative transaction amount   | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Rejected - Limitations on credit party cumulative transaction quantity | | `REJECTED`  | `REJECTED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Rejected - Operator currently unavailable                              | | `REJECTED`  | `REJECTED-INSUFFICIENT-BALANCE`                 | Rejected - Insufficient balance                                        | | `CANCELLED` | `CANCELLED`                                     | Cancelled                                                              | | `SUBMITTED` | `SUBMITTED`                                     | Submitted                                                              | | `COMPLETED` | `COMPLETED`                                     | Completed                                                              | | `REVERSED`  | `REVERSED`                                      | Reversed                                                               | | `DECLINED`  | `DECLINED`                                      | Declined (no additional information available)                         | | `DECLINED`  | `DECLINED-INVALID-CREDIT-PARTY`                 | Declined - Credit party is invalid                                     | | `DECLINED`  | `DECLINED-BARRED-CREDIT-PARTY`                  | Declined - Credit party is barred                                      | | `DECLINED`  | `DECLINED-INELIGIBLE-CREDIT-PARTY`              | Declined - Credit party is ineligible for chosen product               | | `DECLINED`  | `DECLINED-INVALID-DEBIT-PARTY`                  | Declined - Debit party is invalid                                      | | `DECLINED`  | `DECLINED-BARRED-DEBIT-PARTY`                   | Declined - Debit party is barred                                       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-AMOUNT`       | Declined - Limitations on operator cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Declined - Limitations on credit party cumulative transaction amount   | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-AMOUNT`       | Declined - Limitations on customer cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-QUANTITY`     | Declined - Limitations on operator cumulative transaction quantity     | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Declined - Limitations on credit party cumulative transaction quantity | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-QUANTITY`     | Declined - Limitations on customer cumulative transaction quantity     | | `DECLINED`  | `DECLINED-DUPLICATED-TRANSACTION`               | Declined - Duplicated transaction                                      | | `DECLINED`  | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Declined - Operator currently unavailable                              |  `REJECTED` and `DECLINED` status classes both denote unsuccessful transactions. The primary distinction between these two relates to the party that determined the failure:   * `REJECTED` are issued by the DVS platform based on various business rules (e.g. insufficient balance, limitations, etc)   * `DECLINED` are issued by the operators  Separately, it is recommended to define application logic based on **classes**, while additional distinction and/or insight are reflected on the actual **status**.  ## Pagination  API resources supporting bulk fetches via \"list\" API methods will be returned in a paginated fashion.  ### Input Parameters  | Field      | Required | Type    | Description                                      | | ---        | ---      | ---     | ---                                              | | `page`     | No       | Integer | Page number                                      | | `per_page` | No       | Integer | Number of results per page (default 50, max 100) |  ### Output Headers  | Field           | Description                   | | ---             | ---                           | | `X-Total`       | Total number of records       | | `X-Total-Pages` | Total number of pages         | | `X-Per-Page`    | Number of records per page    | | `X-Page`        | Current page number           | | `X-Next-Page`   | Next page number (if any)     | | `X-Prev-Page`   | Previous page number (if any) |  ## Rate Limiting  The API endpoints have rate limiting in place to protect our service from excessive number of requests.  If the limit is reached, an [HTTP error 429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) will be returned by the server.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dvsapi

import (
	"encoding/json"
	"time"
)

// checks if the PostTransactionAsyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTransactionAsyncRequest{}

// PostTransactionAsyncRequest struct for PostTransactionAsyncRequest
type PostTransactionAsyncRequest struct {
	Id                         *int64                                  `json:"id,omitempty"`
	ExternalId                 string                                  `json:"external_id"`
	CreationDate               *time.Time                              `json:"creation_date,omitempty"`
	ConfirmationExpirationDate *time.Time                              `json:"confirmation_expiration_date,omitempty"`
	ConfirmationDate           *time.Time                              `json:"confirmation_date,omitempty"`
	ProductId                  int32                                   `json:"product_id"`
	CalculationMode            *CalculationModes                       `json:"calculation_mode,omitempty"`
	Source                     NullableTransactionSource               `json:"source,omitempty"`
	Destination                NullableTransactionDestination          `json:"destination,omitempty"`
	AutoConfirm                *bool                                   `json:"auto_confirm,omitempty"`
	Status                     *TransactionStatus                      `json:"status,omitempty"`
	OperatorReference          NullableString                          `json:"operator_reference,omitempty"`
	Pin                        NullableTransactionPin                  `json:"pin,omitempty"`
	Product                    *TransactionProduct                     `json:"product,omitempty"`
	Prices                     *TransactionPrices                      `json:"prices,omitempty"`
	Rates                      *TransactionRates                       `json:"rates,omitempty"`
	Benefits                   []TransactionBenefitsInner              `json:"benefits,omitempty"`
	Promotions                 []ProductPromotion                      `json:"promotions,omitempty"`
	RequestedValues            NullableTransactionRequestedValues      `json:"requested_values,omitempty"`
	AdjustedValues             NullableTransactionRequestedValues      `json:"adjusted_values,omitempty"`
	Sender                     NullableTransactionSender               `json:"sender,omitempty"`
	Beneficiary                NullableTransactionSender               `json:"beneficiary,omitempty"`
	DebitPartyIdentifier       NullableTransactionDebitPartyIdentifier `json:"debit_party_identifier,omitempty"`
	CreditPartyIdentifier      *TransactionCreditPartyIdentifier       `json:"credit_party_identifier,omitempty"`
	StatementIdentifier        *TransactionStatementIdentifier         `json:"statement_identifier,omitempty"`
	AdditionalIdentifier       *TransactionAdditionalIdentifier        `json:"additional_identifier,omitempty"`
	CallbackUrl                *string                                 `json:"callback_url,omitempty"`
}

// NewPostTransactionAsyncRequest instantiates a new PostTransactionAsyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTransactionAsyncRequest(externalId string, productId int32) *PostTransactionAsyncRequest {
	this := PostTransactionAsyncRequest{}
	this.ExternalId = externalId
	this.ProductId = productId
	var autoConfirm bool = false
	this.AutoConfirm = &autoConfirm
	return &this
}

// NewPostTransactionAsyncRequestWithDefaults instantiates a new PostTransactionAsyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTransactionAsyncRequestWithDefaults() *PostTransactionAsyncRequest {
	this := PostTransactionAsyncRequest{}
	var autoConfirm bool = false
	this.AutoConfirm = &autoConfirm
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PostTransactionAsyncRequest) SetId(v int64) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value
func (o *PostTransactionAsyncRequest) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *PostTransactionAsyncRequest) SetExternalId(v string) {
	o.ExternalId = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *PostTransactionAsyncRequest) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetConfirmationExpirationDate returns the ConfirmationExpirationDate field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetConfirmationExpirationDate() time.Time {
	if o == nil || IsNil(o.ConfirmationExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ConfirmationExpirationDate
}

// GetConfirmationExpirationDateOk returns a tuple with the ConfirmationExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetConfirmationExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConfirmationExpirationDate) {
		return nil, false
	}
	return o.ConfirmationExpirationDate, true
}

// HasConfirmationExpirationDate returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasConfirmationExpirationDate() bool {
	if o != nil && !IsNil(o.ConfirmationExpirationDate) {
		return true
	}

	return false
}

// SetConfirmationExpirationDate gets a reference to the given time.Time and assigns it to the ConfirmationExpirationDate field.
func (o *PostTransactionAsyncRequest) SetConfirmationExpirationDate(v time.Time) {
	o.ConfirmationExpirationDate = &v
}

// GetConfirmationDate returns the ConfirmationDate field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetConfirmationDate() time.Time {
	if o == nil || IsNil(o.ConfirmationDate) {
		var ret time.Time
		return ret
	}
	return *o.ConfirmationDate
}

// GetConfirmationDateOk returns a tuple with the ConfirmationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetConfirmationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConfirmationDate) {
		return nil, false
	}
	return o.ConfirmationDate, true
}

// HasConfirmationDate returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasConfirmationDate() bool {
	if o != nil && !IsNil(o.ConfirmationDate) {
		return true
	}

	return false
}

// SetConfirmationDate gets a reference to the given time.Time and assigns it to the ConfirmationDate field.
func (o *PostTransactionAsyncRequest) SetConfirmationDate(v time.Time) {
	o.ConfirmationDate = &v
}

// GetProductId returns the ProductId field value
func (o *PostTransactionAsyncRequest) GetProductId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *PostTransactionAsyncRequest) SetProductId(v int32) {
	o.ProductId = v
}

// GetCalculationMode returns the CalculationMode field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetCalculationMode() CalculationModes {
	if o == nil || IsNil(o.CalculationMode) {
		var ret CalculationModes
		return ret
	}
	return *o.CalculationMode
}

// GetCalculationModeOk returns a tuple with the CalculationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetCalculationModeOk() (*CalculationModes, bool) {
	if o == nil || IsNil(o.CalculationMode) {
		return nil, false
	}
	return o.CalculationMode, true
}

// HasCalculationMode returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasCalculationMode() bool {
	if o != nil && !IsNil(o.CalculationMode) {
		return true
	}

	return false
}

// SetCalculationMode gets a reference to the given CalculationModes and assigns it to the CalculationMode field.
func (o *PostTransactionAsyncRequest) SetCalculationMode(v CalculationModes) {
	o.CalculationMode = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetSource() TransactionSource {
	if o == nil || IsNil(o.Source.Get()) {
		var ret TransactionSource
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetSourceOk() (*TransactionSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableTransactionSource and assigns it to the Source field.
func (o *PostTransactionAsyncRequest) SetSource(v TransactionSource) {
	o.Source.Set(&v)
}

// SetSourceNil sets the value for Source to be an explicit nil
func (o *PostTransactionAsyncRequest) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetSource() {
	o.Source.Unset()
}

// GetDestination returns the Destination field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetDestination() TransactionDestination {
	if o == nil || IsNil(o.Destination.Get()) {
		var ret TransactionDestination
		return ret
	}
	return *o.Destination.Get()
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetDestinationOk() (*TransactionDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destination.Get(), o.Destination.IsSet()
}

// HasDestination returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasDestination() bool {
	if o != nil && o.Destination.IsSet() {
		return true
	}

	return false
}

// SetDestination gets a reference to the given NullableTransactionDestination and assigns it to the Destination field.
func (o *PostTransactionAsyncRequest) SetDestination(v TransactionDestination) {
	o.Destination.Set(&v)
}

// SetDestinationNil sets the value for Destination to be an explicit nil
func (o *PostTransactionAsyncRequest) SetDestinationNil() {
	o.Destination.Set(nil)
}

// UnsetDestination ensures that no value is present for Destination, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetDestination() {
	o.Destination.Unset()
}

// GetAutoConfirm returns the AutoConfirm field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetAutoConfirm() bool {
	if o == nil || IsNil(o.AutoConfirm) {
		var ret bool
		return ret
	}
	return *o.AutoConfirm
}

// GetAutoConfirmOk returns a tuple with the AutoConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetAutoConfirmOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoConfirm) {
		return nil, false
	}
	return o.AutoConfirm, true
}

// HasAutoConfirm returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasAutoConfirm() bool {
	if o != nil && !IsNil(o.AutoConfirm) {
		return true
	}

	return false
}

// SetAutoConfirm gets a reference to the given bool and assigns it to the AutoConfirm field.
func (o *PostTransactionAsyncRequest) SetAutoConfirm(v bool) {
	o.AutoConfirm = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetStatus() TransactionStatus {
	if o == nil || IsNil(o.Status) {
		var ret TransactionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetStatusOk() (*TransactionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TransactionStatus and assigns it to the Status field.
func (o *PostTransactionAsyncRequest) SetStatus(v TransactionStatus) {
	o.Status = &v
}

// GetOperatorReference returns the OperatorReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetOperatorReference() string {
	if o == nil || IsNil(o.OperatorReference.Get()) {
		var ret string
		return ret
	}
	return *o.OperatorReference.Get()
}

// GetOperatorReferenceOk returns a tuple with the OperatorReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetOperatorReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatorReference.Get(), o.OperatorReference.IsSet()
}

// HasOperatorReference returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasOperatorReference() bool {
	if o != nil && o.OperatorReference.IsSet() {
		return true
	}

	return false
}

// SetOperatorReference gets a reference to the given NullableString and assigns it to the OperatorReference field.
func (o *PostTransactionAsyncRequest) SetOperatorReference(v string) {
	o.OperatorReference.Set(&v)
}

// SetOperatorReferenceNil sets the value for OperatorReference to be an explicit nil
func (o *PostTransactionAsyncRequest) SetOperatorReferenceNil() {
	o.OperatorReference.Set(nil)
}

// UnsetOperatorReference ensures that no value is present for OperatorReference, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetOperatorReference() {
	o.OperatorReference.Unset()
}

// GetPin returns the Pin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetPin() TransactionPin {
	if o == nil || IsNil(o.Pin.Get()) {
		var ret TransactionPin
		return ret
	}
	return *o.Pin.Get()
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetPinOk() (*TransactionPin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pin.Get(), o.Pin.IsSet()
}

// HasPin returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasPin() bool {
	if o != nil && o.Pin.IsSet() {
		return true
	}

	return false
}

// SetPin gets a reference to the given NullableTransactionPin and assigns it to the Pin field.
func (o *PostTransactionAsyncRequest) SetPin(v TransactionPin) {
	o.Pin.Set(&v)
}

// SetPinNil sets the value for Pin to be an explicit nil
func (o *PostTransactionAsyncRequest) SetPinNil() {
	o.Pin.Set(nil)
}

// UnsetPin ensures that no value is present for Pin, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetPin() {
	o.Pin.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetProduct() TransactionProduct {
	if o == nil || IsNil(o.Product) {
		var ret TransactionProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetProductOk() (*TransactionProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given TransactionProduct and assigns it to the Product field.
func (o *PostTransactionAsyncRequest) SetProduct(v TransactionProduct) {
	o.Product = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetPrices() TransactionPrices {
	if o == nil || IsNil(o.Prices) {
		var ret TransactionPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetPricesOk() (*TransactionPrices, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given TransactionPrices and assigns it to the Prices field.
func (o *PostTransactionAsyncRequest) SetPrices(v TransactionPrices) {
	o.Prices = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetRates() TransactionRates {
	if o == nil || IsNil(o.Rates) {
		var ret TransactionRates
		return ret
	}
	return *o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetRatesOk() (*TransactionRates, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given TransactionRates and assigns it to the Rates field.
func (o *PostTransactionAsyncRequest) SetRates(v TransactionRates) {
	o.Rates = &v
}

// GetBenefits returns the Benefits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetBenefits() []TransactionBenefitsInner {
	if o == nil {
		var ret []TransactionBenefitsInner
		return ret
	}
	return o.Benefits
}

// GetBenefitsOk returns a tuple with the Benefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetBenefitsOk() ([]TransactionBenefitsInner, bool) {
	if o == nil || IsNil(o.Benefits) {
		return nil, false
	}
	return o.Benefits, true
}

// HasBenefits returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasBenefits() bool {
	if o != nil && IsNil(o.Benefits) {
		return true
	}

	return false
}

// SetBenefits gets a reference to the given []TransactionBenefitsInner and assigns it to the Benefits field.
func (o *PostTransactionAsyncRequest) SetBenefits(v []TransactionBenefitsInner) {
	o.Benefits = v
}

// GetPromotions returns the Promotions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetPromotions() []ProductPromotion {
	if o == nil {
		var ret []ProductPromotion
		return ret
	}
	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetPromotionsOk() ([]ProductPromotion, bool) {
	if o == nil || IsNil(o.Promotions) {
		return nil, false
	}
	return o.Promotions, true
}

// HasPromotions returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasPromotions() bool {
	if o != nil && IsNil(o.Promotions) {
		return true
	}

	return false
}

// SetPromotions gets a reference to the given []ProductPromotion and assigns it to the Promotions field.
func (o *PostTransactionAsyncRequest) SetPromotions(v []ProductPromotion) {
	o.Promotions = v
}

// GetRequestedValues returns the RequestedValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetRequestedValues() TransactionRequestedValues {
	if o == nil || IsNil(o.RequestedValues.Get()) {
		var ret TransactionRequestedValues
		return ret
	}
	return *o.RequestedValues.Get()
}

// GetRequestedValuesOk returns a tuple with the RequestedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetRequestedValuesOk() (*TransactionRequestedValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedValues.Get(), o.RequestedValues.IsSet()
}

// HasRequestedValues returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasRequestedValues() bool {
	if o != nil && o.RequestedValues.IsSet() {
		return true
	}

	return false
}

// SetRequestedValues gets a reference to the given NullableTransactionRequestedValues and assigns it to the RequestedValues field.
func (o *PostTransactionAsyncRequest) SetRequestedValues(v TransactionRequestedValues) {
	o.RequestedValues.Set(&v)
}

// SetRequestedValuesNil sets the value for RequestedValues to be an explicit nil
func (o *PostTransactionAsyncRequest) SetRequestedValuesNil() {
	o.RequestedValues.Set(nil)
}

// UnsetRequestedValues ensures that no value is present for RequestedValues, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetRequestedValues() {
	o.RequestedValues.Unset()
}

// GetAdjustedValues returns the AdjustedValues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetAdjustedValues() TransactionRequestedValues {
	if o == nil || IsNil(o.AdjustedValues.Get()) {
		var ret TransactionRequestedValues
		return ret
	}
	return *o.AdjustedValues.Get()
}

// GetAdjustedValuesOk returns a tuple with the AdjustedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetAdjustedValuesOk() (*TransactionRequestedValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdjustedValues.Get(), o.AdjustedValues.IsSet()
}

// HasAdjustedValues returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasAdjustedValues() bool {
	if o != nil && o.AdjustedValues.IsSet() {
		return true
	}

	return false
}

// SetAdjustedValues gets a reference to the given NullableTransactionRequestedValues and assigns it to the AdjustedValues field.
func (o *PostTransactionAsyncRequest) SetAdjustedValues(v TransactionRequestedValues) {
	o.AdjustedValues.Set(&v)
}

// SetAdjustedValuesNil sets the value for AdjustedValues to be an explicit nil
func (o *PostTransactionAsyncRequest) SetAdjustedValuesNil() {
	o.AdjustedValues.Set(nil)
}

// UnsetAdjustedValues ensures that no value is present for AdjustedValues, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetAdjustedValues() {
	o.AdjustedValues.Unset()
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetSender() TransactionSender {
	if o == nil || IsNil(o.Sender.Get()) {
		var ret TransactionSender
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetSenderOk() (*TransactionSender, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableTransactionSender and assigns it to the Sender field.
func (o *PostTransactionAsyncRequest) SetSender(v TransactionSender) {
	o.Sender.Set(&v)
}

// SetSenderNil sets the value for Sender to be an explicit nil
func (o *PostTransactionAsyncRequest) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetSender() {
	o.Sender.Unset()
}

// GetBeneficiary returns the Beneficiary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetBeneficiary() TransactionSender {
	if o == nil || IsNil(o.Beneficiary.Get()) {
		var ret TransactionSender
		return ret
	}
	return *o.Beneficiary.Get()
}

// GetBeneficiaryOk returns a tuple with the Beneficiary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetBeneficiaryOk() (*TransactionSender, bool) {
	if o == nil {
		return nil, false
	}
	return o.Beneficiary.Get(), o.Beneficiary.IsSet()
}

// HasBeneficiary returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasBeneficiary() bool {
	if o != nil && o.Beneficiary.IsSet() {
		return true
	}

	return false
}

// SetBeneficiary gets a reference to the given NullableTransactionSender and assigns it to the Beneficiary field.
func (o *PostTransactionAsyncRequest) SetBeneficiary(v TransactionSender) {
	o.Beneficiary.Set(&v)
}

// SetBeneficiaryNil sets the value for Beneficiary to be an explicit nil
func (o *PostTransactionAsyncRequest) SetBeneficiaryNil() {
	o.Beneficiary.Set(nil)
}

// UnsetBeneficiary ensures that no value is present for Beneficiary, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetBeneficiary() {
	o.Beneficiary.Unset()
}

// GetDebitPartyIdentifier returns the DebitPartyIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransactionAsyncRequest) GetDebitPartyIdentifier() TransactionDebitPartyIdentifier {
	if o == nil || IsNil(o.DebitPartyIdentifier.Get()) {
		var ret TransactionDebitPartyIdentifier
		return ret
	}
	return *o.DebitPartyIdentifier.Get()
}

// GetDebitPartyIdentifierOk returns a tuple with the DebitPartyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransactionAsyncRequest) GetDebitPartyIdentifierOk() (*TransactionDebitPartyIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebitPartyIdentifier.Get(), o.DebitPartyIdentifier.IsSet()
}

// HasDebitPartyIdentifier returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasDebitPartyIdentifier() bool {
	if o != nil && o.DebitPartyIdentifier.IsSet() {
		return true
	}

	return false
}

// SetDebitPartyIdentifier gets a reference to the given NullableTransactionDebitPartyIdentifier and assigns it to the DebitPartyIdentifier field.
func (o *PostTransactionAsyncRequest) SetDebitPartyIdentifier(v TransactionDebitPartyIdentifier) {
	o.DebitPartyIdentifier.Set(&v)
}

// SetDebitPartyIdentifierNil sets the value for DebitPartyIdentifier to be an explicit nil
func (o *PostTransactionAsyncRequest) SetDebitPartyIdentifierNil() {
	o.DebitPartyIdentifier.Set(nil)
}

// UnsetDebitPartyIdentifier ensures that no value is present for DebitPartyIdentifier, not even an explicit nil
func (o *PostTransactionAsyncRequest) UnsetDebitPartyIdentifier() {
	o.DebitPartyIdentifier.Unset()
}

// GetCreditPartyIdentifier returns the CreditPartyIdentifier field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetCreditPartyIdentifier() TransactionCreditPartyIdentifier {
	if o == nil || IsNil(o.CreditPartyIdentifier) {
		var ret TransactionCreditPartyIdentifier
		return ret
	}
	return *o.CreditPartyIdentifier
}

// GetCreditPartyIdentifierOk returns a tuple with the CreditPartyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetCreditPartyIdentifierOk() (*TransactionCreditPartyIdentifier, bool) {
	if o == nil || IsNil(o.CreditPartyIdentifier) {
		return nil, false
	}
	return o.CreditPartyIdentifier, true
}

// HasCreditPartyIdentifier returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasCreditPartyIdentifier() bool {
	if o != nil && !IsNil(o.CreditPartyIdentifier) {
		return true
	}

	return false
}

// SetCreditPartyIdentifier gets a reference to the given TransactionCreditPartyIdentifier and assigns it to the CreditPartyIdentifier field.
func (o *PostTransactionAsyncRequest) SetCreditPartyIdentifier(v TransactionCreditPartyIdentifier) {
	o.CreditPartyIdentifier = &v
}

// GetStatementIdentifier returns the StatementIdentifier field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetStatementIdentifier() TransactionStatementIdentifier {
	if o == nil || IsNil(o.StatementIdentifier) {
		var ret TransactionStatementIdentifier
		return ret
	}
	return *o.StatementIdentifier
}

// GetStatementIdentifierOk returns a tuple with the StatementIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetStatementIdentifierOk() (*TransactionStatementIdentifier, bool) {
	if o == nil || IsNil(o.StatementIdentifier) {
		return nil, false
	}
	return o.StatementIdentifier, true
}

// HasStatementIdentifier returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasStatementIdentifier() bool {
	if o != nil && !IsNil(o.StatementIdentifier) {
		return true
	}

	return false
}

// SetStatementIdentifier gets a reference to the given TransactionStatementIdentifier and assigns it to the StatementIdentifier field.
func (o *PostTransactionAsyncRequest) SetStatementIdentifier(v TransactionStatementIdentifier) {
	o.StatementIdentifier = &v
}

// GetAdditionalIdentifier returns the AdditionalIdentifier field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetAdditionalIdentifier() TransactionAdditionalIdentifier {
	if o == nil || IsNil(o.AdditionalIdentifier) {
		var ret TransactionAdditionalIdentifier
		return ret
	}
	return *o.AdditionalIdentifier
}

// GetAdditionalIdentifierOk returns a tuple with the AdditionalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetAdditionalIdentifierOk() (*TransactionAdditionalIdentifier, bool) {
	if o == nil || IsNil(o.AdditionalIdentifier) {
		return nil, false
	}
	return o.AdditionalIdentifier, true
}

// HasAdditionalIdentifier returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasAdditionalIdentifier() bool {
	if o != nil && !IsNil(o.AdditionalIdentifier) {
		return true
	}

	return false
}

// SetAdditionalIdentifier gets a reference to the given TransactionAdditionalIdentifier and assigns it to the AdditionalIdentifier field.
func (o *PostTransactionAsyncRequest) SetAdditionalIdentifier(v TransactionAdditionalIdentifier) {
	o.AdditionalIdentifier = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *PostTransactionAsyncRequest) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTransactionAsyncRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PostTransactionAsyncRequest) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *PostTransactionAsyncRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o PostTransactionAsyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTransactionAsyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["external_id"] = o.ExternalId
	// skip: creation_date is readOnly
	// skip: confirmation_expiration_date is readOnly
	// skip: confirmation_date is readOnly
	toSerialize["product_id"] = o.ProductId
	if !IsNil(o.CalculationMode) {
		toSerialize["calculation_mode"] = o.CalculationMode
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Destination.IsSet() {
		toSerialize["destination"] = o.Destination.Get()
	}
	if !IsNil(o.AutoConfirm) {
		toSerialize["auto_confirm"] = o.AutoConfirm
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.OperatorReference.IsSet() {
		toSerialize["operator_reference"] = o.OperatorReference.Get()
	}
	if o.Pin.IsSet() {
		toSerialize["pin"] = o.Pin.Get()
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}
	if o.Benefits != nil {
		toSerialize["benefits"] = o.Benefits
	}
	if o.Promotions != nil {
		toSerialize["promotions"] = o.Promotions
	}
	if o.RequestedValues.IsSet() {
		toSerialize["requested_values"] = o.RequestedValues.Get()
	}
	if o.AdjustedValues.IsSet() {
		toSerialize["adjusted_values"] = o.AdjustedValues.Get()
	}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	if o.Beneficiary.IsSet() {
		toSerialize["beneficiary"] = o.Beneficiary.Get()
	}
	if o.DebitPartyIdentifier.IsSet() {
		toSerialize["debit_party_identifier"] = o.DebitPartyIdentifier.Get()
	}
	if !IsNil(o.CreditPartyIdentifier) {
		toSerialize["credit_party_identifier"] = o.CreditPartyIdentifier
	}
	if !IsNil(o.StatementIdentifier) {
		toSerialize["statement_identifier"] = o.StatementIdentifier
	}
	if !IsNil(o.AdditionalIdentifier) {
		toSerialize["additional_identifier"] = o.AdditionalIdentifier
	}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	return toSerialize, nil
}

type NullablePostTransactionAsyncRequest struct {
	value *PostTransactionAsyncRequest
	isSet bool
}

func (v NullablePostTransactionAsyncRequest) Get() *PostTransactionAsyncRequest {
	return v.value
}

func (v *NullablePostTransactionAsyncRequest) Set(val *PostTransactionAsyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTransactionAsyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTransactionAsyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTransactionAsyncRequest(val *PostTransactionAsyncRequest) *NullablePostTransactionAsyncRequest {
	return &NullablePostTransactionAsyncRequest{value: val, isSet: true}
}

func (v NullablePostTransactionAsyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTransactionAsyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
