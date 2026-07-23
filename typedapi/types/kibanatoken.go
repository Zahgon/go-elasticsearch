package types

type KibanaToken struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

func (s *KibanaToken) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewKibanaToken() *KibanaToken { _ = "STUB: not implemented"; return nil }
