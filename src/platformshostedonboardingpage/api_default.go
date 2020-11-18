/*
 * Adyen for Platforms: Hosted Onboarding
 *
 * Hosted Onboarding provides endpoints for using the Hosted Onboarding Page. The related entities include account holders only.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/onboarding-and-verification/hosted-onboarding-page). ## Authentication To connect to the HOP API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The HOP API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Hop/v6/getOnboardingUrl ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformshostedonboardingpage

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v4/src/common"
)

// PlatformsHostedOnboardingPage PlatformsHostedOnboardingPage service
type PlatformsHostedOnboardingPage common.Service

/*
PostGetOnboardingUrl Get a link to a Hosted Onboarding Page.
Returns a link to a Hosted Onboarding Page (HOP) to be used by a specific account holder. Each account holder represents a single sub-merchant.  For more information on how to use HOP, refer to [Hosted Onboarding Page](https://docs.adyen.com/platforms/onboarding-and-verification/hosted-onboarding-page). 
 * @param request GetOnboardingUrlRequest - reference of GetOnboardingUrlRequest). 
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return GetOnboardingUrlResponse
*/
func (a PlatformsHostedOnboardingPage) GetOnboardingUrl(req *GetOnboardingUrlRequest, ctxs ..._context.Context) (GetOnboardingUrlResponse, *_nethttp.Response, error) {
    res := &GetOnboardingUrlResponse{}
    httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath() + "/getOnboardingUrl", ctxs...)
    return *res, httpRes, err
}
