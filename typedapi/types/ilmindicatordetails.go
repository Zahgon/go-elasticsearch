package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/lifecycleoperationmode"
)

type IlmIndicatorDetails struct {
	IlmStatus         lifecycleoperationmode.LifecycleOperationMode `json:"ilm_status"`
	Policies          int64                                         `json:"policies"`
	StagnatingIndices int                                           `json:"stagnating_indices"`
}

func (s *IlmIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIlmIndicatorDetails() *IlmIndicatorDetails { _ = "STUB: not implemented"; return nil }
