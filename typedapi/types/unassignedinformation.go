package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/unassignedinformationreason"
)

type UnassignedInformation struct {
	AllocationStatus         *string                                                 `json:"allocation_status,omitempty"`
	At                       DateTime                                                `json:"at"`
	Delayed                  *bool                                                   `json:"delayed,omitempty"`
	Details                  *string                                                 `json:"details,omitempty"`
	FailedAllocationAttempts *int                                                    `json:"failed_allocation_attempts,omitempty"`
	LastAllocationStatus     *string                                                 `json:"last_allocation_status,omitempty"`
	Reason                   unassignedinformationreason.UnassignedInformationReason `json:"reason"`
}

func (s *UnassignedInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUnassignedInformation() *UnassignedInformation { _ = "STUB: not implemented"; return nil }
