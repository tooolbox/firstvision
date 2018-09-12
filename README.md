# FirstApiSDK

FirstApiSdk - JavaScript client for FirstApiSDK
Payment Gateway API for payment processing. 
This SDK is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: 0.0.1
- Package version: 6.3.0
- Build package: io.swagger.codegen.languages.JavascriptClientCodegen

## Installation

### For [Node.js](https://nodejs.org/)

#### npm

To publish the library as a [npm](https://www.npmjs.com/),
please follow the procedure in ["Publishing npm packages"](https://docs.npmjs.com/getting-started/publishing-npm-packages).

Then install it via:

```shell
npm install FirstApiSDK --save
```

#### git
#
If the library is hosted at a git repository, e.g.
https://github.com/GIT_USER_ID/GIT_REPO_ID
then install it via:

```shell
    npm install GIT_USER_ID/GIT_REPO_ID --save
```

### For browser

The library also works in the browser environment via npm and [browserify](http://browserify.org/). After following
the above steps with Node.js and installing browserify with `npm install -g browserify`,
perform the following (assuming *main.js* is your entry file):

```shell
browserify main.js > bundle.js
```

Then include *bundle.js* in the HTML pages.

### Webpack Configuration

Using Webpack you may encounter the following error: "Module not found: Error:
Cannot resolve module", most certainly you should disable AMD loader. Add/merge
the following section to your webpack config:

```javascript
module: {
  rules: [
    {
      parser: {
        amd: false
      }
    }
  ]
}
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following JS code:

```javascript
var FirstApiSDK = require('FirstApiSDK');
var SimpleClient = require('FirstApiSDK/simple/SimpleClient');

var api = new SimpleClient({
  apiSecret: "apiSecret",
  apiKey: "apiKey",
});

api.requestAccessToken().then(function (data) {
  console.log('API called successfully. Returned data: ' + data);
}, function (error) {
  console.error(error);
});
```

## Documentation for API Endpoints

All URIs are relative to *https://cert.api.firstdata.com/gateway*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SimpleClient* | [**requestAccessToken**](docs/SimpleClient.md#requestAccessToken) | **POST** /v1/authentication/access-tokens | Generate an access token for user authentication
*SimpleClient* | [**performPaymentPostAuthorisationByOrder**](docs/SimpleClient.md#performPaymentPostAuthorisationByOrder) | **POST** /v1/orders/{order-id}/postauth | Use this to capture/complete a transaction. Partial postauths are allowed.
*SimpleClient* | [**returnTransactionByOrder**](docs/SimpleClient.md#returnTransactionByOrder) | **POST** /v1/orders/{order-id}/return | Use this to return/refund on the order. Partial returns are allowed.
*SimpleClient* | [**performPaymentPostAuthorisationByTransaction**](docs/SimpleClient.md#performPaymentPostAuthorisationByTransaction) | **POST** /v1/payments/{transaction-id}/postauth | Use this to capture/complete a transaction. Partial postauths are allowed.
*SimpleClient* | [**primaryPaymentTransaction**](docs/SimpleClient.md#primaryPaymentTransaction) | **POST** /v1/payments | Generate a primary transaction
*SimpleClient* | [**returnTransaction**](docs/SimpleClient.md#returnTransaction) | **POST** /v1/payments/{transaction-id}/return | Return/refund a transaction.
*SimpleClient* | [**transactionInquiry**](docs/SimpleClient.md#transactionInquiry) | **GET** /v1/payments/{transaction-id} | Retrieve the state of a transaction
*SimpleClient* | [**voidTransaction**](docs/SimpleClient.md#voidTransaction) | **POST** /v1/payments/{transaction-id}/void | Reverse a previous action on an existing transaction


## Documentation for Models

 - [FirstApiSdk.AccessTokenResponse](docs/AccessTokenResponse.md)
 - [FirstApiSdk.Address](docs/Address.md)
 - [FirstApiSdk.Airline](docs/Airline.md)
 - [FirstApiSdk.AirlineAncillaryServiceCategory](docs/AirlineAncillaryServiceCategory.md)
 - [FirstApiSdk.AirlineTravelRoute](docs/AirlineTravelRoute.md)
 - [FirstApiSdk.Amount](docs/Amount.md)
 - [FirstApiSdk.AmountComponents](docs/AmountComponents.md)
 - [FirstApiSdk.AuthenticationResponseVerification](docs/AuthenticationResponseVerification.md)
 - [FirstApiSdk.BasketItem](docs/BasketItem.md)
 - [FirstApiSdk.Billing](docs/Billing.md)
 - [FirstApiSdk.CarRental](docs/CarRental.md)
 - [FirstApiSdk.CarRentalExtraCharges](docs/CarRentalExtraCharges.md)
 - [FirstApiSdk.CardVerificationsTransaction](docs/CardVerificationsTransaction.md)
 - [FirstApiSdk.ClientLocale](docs/ClientLocale.md)
 - [FirstApiSdk.Contact](docs/Contact.md)
 - [FirstApiSdk.Error](docs/Error.md)
 - [FirstApiSdk.ErrorDetails](docs/ErrorDetails.md)
 - [FirstApiSdk.ErrorResponse](docs/ErrorResponse.md)
 - [FirstApiSdk.Expiration](docs/Expiration.md)
 - [FirstApiSdk.Frequency](docs/Frequency.md)
 - [FirstApiSdk.IndustrySpecificExtensions](docs/IndustrySpecificExtensions.md)
 - [FirstApiSdk.InstallmentOptions](docs/InstallmentOptions.md)
 - [FirstApiSdk.Lodging](docs/Lodging.md)
 - [FirstApiSdk.LodgingExtraCharges](docs/LodgingExtraCharges.md)
 - [FirstApiSdk.Order](docs/Order.md)
 - [FirstApiSdk.PayPal](docs/PayPal.md)
 - [FirstApiSdk.PaymentCard](docs/PaymentCard.md)
 - [FirstApiSdk.PaymentCardAuthenticationRequest](docs/PaymentCardAuthenticationRequest.md)
 - [FirstApiSdk.PaymentCardAuthenticationResult](docs/PaymentCardAuthenticationResult.md)
 - [FirstApiSdk.PaymentMethod](docs/PaymentMethod.md)
 - [FirstApiSdk.PaymentSchedulesRequest](docs/PaymentSchedulesRequest.md)
 - [FirstApiSdk.PaymentSchedulesResponse](docs/PaymentSchedulesResponse.md)
 - [FirstApiSdk.PaymentUrlRequest](docs/PaymentUrlRequest.md)
 - [FirstApiSdk.PaymentUrlResponse](docs/PaymentUrlResponse.md)
 - [FirstApiSdk.PrimaryTransaction](docs/PrimaryTransaction.md)
 - [FirstApiSdk.PrimaryTransactionAdditionalDetails](docs/PrimaryTransactionAdditionalDetails.md)
 - [FirstApiSdk.ProcessorData](docs/ProcessorData.md)
 - [FirstApiSdk.ResponseType](docs/ResponseType.md)
 - [FirstApiSdk.SecondaryTransaction](docs/SecondaryTransaction.md)
 - [FirstApiSdk.Sepa](docs/Sepa.md)
 - [FirstApiSdk.SepaMandate](docs/SepaMandate.md)
 - [FirstApiSdk.Shipping](docs/Shipping.md)
 - [FirstApiSdk.SplitShipment](docs/SplitShipment.md)
 - [FirstApiSdk.StoredCredential](docs/StoredCredential.md)
 - [FirstApiSdk.TransactionResponse](docs/TransactionResponse.md)
 - [FirstApiSdk.TransactionResponseAuthenticationRedirect](docs/TransactionResponseAuthenticationRedirect.md)
 - [FirstApiSdk.TransactionResponseAuthenticationRedirectParams](docs/TransactionResponseAuthenticationRedirectParams.md)
 - [FirstApiSdk.TransactionType](docs/TransactionType.md)
 - [FirstApiSdk.TransactionErrorResponse](docs/TransactionErrorResponse.md)


## Documentation for Authorization

 All endpoints do not require authorization.

