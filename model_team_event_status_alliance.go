/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TeamEventStatusAlliance struct {
	// Alliance name, may be null.
	Name string `json:"name,omitempty"`
	// Alliance number.
	Number int32 `json:"number"`
	Backup TeamEventStatusAllianceBackup `json:"backup,omitempty"`
	// Order the team was picked in the alliance from 0-2, with 0 being alliance captain.
	Pick int32 `json:"pick"`
}
