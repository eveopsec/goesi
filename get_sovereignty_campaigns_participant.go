/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.1.dev1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// participant object
type GetSovereigntyCampaignsParticipant struct {

	// alliance_id integer
	AllianceId int32 `json:"alliance_id,omitempty"`

	// score number
	Score float32 `json:"score,omitempty"`
}
