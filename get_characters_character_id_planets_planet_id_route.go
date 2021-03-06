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

// route object
type GetCharactersCharacterIdPlanetsPlanetIdRoute struct {

	// content_type_id integer
	ContentTypeId int32 `json:"content_type_id,omitempty"`

	// destination_pin_id integer
	DestinationPinId int64 `json:"destination_pin_id,omitempty"`

	// quantity number
	Quantity float32 `json:"quantity,omitempty"`

	// route_id integer
	RouteId int64 `json:"route_id,omitempty"`

	// source_pin_id integer
	SourcePinId int64 `json:"source_pin_id,omitempty"`

	// waypoints array
	Waypoints []GetCharactersCharacterIdPlanetsPlanetIdWaypoint `json:"waypoints,omitempty"`
}
