package types

type CategorizeTextAnalyzer any

type CategorizeTextAnalyzerVariant interface {
	CategorizeTextAnalyzerCaster() *CategorizeTextAnalyzer
}
