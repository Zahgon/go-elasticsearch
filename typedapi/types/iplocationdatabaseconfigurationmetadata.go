package types

type IpLocationDatabaseConfigurationMetadata struct {
	Database           DatabaseConfigurationFull `json:"database"`
	Id                 string                    `json:"id"`
	ModifiedDate       *int64                    `json:"modified_date,omitempty"`
	ModifiedDateMillis *int64                    `json:"modified_date_millis,omitempty"`
	Version            int64                     `json:"version"`
}

func (s *IpLocationDatabaseConfigurationMetadata) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIpLocationDatabaseConfigurationMetadata() *IpLocationDatabaseConfigurationMetadata {
	_ = "STUB: not implemented"
	return nil
}
