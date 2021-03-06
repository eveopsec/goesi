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

// planet object
type GetUniverseSystemsSystemIdPlanet struct {

	// moons array
	Moons []int32 `json:"moons,omitempty"`

	// planet_id integer
	PlanetId int32 `json:"planet_id,omitempty"`
}
