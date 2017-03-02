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

// 200 ok object
type GetUniverseRaces200Ok struct {

	// The alliance generally associated with this race
	AllianceId int32 `json:"alliance_id,omitempty"`

	// description string
	Description string `json:"description,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	// race_id integer
	RaceId int32 `json:"race_id,omitempty"`
}