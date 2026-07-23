package types

type GeoIpDownloadStatistics struct {
	DatabasesCount int `json:"databases_count"`

	ExpiredDatabases int `json:"expired_databases"`

	FailedDownloads int `json:"failed_downloads"`

	SkippedUpdates int `json:"skipped_updates"`

	SuccessfulDownloads int `json:"successful_downloads"`

	TotalDownloadTime int64 `json:"total_download_time"`
}

func (s *GeoIpDownloadStatistics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoIpDownloadStatistics() *GeoIpDownloadStatistics { _ = "STUB: not implemented"; return nil }
