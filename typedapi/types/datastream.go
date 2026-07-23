package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/managedby"
)

type DataStream struct {
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`

	FailureStore *FailureStore `json:"failure_store,omitempty"`

	Generation int `json:"generation"`

	Hidden bool `json:"hidden"`

	IlmPolicy *string `json:"ilm_policy,omitempty"`

	IndexMode *indexmode.IndexMode `json:"index_mode,omitempty"`

	Indices []DataStreamIndex `json:"indices"`

	Lifecycle *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`

	Mappings *TypeMapping `json:"mappings,omitempty"`

	Meta_ Metadata `json:"_meta,omitempty"`

	Name string `json:"name"`

	NextGenerationManagedBy managedby.ManagedBy `json:"next_generation_managed_by"`

	PreferIlm bool `json:"prefer_ilm"`

	Replicated *bool `json:"replicated,omitempty"`

	RolloverOnWrite bool `json:"rollover_on_write"`

	Settings IndexSettings `json:"settings"`

	Status healthstatus.HealthStatus `json:"status"`

	System *bool `json:"system,omitempty"`

	Template string `json:"template"`

	TimestampField DataStreamTimestampField `json:"timestamp_field"`
}

func (s *DataStream) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataStream() *DataStream { _ = "STUB: not implemented"; return nil }
