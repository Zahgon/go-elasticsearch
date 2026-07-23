package types

type DataStreamSettings struct {
	EffectiveSettings IndexSettings `json:"effective_settings"`

	Name string `json:"name"`

	Settings IndexSettings `json:"settings"`
}

func (s *DataStreamSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamSettings() *DataStreamSettings { _ = "STUB: not implemented"; return nil }
