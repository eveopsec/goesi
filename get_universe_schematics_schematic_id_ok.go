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
type GetUniverseSchematicsSchematicIdOk struct {

	// Time in seconds to process a run
	CycleTime int32 `json:"cycle_time,omitempty"`

	// schematic_name string
	SchematicName string `json:"schematic_name,omitempty"`
}