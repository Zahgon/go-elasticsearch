package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type TDigest struct {
	Compression *int `json:"compression,omitempty"`

	ExecutionHint *tdigestexecutionhint.TDigestExecutionHint `json:"execution_hint,omitempty"`
}

func (s *TDigest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTDigest() *TDigest { _ = "STUB: not implemented"; return nil }

type TDigestVariant interface {
	TDigestCaster() *TDigest
}

func (s *TDigest) TDigestCaster() *TDigest { _ = "STUB: not implemented"; return nil }
