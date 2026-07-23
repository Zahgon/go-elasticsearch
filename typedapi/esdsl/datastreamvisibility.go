package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamVisibility struct {
	v *types.DataStreamVisibility
}

func NewDataStreamVisibility() *_dataStreamVisibility { _ = "STUB: not implemented"; return nil }

func (s *_dataStreamVisibility) AllowCustomRouting(allowcustomrouting bool) *_dataStreamVisibility {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamVisibility) FailureStore(failurestore bool) *_dataStreamVisibility {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamVisibility) Hidden(hidden bool) *_dataStreamVisibility {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamVisibility) DataStreamVisibilityCaster() *types.DataStreamVisibility {
	_ = "STUB: not implemented"
	return nil
}
