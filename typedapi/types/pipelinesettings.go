package types

type PipelineSettings struct {
	PipelineBatchDelay int `json:"pipeline.batch.delay"`

	PipelineBatchSize int `json:"pipeline.batch.size"`

	PipelineWorkers int `json:"pipeline.workers"`

	QueueCheckpointWrites int `json:"queue.checkpoint.writes"`

	QueueMaxBytes string `json:"queue.max_bytes"`

	QueueType string `json:"queue.type"`
}

func (s *PipelineSettings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPipelineSettings() *PipelineSettings { _ = "STUB: not implemented"; return nil }

type PipelineSettingsVariant interface {
	PipelineSettingsCaster() *PipelineSettings
}

func (s *PipelineSettings) PipelineSettingsCaster() *PipelineSettings {
	_ = "STUB: not implemented"
	return nil
}
