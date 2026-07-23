package types

type ElasticsearchVersionInfo struct {
	BuildDate DateTime `json:"build_date"`

	BuildFlavor string `json:"build_flavor"`

	BuildHash string `json:"build_hash"`

	BuildSnapshot bool `json:"build_snapshot"`

	BuildType string `json:"build_type"`

	Int string `json:"number"`

	LuceneVersion string `json:"lucene_version"`

	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`

	MinimumWireCompatibilityVersion string `json:"minimum_wire_compatibility_version"`
}

func (s *ElasticsearchVersionInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewElasticsearchVersionInfo() *ElasticsearchVersionInfo { _ = "STUB: not implemented"; return nil }
