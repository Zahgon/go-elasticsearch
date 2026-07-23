package types

import (
	"encoding/json"
)

type LifecycleExplainManaged struct {
	Action                  *string                         `json:"action,omitempty"`
	ActionTime              DateTime                        `json:"action_time,omitempty"`
	ActionTimeMillis        *int64                          `json:"action_time_millis,omitempty"`
	Age                     Duration                        `json:"age,omitempty"`
	AgeInMillis             *int64                          `json:"age_in_millis,omitempty"`
	FailedStep              *string                         `json:"failed_step,omitempty"`
	FailedStepRetryCount    *int                            `json:"failed_step_retry_count,omitempty"`
	Index                   string                          `json:"index"`
	IndexCreationDate       DateTime                        `json:"index_creation_date,omitempty"`
	IndexCreationDateMillis *int64                          `json:"index_creation_date_millis,omitempty"`
	IsAutoRetryableError    *bool                           `json:"is_auto_retryable_error,omitempty"`
	LifecycleDate           DateTime                        `json:"lifecycle_date,omitempty"`
	LifecycleDateMillis     *int64                          `json:"lifecycle_date_millis,omitempty"`
	Managed                 bool                            `json:"managed,omitempty"`
	Phase                   *string                         `json:"phase,omitempty"`
	PhaseExecution          *LifecycleExplainPhaseExecution `json:"phase_execution,omitempty"`
	PhaseTime               DateTime                        `json:"phase_time,omitempty"`
	PhaseTimeMillis         *int64                          `json:"phase_time_millis,omitempty"`
	Policy                  *string                         `json:"policy,omitempty"`
	PreviousStepInfo        map[string]json.RawMessage      `json:"previous_step_info,omitempty"`
	RepositoryName          *string                         `json:"repository_name,omitempty"`
	ShrinkIndexName         *string                         `json:"shrink_index_name,omitempty"`
	Skip                    bool                            `json:"skip"`
	SnapshotName            *string                         `json:"snapshot_name,omitempty"`
	Step                    *string                         `json:"step,omitempty"`
	StepInfo                map[string]json.RawMessage      `json:"step_info,omitempty"`
	StepTime                DateTime                        `json:"step_time,omitempty"`
	StepTimeMillis          *int64                          `json:"step_time_millis,omitempty"`
	TimeSinceIndexCreation  Duration                        `json:"time_since_index_creation,omitempty"`
}

func (s *LifecycleExplainManaged) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LifecycleExplainManaged) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLifecycleExplainManaged() *LifecycleExplainManaged { _ = "STUB: not implemented"; return nil }
