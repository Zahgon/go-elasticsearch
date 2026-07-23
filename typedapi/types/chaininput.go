package types

type ChainInput struct {
	Inputs []map[string]WatcherInput `json:"inputs"`
}

func NewChainInput() *ChainInput { _ = "STUB: not implemented"; return nil }

type ChainInputVariant interface {
	ChainInputCaster() *ChainInput
}

func (s *ChainInput) ChainInputCaster() *ChainInput { _ = "STUB: not implemented"; return nil }
