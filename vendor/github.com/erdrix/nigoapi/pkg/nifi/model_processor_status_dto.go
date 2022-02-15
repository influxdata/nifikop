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

type ProcessorStatusDto struct {
	// The unique ID of the process group that the Processor belongs to
	GroupId string `json:"groupId,omitempty"`
	// The unique ID of the Processor
	Id string `json:"id,omitempty"`
	// The name of the Processor
	Name string `json:"name,omitempty"`
	// The type of the Processor
	Type_ string `json:"type,omitempty"`
	// The run status of the Processor
	RunStatus string `json:"runStatus,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	// A status snapshot that represents the aggregate stats of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance.
	AggregateSnapshot *ProcessorStatusSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null.
	NodeSnapshots []NodeProcessorStatusSnapshotDto `json:"nodeSnapshots,omitempty"`
}
