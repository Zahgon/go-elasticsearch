package types

type ReadBlobDetails struct {
	BeforeWriteComplete *bool `json:"before_write_complete,omitempty"`

	Elapsed Duration `json:"elapsed,omitempty"`

	ElapsedNanos *int64 `json:"elapsed_nanos,omitempty"`

	FirstByteTime Duration `json:"first_byte_time,omitempty"`

	FirstByteTimeNanos int64 `json:"first_byte_time_nanos"`

	Found bool `json:"found"`

	Node SnapshotNodeInfo `json:"node"`

	Throttled Duration `json:"throttled,omitempty"`

	ThrottledNanos *int64 `json:"throttled_nanos,omitempty"`
}

func (s *ReadBlobDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReadBlobDetails() *ReadBlobDetails { _ = "STUB: not implemented"; return nil }
