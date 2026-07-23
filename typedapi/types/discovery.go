package types

type Discovery struct {
	ClusterApplierStats *ClusterAppliedStats `json:"cluster_applier_stats,omitempty"`

	ClusterStateQueue *ClusterStateQueue `json:"cluster_state_queue,omitempty"`

	ClusterStateUpdate map[string]ClusterStateUpdate `json:"cluster_state_update,omitempty"`

	PublishedClusterStates  *PublishedClusterStates `json:"published_cluster_states,omitempty"`
	SerializedClusterStates *SerializedClusterState `json:"serialized_cluster_states,omitempty"`
}

func NewDiscovery() *Discovery { _ = "STUB: not implemented"; return nil }
