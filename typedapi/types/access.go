package types

type Access struct {
	Replication []ReplicationAccess `json:"replication,omitempty"`

	Search []SearchAccess `json:"search,omitempty"`
}

func NewAccess() *Access { _ = "STUB: not implemented"; return nil }

type AccessVariant interface {
	AccessCaster() *Access
}

func (s *Access) AccessCaster() *Access { _ = "STUB: not implemented"; return nil }
