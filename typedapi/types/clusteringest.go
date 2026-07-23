package types

type ClusterIngest struct {
	NumberOfPipelines int                         `json:"number_of_pipelines"`
	ProcessorStats    map[string]ClusterProcessor `json:"processor_stats"`
}

func (s *ClusterIngest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterIngest() *ClusterIngest { _ = "STUB: not implemented"; return nil }
