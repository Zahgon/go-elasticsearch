package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalyticsDestination struct {
	v *types.DataframeAnalyticsDestination
}

func NewDataframeAnalyticsDestination() *_dataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsDestination) Index(indexname string) *_dataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsDestination) ResultsField(field string) *_dataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalyticsDestination) DataframeAnalyticsDestinationCaster() *types.DataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}
