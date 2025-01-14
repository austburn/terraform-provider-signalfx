/*
 * Detectors API
 *
 * **Detectors** define rules for identifying conditions of interest to the customer, and the notifications to send when the conditions occur or stop occurring.
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package detector

// Contains properties for an alert notification sent via HipChat
type HipChatNotification struct {
	// Tells SignalFx which external system it should use to send the notification. For a HipChat notification, this is always \"HipChat\".
	Type string `json:"type"`
	// The SignalFx ID of the integration profile for HipChat. Use the Integrations API to get the credential ID for your HipChat integration.
	CredentialId string `json:"credentialId"`
	// Name of the HipChat room in which to display the notification message.
	Room string `json:"room,omitempty"`
}
