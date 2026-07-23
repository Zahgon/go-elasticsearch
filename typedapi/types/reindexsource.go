package types

type ReindexSource struct {
	Index []string `json:"index"`

	ProjectRouting *string `json:"project_routing,omitempty"`

	Query *Query `json:"query,omitempty"`

	Remote          *RemoteSource `json:"remote,omitempty"`
	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`

	Size *int `json:"size,omitempty"`

	Slice *SlicedScroll `json:"slice,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`

	SourceFields_ SourceConfig `json:"_source,omitempty"`
}

func (s *ReindexSource) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReindexSource() *ReindexSource { _ = "STUB: not implemented"; return nil }

type ReindexSourceVariant interface {
	ReindexSourceCaster() *ReindexSource
}

func (s *ReindexSource) ReindexSourceCaster() *ReindexSource { _ = "STUB: not implemented"; return nil }
