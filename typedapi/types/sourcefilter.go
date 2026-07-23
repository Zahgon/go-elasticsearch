package types

type SourceFilter struct {
	ExcludeVectors *bool `json:"exclude_vectors,omitempty"`

	Excludes []string `json:"excludes,omitempty"`

	Includes []string `json:"includes,omitempty"`
}

func (s *SourceFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSourceFilter() *SourceFilter { _ = "STUB: not implemented"; return nil }

type SourceFilterVariant interface {
	SourceFilterCaster() *SourceFilter
}

func (s *SourceFilter) SourceFilterCaster() *SourceFilter { _ = "STUB: not implemented"; return nil }

func (s *SourceFilter) SourceConfigCaster() *SourceConfig { _ = "STUB: not implemented"; return nil }
