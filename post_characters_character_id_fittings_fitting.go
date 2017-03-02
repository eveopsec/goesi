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

// fitting object
type PostCharactersCharacterIdFittingsFitting struct {

	// description string
	Description string `json:"description,omitempty"`

	// items array
	Items []PostCharactersCharacterIdFittingsItem `json:"items,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
