package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _access struct {
	v *types.Access
}

func NewAccess() *_access { _ = "STUB: not implemented"; return nil }

func (s *_access) Replication(replications ...types.ReplicationAccessVariant) *_access {
	_ = "STUB: not implemented"
	return nil
}

func (s *_access) ReplicationValues(replicationvalues []types.ReplicationAccess) *_access {
	_ = "STUB: not implemented"
	return nil
}

func (s *_access) Search(searches ...types.SearchAccessVariant) *_access {
	_ = "STUB: not implemented"
	return nil
}

func (s *_access) SearchValues(searchvalues []types.SearchAccess) *_access {
	_ = "STUB: not implemented"
	return nil
}

func (s *_access) AccessCaster() *types.Access { _ = "STUB: not implemented"; return nil }
