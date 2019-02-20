/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Award struct {

	// The name of the award as provided by FIRST. May vary for the same award type.
	Name string `json:"name"`

	// Type of award given. See https://github.com/the-blue-alliance/the-blue-alliance/blob/master/consts/award_type.py#L6
	AwardType int32 `json:"award_type"`

	// The event_key of the event the award was won at.
	EventKey string `json:"event_key"`

	// A list of recipients of the award at the event. Either team_key and/or awardee for individual awards.
	RecipientList []AwardRecipient `json:"recipient_list"`

	// The year this award was won.
	Year int32 `json:"year"`
}