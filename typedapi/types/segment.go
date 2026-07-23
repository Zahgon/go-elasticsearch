package types

type Segment struct {
	Attributes map[string]string `json:"attributes"`

	Committed bool `json:"committed"`

	Compound bool `json:"compound"`

	DeletedDocs int64 `json:"deleted_docs"`

	Generation int `json:"generation"`

	NumDocs int64 `json:"num_docs"`

	Search bool `json:"search"`

	SizeInBytes Float64 `json:"size_in_bytes"`

	Version string `json:"version"`
}

func (s *Segment) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSegment() *Segment { _ = "STUB: not implemented"; return nil }
