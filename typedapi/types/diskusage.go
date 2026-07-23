package types

type DiskUsage struct {
	FreeBytes       int64   `json:"free_bytes"`
	FreeDiskPercent Float64 `json:"free_disk_percent"`
	Path            string  `json:"path"`
	TotalBytes      int64   `json:"total_bytes"`
	UsedBytes       int64   `json:"used_bytes"`
	UsedDiskPercent Float64 `json:"used_disk_percent"`
}

func (s *DiskUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDiskUsage() *DiskUsage { _ = "STUB: not implemented"; return nil }
