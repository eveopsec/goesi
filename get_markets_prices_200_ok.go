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
type GetMarketsPrices200Ok struct {

	// adjusted_price number
	AdjustedPrice float32 `json:"adjusted_price,omitempty"`

	// average_price number
	AveragePrice float32 `json:"average_price,omitempty"`

	// type_id integer
	TypeId int32 `json:"type_id,omitempty"`
}
