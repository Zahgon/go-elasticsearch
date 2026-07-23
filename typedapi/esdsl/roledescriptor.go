package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type _roleDescriptor struct {
	v *types.RoleDescriptor
}

func NewRoleDescriptor() *_roleDescriptor { _ = "STUB: not implemented"; return nil }

func (s *_roleDescriptor) Applications(applications ...types.ApplicationPrivilegesVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) ApplicationsValues(applicationsvalues []types.ApplicationPrivileges) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Description(description string) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Global(globals ...types.GlobalPrivilegeVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Indices(indices ...types.IndicesPrivilegesVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) IndicesValues(indicesvalues []types.IndicesPrivileges) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Metadata(metadata types.MetadataVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RemoteCluster(remoteclusters ...types.RemoteClusterPrivilegesVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RemoteClusterValues(remoteclustervalues []types.RemoteClusterPrivileges) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RemoteIndices(remoteindices ...types.RemoteIndicesPrivilegesVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RemoteIndicesValues(remoteindicesvalues []types.RemoteIndicesPrivileges) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) Restriction(restriction types.RestrictionVariant) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RunAs(runas ...string) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) TransientMetadata(transientmetadata map[string]json.RawMessage) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) AddTransientMetadatum(key string, value json.RawMessage) *_roleDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleDescriptor) RoleDescriptorCaster() *types.RoleDescriptor {
	_ = "STUB: not implemented"
	return nil
}
