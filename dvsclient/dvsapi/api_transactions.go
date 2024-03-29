/*
Digital Value Services API

# Overview  Welcome to the Digital Value Services (DVS) API reference.  This API serves as the primary gateway to facilitate digital value transfers through [DT One](https://www.dtone.com), a leading global network covering more than 160 countries and 550 mobile operators.  The Digital Value Services API is organized according to [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, using [JSON](https://en.wikipedia.org/wiki/JSON) as format for data interchange, and provides the following services:   - [Discovery Services](#tag/Services)   - [Transaction Services](#tag/Transactions)   - [Account Services](#tag/Balances)   - [Look-Up Services](#tag/Mobile-Number)  If you have any questions and/or suggestions related to our API, please do not hesitate to send an email to the [DVS API support team](mailto:dvs-api-support@dtone.com).  ## Integration Libraries  Officially supported [SDK](https://en.wikipedia.org/wiki/Software_development_kit)s are available for the following languages:  * [Java](https://github.com/dtone/dtone-dvs-api-java-client) * [Node.js](https://www.npmjs.com/package/@dtone/dvs)  These SDKs offer an accelerated path to developing your applications as an alternative to accessing the REST API directly.  Separately, we would love to hear from you! If you have any questions and/or suggestions related to our SDKs, please do not hesitate to create corresponding [GitHub issues](https://guides.github.com/features/issues/) or send an email to the [DVS Open Source team](mailto:opensource@dtone.com).  ## Sandbox  A sandbox environment is available for testing integrations with the DVS API. It is available at [https://preprod-dvs-api.dtone.com/v1/](https://preprod-dvs-api.dtone.com/v1/).  You can generate sandbox API keys from your [DT Shop](https://dtshop.dtone.com/account?tab=developer) account, under the **Pre-Production API Keys** section.  All transactions on the sandbox environment are simulated: no real transaction goes through. To simulate different responses, the last three digits of the `credit_party_identifier` (i.e. `mobile_number` or `account_number`, depending on the `required_credit_party_identifier_fields` of a given [Product](/#tag/Products)) will have to be replaced with one of the following suffixes:  | Suffix              | Transaction Status                        | Example       | | ---                 | ---                                       | ---           | | `100`, `200`, `300` | `COMPLETED` (PIN-less)                    | `+6595123100` | | `101`, `201`, `301` | `COMPLETED` (PIN-based)                   | `+6595123201` | | `102`, `202`, `302` | `DECLINED-INVALID-CREDIT-PARTY`           | `+6595123102` | | `103`, `203`, `303` | `DECLINED-BARRED-CREDIT-PARTY`            | `+6595123103` | | `104`, `204`, `304` | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE` | `+6595123204` | | `105`, `205`, `305` | `DECLINED-DUPLICATED-TRANSACTION`         | `+6595123105` | | `106`, `206`, `306` | `DECLINED`                                | `+6595123206` | | `107`, `207`, `307` | `DECLINED-EXCEPTION`                      | `+6595123107` |  The different suffixes for a given transaction status can be used to simulate delays, as follows:   * `10X` suffix will take at least 3 seconds to finish   * `20X` suffix will take at least 20 seconds to finish   * `30X` suffix will take at least 5 minutes to finish  Please note that there are products that do not require any credit party identifier such as Gift Cards. For these products, the simulated transaction will always be in `COMPLETED` status.  ## Versioning  Endpoints of the API are prefixed with a corresponding version number.  This is done to provide complete isolation between implementations and to ensure that subsequent major changes to the API will never affect existing integrations.  When a new version of the API is available and you are keen to upgrade, testing in the sandbox environment to ensure that everything works well with your implementation before switching to the production environment comes highly recommended.  Feel free to [contact us](mailto:dvs-api-support@dtone.com) should you have any questions and/or are in need of assistance during your tests.  ## Transactions  The main purpose of this API is to deliver value (e.g. mobile airtime top-up, data bundles, etc.) to a beneficiary. This is what we call a \"transaction\".  During the course of a transfer, a transaction undergoes various status changes (or transitions) as illustrated below.  ![transaction states](/images/transaction_states.png)  As changes in transaction status occur, updates are sent in real-time when a callback URL is provided. In conjunction, transaction status can be queried through one of two means: via the returned `id` or a provided `external_id`.  The latter serves as your unique reference and provides a utility to retrieve transaction details when exceptions occur, such as when the supposed API response was not received, as an example.  ## Balances  Transactions can be created through the platform as long as there is enough balance available in your account. A given balance is composed of the following:  | Balance   | Description                                              | | ---       | ---                                                      | | Available | Balance amount available for use                         | | Holding   | Amount being held while transactions are being processed |  As a given transaction goes through various changes in status as outlined [here](#section/Overview/Transactions), corresponding balance movements will be made. The following table illustrates the relationship between transaction status and balance movements:  | Transaction Status                  | Balance Movement | Description                                                | | ---                                 | ---              | ---                                                        | | `CREATED`                           | Authorize        | Transfer wholesale price and fee from available to holding | | `CANCELLED`, `REJECTED`, `DECLINED` | Void             | Amount authorized in holding moves back to available       | | `COMPLETED`                         | Capture          | Amount authorized in holding is captured, i.e. debited     | | `REVERSED`                          | Reverse          | Debited amount is reversed back into available             |  ## Flow  Once a product has been selected through one of the [discovery methods](/#tag/Services) provided by the API, the actual transfer (i.e. transaction) can be performed in either one of the following modes:   - Asynchronous (recommended)   - Synchronous  Each mode is accessible via a specific endpoint.  As soon as a transaction is confirmed, the transfer order will be sent to the operator for immediate processing. During this time, the transaction will remain in a `CONFIRMED` status until the final status is received from the operator.  ### Asynchronous Mode  When a transaction is created and confirmed in an asynchronous fashion, the HTTP connection won't have to be kept open. This preserves system resources on your applications. As such, performing transactions **asynchronously** is **recommended**.  ### Synchronous Mode  When a transaction is created and confirmed in a synchronous fashion, the HTTP connection will be kept open in an attempt to capture the final status from the receiving operator so it can be returned in the API response. The processing time usually takes just a few seconds. However, with some receiving operators, it may take longer.  Our system will keep the HTTP connection open for up to 180 seconds (this value can be configured upon request) and will return a status before closing this connection. This status can be in one of the final status (e.g. `COMPLETED`, `DECLINED`) or not (e.g. `SUBMITTED`). In the latter case, this denotes the transaction is still being processed by the receiving operator.  **Note:** your application does not have to wait for the connection to close, it can listen for a shorter period of time and query the final status later on (refer to the \"Final Status\" section below for more details).  ### Final Status  Regardless of the processing mode, the application should be designed to capture the final status of a transaction. This can be done through one of the following means:   - Checking the status of a specific transaction via the corresponding API method (\"pull\" mechanism)   - Configuring a callback URL passed in the request when creating a transaction (\"push\" mechanism)  ## Callbacks  As a transaction is being processed, changes in status will be notified in real-time if a callback URL was provided. Even though one callback per transaction is expected (i.e. change to either `COMPLETED` or `DECLINED`), a manual reversal from the [DT One](https://dtone.com/) team, which may happen in very rare occasions, will also trigger a callback to inform your application of a change in transaction status to `REVERSED`.  This callback endpoint must be implemented by the sending partner, which should expect an HTTP `POST` request containing a transaction object represented in [JSON](https://en.wikipedia.org/wiki/JSON). As callbacks will be sent from the [DT One](https://dtone.com/) servers, these endpoints will have to be publicly-accessible in most cases. During development, a service such as [ngrok](https://ngrok.com/) can be used to expose local servers to the internet.  Upon successful receipt of data, the callback endpoint should respond with an HTTP `2XX` status. In the event that the platform did not receive a successful status, callback notifications will be retried several times, beyond which, the transaction status will have to be queried through the API.  ## Status and Errors  ### HTTP Status Codes  [DT One](https://dtone.com/) uses standard HTTP response codes to indicate whether an API request was successful or not.  | Status | Description                                        | | ---    | ---                                                | | `200`  | OK                                                 | | `201`  | Created: Resource created                          | | `202`  | Accepted: Request has been accepted for processing | | `400`  | Bad Request: Request was malformed                 | | `401`  | Unauthorized: Credentials missing or invalid       | | `404`  | Not Found: Resource doesn't exist                  | | `429`  | Too Many Requests                                  | | `500`  | Server Error: Error occurred on DT One             | | `503`  | Service Unavailable                                |  ### API Error Codes  | Code      | Description                                       | HTTP Status | | ---       | ---                                               | ---         | | `1000400` | Bad Request                                       | `400`       | | `1000401` | Unauthorized                                      | `401`       | | `1000404` | Resource not found                                | `404`       | | `1000429` | Too many requests                                 | `429`       | | `1003001` | Product is not available in your account          | `404`       | | `1003002` | Requested product amount is out of range          | `400`       | | `1003003` | Requested product unit is invalid                 | `400`       | | `1003101` | Benefits not defined for available products       | `404`       | | `1003201` | Promotion not found                               | `404`       | | `1003301` | Campaign not found                                | `404`       | | `1005003` | Credit party mobile number is invalid             | `400`       | | `1005004` | Service not found                                 | `404`       | | `1005005` | Country not found                                 | `404`       | | `1005006` | Operator not found                                | `404`       | | `1005503` | Sender mobile number is invalid                   | `400`       | | `1006001` | Insufficient balance                              | `400`       | | `1006003` | Debit party mobile number is invalid              | `400`       | | `1006009` | Account balance not found                         | `404`       | | `1006503` | Beneficiary mobile number is invalid              | `400`       | | `1007001` | Transaction external ID has already been used     | `400`       | | `1007002` | Transaction has already been confirmed            | `400`       | | `1007004` | Transaction can no longer be confirmed            | `400`       | | `1007005` | Transaction has already been cancelled            | `400`       | | `1007007` | Transaction can no longer be cancelled            | `400`       | | `1007500` | Method not supported by operator                  | `400`       | | `1008004` | Transaction not found                             | `404`       | | `1009001` | Unexpected error, please contact our support team | `500`       | | `1009503` | Service unavailable, please retry later           | `503`       |  ### Transaction Status  | Class       | Status                                          | Description                                                            | | ---         | ---                                             | ---                                                                    | | `CREATED`   | `CREATED`                                       | Created                                                                | | `CONFIRMED` | `CONFIRMED`                                     | Confirmed                                                              | | `REJECTED`  | `REJECTED`                                      | Rejected                                                               | | `REJECTED`  | `REJECTED-INVALID-CREDIT-PARTY`                 | Rejected - Credit party is invalid                                     | | `REJECTED`  | `REJECTED-BARRED-CREDIT-PARTY`                  | Rejected - Credit party is barred                                      | | `REJECTED`  | `REJECTED-INELIGIBLE-CREDIT-PARTY`              | Rejected - Credit party is ineligible for chosen product               | | `REJECTED`  | `REJECTED-INVALID-DEBIT-PARTY`                  | Rejected - Debit party is invalid                                      | | `REJECTED`  | `REJECTED-BARRED-DEBIT-PARTY`                   | Rejected - Debit party is barred                                       | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Rejected - Limitations on credit party cumulative transaction amount   | | `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Rejected - Limitations on credit party cumulative transaction quantity | | `REJECTED`  | `REJECTED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Rejected - Operator currently unavailable                              | | `REJECTED`  | `REJECTED-INSUFFICIENT-BALANCE`                 | Rejected - Insufficient balance                                        | | `CANCELLED` | `CANCELLED`                                     | Cancelled                                                              | | `SUBMITTED` | `SUBMITTED`                                     | Submitted                                                              | | `COMPLETED` | `COMPLETED`                                     | Completed                                                              | | `REVERSED`  | `REVERSED`                                      | Reversed                                                               | | `DECLINED`  | `DECLINED`                                      | Declined (no additional information available)                         | | `DECLINED`  | `DECLINED-INVALID-CREDIT-PARTY`                 | Declined - Credit party is invalid                                     | | `DECLINED`  | `DECLINED-BARRED-CREDIT-PARTY`                  | Declined - Credit party is barred                                      | | `DECLINED`  | `DECLINED-INELIGIBLE-CREDIT-PARTY`              | Declined - Credit party is ineligible for chosen product               | | `DECLINED`  | `DECLINED-INVALID-DEBIT-PARTY`                  | Declined - Debit party is invalid                                      | | `DECLINED`  | `DECLINED-BARRED-DEBIT-PARTY`                   | Declined - Debit party is barred                                       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-AMOUNT`       | Declined - Limitations on operator cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Declined - Limitations on credit party cumulative transaction amount   | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-AMOUNT`       | Declined - Limitations on customer cumulative transaction amount       | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-QUANTITY`     | Declined - Limitations on operator cumulative transaction quantity     | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Declined - Limitations on credit party cumulative transaction quantity | | `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-QUANTITY`     | Declined - Limitations on customer cumulative transaction quantity     | | `DECLINED`  | `DECLINED-DUPLICATED-TRANSACTION`               | Declined - Duplicated transaction                                      | | `DECLINED`  | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Declined - Operator currently unavailable                              |  `REJECTED` and `DECLINED` status classes both denote unsuccessful transactions. The primary distinction between these two relates to the party that determined the failure:   * `REJECTED` are issued by the DVS platform based on various business rules (e.g. insufficient balance, limitations, etc)   * `DECLINED` are issued by the operators  Separately, it is recommended to define application logic based on **classes**, while additional distinction and/or insight are reflected on the actual **status**.  ## Pagination  API resources supporting bulk fetches via \"list\" API methods will be returned in a paginated fashion.  ### Input Parameters  | Field      | Required | Type    | Description                                      | | ---        | ---      | ---     | ---                                              | | `page`     | No       | Integer | Page number                                      | | `per_page` | No       | Integer | Number of results per page (default 50, max 100) |  ### Output Headers  | Field           | Description                   | | ---             | ---                           | | `X-Total`       | Total number of records       | | `X-Total-Pages` | Total number of pages         | | `X-Per-Page`    | Number of records per page    | | `X-Page`        | Current page number           | | `X-Next-Page`   | Next page number (if any)     | | `X-Prev-Page`   | Previous page number (if any) |  ## Rate Limiting  The API endpoints have rate limiting in place to protect our service from excessive number of requests.  If the limit is reached, an [HTTP error 429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) will be returned by the server.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dvsapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TransactionsApiService TransactionsApi service
type TransactionsApiService service

type TransactionsApiGetTransactionByIdRequest struct {
	ctx           context.Context
	ApiService    *TransactionsApiService
	transactionId int64
}

func (r TransactionsApiGetTransactionByIdRequest) Execute() (*PostTransactionAsyncRequest, *http.Response, error) {
	return r.ApiService.GetTransactionByIdExecute(r)
}

/*
GetTransactionById Query a transaction by ID

This endpoint will return the details of the requested transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId
	@return TransactionsApiGetTransactionByIdRequest
*/
func (a *TransactionsApiService) GetTransactionById(ctx context.Context, transactionId int64) TransactionsApiGetTransactionByIdRequest {
	return TransactionsApiGetTransactionByIdRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return PostTransactionAsyncRequest
func (a *TransactionsApiService) GetTransactionByIdExecute(r TransactionsApiGetTransactionByIdRequest) (*PostTransactionAsyncRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PostTransactionAsyncRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetTransactionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{transaction_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transactionId < 1 {
		return localVarReturnValue, nil, reportError("transactionId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiGetTransactionsRequest struct {
	ctx                      context.Context
	ApiService               *TransactionsApiService
	externalId               *string
	productType              *ProductTypes
	serviceId                *int32
	subserviceId             *int32
	countryIsoCode           *string
	operatorId               *int32
	statusId                 *int32
	creditPartyMobileNumber  *string
	creditPartyAccountNumber *string
	fromDate                 *time.Time
	toDate                   *time.Time
	page                     *int32
	perPage                  *int32
}

func (r TransactionsApiGetTransactionsRequest) ExternalId(externalId string) TransactionsApiGetTransactionsRequest {
	r.externalId = &externalId
	return r
}

func (r TransactionsApiGetTransactionsRequest) ProductType(productType ProductTypes) TransactionsApiGetTransactionsRequest {
	r.productType = &productType
	return r
}

func (r TransactionsApiGetTransactionsRequest) ServiceId(serviceId int32) TransactionsApiGetTransactionsRequest {
	r.serviceId = &serviceId
	return r
}

func (r TransactionsApiGetTransactionsRequest) SubserviceId(subserviceId int32) TransactionsApiGetTransactionsRequest {
	r.subserviceId = &subserviceId
	return r
}

func (r TransactionsApiGetTransactionsRequest) CountryIsoCode(countryIsoCode string) TransactionsApiGetTransactionsRequest {
	r.countryIsoCode = &countryIsoCode
	return r
}

func (r TransactionsApiGetTransactionsRequest) OperatorId(operatorId int32) TransactionsApiGetTransactionsRequest {
	r.operatorId = &operatorId
	return r
}

func (r TransactionsApiGetTransactionsRequest) StatusId(statusId int32) TransactionsApiGetTransactionsRequest {
	r.statusId = &statusId
	return r
}

func (r TransactionsApiGetTransactionsRequest) CreditPartyMobileNumber(creditPartyMobileNumber string) TransactionsApiGetTransactionsRequest {
	r.creditPartyMobileNumber = &creditPartyMobileNumber
	return r
}

func (r TransactionsApiGetTransactionsRequest) CreditPartyAccountNumber(creditPartyAccountNumber string) TransactionsApiGetTransactionsRequest {
	r.creditPartyAccountNumber = &creditPartyAccountNumber
	return r
}

// Starting date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between &#x60;from_date&#x60; and &#x60;to_date&#x60; can not exceed 24 hours.
func (r TransactionsApiGetTransactionsRequest) FromDate(fromDate time.Time) TransactionsApiGetTransactionsRequest {
	r.fromDate = &fromDate
	return r
}

// Ending date to filter transactions based on creation date, inclusive of the provided date. Please note that the window between &#x60;from_date&#x60; and &#x60;to_date&#x60; can not exceed 24 hours.
func (r TransactionsApiGetTransactionsRequest) ToDate(toDate time.Time) TransactionsApiGetTransactionsRequest {
	r.toDate = &toDate
	return r
}

// Page number
func (r TransactionsApiGetTransactionsRequest) Page(page int32) TransactionsApiGetTransactionsRequest {
	r.page = &page
	return r
}

// Number of records per page
func (r TransactionsApiGetTransactionsRequest) PerPage(perPage int32) TransactionsApiGetTransactionsRequest {
	r.perPage = &perPage
	return r
}

func (r TransactionsApiGetTransactionsRequest) Execute() ([]PostTransactionAsyncRequest, *http.Response, error) {
	return r.ApiService.GetTransactionsExecute(r)
}

/*
GetTransactions Query list of transactions

This endpoint will return a list of transactions matching the search criteria. Please note that when this endpoint is called without any parameters and/or if neither date ranges (i.e. `from_date`, `to_date`) nor `external_id` are specified, transactions created within the last 24 hours will be returned by default.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TransactionsApiGetTransactionsRequest
*/
func (a *TransactionsApiService) GetTransactions(ctx context.Context) TransactionsApiGetTransactionsRequest {
	return TransactionsApiGetTransactionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []PostTransactionAsyncRequest
func (a *TransactionsApiService) GetTransactionsExecute(r TransactionsApiGetTransactionsRequest) ([]PostTransactionAsyncRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []PostTransactionAsyncRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.externalId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_id", r.externalId, "")
	}
	if r.productType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_type", r.productType, "")
	}
	if r.serviceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service_id", r.serviceId, "")
	}
	if r.subserviceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subservice_id", r.subserviceId, "")
	}
	if r.countryIsoCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_iso_code", r.countryIsoCode, "")
	}
	if r.operatorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operator_id", r.operatorId, "")
	}
	if r.statusId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status_id", r.statusId, "")
	}
	if r.creditPartyMobileNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credit_party_mobile_number", r.creditPartyMobileNumber, "")
	}
	if r.creditPartyAccountNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credit_party_account_number", r.creditPartyAccountNumber, "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_date", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to_date", r.toDate, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiPostTransactionAsyncRequest struct {
	ctx                         context.Context
	ApiService                  *TransactionsApiService
	postTransactionAsyncRequest *PostTransactionAsyncRequest
}

