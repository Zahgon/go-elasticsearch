package types

type FunctionScore struct {
	Exp DecayFunction `json:"exp,omitempty"`

	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *Query                         `json:"filter,omitempty"`

	Gauss DecayFunction `json:"gauss,omitempty"`

	Linear DecayFunction `json:"linear,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	RandomScore *RandomScoreFunction `json:"random_score,omitempty"`

	ScriptScore *ScriptScoreFunction `json:"script_score,omitempty"`
	Weight      *Float64             `json:"weight,omitempty"`
}

func (s *FunctionScore) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFunctionScore() *FunctionScore { _ = "STUB: not implemented"; return nil }

type FunctionScoreVariant interface {
	FunctionScoreCaster() *FunctionScore
}

func (s *FunctionScore) FunctionScoreCaster() *FunctionScore { _ = "STUB: not implemented"; return nil }
