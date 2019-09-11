/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EliminationAlliance struct for EliminationAlliance
type EliminationAlliance struct {
	// Alliance name, may be null.
	Name string `json:"name,omitempty"`
	Backup EliminationAllianceBackup `json:"backup,omitempty"`
	// List of teams that declined the alliance.
	Declines []string `json:"declines,omitempty"`
	// List of team keys picked for the alliance. First pick is captain.
	Picks []string `json:"picks"`
	Status EliminationAllianceStatus `json:"status,omitempty"`
}
