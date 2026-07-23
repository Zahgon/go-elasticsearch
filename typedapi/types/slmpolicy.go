package types

type SLMPolicy struct {
	Config     *Configuration `json:"config,omitempty"`
	Name       string         `json:"name"`
	Repository string         `json:"repository"`
	Retention  *Retention     `json:"retention,omitempty"`
	Schedule   string         `json:"schedule"`
}

func (s *SLMPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSLMPolicy() *SLMPolicy { _ = "STUB: not implemented"; return nil }
