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

// level object
type GetInsurancePricesLevel struct {

	// cost number
	Cost float32 `json:"cost,omitempty"`

	// Localized insurance level
	Name string `json:"name,omitempty"`

	// payout number
	Payout float32 `json:"payout,omitempty"`
}