package types

type TransformProgress struct {
	DocsIndexed     int64    `json:"docs_indexed"`
	DocsProcessed   int64    `json:"docs_processed"`
	DocsRemaining   *int64   `json:"docs_remaining,omitempty"`
	PercentComplete *Float64 `json:"percent_complete,omitempty"`
	TotalDocs       *int64   `json:"total_docs,omitempty"`
}

func (s *TransformProgress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransformProgress() *TransformProgress { _ = "STUB: not implemented"; return nil }
