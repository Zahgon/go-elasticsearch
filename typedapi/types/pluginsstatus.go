package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shutdownstatus"
)

type PluginsStatus struct {
	Status shutdownstatus.ShutdownStatus `json:"status"`
}

func NewPluginsStatus() *PluginsStatus { _ = "STUB: not implemented"; return nil }
