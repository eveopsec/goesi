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
type GetUniverseConstellationsConstellationIdOk struct {

	// constellation_id integer
	ConstellationId int32 `json:"constellation_id,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	Position GetUniverseConstellationsConstellationIdPosition `json:"position,omitempty"`

	// The region this constellation is in
	RegionId int32 `json:"region_id,omitempty"`

	// systems array
	Systems []int32 `json:"systems,omitempty"`
}
