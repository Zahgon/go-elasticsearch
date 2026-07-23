package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _replicationAccess struct {
	v *types.ReplicationAccess
}

func NewReplicationAccess() *_replicationAccess { _ = "STUB: not implemented"; return nil }

func (s *_replicationAccess) AllowRestrictedIndices(allowrestrictedindices bool) *_replicationAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_replicationAccess) Names(names ...string) *_replicationAccess {
	_ = "STUB: not implemented"
	return nil
}

func (s *_replicationAccess) ReplicationAccessCaster() *types.ReplicationAccess {
	_ = "STUB: not implemented"
	return nil
}
