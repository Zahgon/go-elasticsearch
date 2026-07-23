package types

type IoStatDevice struct {
	DeviceName *string `json:"device_name,omitempty"`

	Operations *int64 `json:"operations,omitempty"`

	ReadKilobytes *int64 `json:"read_kilobytes,omitempty"`

	ReadOperations *int64 `json:"read_operations,omitempty"`

	WriteKilobytes *int64 `json:"write_kilobytes,omitempty"`

	WriteOperations *int64 `json:"write_operations,omitempty"`
}

func (s *IoStatDevice) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIoStatDevice() *IoStatDevice { _ = "STUB: not implemented"; return nil }
