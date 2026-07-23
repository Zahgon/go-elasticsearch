package types

type ShardsAvailabilityIndicatorDetails struct {
	CreatingPrimaries     int64 `json:"creating_primaries"`
	CreatingReplicas      int64 `json:"creating_replicas"`
	InitializingPrimaries int64 `json:"initializing_primaries"`
	InitializingReplicas  int64 `json:"initializing_replicas"`
	RestartingPrimaries   int64 `json:"restarting_primaries"`
	RestartingReplicas    int64 `json:"restarting_replicas"`
	StartedPrimaries      int64 `json:"started_primaries"`
	StartedReplicas       int64 `json:"started_replicas"`
	UnassignedPrimaries   int64 `json:"unassigned_primaries"`
	UnassignedReplicas    int64 `json:"unassigned_replicas"`
}

func (s *ShardsAvailabilityIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardsAvailabilityIndicatorDetails() *ShardsAvailabilityIndicatorDetails {
	_ = "STUB: not implemented"
	return nil
}
