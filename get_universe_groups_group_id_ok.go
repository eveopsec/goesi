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
type GetUniverseGroupsGroupIdOk struct {

	// category_id number
	CategoryId float32 `json:"category_id,omitempty"`

	// group_id integer
	GroupId int32 `json:"group_id,omitempty"`

	// name string
	Name string `json:"name,omitempty"`

	// published boolean
	Published bool `json:"published,omitempty"`

	// types array
	Types []int32 `json:"types,omitempty"`
}
