package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sourcefieldmode"
)

type _sourceField struct {
	v *types.SourceField
}

func NewSourceField() *_sourceField { _ = "STUB: not implemented"; return nil }

func (s *_sourceField) Compress(compress bool) *_sourceField { _ = "STUB: not implemented"; return nil }

func (s *_sourceField) CompressThreshold(compressthreshold string) *_sourceField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceField) Enabled(enabled bool) *_sourceField { _ = "STUB: not implemented"; return nil }

func (s *_sourceField) Excludes(excludes ...string) *_sourceField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceField) Includes(includes ...string) *_sourceField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceField) Mode(mode sourcefieldmode.SourceFieldMode) *_sourceField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sourceField) SourceFieldCaster() *types.SourceField {
	_ = "STUB: not implemented"
	return nil
}
