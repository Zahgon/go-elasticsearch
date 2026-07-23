package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/excludefrequent"
)

type _detector struct {
	v *types.Detector
}

func NewDetector() *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) ByFieldName(field string) *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) CustomRules(customrules ...types.DetectionRuleVariant) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) CustomRulesValues(customrulesvalues []types.DetectionRule) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) DetectorDescription(detectordescription string) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) DetectorIndex(detectorindex int) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) FieldName(field string) *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) Function(function string) *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) OverFieldName(field string) *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) PartitionFieldName(field string) *_detector {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detector) UseNull(usenull bool) *_detector { _ = "STUB: not implemented"; return nil }

func (s *_detector) DetectorCaster() *types.Detector { _ = "STUB: not implemented"; return nil }
