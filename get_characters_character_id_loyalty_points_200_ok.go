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
type GetCharactersCharacterIdLoyaltyPoints200Ok struct {

	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`

	// loyalty_points integer
	LoyaltyPoints int32 `json:"loyalty_points,omitempty"`
}