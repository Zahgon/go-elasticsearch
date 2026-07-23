package types

type DataStreamSettingsError struct {
	Error string `json:"error"`
	Index string `json:"index"`
}

func (s *DataStreamSettingsError) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamSettingsError() *DataStreamSettingsError { _ = "STUB: not implemented"; return nil }
