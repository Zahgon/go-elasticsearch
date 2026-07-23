package types

type SerializedClusterState struct {
	Diffs *SerializedClusterStateDetail `json:"diffs,omitempty"`

	FullStates *SerializedClusterStateDetail `json:"full_states,omitempty"`
}

func NewSerializedClusterState() *SerializedClusterState { _ = "STUB: not implemented"; return nil }
