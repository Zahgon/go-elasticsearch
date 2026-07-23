package types

type TransformSummary struct {
	Authorization *TransformAuthorization `json:"authorization,omitempty"`

	CreateTime       *int64   `json:"create_time,omitempty"`
	CreateTimeString DateTime `json:"create_time_string,omitempty"`

	Description *string `json:"description,omitempty"`

	Dest      ReindexDestination `json:"dest"`
	Frequency Duration           `json:"frequency,omitempty"`
	Id        string             `json:"id"`
	Latest    *Latest            `json:"latest,omitempty"`
	Meta_     Metadata           `json:"_meta,omitempty"`

	Pivot           *Pivot                    `json:"pivot,omitempty"`
	RetentionPolicy *RetentionPolicyContainer `json:"retention_policy,omitempty"`

	Settings *Settings `json:"settings,omitempty"`

	Source TransformSource `json:"source"`

	Sync *SyncContainer `json:"sync,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *TransformSummary) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransformSummary() *TransformSummary { _ = "STUB: not implemented"; return nil }
