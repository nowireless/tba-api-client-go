/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MatchSimpleAlliances A list of alliances, the teams on the alliances, and their score.
type MatchSimpleAlliances struct {
	Red MatchAlliance `json:"red,omitempty"`
	Blue MatchAlliance `json:"blue,omitempty"`
}
