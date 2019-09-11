/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EventDistrictPointsTiebreakers struct for EventDistrictPointsTiebreakers
type EventDistrictPointsTiebreakers struct {
	HighestQualScores []int32 `json:"highest_qual_scores,omitempty"`
	QualWins int32 `json:"qual_wins,omitempty"`
}
