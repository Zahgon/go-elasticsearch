package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/modeenum"
)

type _migrateReindex struct {
	v *types.MigrateReindex
}

func NewMigrateReindex(mode modeenum.ModeEnum, source types.SourceIndexVariant) *_migrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (s *_migrateReindex) Mode(mode modeenum.ModeEnum) *_migrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (s *_migrateReindex) Source(source types.SourceIndexVariant) *_migrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (s *_migrateReindex) MigrateReindexCaster() *types.MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}
