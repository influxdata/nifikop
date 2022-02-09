/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ParameterContextDto struct {
	// The Name of the Parameter Context.
	Name string `json:"name,omitempty"`
	// The Description of the Parameter Context.
	Description string `json:"description,omitempty"`
	// The Parameters for the Parameter Context
	Parameters []ParameterEntity `json:"parameters"`
	// The Process Groups that are bound to this Parameter Context
	BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups,omitempty"`
	// The ID the Parameter Context.
	Id string `json:"id,omitempty"`
}
