package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _privilegesActions struct {
	v *types.PrivilegesActions
}

func NewPrivilegesActions() *_privilegesActions { _ = "STUB: not implemented"; return nil }

func (s *_privilegesActions) Actions(actions ...string) *_privilegesActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesActions) Application(application string) *_privilegesActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesActions) Metadata(metadata types.MetadataVariant) *_privilegesActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesActions) Name(name string) *_privilegesActions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_privilegesActions) PrivilegesActionsCaster() *types.PrivilegesActions {
	_ = "STUB: not implemented"
	return nil
}
