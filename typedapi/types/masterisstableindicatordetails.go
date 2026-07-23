package types

type MasterIsStableIndicatorDetails struct {
	ClusterFormation         []MasterIsStableIndicatorClusterFormationNode    `json:"cluster_formation,omitempty"`
	CurrentMaster            IndicatorNode                                    `json:"current_master"`
	ExceptionFetchingHistory *MasterIsStableIndicatorExceptionFetchingHistory `json:"exception_fetching_history,omitempty"`
	RecentMasters            []IndicatorNode                                  `json:"recent_masters"`
}

func NewMasterIsStableIndicatorDetails() *MasterIsStableIndicatorDetails {
	_ = "STUB: not implemented"
	return nil
}
