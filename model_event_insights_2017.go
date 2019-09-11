/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EventInsights2017 Insights for FIRST STEAMWORKS qualification and elimination matches.
type EventInsights2017 struct {
	// Average foul score.
	AverageFoulScore float32 `json:"average_foul_score"`
	// Average fuel points scored.
	AverageFuelPoints float32 `json:"average_fuel_points"`
	// Average fuel points scored during auto.
	AverageFuelPointsAuto float32 `json:"average_fuel_points_auto"`
	// Average fuel points scored during teleop.
	AverageFuelPointsTeleop float32 `json:"average_fuel_points_teleop"`
	// Average points scored in the high goal.
	AverageHighGoals float32 `json:"average_high_goals"`
	// Average points scored in the high goal during auto.
	AverageHighGoalsAuto float32 `json:"average_high_goals_auto"`
	// Average points scored in the high goal during teleop.
	AverageHighGoalsTeleop float32 `json:"average_high_goals_teleop"`
	// Average points scored in the low goal.
	AverageLowGoals float32 `json:"average_low_goals"`
	// Average points scored in the low goal during auto.
	AverageLowGoalsAuto float32 `json:"average_low_goals_auto"`
	// Average points scored in the low goal during teleop.
	AverageLowGoalsTeleop float32 `json:"average_low_goals_teleop"`
	// Average mobility points scored during auto.
	AverageMobilityPointsAuto float32 `json:"average_mobility_points_auto"`
	// Average points scored during auto.
	AveragePointsAuto float32 `json:"average_points_auto"`
	// Average points scored during teleop.
	AveragePointsTeleop float32 `json:"average_points_teleop"`
	// Average rotor points scored.
	AverageRotorPoints float32 `json:"average_rotor_points"`
	// Average rotor points scored during auto.
	AverageRotorPointsAuto float32 `json:"average_rotor_points_auto"`
	// Average rotor points scored during teleop.
	AverageRotorPointsTeleop float32 `json:"average_rotor_points_teleop"`
	// Average score.
	AverageScore float32 `json:"average_score"`
	// Average takeoff points scored during teleop.
	AverageTakeoffPointsTeleop float32 `json:"average_takeoff_points_teleop"`
	// Average margin of victory.
	AverageWinMargin float32 `json:"average_win_margin"`
	// Average winning score.
	AverageWinScore float32 `json:"average_win_score"`
	// An array with three values, kPa scored, match key from the match with the high kPa, and the name of the match
	HighKpa []string `json:"high_kpa"`
	// An array with three values, high score, match key from the match with the high score, and the name of the match
	HighScore []string `json:"high_score"`
	// An array with three values, number of times kPa bonus achieved, number of opportunities to bonus, and percentage.
	KpaAchieved []float32 `json:"kpa_achieved"`
	// An array with three values, number of times mobility bonus achieved, number of opportunities to bonus, and percentage.
	MobilityCounts []float32 `json:"mobility_counts"`
	// An array with three values, number of times rotor 1 engaged, number of opportunities to engage, and percentage.
	Rotor1Engaged []float32 `json:"rotor_1_engaged"`
	// An array with three values, number of times rotor 1 engaged in auto, number of opportunities to engage in auto, and percentage.
	Rotor1EngagedAuto []float32 `json:"rotor_1_engaged_auto"`
	// An array with three values, number of times rotor 2 engaged, number of opportunities to engage, and percentage.
	Rotor2Engaged []float32 `json:"rotor_2_engaged"`
	// An array with three values, number of times rotor 2 engaged in auto, number of opportunities to engage in auto, and percentage.
	Rotor2EngagedAuto []float32 `json:"rotor_2_engaged_auto"`
	// An array with three values, number of times rotor 3 engaged, number of opportunities to engage, and percentage.
	Rotor3Engaged []float32 `json:"rotor_3_engaged"`
	// An array with three values, number of times rotor 4 engaged, number of opportunities to engage, and percentage.
	Rotor4Engaged []float32 `json:"rotor_4_engaged"`
	// An array with three values, number of times takeoff was counted, number of opportunities to takeoff, and percentage.
	TakeoffCounts []float32 `json:"takeoff_counts"`
	// An array with three values, number of times a unicorn match (Win + kPa & Rotor Bonuses) occured, number of opportunities to have a unicorn match, and percentage.
	UnicornMatches []float32 `json:"unicorn_matches"`
}
