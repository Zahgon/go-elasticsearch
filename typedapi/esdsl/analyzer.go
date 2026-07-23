package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _analyzer struct {
	v types.Analyzer
}

func NewAnalyzer() *_analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) UnknownAnalyzer(unknown json.RawMessage) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) CustomAnalyzer(customanalyzer types.CustomAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) FingerprintAnalyzer(fingerprintanalyzer types.FingerprintAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_fingerprintAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) KeywordAnalyzer(keywordanalyzer types.KeywordAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_keywordAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) NoriAnalyzer(norianalyzer types.NoriAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_noriAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) PatternAnalyzer(patternanalyzer types.PatternAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_patternAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SimpleAnalyzer(simpleanalyzer types.SimpleAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_simpleAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) StandardAnalyzer(standardanalyzer types.StandardAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_standardAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) StopAnalyzer(stopanalyzer types.StopAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_stopAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) WhitespaceAnalyzer(whitespaceanalyzer types.WhitespaceAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_whitespaceAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) IcuAnalyzer(icuanalyzer types.IcuAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_icuAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) KuromojiAnalyzer(kuromojianalyzer types.KuromojiAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_kuromojiAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SnowballAnalyzer(snowballanalyzer types.SnowballAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_snowballAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) ArabicAnalyzer(arabicanalyzer types.ArabicAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_arabicAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) ArmenianAnalyzer(armeniananalyzer types.ArmenianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_armenianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) BasqueAnalyzer(basqueanalyzer types.BasqueAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_basqueAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) BengaliAnalyzer(bengalianalyzer types.BengaliAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bengaliAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) BrazilianAnalyzer(braziliananalyzer types.BrazilianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_brazilianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) BulgarianAnalyzer(bulgariananalyzer types.BulgarianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bulgarianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) CatalanAnalyzer(catalananalyzer types.CatalanAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_catalanAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) ChineseAnalyzer(chineseanalyzer types.ChineseAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_chineseAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) CjkAnalyzer(cjkanalyzer types.CjkAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_cjkAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) CzechAnalyzer(czechanalyzer types.CzechAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_czechAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) DanishAnalyzer(danishanalyzer types.DanishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_danishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) DutchAnalyzer(dutchanalyzer types.DutchAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_dutchAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) EnglishAnalyzer(englishanalyzer types.EnglishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_englishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) EstonianAnalyzer(estoniananalyzer types.EstonianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_estonianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) FinnishAnalyzer(finnishanalyzer types.FinnishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_finnishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) FrenchAnalyzer(frenchanalyzer types.FrenchAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_frenchAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) GalicianAnalyzer(galiciananalyzer types.GalicianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_galicianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) GermanAnalyzer(germananalyzer types.GermanAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_germanAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) GreekAnalyzer(greekanalyzer types.GreekAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_greekAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) HindiAnalyzer(hindianalyzer types.HindiAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_hindiAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) HungarianAnalyzer(hungariananalyzer types.HungarianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_hungarianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) IndonesianAnalyzer(indonesiananalyzer types.IndonesianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_indonesianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) IrishAnalyzer(irishanalyzer types.IrishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_irishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) ItalianAnalyzer(italiananalyzer types.ItalianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_italianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) LatvianAnalyzer(latviananalyzer types.LatvianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_latvianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) LithuanianAnalyzer(lithuaniananalyzer types.LithuanianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_lithuanianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) NorwegianAnalyzer(norwegiananalyzer types.NorwegianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_norwegianAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) PersianAnalyzer(persiananalyzer types.PersianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_persianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) PortugueseAnalyzer(portugueseanalyzer types.PortugueseAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_portugueseAnalyzer) AnalyzerCaster() *types.Analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_analyzer) RomanianAnalyzer(romaniananalyzer types.RomanianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_romanianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) RussianAnalyzer(russiananalyzer types.RussianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_russianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SerbianAnalyzer(serbiananalyzer types.SerbianAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_serbianAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SoraniAnalyzer(soranianalyzer types.SoraniAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_soraniAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SpanishAnalyzer(spanishanalyzer types.SpanishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_spanishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) SwedishAnalyzer(swedishanalyzer types.SwedishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_swedishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) TurkishAnalyzer(turkishanalyzer types.TurkishAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_turkishAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) ThaiAnalyzer(thaianalyzer types.ThaiAnalyzerVariant) *_analyzer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_thaiAnalyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }

func (u *_analyzer) AnalyzerCaster() *types.Analyzer { _ = "STUB: not implemented"; return nil }
