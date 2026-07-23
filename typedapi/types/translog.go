package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/translogdurability"
)

type Translog struct {
	Durability *translogdurability.TranslogDurability `json:"durability,omitempty"`

	FlushThresholdSize ByteSize           `json:"flush_threshold_size,omitempty"`
	Retention          *TranslogRetention `json:"retention,omitempty"`

	SyncInterval Duration `json:"sync_interval,omitempty"`
}

func (s *Translog) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTranslog() *Translog { _ = "STUB: not implemented"; return nil }

type TranslogVariant interface {
	TranslogCaster() *Translog
}

func (s *Translog) TranslogCaster() *Translog { _ = "STUB: not implemented"; return nil }
