package updatetransform

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Authorization   *types.TransformAuthorization   `json:"authorization,omitempty"`
	CreateTime      int64                           `json:"create_time"`
	Description     string                          `json:"description"`
	Dest            types.ReindexDestination        `json:"dest"`
	Frequency       types.Duration                  `json:"frequency,omitempty"`
	Id              string                          `json:"id"`
	Latest          *types.Latest                   `json:"latest,omitempty"`
	Meta_           types.Metadata                  `json:"_meta,omitempty"`
	Pivot           *types.Pivot                    `json:"pivot,omitempty"`
	RetentionPolicy *types.RetentionPolicyContainer `json:"retention_policy,omitempty"`
	Settings        types.Settings                  `json:"settings"`
	Source          types.ReindexSource             `json:"source"`
	Sync            *types.SyncContainer            `json:"sync,omitempty"`
	Version         string                          `json:"version"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
