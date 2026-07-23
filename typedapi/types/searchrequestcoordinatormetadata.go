package types

type SearchRequestCoordinatorMetadata struct {
	Indices []string `json:"indices,omitempty"`

	Source *SearchRequestBody `json:"source,omitempty"`
}

func NewSearchRequestCoordinatorMetadata() *SearchRequestCoordinatorMetadata {
	_ = "STUB: not implemented"
	return nil
}
