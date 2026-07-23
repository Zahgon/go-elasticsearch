package include

type Include struct {
	Name string
}

var (
	Definition = Include{"definition"}

	Featureimportancebaseline = Include{"feature_importance_baseline"}

	Hyperparameters = Include{"hyperparameters"}

	Totalfeatureimportance = Include{"total_feature_importance"}

	Definitionstatus = Include{"definition_status"}
)

func (i Include) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Include) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i Include) String() string { _ = "STUB: not implemented"; return "" }
