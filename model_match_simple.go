/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MatchSimple struct {
	// TBA match key with the format `yyyy[EVENT_CODE]_[COMP_LEVEL]m[MATCH_NUMBER]`, where `yyyy` is the year, and `EVENT_CODE` is the event code of the event, `COMP_LEVEL` is (qm, ef, qf, sf, f), and `MATCH_NUMBER` is the match number in the competition level. A set number may append the competition level if more than one match in required per set.
	Key string `json:"key"`
	// The competition level the match was played at.
	CompLevel string `json:"comp_level"`
	// The set number in a series of matches where more than one match is required in the match series.
	SetNumber int32 `json:"set_number"`
	// The match number of the match in the competition level.
	MatchNumber int32 `json:"match_number"`
	Alliances MatchAlliances `json:"alliances,omitempty"`
	// The color (red/blue) of the winning alliance. Will contain an empty string in the event of no winner, or a tie.
	WinningAlliance string `json:"winning_alliance,omitempty"`
	// Event key of the event the match was played at.
	EventKey string `json:"event_key"`
	// UNIX timestamp (seconds since 1-Jan-1970 00:00:00) of the scheduled match time, as taken from the published schedule.
	Time int64 `json:"time,omitempty"`
	// UNIX timestamp (seconds since 1-Jan-1970 00:00:00) of the TBA predicted match start time.
	PredictedTime int64 `json:"predicted_time,omitempty"`
	// UNIX timestamp (seconds since 1-Jan-1970 00:00:00) of actual match start time.
	ActualTime int64 `json:"actual_time,omitempty"`
}
