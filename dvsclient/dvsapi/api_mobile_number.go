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
)

// MobileNumberApiService MobileNumberApi service
type MobileNumberApiService service

type MobileNumberApiGetLookupMobileNumberRequest struct {
	ctx          context.Context
	ApiService   *MobileNumberApiService
	mobileNumber string
	page         *int32
	perPage      *int32
}

// Page number
func (r MobileNumberApiGetLookupMobileNumberRequest) Page(page int32) MobileNumberApiGetLookupMobileNumberRequest {
	r.page = &page
	return r
}

// Number of records per page
func (r MobileNumberApiGetLookupMobileNumberRequest) PerPage(perPage int32) MobileNumberApiGetLookupMobileNumberRequest {
	r.perPage = &perPage
	return r
}

func (r MobileNumberApiGetLookupMobileNumberRequest) Execute() ([]GetLookupMobileNumber200ResponseInner, *http.Response, error) {
	return r.ApiService.GetLookupMobileNumberExecute(r)
}

/*
GetLookupMobileNumber Look up operators for a given mobile number

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param mobileNumber
	@return MobileNumberApiGetLookupMobileNumberRequest
*/
func (a *MobileNumberApiService) GetLookupMobileNumber(ctx context.Context, mobileNumber string) MobileNumberApiGetLookupMobileNumberRequest {
	return MobileNumberApiGetLookupMobileNumberRequest{
		ApiService:   a,
		ctx:          ctx,
		mobileNumber: mobileNumber,
	}
}

// Execute executes the request
//
//	@return []GetLookupMobileNumber200ResponseInner
func (a *MobileNumberApiService) GetLookupMobileNumberExecute(r MobileNumberApiGetLookupMobileNumberRequest) ([]GetLookupMobileNumber200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []GetLookupMobileNumber200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MobileNumberApiService.GetLookupMobileNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/lookup/mobile-number/{mobile_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"mobile_number"+"}", url.PathEscape(parameterValueToString(r.mobileNumber, "mobileNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type MobileNumberApiPostLookupMobileNumberRequest struct {
	ctx                           context.Context
	ApiService                    *MobileNumberApiService
	postLookupMobileNumberRequest *PostLookupMobileNumberRequest
}

func (r MobileNumberApiPostLookupMobileNumberRequest) PostLookupMobileNumberRequest(postLookupMobileNumberRequest PostLookupMobileNumberRequest) MobileNumberApiPostLookupMobileNumberRequest {
	r.postLookupMobileNumberRequest = &postLookupMobileNumberRequest
	return r
}

func (r MobileNumberApiPostLookupMobileNumberRequest) Execute() ([]GetLookupMobileNumber200ResponseInner, *http.Response, error) {
	return r.ApiService.PostLookupMobileNumberExecute(r)
}

/*
PostLookupMobileNumber Look up operators for a given mobile number

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MobileNumberApiPostLookupMobileNumberRequest
*/
func (a *MobileNumberApiService) PostLookupMobileNumber(ctx context.Context) MobileNumberApiPostLookupMobileNumberRequest {
	return MobileNumberApiPostLookupMobileNumberRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []GetLookupMobileNumber200ResponseInner
func (a *MobileNumberApiService) PostLookupMobileNumberExecute(r MobileNumberApiPostLookupMobileNumberRequest) ([]GetLookupMobileNumber200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []GetLookupMobileNumber200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MobileNumberApiService.PostLookupMobileNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/lookup/mobile-number"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postLookupMobileNumberRequest == nil {
		return localVarReturnValue, nil, reportError("postLookupMobileNumberRequest is required and must be specified")
	}

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
	localVarPostBody = r.postLookupMobileNumberRequest
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
