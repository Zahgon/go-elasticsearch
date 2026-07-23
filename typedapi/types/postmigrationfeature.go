package types

type PostMigrationFeature struct {
	FeatureName string `json:"feature_name"`
}

func (s *PostMigrationFeature) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPostMigrationFeature() *PostMigrationFeature { _ = "STUB: not implemented"; return nil }
