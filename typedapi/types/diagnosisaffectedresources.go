package types

type DiagnosisAffectedResources struct {
	FeatureStates        []string        `json:"feature_states,omitempty"`
	Indices              []string        `json:"indices,omitempty"`
	Nodes                []IndicatorNode `json:"nodes,omitempty"`
	SlmPolicies          []string        `json:"slm_policies,omitempty"`
	SnapshotRepositories []string        `json:"snapshot_repositories,omitempty"`
}

func (s *DiagnosisAffectedResources) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDiagnosisAffectedResources() *DiagnosisAffectedResources {
	_ = "STUB: not implemented"
	return nil
}
