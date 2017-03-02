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
type GetUniverseSystemsSystemIdOk struct {

	// The constellation this solar system is in
	ConstellationId int32 `json:"constellation_id,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	// planets array
	Planets []GetUniverseSystemsSystemIdPlanet `json:"planets,omitempty"`

	Position GetUniverseSystemsSystemIdPosition `json:"position,omitempty"`

	// security_class string
	SecurityClass string `json:"security_class,omitempty"`

	// security_status number
	SecurityStatus float32 `json:"security_status,omitempty"`

	// stargates array
	Stargates []int32 `json:"stargates,omitempty"`

	// system_id integer
	SystemId int32 `json:"system_id,omitempty"`
}
