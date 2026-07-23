package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdownstatus"
)

type PersistentTaskStatus struct {
	Status shutdownstatus.ShutdownStatus `json:"status"`
}

func NewPersistentTaskStatus() *PersistentTaskStatus { _ = "STUB: not implemented"; return nil }
