package types

type Ssl struct {
	Http      FeatureToggle `json:"http"`
	Transport FeatureToggle `json:"transport"`
}

func NewSsl() *Ssl { _ = "STUB: not implemented"; return nil }
