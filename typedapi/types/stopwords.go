package types

type StopWords any

type StopWordsVariant interface {
	StopWordsCaster() *StopWords
}
