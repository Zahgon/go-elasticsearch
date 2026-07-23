package stopwordlanguage

type StopWordLanguage struct {
	Name string
}

var (
	Arabic = StopWordLanguage{"_arabic_"}

	Armenian = StopWordLanguage{"_armenian_"}

	Basque = StopWordLanguage{"_basque_"}

	Bengali = StopWordLanguage{"_bengali_"}

	Brazilian = StopWordLanguage{"_brazilian_"}

	Bulgarian = StopWordLanguage{"_bulgarian_"}

	Catalan = StopWordLanguage{"_catalan_"}

	Cjk = StopWordLanguage{"_cjk_"}

	Czech = StopWordLanguage{"_czech_"}

	Danish = StopWordLanguage{"_danish_"}

	Dutch = StopWordLanguage{"_dutch_"}

	English = StopWordLanguage{"_english_"}

	Estonian = StopWordLanguage{"_estonian_"}

	Finnish = StopWordLanguage{"_finnish_"}

	French = StopWordLanguage{"_french_"}

	Galician = StopWordLanguage{"_galician_"}

	German = StopWordLanguage{"_german_"}

	Greek = StopWordLanguage{"_greek_"}

	Hindi = StopWordLanguage{"_hindi_"}

	Hungarian = StopWordLanguage{"_hungarian_"}

	Indonesian = StopWordLanguage{"_indonesian_"}

	Irish = StopWordLanguage{"_irish_"}

	Italian = StopWordLanguage{"_italian_"}

	Latvian = StopWordLanguage{"_latvian_"}

	Lithuanian = StopWordLanguage{"_lithuanian_"}

	Norwegian = StopWordLanguage{"_norwegian_"}

	Persian = StopWordLanguage{"_persian_"}

	Portuguese = StopWordLanguage{"_portuguese_"}

	Romanian = StopWordLanguage{"_romanian_"}

	Russian = StopWordLanguage{"_russian_"}

	Serbian = StopWordLanguage{"_serbian_"}

	Sorani = StopWordLanguage{"_sorani_"}

	Spanish = StopWordLanguage{"_spanish_"}

	Swedish = StopWordLanguage{"_swedish_"}

	Thai = StopWordLanguage{"_thai_"}

	Turkish = StopWordLanguage{"_turkish_"}

	None = StopWordLanguage{"_none_"}
)

func (s StopWordLanguage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StopWordLanguage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StopWordLanguage) String() string { _ = "STUB: not implemented"; return "" }
