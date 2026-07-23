package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/migrationstatus"
)

type GetMigrationFeature struct {
	FeatureName         string                          `json:"feature_name"`
	Indices             []MigrationFeatureIndexInfo     `json:"indices"`
	MigrationStatus     migrationstatus.MigrationStatus `json:"migration_status"`
	MinimumIndexVersion string                          `json:"minimum_index_version"`
}

func (s *GetMigrationFeature) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGetMigrationFeature() *GetMigrationFeature { _ = "STUB: not implemented"; return nil }
