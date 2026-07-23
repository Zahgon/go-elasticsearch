package types

type NativeCode struct {
	BuildHash string `json:"build_hash"`
	Version   string `json:"version"`
}

func (s *NativeCode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNativeCode() *NativeCode { _ = "STUB: not implemented"; return nil }
