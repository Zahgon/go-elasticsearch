package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settings struct {
	v *types.Settings
}

func NewSettings() *_settings { _ = "STUB: not implemented"; return nil }

func (s *_settings) AlignCheckpoints(aligncheckpoints bool) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) DatesAsEpochMillis(datesasepochmillis bool) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) DeduceMappings(deducemappings bool) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) DocsPerSecond(docspersecond float32) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) MaxPageSearchSize(maxpagesearchsize int) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) NumFailureRetries(numfailureretries int) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) Unattended(unattended bool) *_settings { _ = "STUB: not implemented"; return nil }

func (s *_settings) UsePointInTime(usepointintime bool) *_settings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settings) SettingsCaster() *types.Settings { _ = "STUB: not implemented"; return nil }
