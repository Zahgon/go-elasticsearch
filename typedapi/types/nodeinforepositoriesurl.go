package types

type NodeInfoRepositoriesUrl struct {
	AllowedUrls string `json:"allowed_urls"`
}

func (s *NodeInfoRepositoriesUrl) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoRepositoriesUrl() *NodeInfoRepositoriesUrl { _ = "STUB: not implemented"; return nil }
