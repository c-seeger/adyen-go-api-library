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
	"github.com/adyen/adyen-go-api-library/v6/src/common"
)

// PaymentDetailsResponse struct for PaymentDetailsResponse
type PaymentDetailsResponse struct {
	// Contains additional information about the payment. Some data fields are included only if you select them first: Go to **Customer Area** > **Account** > **API URLs** > **Additional data settings**.
	AdditionalData map[string]string `json:"additionalData,omitempty"`
	Amount *Amount `json:"amount,omitempty"`
	// Donation Token containing payment details for Adyen Giving.
	DonationToken string `json:"donationToken,omitempty"`
	FraudResult *FraudResult `json:"fraudResult,omitempty"`
	// The reference used during the /payments request.
	MerchantReference string `json:"merchantReference,omitempty"`
	Order *CheckoutOrderResponse `json:"order,omitempty"`
	PaymentMethod *ResponsePaymentMethod `json:"paymentMethod,omitempty"`
	// Adyen's 16-character string reference associated with the transaction/request. This value is globally unique; quote it when communicating with us about this request.
	PspReference string `json:"pspReference,omitempty"`
	// If the payment's authorisation is refused or an error occurs during authorisation, this field holds Adyen's mapped reason for the refusal or a description of the error. When a transaction fails, the authorisation response includes `resultCode` and `refusalReason` values.  For more information, see [Refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReason string `json:"refusalReason,omitempty"`
	// Code that specifies the refusal reason. For more information, see [Authorisation refusal reasons](https://docs.adyen.com/development-resources/refusal-reasons).
	RefusalReasonCode string `json:"refusalReasonCode,omitempty"`
	// The result of the payment. For more information, see [Result codes](https://docs.adyen.com/online-payments/payment-result-codes).  Possible values:  * **AuthenticationFinished** – The payment has been successfully authenticated with 3D Secure 2. Returned for 3D Secure 2 authentication-only transactions. * **AuthenticationNotRequired** – The transaction does not require 3D Secure authentication. Returned for [standalone authentication-only integrations](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only). * **Authorised** – The payment was successfully authorised. This state serves as an indicator to proceed with the delivery of goods and services. This is a final state. * **Cancelled** – Indicates the payment has been cancelled (either by the shopper or the merchant) before processing was completed. This is a final state. * **ChallengeShopper** – The issuer requires further shopper interaction before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Error** – There was an error when the payment was being processed. The reason is given in the `refusalReason` field. This is a final state. * **IdentifyShopper** – The issuer requires the shopper's device fingerprint before the payment can be authenticated. Returned for 3D Secure 2 transactions. * **Pending** – Indicates that it is not possible to obtain the final status of the payment. This can happen if the systems providing final status information for the payment are unavailable, or if the shopper needs to take further action to complete the payment. * **PresentToShopper** – Indicates that the response contains additional information that you need to present to a shopper, so that they can use it to complete a payment. * **Received** – Indicates the payment has successfully been received by Adyen, and will be processed. This is the initial state for all payments. * **RedirectShopper** – Indicates the shopper should be redirected to an external web page or app to complete the authorisation. * **Refused** – Indicates the payment was refused. The reason is given in the `refusalReason` field. This is a final state.
	ResultCode *common.ResultCode `json:"resultCode,omitempty"`
	//ResultCode string `json:"resultCode,omitempty"`
	// The shopperLocale.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	ThreeDS2ResponseData *ThreeDS2ResponseData `json:"threeDS2ResponseData,omitempty"`
	ThreeDS2Result *ThreeDS2Result `json:"threeDS2Result,omitempty"`
	// When non-empty, contains a value that you must submit to the `/payments/details` endpoint as `paymentData`.
	ThreeDSPaymentData string `json:"threeDSPaymentData,omitempty"`
}
