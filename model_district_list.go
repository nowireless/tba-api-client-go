/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DistrictList struct for DistrictList
type DistrictList struct {
	// The short identifier for the district.
	Abbreviation string `json:"abbreviation"`
	// The long name for the district.
	DisplayName string `json:"display_name"`
	// Key for this district, e.g. `2016ne`.
	Key string `json:"key"`
	// Year this district participated.
	Year int32 `json:"year"`
}
