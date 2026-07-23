package types

type ElasticsearchVersionMinInfo struct {
	BuildFlavor                      string `json:"build_flavor"`
	Int                              string `json:"number"`
	MinimumIndexCompatibilityVersion string `json:"minimum_index_compatibility_version"`
	MinimumWireCompatibilityVersion  string `json:"minimum_wire_compatibility_version"`
}

func (s *ElasticsearchVersionMinInfo) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewElasticsearchVersionMinInfo() *ElasticsearchVersionMinInfo {
	_ = "STUB: not implemented"
	return nil
}
