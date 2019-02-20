/*
 * The Blue Alliance API v3
 *
 * # Overview    Information and statistics about FIRST Robotics Competition teams and events.   # Authentication   All endpoints require an Auth Key to be passed in the header `X-TBA-Auth-Key`. If you do not have an auth key yet, you can obtain one from your [Account Page](/account).    A `User-Agent` header may need to be set to prevent a 403 Unauthorized error.
 *
 * API version: 3.04.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MatchScoreBreakdown2019Alliance struct {

	AdjustPoints int32 `json:"adjustPoints,omitempty"`

	AutoPoints int32 `json:"autoPoints,omitempty"`

	Bay1 string `json:"bay1,omitempty"`

	Bay2 string `json:"bay2,omitempty"`

	Bay3 string `json:"bay3,omitempty"`

	Bay4 string `json:"bay4,omitempty"`

	Bay5 string `json:"bay5,omitempty"`

	Bay6 string `json:"bay6,omitempty"`

	Bay7 string `json:"bay7,omitempty"`

	Bay8 string `json:"bay8,omitempty"`

	CargoPoints int32 `json:"cargoPoints,omitempty"`

	CompleteRocketRankingPoint bool `json:"completeRocketRankingPoint,omitempty"`

	CompletedRocketFar bool `json:"completedRocketFar,omitempty"`

	CompletedRocketNear bool `json:"completedRocketNear,omitempty"`

	EndgameRobot1 string `json:"endgameRobot1,omitempty"`

	EndgameRobot2 string `json:"endgameRobot2,omitempty"`

	EndgameRobot3 string `json:"endgameRobot3,omitempty"`

	FoulCount int32 `json:"foulCount,omitempty"`

	FoulPoints int32 `json:"foulPoints,omitempty"`

	HabClimbPoints int32 `json:"habClimbPoints,omitempty"`

	HabDockingRankingPoint bool `json:"habDockingRankingPoint,omitempty"`

	HabLineRobot1 string `json:"habLineRobot1,omitempty"`

	HabLineRobot2 string `json:"habLineRobot2,omitempty"`

	HabLineRobot3 string `json:"habLineRobot3,omitempty"`

	HatchPanelPoints int32 `json:"hatchPanelPoints,omitempty"`

	LowLeftRocketFar string `json:"lowLeftRocketFar,omitempty"`

	LowLeftRocketNear string `json:"lowLeftRocketNear,omitempty"`

	LowRightRocketFar string `json:"lowRightRocketFar,omitempty"`

	LowRightRocketNear string `json:"lowRightRocketNear,omitempty"`

	MidLeftRocketFar string `json:"midLeftRocketFar,omitempty"`

	MidLeftRocketNear string `json:"midLeftRocketNear,omitempty"`

	MidRightRocketFar string `json:"midRightRocketFar,omitempty"`

	MidRightRocketNear string `json:"midRightRocketNear,omitempty"`

	PreMatchBay1 string `json:"preMatchBay1,omitempty"`

	PreMatchBay2 string `json:"preMatchBay2,omitempty"`

	PreMatchBay3 string `json:"preMatchBay3,omitempty"`

	PreMatchBay6 string `json:"preMatchBay6,omitempty"`

	PreMatchBay7 string `json:"preMatchBay7,omitempty"`

	PreMatchBay8 string `json:"preMatchBay8,omitempty"`

	PreMatchLevelRobot1 string `json:"preMatchLevelRobot1,omitempty"`

	PreMatchLevelRobot2 string `json:"preMatchLevelRobot2,omitempty"`

	PreMatchLevelRobot3 string `json:"preMatchLevelRobot3,omitempty"`

	Rp int32 `json:"rp,omitempty"`

	SandStormBonusPoints int32 `json:"sandStormBonusPoints,omitempty"`

	TechFoulCount int32 `json:"techFoulCount,omitempty"`

	TeleopPoints int32 `json:"teleopPoints,omitempty"`

	TopLeftRocketFar string `json:"topLeftRocketFar,omitempty"`

	TopLeftRocketNear string `json:"topLeftRocketNear,omitempty"`

	TopRightRocketFar string `json:"topRightRocketFar,omitempty"`

	TopRightRocketNear string `json:"topRightRocketNear,omitempty"`

	TotalPoints int32 `json:"totalPoints,omitempty"`
}