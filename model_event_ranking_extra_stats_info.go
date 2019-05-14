/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EventRankingExtraStatsInfo struct {
	// Integer expressing the number of digits of precision in the number provided in `sort_orders`.
	Precision float32 `json:"precision"`
	// Name of the field used in the `extra_stats` array.
	Name string `json:"name"`
}