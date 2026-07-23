package types

type SpanGapQuery map[string]int

type SpanGapQueryVariant interface {
	SpanGapQueryCaster() *SpanGapQuery
}
