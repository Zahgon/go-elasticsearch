package snowballlanguage

type SnowballLanguage struct {
	Name string
}

var (
	Arabic = SnowballLanguage{"Arabic"}

	Armenian = SnowballLanguage{"Armenian"}

	Basque = SnowballLanguage{"Basque"}

	Catalan = SnowballLanguage{"Catalan"}

	Danish = SnowballLanguage{"Danish"}

	Dutch = SnowballLanguage{"Dutch"}

	English = SnowballLanguage{"English"}

	Estonian = SnowballLanguage{"Estonian"}

	Finnish = SnowballLanguage{"Finnish"}

	French = SnowballLanguage{"French"}

	German = SnowballLanguage{"German"}

	German2 = SnowballLanguage{"German2"}

	Hungarian = SnowballLanguage{"Hungarian"}

	Italian = SnowballLanguage{"Italian"}

	Irish = SnowballLanguage{"Irish"}

	Kp = SnowballLanguage{"Kp"}

	Lithuanian = SnowballLanguage{"Lithuanian"}

	Lovins = SnowballLanguage{"Lovins"}

	Norwegian = SnowballLanguage{"Norwegian"}

	Porter = SnowballLanguage{"Porter"}

	Portuguese = SnowballLanguage{"Portuguese"}

	Romanian = SnowballLanguage{"Romanian"}

	Russian = SnowballLanguage{"Russian"}

	Serbian = SnowballLanguage{"Serbian"}

	Spanish = SnowballLanguage{"Spanish"}

	Swedish = SnowballLanguage{"Swedish"}

	Turkish = SnowballLanguage{"Turkish"}
)

func (s SnowballLanguage) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SnowballLanguage) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SnowballLanguage) String() string { _ = "STUB: not implemented"; return "" }
