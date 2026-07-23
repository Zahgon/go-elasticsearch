package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/matchtype"
)

type _dynamicTemplate struct {
	v *types.DynamicTemplate
}

func NewDynamicTemplate() *_dynamicTemplate { _ = "STUB: not implemented"; return nil }

func (s *_dynamicTemplate) Mapping(property types.PropertyVariant) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) Match(matches ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) MatchMappingType(matchmappingtypes ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) MatchPattern(matchpattern matchtype.MatchType) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) PathMatch(pathmatches ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) PathUnmatch(pathunmatches ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) Runtime(runtime types.RuntimeFieldVariant) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) Unmatch(unmatches ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) UnmatchMappingType(unmatchmappingtypes ...string) *_dynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dynamicTemplate) DynamicTemplateCaster() *types.DynamicTemplate {
	_ = "STUB: not implemented"
	return nil
}
