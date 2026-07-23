package types

type Checkpointing struct {
	ChangesLastDetectedAt       *int64           `json:"changes_last_detected_at,omitempty"`
	ChangesLastDetectedAtString DateTime         `json:"changes_last_detected_at_string,omitempty"`
	Last                        CheckpointStats  `json:"last"`
	LastSearchTime              *int64           `json:"last_search_time,omitempty"`
	LastSearchTimeString        DateTime         `json:"last_search_time_string,omitempty"`
	Next                        *CheckpointStats `json:"next,omitempty"`
	OperationsBehind            *int64           `json:"operations_behind,omitempty"`
}

func (s *Checkpointing) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCheckpointing() *Checkpointing { _ = "STUB: not implemented"; return nil }
