package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

type IndexResultSummary struct {
	Created bool          `json:"created"`
	Id      string        `json:"id"`
	Index   string        `json:"index"`
	Result  result.Result `json:"result"`
	Version int64         `json:"version"`
}

func (s *IndexResultSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexResultSummary() *IndexResultSummary { _ = "STUB: not implemented"; return nil }
