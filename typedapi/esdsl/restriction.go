package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/restrictionworkflow"
)

type _restriction struct {
	v *types.Restriction
}

func NewRestriction() *_restriction { _ = "STUB: not implemented"; return nil }

func (s *_restriction) Workflows(workflows ...restrictionworkflow.RestrictionWorkflow) *_restriction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_restriction) RestrictionCaster() *types.Restriction {
	_ = "STUB: not implemented"
	return nil
}
