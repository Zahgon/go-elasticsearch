package types

type GeoipDatabaseConfigurationMetadata struct {
	Database           DatabaseConfiguration `json:"database"`
	Id                 string                `json:"id"`
	ModifiedDateMillis int64                 `json:"modified_date_millis"`
	Version            int64                 `json:"version"`
}

func (s *GeoipDatabaseConfigurationMetadata) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoipDatabaseConfigurationMetadata() *GeoipDatabaseConfigurationMetadata {
	_ = "STUB: not implemented"
	return nil
}
