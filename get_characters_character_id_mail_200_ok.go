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

import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdMail200Ok struct {

	// From whom the mail was sent
	From int32 `json:"from,omitempty"`

	// is_read boolean
	IsRead bool `json:"is_read,omitempty"`

	// labels array
	Labels []int64 `json:"labels,omitempty"`

	// mail_id integer
	MailId int64 `json:"mail_id,omitempty"`

	// Recipients of the mail
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`

	// Mail subject
	Subject string `json:"subject,omitempty"`

	// When the mail was sent
	Timestamp time.Time `json:"timestamp,omitempty"`
}