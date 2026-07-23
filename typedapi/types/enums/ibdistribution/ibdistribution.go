package ibdistribution

type IBDistribution struct {
	Name string
}

var (
	Ll = IBDistribution{"ll"}

	Spl = IBDistribution{"spl"}
)

func (i IBDistribution) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *IBDistribution) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i IBDistribution) String() string { _ = "STUB: not implemented"; return "" }
