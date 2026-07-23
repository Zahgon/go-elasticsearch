package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/modeenum"
)

type MigrateReindex struct {
	Mode modeenum.ModeEnum `json:"mode"`

	Source SourceIndex `json:"source"`
}

func NewMigrateReindex() *MigrateReindex { _ = "STUB: not implemented"; return nil }

type MigrateReindexVariant interface {
	MigrateReindexCaster() *MigrateReindex
}

func (s *MigrateReindex) MigrateReindexCaster() *MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}
