package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalyticsSource struct {
	v *types.DataframeAnalyticsSource
}

func NewDataframeAnalyticsSource() *_dataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsSource) Index(indices ...string) *_dataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsSource) Query(query types.QueryVariant) *_dataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_dataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsSource) Source_(source_ types.DataframeAnalysisAnalyzedFieldsVariant) *_dataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsSource) DataframeAnalyticsSourceCaster() *types.DataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}
