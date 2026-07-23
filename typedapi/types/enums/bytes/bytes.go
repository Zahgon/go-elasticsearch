package bytes

type Bytes struct {
	Name string
}

var (
	B = Bytes{"b"}

	Kb = Bytes{"kb"}

	Mb = Bytes{"mb"}

	Gb = Bytes{"gb"}

	Tb = Bytes{"tb"}

	Pb = Bytes{"pb"}
)

func (b Bytes) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Bytes) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (b Bytes) String() string { _ = "STUB: not implemented"; return "" }
