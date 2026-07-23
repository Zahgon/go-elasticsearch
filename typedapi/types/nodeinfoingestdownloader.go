package types

type NodeInfoIngestDownloader struct {
	Enabled string `json:"enabled"`
}

func (s *NodeInfoIngestDownloader) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoIngestDownloader() *NodeInfoIngestDownloader { _ = "STUB: not implemented"; return nil }
