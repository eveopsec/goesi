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
type GetCharactersCharacterIdBookmarks200Ok struct {

	// bookmark_id integer
	BookmarkId int64 `json:"bookmark_id,omitempty"`

	// create_date string
	CreateDate time.Time `json:"create_date,omitempty"`

	// creator_id integer
	CreatorId int32 `json:"creator_id,omitempty"`

	// folder_id integer
	FolderId int32 `json:"folder_id,omitempty"`

	// memo string
	Memo string `json:"memo,omitempty"`

	// note string
	Note string `json:"note,omitempty"`

	// owner_id integer
	OwnerId int32 `json:"owner_id,omitempty"`

	Target GetCharactersCharacterIdBookmarksTarget `json:"target,omitempty"`
}
