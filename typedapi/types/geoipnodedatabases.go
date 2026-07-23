package types

type GeoIpNodeDatabases struct {
	Databases []GeoIpNodeDatabaseName `json:"databases"`

	FilesInTemp []string `json:"files_in_temp"`
}

func NewGeoIpNodeDatabases() *GeoIpNodeDatabases { _ = "STUB: not implemented"; return nil }
