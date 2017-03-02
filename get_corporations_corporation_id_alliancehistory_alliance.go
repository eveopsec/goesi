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

// alliance object
type GetCorporationsCorporationIdAlliancehistoryAlliance struct {

	// alliance_id integer
	AllianceId int32 `json:"alliance_id,omitempty"`

	// True if the alliance has been deleted
	IsDeleted bool `json:"is_deleted,omitempty"`
}