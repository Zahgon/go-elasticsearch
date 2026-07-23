package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type TransformStatsHealth struct {
	Issues []TransformHealthIssue    `json:"issues,omitempty"`
	Status healthstatus.HealthStatus `json:"status"`
}

func NewTransformStatsHealth() *TransformStatsHealth { _ = "STUB: not implemented"; return nil }
