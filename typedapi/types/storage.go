package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/storagetype"
)

type Storage struct {
	AllowMmap *bool `json:"allow_mmap,omitempty"`

	StatsRefreshInterval Duration                `json:"stats_refresh_interval,omitempty"`
	Type                 storagetype.StorageType `json:"type"`
}

func (s *Storage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStorage() *Storage { _ = "STUB: not implemented"; return nil }

type StorageVariant interface {
	StorageCaster() *Storage
}

func (s *Storage) StorageCaster() *Storage { _ = "STUB: not implemented"; return nil }
