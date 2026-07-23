package restrictionworkflow

type RestrictionWorkflow struct {
	Name string
}

var (
	Searchapplicationquery = RestrictionWorkflow{"search_application_query"}
)

func (r RestrictionWorkflow) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RestrictionWorkflow) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r RestrictionWorkflow) String() string { _ = "STUB: not implemented"; return "" }
