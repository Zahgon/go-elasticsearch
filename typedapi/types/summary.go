package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/policytype"
)

type Summary struct {
	Config map[policytype.PolicyType]EnrichPolicy `json:"config"`
}

func NewSummary() *Summary { _ = "STUB: not implemented"; return nil }
