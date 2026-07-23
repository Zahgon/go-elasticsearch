package types

type IoStats struct {
	Devices []IoStatDevice `json:"devices,omitempty"`

	Total *IoStatDevice `json:"total,omitempty"`
}

func NewIoStats() *IoStats { _ = "STUB: not implemented"; return nil }
