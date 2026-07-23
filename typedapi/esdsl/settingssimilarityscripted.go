package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsSimilarityScripted struct {
	v *types.SettingsSimilarityScripted
}

func NewSettingsSimilarityScripted(script types.ScriptVariant) *_settingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityScripted) Script(script types.ScriptVariant) *_settingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityScripted) WeightScript(weightscript types.ScriptVariant) *_settingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityScripted) SettingsSimilarityScriptedCaster() *types.SettingsSimilarityScripted {
	_ = "STUB: not implemented"
	return nil
}
