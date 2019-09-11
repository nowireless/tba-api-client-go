/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// TeamEventStatusPlayoff Playoff status for this team, may be null if the team did not make playoffs, or playoffs have not begun.
type TeamEventStatusPlayoff struct {
	// The highest playoff level the team reached.
	Level string `json:"level,omitempty"`
	CurrentLevelRecord WltRecord `json:"current_level_record,omitempty"`
	Record WltRecord `json:"record,omitempty"`
	// Current competition status for the playoffs.
	Status string `json:"status,omitempty"`
	// The average match score during playoffs. Year specific. May be null if not relevant for a given year.
	PlayoffAverage int32 `json:"playoff_average,omitempty"`
}
