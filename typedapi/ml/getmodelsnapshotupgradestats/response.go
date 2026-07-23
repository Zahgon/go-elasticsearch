package getmodelsnapshotupgradestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count                 int64                        `json:"count"`
	ModelSnapshotUpgrades []types.ModelSnapshotUpgrade `json:"model_snapshot_upgrades"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
