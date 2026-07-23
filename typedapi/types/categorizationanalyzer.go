package types

type CategorizationAnalyzer any

type CategorizationAnalyzerVariant interface {
	CategorizationAnalyzerCaster() *CategorizationAnalyzer
}
