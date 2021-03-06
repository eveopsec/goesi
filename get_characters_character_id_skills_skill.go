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

// skill object
type GetCharactersCharacterIdSkillsSkill struct {

	// current_skill_level integer
	CurrentSkillLevel int32 `json:"current_skill_level,omitempty"`

	// skill_id integer
	SkillId int32 `json:"skill_id,omitempty"`

	// skillpoints_in_skill integer
	SkillpointsInSkill int64 `json:"skillpoints_in_skill,omitempty"`
}
