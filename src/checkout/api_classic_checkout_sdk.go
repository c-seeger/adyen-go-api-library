/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports [versioning](https://docs.adyen.com/development-resources/versioning) using a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v69/payments ```  ## Release notes Have a look at the [release notes](https://docs.adyen.com/online-payments/release-notes?integration_type=api&version=69) to find out what changed in this version!
 *
 * API version: 69
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

import (
	_context "context"
    _nethttp "net/http"
)


/*
PaymentSession Create a payment session
Provides the data object that can be used to start the Checkout SDK. To set up the payment, pass its amount, currency, and other required parameters. We use this to optimise the payment flow and perform better risk assessment of the transaction.  For more information, refer to [How it works](https://docs.adyen.com/online-payments#howitworks).
 * @param req PaymentSetupRequest - reference of PaymentSetupRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentSetupResponse
*/
func (a Checkout) PaymentSession(req *PaymentSetupRequest, ctxs ..._context.Context) (PaymentSetupResponse, *_nethttp.Response, error) {
    res := &PaymentSetupResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/paymentSession", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsResult Verify a payment result
Verifies the payment result using the payload returned from the Checkout SDK.  For more information, refer to [How it works](https://docs.adyen.com/online-payments#howitworks).
 * @param req PaymentVerificationRequest - reference of PaymentVerificationRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentVerificationResponse
*/
func (a Checkout) PaymentsResult(req *PaymentVerificationRequest, ctxs ..._context.Context) (PaymentVerificationResponse, *_nethttp.Response, error) {
    res := &PaymentVerificationResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/result", ctxs...)
    return *res, httpRes, err
}

