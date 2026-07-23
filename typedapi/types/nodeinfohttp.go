package types

type NodeInfoHttp struct {
	BoundAddress            []string `json:"bound_address"`
	MaxContentLength        ByteSize `json:"max_content_length,omitempty"`
	MaxContentLengthInBytes int64    `json:"max_content_length_in_bytes"`
	PublishAddress          string   `json:"publish_address"`
}

func (s *NodeInfoHttp) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoHttp() *NodeInfoHttp { _ = "STUB: not implemented"; return nil }
