package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _storedScript struct {
	v *types.StoredScript
}

func NewStoredScript(lang scriptlanguage.ScriptLanguage) *_storedScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storedScript) Lang(lang scriptlanguage.ScriptLanguage) *_storedScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storedScript) Options(options map[string]string) *_storedScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storedScript) AddOption(key string, value string) *_storedScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storedScript) Source(scriptsource types.ScriptSourceVariant) *_storedScript {
	_ = "STUB: not implemented"
	return nil
}

func (s *_storedScript) StoredScriptCaster() *types.StoredScript {
	_ = "STUB: not implemented"
	return nil
}
