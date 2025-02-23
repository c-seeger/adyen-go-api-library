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
Cancels Cancel an authorised payment
Cancels the authorisation on a payment that has not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**TECHNICAL_CANCEL** webhook](https://docs.adyen.com/online-payments/cancel#cancellation-webhook).  If you want to cancel a payment using the [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference), use the [&#x60;/payments/{paymentPspReference}/cancels&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/cancels) endpoint instead.  If you want to cancel a payment but are not sure whether it has been captured, use the [&#x60;/payments/{paymentPspReference}/reversals&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.  For more information, refer to [Cancel](https://docs.adyen.com/online-payments/cancel).
 * @param req CreateStandalonePaymentCancelRequest - reference of CreateStandalonePaymentCancelRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return StandalonePaymentCancelResource
*/
func (a Checkout) Cancels(req *CreateStandalonePaymentCancelRequest, ctxs ..._context.Context) (StandalonePaymentCancelResource, *_nethttp.Response, error) {
    res := &StandalonePaymentCancelResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/cancels", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsPaymentPspReferenceAmountUpdates Update an authorised amount
Increases or decreases the authorised payment amount and returns a unique reference for this request. You get the outcome of the request asynchronously, in an [**AUTHORISATION_ADJUSTMENT** webhook](https://docs.adyen.com/development-resources/webhooks/understand-notifications#event-codes).  You can only update authorised amounts that have not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures).  The amount you specify in the request is the updated amount, which is larger or smaller than the initial authorised amount.  For more information, refer to [Authorisation adjustment](https://docs.adyen.com/online-payments/adjust-authorisation#use-cases).
 * @param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment.
 * @param req CreatePaymentAmountUpdateRequest - reference of CreatePaymentAmountUpdateRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentAmountUpdateResource
*/
func (a Checkout) PaymentsPaymentPspReferenceAmountUpdates(paymentPspReference *string, req *CreatePaymentAmountUpdateRequest, ctxs ..._context.Context) (PaymentAmountUpdateResource, *_nethttp.Response, error) {
    res := &PaymentAmountUpdateResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/{paymentPspReference}/amountUpdates", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsPaymentPspReferenceCancels Cancel an authorised payment
Cancels the authorisation on a payment that has not yet been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/paymentPspReference/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CANCELLATION** webhook](https://docs.adyen.com/online-payments/cancel#cancellation-webhook).  If you want to cancel a payment but don&#39;t have the [&#x60;pspReference&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference), use the [&#x60;/cancels&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/cancels) endpoint instead.  If you want to cancel a payment but are not sure whether it has been captured, use the [&#x60;/payments/{paymentPspReference}/reversals&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.  For more information, refer to [Cancel](https://docs.adyen.com/online-payments/cancel).
 * @param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to cancel. 
 * @param req CreatePaymentCancelRequest - reference of CreatePaymentCancelRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentCancelResource
*/
func (a Checkout) PaymentsPaymentPspReferenceCancels(paymentPspReference *string, req *CreatePaymentCancelRequest, ctxs ..._context.Context) (PaymentCancelResource, *_nethttp.Response, error) {
    res := &PaymentCancelResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/{paymentPspReference}/cancels", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsPaymentPspReferenceCaptures Capture an authorised payment
 Captures an authorised payment and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CAPTURE** webhook](https://docs.adyen.com/online-payments/capture#capture-notification).  You can capture either the full authorised amount or a part of the authorised amount. By default, any unclaimed amount after a partial capture gets cancelled. This does not apply if you enabled multiple partial captures on your account and the payment method supports multiple partial captures.   [Automatic capture](https://docs.adyen.com/online-payments/capture#automatic-capture) is the default setting for most payment methods. In these cases, you don&#39;t need to make capture requests. However, making capture requests for payments that are captured automatically does not result in double charges.  For more information, refer to [Capture](https://docs.adyen.com/online-payments/capture).
 * @param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to capture.
 * @param req CreatePaymentCaptureRequest - reference of CreatePaymentCaptureRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentCaptureResource
*/
func (a Checkout) PaymentsPaymentPspReferenceCaptures(paymentPspReference *string, req *CreatePaymentCaptureRequest, ctxs ..._context.Context) (PaymentCaptureResource, *_nethttp.Response, error) {
    res := &PaymentCaptureResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/{paymentPspReference}/captures", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsPaymentPspReferenceRefunds Refund a captured payment
Refunds a payment that has been [captured](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/captures), and returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**REFUND** webhook](https://docs.adyen.com/online-payments/refund#refund-webhook).  You can refund either the full captured amount or a part of the captured amount. You can also perform multiple partial refunds, as long as their sum doesn&#39;t exceed the captured amount.  &gt; Some payment methods do not support partial refunds. To learn if a payment method supports partial refunds, refer to the payment method page such as [cards](https://docs.adyen.com/payment-methods/cards#supported-cards), [iDEAL](https://docs.adyen.com/payment-methods/ideal), or [Klarna](https://docs.adyen.com/payment-methods/klarna).   If you want to refund a payment but are not sure whether it has been captured, use the [&#x60;/payments/{paymentPspReference}/reversals&#x60;](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/reversals) endpoint instead.  For more information, refer to [Refund](https://docs.adyen.com/online-payments/refund).
 * @param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to refund.
 * @param req CreatePaymentRefundRequest - reference of CreatePaymentRefundRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentRefundResource
*/
func (a Checkout) PaymentsPaymentPspReferenceRefunds(paymentPspReference *string, req *CreatePaymentRefundRequest, ctxs ..._context.Context) (PaymentRefundResource, *_nethttp.Response, error) {
    res := &PaymentRefundResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/{paymentPspReference}/refunds", ctxs...)
    return *res, httpRes, err
}


/*
PaymentsPaymentPspReferenceReversals Refund or cancel a payment
[Refunds](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/refunds) a payment if it has already been captured, and [cancels](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments/{paymentPspReference}/cancels) a payment if it has not yet been captured. Returns a unique reference for this request. You get the outcome of the request asynchronously, in a [**CANCEL_OR_REFUND** webhook](https://docs.adyen.com/online-payments/reverse#cancel-or-refund-webhook).  The reversed amount is always the full payment amount. &gt; Do not use this request for payments that involve multiple partial captures.  For more information, refer to [Reversal](https://docs.adyen.com/online-payments/reversal).
 * @param paymentPspReference The [`pspReference`](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/payments__resParam_pspReference) of the payment that you want to reverse. 
 * @param req CreatePaymentReversalRequest - reference of CreatePaymentReversalRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentReversalResource
*/
func (a Checkout) PaymentsPaymentPspReferenceReversals(paymentPspReference *string, req *CreatePaymentReversalRequest, ctxs ..._context.Context) (PaymentReversalResource, *_nethttp.Response, error) {
    res := &PaymentReversalResource{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/payments/{paymentPspReference}/reversals", ctxs...)
    return *res, httpRes, err
}

