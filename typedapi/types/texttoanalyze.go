package types

type TextToAnalyze []string

type TextToAnalyzeVariant interface {
	TextToAnalyzeCaster() *TextToAnalyze
}
