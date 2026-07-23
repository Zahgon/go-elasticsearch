package types

type PointInTimeReference struct {
	Id        string   `json:"id"`
	KeepAlive Duration `json:"keep_alive,omitempty"`
}

func (s *PointInTimeReference) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPointInTimeReference() *PointInTimeReference { _ = "STUB: not implemented"; return nil }

type PointInTimeReferenceVariant interface {
	PointInTimeReferenceCaster() *PointInTimeReference
}

func (s *PointInTimeReference) PointInTimeReferenceCaster() *PointInTimeReference {
	_ = "STUB: not implemented"
	return nil
}
