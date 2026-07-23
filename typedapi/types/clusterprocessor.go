package types

type ClusterProcessor struct {
	Count        int64    `json:"count"`
	Current      int64    `json:"current"`
	Failed       int64    `json:"failed"`
	Time         Duration `json:"time,omitempty"`
	TimeInMillis int64    `json:"time_in_millis"`
}

func (s *ClusterProcessor) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterProcessor() *ClusterProcessor { _ = "STUB: not implemented"; return nil }
