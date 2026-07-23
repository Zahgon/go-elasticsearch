package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noridecompoundmode"
)

type _noriAnalyzer struct {
	v *types.NoriAnalyzer
}

func NewNoriAnalyzer() *_noriAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_noriAnalyzer) DecompoundMode(decompoundmode noridecompoundmode.NoriDecompoundMode) *_noriAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriAnalyzer) Stoptags(stoptags ...string) *_noriAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriAnalyzer) UserDictionary(userdictionary string) *_noriAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriAnalyzer) Version(versionstring string) *_noriAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_noriAnalyzer) NoriAnalyzerCaster() *types.NoriAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
