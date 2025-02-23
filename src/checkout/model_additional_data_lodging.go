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
// AdditionalDataLodging struct for AdditionalDataLodging
type AdditionalDataLodging struct {
	// The arrival date. * Date format: `yyyyMMdd`
	LodgingCheckInDate string `json:"lodging.checkInDate,omitempty"`
	// The departure date. * Date format: `yyyyMMdd`
	LodgingCheckOutDate string `json:"lodging.checkOutDate,omitempty"`
	// The toll free phone number for the hotel/lodgings. * Format: Alphanumeric * maxLength: 17
	LodgingCustomerServiceTollFreeNumber string `json:"lodging.customerServiceTollFreeNumber,omitempty"`
	// Identifies that the facility complies with the Hotel and Motel Fire Safety Act of 1990. Values can be: 'Y' or 'N'. * Format: Alphabetic * maxLength: 1
	LodgingFireSafetyActIndicator string `json:"lodging.fireSafetyActIndicator,omitempty"`
	// The folio cash advances. * Format: Numeric * maxLength: 12
	LodgingFolioCashAdvances string `json:"lodging.folioCashAdvances,omitempty"`
	// Card acceptor’s internal invoice or billing ID reference number. * Format: Alphanumeric * maxLength: 25
	LodgingFolioNumber string `json:"lodging.folioNumber,omitempty"`
	// Any charges for food and beverages associated with the booking. * Format: Numeric * maxLength: 12
	LodgingFoodBeverageCharges string `json:"lodging.foodBeverageCharges,omitempty"`
	// Indicates if the customer was a \"no-show\" (neither keeps nor cancels their booking).  Value should be Y or N. * Format: Numeric * maxLength: 1
	LodgingNoShowIndicator string `json:"lodging.noShowIndicator,omitempty"`
	// Prepaid expenses for the booking. * Format: Numeric * maxLength: 12
	LodgingPrepaidExpenses string `json:"lodging.prepaidExpenses,omitempty"`
	// Identifies specific lodging property location by its local phone number. * Format: Alphanumeric * maxLength: 17
	LodgingPropertyPhoneNumber string `json:"lodging.propertyPhoneNumber,omitempty"`
	// Total number of nights the room will be rented. * Format: Numeric * maxLength: 4
	LodgingRoom1NumberOfNights string `json:"lodging.room1.numberOfNights,omitempty"`
	// The rate of the room. * Format: Numeric * maxLength: 12
	LodgingRoom1Rate string `json:"lodging.room1.rate,omitempty"`
	// The total amount of tax to be paid. * Format: Numeric * maxLength: 12
	LodgingRoom1Tax string `json:"lodging.room1.tax,omitempty"`
	// Total room tax amount. * Format: Numeric * maxLength: 12
	LodgingTotalRoomTax string `json:"lodging.totalRoomTax,omitempty"`
	// Total tax amount. * Format: Numeric * maxLength: 12
	LodgingTotalTax string `json:"lodging.totalTax,omitempty"`
	// Number of nights. This should be included in the auth message. * Format: Numeric * maxLength: 2
	TravelEntertainmentAuthDataDuration string `json:"travelEntertainmentAuthData.duration,omitempty"`
	// Indicates what market-specific dataset will be submitted or is being submitted. Value should be \"H\" for Hotel. This should be included in the auth message.  * Format: Alphanumeric * maxLength: 1
	TravelEntertainmentAuthDataMarket string `json:"travelEntertainmentAuthData.market,omitempty"`
}
