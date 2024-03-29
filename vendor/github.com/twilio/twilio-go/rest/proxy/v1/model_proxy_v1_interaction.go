/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ProxyV1Interaction struct for ProxyV1Interaction
type ProxyV1Interaction struct {
	// The unique string that we created to identify the Interaction resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource.
	SessionSid *string `json:"session_sid,omitempty"`
	// The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Interaction resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// A JSON string that includes the message body of message interactions (e.g. `{\"body\": \"hello\"}`) or the call duration (when available) of a call (e.g. `{\"duration\": \"5\"}`).
	Data *string `json:"data,omitempty"`
	Type *string `json:"type,omitempty"`
	// The SID of the inbound [Participant](https://www.twilio.com/docs/proxy/api/participant) resource.
	InboundParticipantSid *string `json:"inbound_participant_sid,omitempty"`
	// The SID of the inbound resource; either the [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource).
	InboundResourceSid    *string `json:"inbound_resource_sid,omitempty"`
	InboundResourceStatus *string `json:"inbound_resource_status,omitempty"`
	// The inbound resource type. Can be [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource).
	InboundResourceType *string `json:"inbound_resource_type,omitempty"`
	// The URL of the Twilio inbound resource
	InboundResourceUrl *string `json:"inbound_resource_url,omitempty"`
	// The SID of the outbound [Participant](https://www.twilio.com/docs/proxy/api/participant)).
	OutboundParticipantSid *string `json:"outbound_participant_sid,omitempty"`
	// The SID of the outbound resource; either the [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource).
	OutboundResourceSid    *string `json:"outbound_resource_sid,omitempty"`
	OutboundResourceStatus *string `json:"outbound_resource_status,omitempty"`
	// The outbound resource type. Can be: [Call](https://www.twilio.com/docs/voice/api/call-resource) or [Message](https://www.twilio.com/docs/sms/api/message-resource).
	OutboundResourceType *string `json:"outbound_resource_type,omitempty"`
	// The URL of the Twilio outbound resource.
	OutboundResourceUrl *string `json:"outbound_resource_url,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the Interaction was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Interaction resource.
	Url *string `json:"url,omitempty"`
}
