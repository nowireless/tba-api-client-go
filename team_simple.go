/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TeamSimple struct {

	// TBA team key with the format `frcXXXX` with `XXXX` representing the team number.
	Key string `json:"key"`

	// Official team number issued by FIRST.
	TeamNumber int32 `json:"team_number"`

	// Team nickname provided by FIRST.
	Nickname string `json:"nickname,omitempty"`

	// Official long name registered with FIRST.
	Name string `json:"name"`

	// City of team derived from parsing the address registered with FIRST.
	City string `json:"city,omitempty"`

	// State of team derived from parsing the address registered with FIRST.
	StateProv string `json:"state_prov,omitempty"`

	// Country of team derived from parsing the address registered with FIRST.
	Country string `json:"country,omitempty"`
}