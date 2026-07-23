package types

type DataFrameAnalyticsRecord struct {
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	Description *string `json:"description,omitempty"`

	DestIndex *string `json:"dest_index,omitempty"`

	FailureReason *string `json:"failure_reason,omitempty"`

	Id *string `json:"id,omitempty"`

	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`

	NodeAddress *string `json:"node.address,omitempty"`

	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`

	NodeId *string `json:"node.id,omitempty"`

	NodeName *string `json:"node.name,omitempty"`

	Progress *string `json:"progress,omitempty"`

	SourceIndex *string `json:"source_index,omitempty"`

	State *string `json:"state,omitempty"`

	Type *string `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *DataFrameAnalyticsRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataFrameAnalyticsRecord() *DataFrameAnalyticsRecord { _ = "STUB: not implemented"; return nil }
