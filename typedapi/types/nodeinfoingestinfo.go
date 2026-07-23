package types

type NodeInfoIngestInfo struct {
	Downloader NodeInfoIngestDownloader `json:"downloader"`
}

func NewNodeInfoIngestInfo() *NodeInfoIngestInfo { _ = "STUB: not implemented"; return nil }
