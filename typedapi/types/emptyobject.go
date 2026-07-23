package types

type EmptyObject struct {
}

func NewEmptyObject() *EmptyObject { _ = "STUB: not implemented"; return nil }

type EmptyObjectVariant interface {
	EmptyObjectCaster() *EmptyObject
}

func (s *EmptyObject) EmptyObjectCaster() *EmptyObject { _ = "STUB: not implemented"; return nil }
