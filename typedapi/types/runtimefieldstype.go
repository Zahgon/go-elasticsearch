package types

type RuntimeFieldsType struct {
	CharsMax        int64    `json:"chars_max"`
	CharsTotal      int64    `json:"chars_total"`
	Count           int64    `json:"count"`
	DocMax          int64    `json:"doc_max"`
	DocTotal        int64    `json:"doc_total"`
	IndexCount      int64    `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int64    `json:"lines_max"`
	LinesTotal      int64    `json:"lines_total"`
	Name            string   `json:"name"`
	ScriptlessCount int64    `json:"scriptless_count"`
	ShadowedCount   int64    `json:"shadowed_count"`
	SourceMax       int64    `json:"source_max"`
	SourceTotal     int64    `json:"source_total"`
}

func (s *RuntimeFieldsType) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRuntimeFieldsType() *RuntimeFieldsType { _ = "STUB: not implemented"; return nil }
