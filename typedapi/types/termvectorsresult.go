package types

type TermVectorsResult struct {
	Error       *ErrorCause           `json:"error,omitempty"`
	Found       *bool                 `json:"found,omitempty"`
	Id_         *string               `json:"_id,omitempty"`
	Index_      string                `json:"_index"`
	TermVectors map[string]TermVector `json:"term_vectors,omitempty"`
	Took        *int64                `json:"took,omitempty"`
	Version_    *int64                `json:"_version,omitempty"`
}

func (s *TermVectorsResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermVectorsResult() *TermVectorsResult { _ = "STUB: not implemented"; return nil }
