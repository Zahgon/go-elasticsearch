package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/lifecycleoperationmode"
)

type SlmIndicatorDetails struct {
	Policies          int64                                         `json:"policies"`
	SlmStatus         lifecycleoperationmode.LifecycleOperationMode `json:"slm_status"`
	UnhealthyPolicies *SlmIndicatorUnhealthyPolicies                `json:"unhealthy_policies,omitempty"`
}

func (s *SlmIndicatorDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlmIndicatorDetails() *SlmIndicatorDetails { _ = "STUB: not implemented"; return nil }
