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

// head object
type GetCharactersCharacterIdPlanetsPlanetIdHead struct {

	// head_id integer
	HeadId int32 `json:"head_id,omitempty"`

	// latitude number
	Latitude float32 `json:"latitude,omitempty"`

	// longitude number
	Longitude float32 `json:"longitude,omitempty"`
}
