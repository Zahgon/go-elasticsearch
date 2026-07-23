package types

type DataStreamLifecycleExplain struct {
	Error                   *string                          `json:"error,omitempty"`
	GenerationTime          Duration                         `json:"generation_time,omitempty"`
	Index                   string                           `json:"index"`
	IndexCreationDateMillis *int64                           `json:"index_creation_date_millis,omitempty"`
	Lifecycle               *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	ManagedByLifecycle      bool                             `json:"managed_by_lifecycle"`
	RolloverDateMillis      *int64                           `json:"rollover_date_millis,omitempty"`
	TimeSinceIndexCreation  Duration                         `json:"time_since_index_creation,omitempty"`
	TimeSinceRollover       Duration                         `json:"time_since_rollover,omitempty"`
}

func (s *DataStreamLifecycleExplain) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamLifecycleExplain() *DataStreamLifecycleExplain {
	_ = "STUB: not implemented"
	return nil
}
