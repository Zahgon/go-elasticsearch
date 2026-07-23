package types

type Memory struct {
	Attributes  map[string]string `json:"attributes"`
	EphemeralId string            `json:"ephemeral_id"`

	Jvm JvmStats `json:"jvm"`

	Mem MemStats `json:"mem"`

	Name string `json:"name"`

	Roles []string `json:"roles"`

	TransportAddress string `json:"transport_address"`
}

func (s *Memory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMemory() *Memory { _ = "STUB: not implemented"; return nil }
