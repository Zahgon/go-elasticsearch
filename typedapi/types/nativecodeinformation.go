package types

type NativeCodeInformation struct {
	BuildHash string `json:"build_hash"`
	Version   string `json:"version"`
}

func (s *NativeCodeInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNativeCodeInformation() *NativeCodeInformation { _ = "STUB: not implemented"; return nil }
