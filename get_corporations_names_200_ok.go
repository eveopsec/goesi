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
type GetCorporationsNames200Ok struct {

	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`

	// corporation_name string
	CorporationName string `json:"corporation_name,omitempty"`
}
