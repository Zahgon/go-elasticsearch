package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _jobConfig struct {
	v *types.JobConfig
}

func NewJobConfig(analysisconfig types.AnalysisConfigVariant, datadescription types.DataDescriptionVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) AllowLazyOpen(allowlazyopen bool) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) AnalysisLimits(analysislimits types.AnalysisLimitsVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) BackgroundPersistInterval(duration types.DurationVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) CustomSettings(customsettings json.RawMessage) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) DataDescription(datadescription types.DataDescriptionVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) DatafeedConfig(datafeedconfig types.DatafeedConfigVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) Description(description string) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) Groups(groups ...string) *_jobConfig { _ = "STUB: not implemented"; return nil }

func (s *_jobConfig) JobId(id string) *_jobConfig { _ = "STUB: not implemented"; return nil }

func (s *_jobConfig) JobType(jobtype string) *_jobConfig { _ = "STUB: not implemented"; return nil }

func (s *_jobConfig) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) RenormalizationWindowDays(renormalizationwindowdays int64) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) ResultsIndexName(indexname string) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) ResultsRetentionDays(resultsretentiondays int64) *_jobConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jobConfig) JobConfigCaster() *types.JobConfig { _ = "STUB: not implemented"; return nil }
