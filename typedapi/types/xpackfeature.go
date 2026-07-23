package types

type XpackFeature struct {
	Available      bool                   `json:"available"`
	Description    *string                `json:"description,omitempty"`
	Enabled        bool                   `json:"enabled"`
	NativeCodeInfo *NativeCodeInformation `json:"native_code_info,omitempty"`
}

func (s *XpackFeature) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewXpackFeature() *XpackFeature { _ = "STUB: not implemented"; return nil }
