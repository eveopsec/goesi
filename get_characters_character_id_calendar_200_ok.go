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

// event
type GetCharactersCharacterIdCalendar200Ok struct {

	// event_date string
	EventDate time.Time `json:"event_date,omitempty"`

	// event_id integer
	EventId int32 `json:"event_id,omitempty"`

	// event_response string
	EventResponse string `json:"event_response,omitempty"`

	// importance integer
	Importance int32 `json:"importance,omitempty"`

	// title string
	Title string `json:"title,omitempty"`
}
