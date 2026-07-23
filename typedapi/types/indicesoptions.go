package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

type IndicesOptions struct {
	AllowNoIndices *bool `json:"allow_no_indices,omitempty"`

	ExpandWildcards []expandwildcard.ExpandWildcard `json:"expand_wildcards,omitempty"`

	IgnoreThrottled *bool `json:"ignore_throttled,omitempty"`

	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
}

func (s *IndicesOptions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicesOptions() *IndicesOptions { _ = "STUB: not implemented"; return nil }

type IndicesOptionsVariant interface {
	IndicesOptionsCaster() *IndicesOptions
}

func (s *IndicesOptions) IndicesOptionsCaster() *IndicesOptions {
	_ = "STUB: not implemented"
	return nil
}
