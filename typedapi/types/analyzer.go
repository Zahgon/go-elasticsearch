package types

type Analyzer any

type AnalyzerVariant interface {
	AnalyzerCaster() *Analyzer
}
