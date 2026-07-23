package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsUnassigned struct {
	v *types.IndexSettingsUnassigned
}

func NewIndexSettingsUnassigned() *_indexSettingsUnassigned { _ = "STUB: not implemented"; return nil }

func (s *_indexSettingsUnassigned) NodeLeft(nodeleft types.IndexSettingsUnassignedNodeLeftVariant) *_indexSettingsUnassigned {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsUnassigned) IndexSettingsUnassignedCaster() *types.IndexSettingsUnassigned {
	_ = "STUB: not implemented"
	return nil
}
