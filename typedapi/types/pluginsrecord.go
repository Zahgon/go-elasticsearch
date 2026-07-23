package types

type PluginsRecord struct {
	Component *string `json:"component,omitempty"`

	Description *string `json:"description,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *PluginsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPluginsRecord() *PluginsRecord { _ = "STUB: not implemented"; return nil }
