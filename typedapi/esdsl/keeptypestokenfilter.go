package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/keeptypesmode"
)

type _keepTypesTokenFilter struct {
	v *types.KeepTypesTokenFilter
}

func NewKeepTypesTokenFilter() *_keepTypesTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_keepTypesTokenFilter) Mode(mode keeptypesmode.KeepTypesMode) *_keepTypesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepTypesTokenFilter) Types(types ...string) *_keepTypesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepTypesTokenFilter) Version(versionstring string) *_keepTypesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keepTypesTokenFilter) KeepTypesTokenFilterCaster() *types.KeepTypesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
