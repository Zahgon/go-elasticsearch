package types

type Dip struct {
	ChangePoint int     `json:"change_point"`
	PValue      Float64 `json:"p_value"`
}

func (s *Dip) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDip() *Dip { _ = "STUB: not implemented"; return nil }
