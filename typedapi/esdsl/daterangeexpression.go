package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateRangeExpression struct {
	v *types.DateRangeExpression
}

func NewDateRangeExpression() *_dateRangeExpression { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeExpression) From(fielddatemath types.FieldDateMathVariant) *_dateRangeExpression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeExpression) Key(key string) *_dateRangeExpression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeExpression) To(fielddatemath types.FieldDateMathVariant) *_dateRangeExpression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeExpression) DateRangeExpressionCaster() *types.DateRangeExpression {
	_ = "STUB: not implemented"
	return nil
}
