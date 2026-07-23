package types

type ClusterRuntimeFieldTypes struct {
	CharsMax int `json:"chars_max"`

	CharsTotal int `json:"chars_total"`

	Count int `json:"count"`

	DocMax int `json:"doc_max"`

	DocTotal int `json:"doc_total"`

	IndexCount int `json:"index_count"`

	Lang []string `json:"lang"`

	LinesMax int `json:"lines_max"`

	LinesTotal int `json:"lines_total"`

	Name string `json:"name"`

	ScriptlessCount int `json:"scriptless_count"`

	ShadowedCount int `json:"shadowed_count"`

	SourceMax int `json:"source_max"`

	SourceTotal int `json:"source_total"`
}

func (s *ClusterRuntimeFieldTypes) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterRuntimeFieldTypes() *ClusterRuntimeFieldTypes { _ = "STUB: not implemented"; return nil }
