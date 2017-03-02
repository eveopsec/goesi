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

// destination object
type GetUniverseStargatesStargateIdDestination struct {

	// The stargate this stargate connects to
	StargateId int32 `json:"stargate_id,omitempty"`

	// The solar system this stargate connects to
	SystemId int32 `json:"system_id,omitempty"`
}