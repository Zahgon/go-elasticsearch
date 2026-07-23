package types

type Indicators struct {
	DataStreamLifecycle  *DataStreamLifecycleIndicator  `json:"data_stream_lifecycle,omitempty"`
	Disk                 *DiskIndicator                 `json:"disk,omitempty"`
	FileSettings         *FileSettingsIndicator         `json:"file_settings,omitempty"`
	Ilm                  *IlmIndicator                  `json:"ilm,omitempty"`
	MasterIsStable       *MasterIsStableIndicator       `json:"master_is_stable,omitempty"`
	ProjectEncryptionKey *ProjectEncryptionKeyIndicator `json:"project_encryption_key,omitempty"`
	RepositoryIntegrity  *RepositoryIntegrityIndicator  `json:"repository_integrity,omitempty"`
	ShardsAvailability   *ShardsAvailabilityIndicator   `json:"shards_availability,omitempty"`
	ShardsCapacity       *ShardsCapacityIndicator       `json:"shards_capacity,omitempty"`
	Slm                  *SlmIndicator                  `json:"slm,omitempty"`
}

func NewIndicators() *Indicators { _ = "STUB: not implemented"; return nil }
