package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snowballlanguage"
)

type _snowballAnalyzer struct {
	v *types.SnowballAnalyzer
}

func NewSnowballAnalyzer(language snowballlanguage.SnowballLanguage) *_snowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballAnalyzer) Language(language snowballlanguage.SnowballLanguage) *_snowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballAnalyzer) Stopwords(stopwords types.StopWordsVariant) *_snowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballAnalyzer) Version(versionstring string) *_snowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_snowballAnalyzer) SnowballAnalyzerCaster() *types.SnowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
