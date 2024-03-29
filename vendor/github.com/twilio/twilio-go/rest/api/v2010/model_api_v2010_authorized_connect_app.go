/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010AuthorizedConnectApp struct for ApiV2010AuthorizedConnectApp
type ApiV2010AuthorizedConnectApp struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The company name set for the Connect App.
	ConnectAppCompanyName *string `json:"connect_app_company_name,omitempty"`
	// A detailed description of the Connect App.
	ConnectAppDescription *string `json:"connect_app_description,omitempty"`
	// The name of the Connect App.
	ConnectAppFriendlyName *string `json:"connect_app_friendly_name,omitempty"`
	// The public URL for the Connect App.
	ConnectAppHomepageUrl *string `json:"connect_app_homepage_url,omitempty"`
	// The SID that we assigned to the Connect App.
	ConnectAppSid *string `json:"connect_app_sid,omitempty"`
	// The set of permissions that you authorized for the Connect App.  Can be: `get-all` or `post-all`.
	Permissions *[]string `json:"permissions,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
}
