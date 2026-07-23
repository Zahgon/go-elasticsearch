package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereinputtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
)

type _cohereTaskSettings struct {
	v *types.CohereTaskSettings
}

func NewCohereTaskSettings(inputtype cohereinputtype.CohereInputType) *_cohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereTaskSettings) InputType(inputtype cohereinputtype.CohereInputType) *_cohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereTaskSettings) ReturnDocuments(returndocuments bool) *_cohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereTaskSettings) TopN(topn int) *_cohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereTaskSettings) Truncate(truncate coheretruncatetype.CohereTruncateType) *_cohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cohereTaskSettings) CohereTaskSettingsCaster() *types.CohereTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
