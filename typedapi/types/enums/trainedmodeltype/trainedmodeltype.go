package trainedmodeltype

type TrainedModelType struct {
	Name string
}

var (
	Treeensemble = TrainedModelType{"tree_ensemble"}

	Langident = TrainedModelType{"lang_ident"}

	Pytorch = TrainedModelType{"pytorch"}
)

func (t TrainedModelType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TrainedModelType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (t TrainedModelType) String() string { _ = "STUB: not implemented"; return "" }
