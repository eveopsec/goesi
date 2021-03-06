/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.2.2
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesi

import (
	"time"
)

// 200 ok object object
type GetCharactersCharacterIdBookmarks200OkObject struct {

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
