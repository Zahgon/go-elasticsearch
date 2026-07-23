package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sorttype"
)

type ExtendedSectionSearchUsage struct {
	Sort map[sorttype.SortType]int64 `json:"sort,omitempty"`
}

func NewExtendedSectionSearchUsage() *ExtendedSectionSearchUsage {
	_ = "STUB: not implemented"
	return nil
}
