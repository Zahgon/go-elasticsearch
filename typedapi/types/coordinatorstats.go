package types

type CoordinatorStats struct {
	ExecutedSearchesTotal int64  `json:"executed_searches_total"`
	NodeId                string `json:"node_id"`
	QueueSize             int    `json:"queue_size"`
	RemoteRequestsCurrent int    `json:"remote_requests_current"`
	RemoteRequestsTotal   int64  `json:"remote_requests_total"`
}

func (s *CoordinatorStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCoordinatorStats() *CoordinatorStats { _ = "STUB: not implemented"; return nil }
