# Go API client for dvsapi

# Overview

Welcome to the Digital Value Services (DVS) API reference.

This API serves as the primary gateway to facilitate digital value transfers through [DT One](https://www.dtone.com), a leading global network covering more than 160 countries and 550 mobile operators.

The Digital Value Services API is organized according to [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) principles, using [JSON](https://en.wikipedia.org/wiki/JSON) as format for data interchange, and provides the following services:
  - [Discovery Services](#tag/Services)
  - [Transaction Services](#tag/Transactions)
  - [Account Services](#tag/Balances)
  - [Look-Up Services](#tag/Mobile-Number)

If you have any questions and/or suggestions related to our API, please do not hesitate to send an email to the [DVS API support team](mailto:dvs-api-support@dtone.com).

## Integration Libraries

Officially supported [SDK](https://en.wikipedia.org/wiki/Software_development_kit)s are available for the following languages:

* [Java](https://github.com/dtone/dtone-dvs-api-java-client)
* [Node.js](https://www.npmjs.com/package/@dtone/dvs)

These SDKs offer an accelerated path to developing your applications as an alternative to accessing the REST API directly.

Separately, we would love to hear from you! If you have any questions and/or suggestions related to our SDKs, please do not hesitate to create corresponding [GitHub issues](https://guides.github.com/features/issues/) or send an email to the [DVS Open Source team](mailto:opensource@dtone.com).

## Sandbox

A sandbox environment is available for testing integrations with the DVS API. It is available at [https://preprod-dvs-api.dtone.com/v1/](https://preprod-dvs-api.dtone.com/v1/).

You can generate sandbox API keys from your [DT Shop](https://dtshop.dtone.com/account?tab=developer) account, under the **Pre-Production API Keys** section.

All transactions on the sandbox environment are simulated: no real transaction goes through. To simulate different responses, the last three digits of the recipient mobile number (i.e. `credit_party_identifier.mobile_number`) will have to be replaced with one of the following suffixes:

| Suffix              | Transaction Status                        | Example       |
| ---                 | ---                                       | ---           |
| `100`, `200`, `300` | `COMPLETED` (PIN-less)                    | `+6595123100` |
| `101`, `201`, `301` | `COMPLETED` (PIN-based)                   | `+6595123201` |
| `102`, `202`, `302` | `DECLINED-INVALID-CREDIT-PARTY`           | `+6595123102` |
| `103`, `203`, `303` | `DECLINED-BARRED-CREDIT-PARTY`            | `+6595123103` |
| `104`, `204`, `304` | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE` | `+6595123205` |
| `105`, `205`, `305` | `DECLINED-DUPLICATED-TRANSACTION`         | `+6595123105` |
| `106`, `206`, `306` | `DECLINED`                                | `+6595123206` |
| `107`, `207`, `307` | `DECLINED-EXCEPTION`                      | `+6595123107` |

The different suffixes for a given transaction status can be used to simulate delays, as follows:
  * `10X` suffix will take at least 3 seconds to finish
  * `20X` suffix will take at least 20 seconds to finish
  * `30X` suffix will take at least 5 minutes to finish

Please note that there are products that do not require any credit party identifier such as Gift Cards. For these products, the simulated transaction will always be in `COMPLETED` status.

## Versioning

Endpoints of the API are prefixed with a corresponding version number.

This is done to provide complete isolation between implementations and to ensure that subsequent major changes to the API will never affect existing integrations.

When a new version of the API is available and you are keen to upgrade, testing in the sandbox environment to ensure that everything works well with your implementation before switching to the production environment comes highly recommended.

Feel free to [contact us](mailto:dvs-api-support@dtone.com) should you have any questions and/or are in need of assistance during your tests.

## Transactions

The main purpose of this API is to deliver value (e.g. mobile airtime top-up, data bundles, etc.) to a beneficiary. This is what we call a \"transaction\".

During the course of a transfer, a transaction undergoes various status changes (or transitions) as illustrated below.

![transaction states](/images/transaction_states.png)

As changes in transaction status occur, updates are sent in real-time when a callback URL is provided. In conjunction, transaction status can be queried through one of two means: via the returned `id` or a provided `external_id`.

The latter serves as your unique reference and provides a utility to retrieve transaction details when exceptions occur, such as when the supposed API response was not received, as an example.

## Balances

Transactions can be created through the platform as long as there is enough balance available in your account. A given balance is composed of the following:

| Balance   | Description                                              |
| ---       | ---                                                      |
| Available | Balance amount available for use                         |
| Holding   | Amount being held while transactions are being processed |

As a given transaction goes through various changes in status as outlined [here](#section/Overview/Transactions), corresponding balance movements will be made. The following table illustrates the relationship between transaction status and balance movements:

| Transaction Status                  | Balance Movement | Description                                                |
| ---                                 | ---              | ---                                                        |
| `CREATED`                           | Authorize        | Transfer wholesale price and fee from available to holding |
| `CANCELLED`, `REJECTED`, `DECLINED` | Void             | Amount authorized in holding moves back to available       |
| `COMPLETED`                         | Capture          | Amount authorized in holding is captured, i.e. debited     |
| `REVERSED`                          | Reverse          | Debited amount is reversed back into available             |

## Flow

Once a product has been selected through one of the [discovery methods](/#tag/Services) provided by the API, the actual transfer (i.e. transaction) can be performed in either one of the following modes:
  - Asynchronous (recommended)
  - Synchronous

Each mode is accessible via a specific endpoint.

As soon as a transaction is confirmed, the transfer order will be sent to the operator for immediate processing. During this time, the transaction will remain in a `CONFIRMED` status until the final status is received from the operator.

### Asynchronous Mode

When a transaction is created and confirmed in an asynchronous fashion, the HTTP connection won't have to be kept open. This preserves system resources on your applications. As such, performing transactions **asynchronously** is **recommended**.

### Synchronous Mode

When a transaction is created and confirmed in a synchronous fashion, the HTTP connection will be kept open in an attempt to capture the final status from the receiving operator so it can be returned in the API response. The processing time usually takes just a few seconds. However, with some receiving operators, it may take longer.

Our system will keep the HTTP connection open for up to 180 seconds (this value can be configured upon request) and will return a status before closing this connection. This status can be in one of the final status (e.g. `COMPLETED`, `DECLINED`) or not (e.g. `SUBMITTED`). In the latter case, this denotes the transaction is still being processed by the receiving operator.

**Note:** your application does not have to wait for the connection to close, it can listen for a shorter period of time and query the final status later on (refer to the \"Final Status\" section below for more details).

### Final Status

Regardless of the processing mode, the application should be designed to capture the final status of a transaction. This can be done through one of the following means:
  - Checking the status of a specific transaction via the corresponding API method (\"pull\" mechanism)
  - Configuring a callback URL passed in the request when creating a transaction (\"push\" mechanism)

## Callbacks

As a transaction is being processed, changes in status will be notified in real-time if a callback URL was provided. Even though one callback per transaction is expected (i.e. change to either `COMPLETED` or `DECLINED`), a manual reversal from the [DT One](https://dtone.com/) team, which may happen in very rare occasions, will also trigger a callback to inform your application of a change in transaction status to `REVERSED`.

This callback endpoint must be implemented by the sending partner, which should expect an HTTP `POST` request containing a transaction object represented in [JSON](https://en.wikipedia.org/wiki/JSON). As callbacks will be sent from the [DT One](https://dtone.com/) servers, these endpoints will have to be publicly-accessible in most cases. During development, a service such as [ngrok](https://ngrok.com/) can be used to expose local servers to the internet.

Upon successful receipt of data, the callback endpoint should respond with an HTTP `2XX` status. In the event that the platform did not receive a successful status, callback notifications will be retried several times, beyond which, the transaction status will have to be queried through the API.

## Status and Errors

### HTTP Status Codes

[DT One](https://dtone.com/) uses standard HTTP response codes to indicate whether an API request was successful or not.

| Status | Description                                        |
| ---    | ---                                                |
| `200`  | OK                                                 |
| `201`  | Created: Resource created                          |
| `202`  | Accepted: Request has been accepted for processing |
| `400`  | Bad Request: Request was malformed                 |
| `401`  | Unauthorized: Credentials missing or invalid       |
| `404`  | Not Found: Resource doesn't exist                  |
| `429`  | Too Many Requests                                  |
| `500`  | Server Error: Error occurred on DT One             |

### API Error Codes

| Code      | Description                                       |
| ---       | ---                                               |
| `1000400` | Bad Request                                       |
| `1000401` | Unauthorized                                      |
| `1000404` | Resource not found                                |
| `1000429` | Too many requests                                 |
| `1003001` | Product is not available in your account          |
| `1003002` | Requested product amount is out of range          |
| `1003003` | Requested product unit is invalid                 |
| `1003101` | Benefits not defined for available products       |
| `1003201` | Promotion not found                               |
| `1003301` | Campaign not found                                |
| `1005003` | Credit party mobile number is invalid             |
| `1005004` | Service not found                                 |
| `1005005` | Country not found                                 |
| `1005006` | Operator not found                                |
| `1005503` | Sender mobile number is invalid                   |
| `1006001` | Insufficient balance                              |
| `1006003` | Debit party mobile number is invalid              |
| `1006009` | Account balance not found                         |
| `1006503` | Beneficiary mobile number is invalid              |
| `1007001` | Transaction external ID has already been used     |
| `1007002` | Transaction has already been confirmed            |
| `1007004` | Transaction can no longer be confirmed            |
| `1007005` | Transaction has already been cancelled            |
| `1007007` | Transaction can no longer be cancelled            |
| `1007500` | Method not supported by operator                  |
| `1008004` | Transaction not found                             |
| `1009001` | Unexpected error, please contact our support team |

### Transaction Status

| Class       | Status                                          | Description                                                            |
| ---         | ---                                             | ---                                                                    |
| `CREATED`   | `CREATED`                                       | Created                                                                |
| `CONFIRMED` | `CONFIRMED`                                     | Confirmed                                                              |
| `REJECTED`  | `REJECTED`                                      | Rejected                                                               |
| `REJECTED`  | `REJECTED-INVALID-CREDIT-PARTY`                 | Rejected - Credit party is invalid                                     |
| `REJECTED`  | `REJECTED-BARRED-CREDIT-PARTY`                  | Rejected - Credit party is barred                                      |
| `REJECTED`  | `REJECTED-INELIGIBLE-CREDIT-PARTY`              | Rejected - Credit party is ineligible for chosen product               |
| `REJECTED`  | `REJECTED-INVALID-DEBIT-PARTY`                  | Rejected - Debit party is invalid                                      |
| `REJECTED`  | `REJECTED-BARRED-DEBIT-PARTY`                   | Rejected - Debit party is barred                                       |
| `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Rejected - Limitations on credit party cumulative transaction amount   |
| `REJECTED`  | `REJECTED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Rejected - Limitations on credit party cumulative transaction quantity |
| `REJECTED`  | `REJECTED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Rejected - Operator currently unavailable                              |
| `REJECTED`  | `REJECTED-INSUFFICIENT-BALANCE`                 | Rejected - Insufficient balance                                        |
| `CANCELLED` | `CANCELLED`                                     | Cancelled                                                              |
| `SUBMITTED` | `SUBMITTED`                                     | Submitted                                                              |
| `COMPLETED` | `COMPLETED`                                     | Completed                                                              |
| `REVERSED`  | `REVERSED`                                      | Reversed                                                               |
| `DECLINED`  | `DECLINED`                                      | Declined (no additional information available)                         |
| `DECLINED`  | `DECLINED-INVALID-CREDIT-PARTY`                 | Declined - Credit party is invalid                                     |
| `DECLINED`  | `DECLINED-BARRED-CREDIT-PARTY`                  | Declined - Credit party is barred                                      |
| `DECLINED`  | `DECLINED-INELIGIBLE-CREDIT-PARTY`              | Declined - Credit party is ineligible for chosen product               |
| `DECLINED`  | `DECLINED-INVALID-DEBIT-PARTY`                  | Declined - Debit party is invalid                                      |
| `DECLINED`  | `DECLINED-BARRED-DEBIT-PARTY`                   | Declined - Debit party is barred                                       |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-AMOUNT`       | Declined - Limitations on operator cumulative transaction amount       |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-AMOUNT`   | Declined - Limitations on credit party cumulative transaction amount   |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-AMOUNT`       | Declined - Limitations on customer cumulative transaction amount       |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-OPERATOR-QUANTITY`     | Declined - Limitations on operator cumulative transaction quantity     |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-CREDIT-PARTY-QUANTITY` | Declined - Limitations on credit party cumulative transaction quantity |
| `DECLINED`  | `DECLINED-LIMITATIONS-ON-CUSTOMER-QUANTITY`     | Declined - Limitations on customer cumulative transaction quantity     |
| `DECLINED`  | `DECLINED-DUPLICATED-TRANSACTION`               | Declined - Duplicated transaction                                      |
| `DECLINED`  | `DECLINED-OPERATOR-CURRENTLY-UNAVAILABLE`       | Declined - Operator currently unavailable                              |

`REJECTED` and `DECLINED` status classes both denote unsuccessful transactions. The primary distinction between these two relates to the party that determined the failure:
  * `REJECTED` are issued by the DVS platform based on various business rules (e.g. insufficient balance, limitations, etc)
  * `DECLINED` are issued by the operators

Separately, it is recommended to define application logic based on **classes**, while additional distinction and/or insight are reflected on the actual **status**.

## Pagination

API resources supporting bulk fetches via \"list\" API methods will be returned in a paginated fashion.

### Input Parameters

| Field      | Required | Type    | Description                                      |
| ---        | ---      | ---     | ---                                              |
| `page`     | No       | Integer | Page number                                      |
| `per_page` | No       | Integer | Number of results per page (default 50, max 100) |

### Output Headers

| Field           | Description                   |
| ---             | ---                           |
| `X-Total`       | Total number of records       |
| `X-Total-Pages` | Total number of pages         |
| `X-Per-Page`    | Number of records per page    |
| `X-Page`        | Current page number           |
| `X-Next-Page`   | Next page number (if any)     |
| `X-Prev-Page`   | Previous page number (if any) |

## Rate Limiting

The API endpoints have rate limiting in place to protect our service from excessive number of requests.

If the limit is reached, an [HTTP error 429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) will be returned by the server.

# Authentication

<!-- ReDoc-Inject: <security-definitions> -->

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.11.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import dvsapi "github.com/josnidhin/go-openapi-codegen-examples/dvsclient/dvsapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), dvsapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), dvsapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), dvsapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dvsapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://staging-dvs-api.transferto.dtone.com:8443/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BalancesApi* | [**BalancesGet**](docs/BalancesApi.md#balancesget) | **Get** /balances | Retrieve balances
*BenefitsApi* | [**BenefitTypesGet**](docs/BenefitsApi.md#benefittypesget) | **Get** /benefit-types | Retrieve list of benefit types
*CampaignsApi* | [**CampaignsCampaignIdGet**](docs/CampaignsApi.md#campaignscampaignidget) | **Get** /campaigns/{campaign_id} | Retrieve campaign by ID
*CampaignsApi* | [**CampaignsGet**](docs/CampaignsApi.md#campaignsget) | **Get** /campaigns | Retrieve list of active campaigns
*CountriesApi* | [**CountriesCountryIsoCodeGet**](docs/CountriesApi.md#countriescountryisocodeget) | **Get** /countries/{country_iso_code} | Retrieve country by ISO code
*CountriesApi* | [**CountriesGet**](docs/CountriesApi.md#countriesget) | **Get** /countries | Retrieve list of countries
*MobileNumberApi* | [**LookupMobileNumberMobileNumberGet**](docs/MobileNumberApi.md#lookupmobilenumbermobilenumberget) | **Get** /lookup/mobile-number/{mobile_number} | Look up operators for a given mobile number
*MobileNumberApi* | [**LookupMobileNumberPost**](docs/MobileNumberApi.md#lookupmobilenumberpost) | **Post** /lookup/mobile-number | Look up operators for a given mobile number
*OperatorsApi* | [**OperatorsGet**](docs/OperatorsApi.md#operatorsget) | **Get** /operators | Retrieve list of operators
*OperatorsApi* | [**OperatorsOperatorIdGet**](docs/OperatorsApi.md#operatorsoperatoridget) | **Get** /operators/{operator_id} | Retrieve operator by ID
*ProductsApi* | [**ProductsGet**](docs/ProductsApi.md#productsget) | **Get** /products | Retrieve list of products
*ProductsApi* | [**ProductsProductIdGet**](docs/ProductsApi.md#productsproductidget) | **Get** /products/{product_id} | Retrieve product by ID
*PromotionsApi* | [**PromotionsGet**](docs/PromotionsApi.md#promotionsget) | **Get** /promotions | Retrieve list of promotions
*PromotionsApi* | [**PromotionsPromotionIdGet**](docs/PromotionsApi.md#promotionspromotionidget) | **Get** /promotions/{promotion_id} | Retrieve promotion by ID
*ServicesApi* | [**ServicesGet**](docs/ServicesApi.md#servicesget) | **Get** /services | Retrieve list of services
*ServicesApi* | [**ServicesServiceIdGet**](docs/ServicesApi.md#servicesserviceidget) | **Get** /services/{service_id} | Retrieve service by ID
*StatementInquiryApi* | [**LookupStatementInquiryPost**](docs/StatementInquiryApi.md#lookupstatementinquirypost) | **Post** /lookup/statement-inquiry | Inquire statements for a given account number
*TransactionsApi* | [**AsyncTransactionsPost**](docs/TransactionsApi.md#asynctransactionspost) | **Post** /async/transactions | Create a transaction asynchronously
*TransactionsApi* | [**AsyncTransactionsTransactionIdConfirmPost**](docs/TransactionsApi.md#asynctransactionstransactionidconfirmpost) | **Post** /async/transactions/{transaction_id}/confirm | Confirm a transaction asynchronously
*TransactionsApi* | [**SyncTransactionsPost**](docs/TransactionsApi.md#synctransactionspost) | **Post** /sync/transactions | Create a transaction synchronously
*TransactionsApi* | [**SyncTransactionsTransactionIdConfirmPost**](docs/TransactionsApi.md#synctransactionstransactionidconfirmpost) | **Post** /sync/transactions/{transaction_id}/confirm | Confirm a transaction synchronously
*TransactionsApi* | [**TransactionsGet**](docs/TransactionsApi.md#transactionsget) | **Get** /transactions | Query list of transactions
*TransactionsApi* | [**TransactionsTransactionIdCancelPost**](docs/TransactionsApi.md#transactionstransactionidcancelpost) | **Post** /transactions/{transaction_id}/cancel | Cancel a transaction
*TransactionsApi* | [**TransactionsTransactionIdGet**](docs/TransactionsApi.md#transactionstransactionidget) | **Get** /transactions/{transaction_id} | Query a transaction by ID


## Documentation For Models

 - [AsyncTransactionsPostRequest](docs/AsyncTransactionsPostRequest.md)
 - [AsyncTransactionsPostRequestAllOf](docs/AsyncTransactionsPostRequestAllOf.md)
 - [AvailabilityZones](docs/AvailabilityZones.md)
 - [Balance](docs/Balance.md)
 - [BenefitType](docs/BenefitType.md)
 - [BenefitTypes](docs/BenefitTypes.md)
 - [BenefitUnitTypes](docs/BenefitUnitTypes.md)
 - [CalculationModes](docs/CalculationModes.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignProductsInner](docs/CampaignProductsInner.md)
 - [Country](docs/Country.md)
 - [Errors](docs/Errors.md)
 - [ErrorsErrorsInner](docs/ErrorsErrorsInner.md)
 - [FixedBenefit](docs/FixedBenefit.md)
 - [FixedBenefitAmount](docs/FixedBenefitAmount.md)
 - [LookupMobileNumberMobileNumberGet200ResponseInner](docs/LookupMobileNumberMobileNumberGet200ResponseInner.md)
 - [LookupMobileNumberMobileNumberGet200ResponseInnerAllOf](docs/LookupMobileNumberMobileNumberGet200ResponseInnerAllOf.md)
 - [LookupMobileNumberPostRequest](docs/LookupMobileNumberPostRequest.md)
 - [LookupStatementInquiryPostRequest](docs/LookupStatementInquiryPostRequest.md)
 - [Operator](docs/Operator.md)
 - [OperatorsGet200ResponseInner](docs/OperatorsGet200ResponseInner.md)
 - [PIN](docs/PIN.md)
 - [PINValidity](docs/PINValidity.md)
 - [Payment](docs/Payment.md)
 - [PaymentModes](docs/PaymentModes.md)
 - [PaymentPostingPeriod](docs/PaymentPostingPeriod.md)
 - [Product](docs/Product.md)
 - [ProductBase](docs/ProductBase.md)
 - [ProductFixedValuePinPurchase](docs/ProductFixedValuePinPurchase.md)
 - [ProductFixedValuePinPurchaseAllOf](docs/ProductFixedValuePinPurchaseAllOf.md)
 - [ProductFixedValueRecharge](docs/ProductFixedValueRecharge.md)
 - [ProductFixedValueRechargeAllOf](docs/ProductFixedValueRechargeAllOf.md)
 - [ProductFixedValueRechargeAllOfPrices](docs/ProductFixedValueRechargeAllOfPrices.md)
 - [ProductFixedValueRechargeAllOfPricesRetail](docs/ProductFixedValueRechargeAllOfPricesRetail.md)
 - [ProductFixedValueRechargeAllOfPricesWholesale](docs/ProductFixedValueRechargeAllOfPricesWholesale.md)
 - [ProductFixedValueRechargeAllOfSource](docs/ProductFixedValueRechargeAllOfSource.md)
 - [ProductFixedValueRechargeAllOfValidity](docs/ProductFixedValueRechargeAllOfValidity.md)
 - [ProductPromotion](docs/ProductPromotion.md)
 - [ProductRangedValuePayment](docs/ProductRangedValuePayment.md)
 - [ProductRangedValuePaymentAllOf](docs/ProductRangedValuePaymentAllOf.md)
 - [ProductRangedValuePinPurchase](docs/ProductRangedValuePinPurchase.md)
 - [ProductRangedValuePinPurchaseAllOf](docs/ProductRangedValuePinPurchaseAllOf.md)
 - [ProductRangedValueRecharge](docs/ProductRangedValueRecharge.md)
 - [ProductRangedValueRechargeAllOf](docs/ProductRangedValueRechargeAllOf.md)
 - [ProductRangedValueRechargeAllOfDestination](docs/ProductRangedValueRechargeAllOfDestination.md)
 - [ProductRangedValueRechargeAllOfPrices](docs/ProductRangedValueRechargeAllOfPrices.md)
 - [ProductRangedValueRechargeAllOfPricesRetail](docs/ProductRangedValueRechargeAllOfPricesRetail.md)
 - [ProductRangedValueRechargeAllOfPricesWholesale](docs/ProductRangedValueRechargeAllOfPricesWholesale.md)
 - [ProductRangedValueRechargeAllOfSource](docs/ProductRangedValueRechargeAllOfSource.md)
 - [ProductTypes](docs/ProductTypes.md)
 - [Promotion](docs/Promotion.md)
 - [PromotionOperator](docs/PromotionOperator.md)
 - [PromotionOperatorCountry](docs/PromotionOperatorCountry.md)
 - [PromotionProductsInner](docs/PromotionProductsInner.md)
 - [Range](docs/Range.md)
 - [RangeBenefitAmountWithoutIncrement](docs/RangeBenefitAmountWithoutIncrement.md)
 - [RangeBenefitAmountWithoutIncrementBase](docs/RangeBenefitAmountWithoutIncrementBase.md)
 - [RangeBenefitAmountWithoutIncrementPromotionBonus](docs/RangeBenefitAmountWithoutIncrementPromotionBonus.md)
 - [RangeWithoutIncrement](docs/RangeWithoutIncrement.md)
 - [RangedBenefit](docs/RangedBenefit.md)
 - [Rates](docs/Rates.md)
 - [Service](docs/Service.md)
 - [ServiceRegion](docs/ServiceRegion.md)
 - [Statement](docs/Statement.md)
 - [StatementBalance](docs/StatementBalance.md)
 - [StatementDates](docs/StatementDates.md)
 - [TimeUnitTypes](docs/TimeUnitTypes.md)
 - [TimeUnits](docs/TimeUnits.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionBenefitsInner](docs/TransactionBenefitsInner.md)
 - [TransactionCreditPartyIdentifier](docs/TransactionCreditPartyIdentifier.md)
 - [TransactionDebitPartyIdentifier](docs/TransactionDebitPartyIdentifier.md)
 - [TransactionDestination](docs/TransactionDestination.md)
 - [TransactionPin](docs/TransactionPin.md)
 - [TransactionPrices](docs/TransactionPrices.md)
 - [TransactionPricesRetail](docs/TransactionPricesRetail.md)
 - [TransactionPricesWholesale](docs/TransactionPricesWholesale.md)
 - [TransactionProduct](docs/TransactionProduct.md)
 - [TransactionProductAllOf](docs/TransactionProductAllOf.md)
 - [TransactionProductAllOfPin](docs/TransactionProductAllOfPin.md)
 - [TransactionRates](docs/TransactionRates.md)
 - [TransactionRequestedValues](docs/TransactionRequestedValues.md)
 - [TransactionRequestedValuesSource](docs/TransactionRequestedValuesSource.md)
 - [TransactionSender](docs/TransactionSender.md)
 - [TransactionSource](docs/TransactionSource.md)
 - [TransactionStatementIdentifier](docs/TransactionStatementIdentifier.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionStatusClass](docs/TransactionStatusClass.md)
 - [UnitTypes](docs/UnitTypes.md)


## Documentation For Authorization



### BasicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



