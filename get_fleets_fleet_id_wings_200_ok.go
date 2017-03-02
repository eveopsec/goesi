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
type GetFleetsFleetIdWings200Ok struct {

	// id integer
	Id int64 `json:"id,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	// squads array
	Squads []GetFleetsFleetIdWingsSquad `json:"squads,omitempty"`
}