package types

type InputConfig struct {
	InputField  string `json:"input_field"`
	OutputField string `json:"output_field"`
}

func (s *InputConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInputConfig() *InputConfig { _ = "STUB: not implemented"; return nil }

type InputConfigVariant interface {
	InputConfigCaster() *InputConfig
}

func (s *InputConfig) InputConfigCaster() *InputConfig { _ = "STUB: not implemented"; return nil }
