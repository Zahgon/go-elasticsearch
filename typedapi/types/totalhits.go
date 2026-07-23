package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/totalhitsrelation"
)

type TotalHits struct {
	Relation totalhitsrelation.TotalHitsRelation `json:"relation"`
	Value    int64                               `json:"value"`
}

func (t *TotalHits) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTotalHits() *TotalHits { _ = "STUB: not implemented"; return nil }
