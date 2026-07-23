package types

type TokenizationConfigContainer struct {
	Bert *NlpBertTokenizationConfig `json:"bert,omitempty"`

	BertJa *NlpBertTokenizationConfig `json:"bert_ja,omitempty"`

	Mpnet *NlpBertTokenizationConfig `json:"mpnet,omitempty"`

	Roberta    *NlpRobertaTokenizationConfig `json:"roberta,omitempty"`
	XlmRoberta *XlmRobertaTokenizationConfig `json:"xlm_roberta,omitempty"`
}

func NewTokenizationConfigContainer() *TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

type TokenizationConfigContainerVariant interface {
	TokenizationConfigContainerCaster() *TokenizationConfigContainer
}

func (s *TokenizationConfigContainer) TokenizationConfigContainerCaster() *TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}
