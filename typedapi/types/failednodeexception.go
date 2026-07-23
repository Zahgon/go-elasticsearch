package types

type FailedNodeException struct {
	NodeId string `json:"node_id"`
}

func (s *FailedNodeException) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFailedNodeException() *FailedNodeException { _ = "STUB: not implemented"; return nil }