func (r TransactionsApiPostTransactionAsyncRequest) PostTransactionAsyncRequest(postTransactionAsyncRequest PostTransactionAsyncRequest) TransactionsApiPostTransactionAsyncRequest {
	r.postTransactionAsyncRequest = &postTransactionAsyncRequest
	return r
}

func (r TransactionsApiPostTransactionAsyncRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.PostTransactionAsyncExecute(r)
}

/*
PostTransactionAsync Create a transaction asynchronously

Two transaction modes (asynchronous and synchronous) are available. This endpoint lets you create a transaction in the **asynchronous** mode. Note that the `auto_confirm` flag can be set to simultaneously create and confirm a transaction in one step (i.e. HTTP request).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TransactionsApiPostTransactionAsyncRequest
*/
func (a *TransactionsApiService) PostTransactionAsync(ctx context.Context) TransactionsApiPostTransactionAsyncRequest {
	return TransactionsApiPostTransactionAsyncRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Transaction
func (a *TransactionsApiService) PostTransactionAsyncExecute(r TransactionsApiPostTransactionAsyncRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.PostTransactionAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/async/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postTransactionAsyncRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiPostTransactionCancelRequest struct {
	ctx           context.Context
	ApiService    *TransactionsApiService
	transactionId int64
}

func (r TransactionsApiPostTransactionCancelRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.PostTransactionCancelExecute(r)
}

/*
PostTransactionCancel Cancel a transaction

If a transaction is still in the `CREATED` state, it has not yet been submitted to the receiving operator for processing. You can thus request to cancel such transactions by calling this endpoint.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId
	@return TransactionsApiPostTransactionCancelRequest
*/
func (a *TransactionsApiService) PostTransactionCancel(ctx context.Context, transactionId int64) TransactionsApiPostTransactionCancelRequest {
	return TransactionsApiPostTransactionCancelRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return Transaction
func (a *TransactionsApiService) PostTransactionCancelExecute(r TransactionsApiPostTransactionCancelRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.PostTransactionCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/{transaction_id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transactionId < 1 {
		return localVarReturnValue, nil, reportError("transactionId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiPostTransactionConfirmAsyncRequest struct {
	ctx           context.Context
	ApiService    *TransactionsApiService
	transactionId int64
}

func (r TransactionsApiPostTransactionConfirmAsyncRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.PostTransactionConfirmAsyncExecute(r)
}

/*
PostTransactionConfirmAsync Confirm a transaction asynchronously

If an **asynchronous** transaction was created without setting the `auto_confirm` flag, this endpoint will have to be called to confirm the transaction. Once successfully confirmed, the transfer order will be submitted to the operator to be processed.

Please note that only unexpired transactions can be confirmed, as denoted in the `confirmation_expiration_date` field of the transaction. Beyond this, the only allowed change is to [cancel the transaction](/#tag/Transactions/paths/~1transactions~1{transaction_id}~1cancel/post), so as to release the held balance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId
	@return TransactionsApiPostTransactionConfirmAsyncRequest
*/
func (a *TransactionsApiService) PostTransactionConfirmAsync(ctx context.Context, transactionId int64) TransactionsApiPostTransactionConfirmAsyncRequest {
	return TransactionsApiPostTransactionConfirmAsyncRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return Transaction
func (a *TransactionsApiService) PostTransactionConfirmAsyncExecute(r TransactionsApiPostTransactionConfirmAsyncRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.PostTransactionConfirmAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/async/transactions/{transaction_id}/confirm"
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transactionId < 1 {
		return localVarReturnValue, nil, reportError("transactionId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiPostTransactionConfirmSyncRequest struct {
	ctx           context.Context
	ApiService    *TransactionsApiService
	transactionId int64
}

func (r TransactionsApiPostTransactionConfirmSyncRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.PostTransactionConfirmSyncExecute(r)
}

/*
PostTransactionConfirmSync Confirm a transaction synchronously

If a **synchronous** transaction was created without setting the `auto_confirm` flag, this endpoint will have to be called to confirm the transaction. Once successfully confirmed, the transfer order will be submitted to the operator to be processed.

Please note that only unexpired transactions can be confirmed, as denoted in the `confirmation_expiration_date` field of the transaction. Beyond this, the only allowed change is to [cancel the transaction](/#tag/Transactions/paths/~1transactions~1{transaction_id}~1cancel/post), so as to release the held balance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId
	@return TransactionsApiPostTransactionConfirmSyncRequest

Deprecated
*/
func (a *TransactionsApiService) PostTransactionConfirmSync(ctx context.Context, transactionId int64) TransactionsApiPostTransactionConfirmSyncRequest {
	return TransactionsApiPostTransactionConfirmSyncRequest{
		ApiService:    a,
		ctx:           ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//
//	@return Transaction
//
// Deprecated
func (a *TransactionsApiService) PostTransactionConfirmSyncExecute(r TransactionsApiPostTransactionConfirmSyncRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.PostTransactionConfirmSync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sync/transactions/{transaction_id}/confirm"
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transactionId < 1 {
		return localVarReturnValue, nil, reportError("transactionId must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TransactionsApiPostTransactionSyncRequest struct {
	ctx         context.Context
	ApiService  *TransactionsApiService
	transaction *Transaction
}

func (r TransactionsApiPostTransactionSyncRequest) Transaction(transaction Transaction) TransactionsApiPostTransactionSyncRequest {
	r.transaction = &transaction
	return r
}

func (r TransactionsApiPostTransactionSyncRequest) Execute() (*Transaction, *http.Response, error) {
	return r.ApiService.PostTransactionSyncExecute(r)
}

/*
PostTransactionSync Create a transaction synchronously

Two transaction modes (asynchronous and synchronous) are available. This endpoint lets you create a transaction in the **synchronous** mode. Note that the `auto_confirm` flag can be set to simultaneously create and confirm a transaction in one step (i.e. HTTP request).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TransactionsApiPostTransactionSyncRequest

Deprecated
*/
func (a *TransactionsApiService) PostTransactionSync(ctx context.Context) TransactionsApiPostTransactionSyncRequest {
	return TransactionsApiPostTransactionSyncRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Transaction
//
// Deprecated
func (a *TransactionsApiService) PostTransactionSyncExecute(r TransactionsApiPostTransactionSyncRequest) (*Transaction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.PostTransactionSync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sync/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transaction
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
