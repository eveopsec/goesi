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

// mail schema
type PostCharactersCharacterIdMailMail struct {

	// approved_cost integer
	ApprovedCost int64 `json:"approved_cost,omitempty"`

	// body string
	Body string `json:"body,omitempty"`

	// recipients array
	Recipients []PostCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`

	// subject string
	Subject string `json:"subject,omitempty"`
}
