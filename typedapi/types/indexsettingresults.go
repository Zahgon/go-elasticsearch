package types

type IndexSettingResults struct {
	AppliedToDataStreamAndBackingIndices []string `json:"applied_to_data_stream_and_backing_indices"`

	AppliedToDataStreamOnly []string                  `json:"applied_to_data_stream_only"`
	Errors                  []DataStreamSettingsError `json:"errors,omitempty"`
}

func NewIndexSettingResults() *IndexSettingResults { _ = "STUB: not implemented"; return nil }
