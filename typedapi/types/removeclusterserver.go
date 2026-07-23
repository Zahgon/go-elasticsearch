package types

type RemoveClusterServer struct {
	BoundAddress   []string `json:"bound_address"`
	PublishAddress string   `json:"publish_address"`
}

func (s *RemoveClusterServer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRemoveClusterServer() *RemoveClusterServer { _ = "STUB: not implemented"; return nil }
