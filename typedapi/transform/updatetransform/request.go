package updatetransform

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Description *string `json:"description,omitempty"`

	Dest *types.TransformDestination `json:"dest,omitempty"`

	Frequency types.Duration `json:"frequency,omitempty"`

	Meta_ types.Metadata `json:"_meta,omitempty"`

	RetentionPolicy *types.RetentionPolicyContainer `json:"retention_policy,omitempty"`

	Settings *types.Settings `json:"settings,omitempty"`

	Source *types.TransformSource `json:"source,omitempty"`

	Sync *types.SyncContainer `json:"sync,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
