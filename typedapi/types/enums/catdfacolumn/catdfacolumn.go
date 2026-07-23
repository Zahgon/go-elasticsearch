package catdfacolumn

type CatDfaColumn struct {
	Name string
}

var (
	Assignmentexplanation = CatDfaColumn{"assignment_explanation"}

	Createtime = CatDfaColumn{"create_time"}

	Description = CatDfaColumn{"description"}

	Destindex = CatDfaColumn{"dest_index"}

	Failurereason = CatDfaColumn{"failure_reason"}

	Id = CatDfaColumn{"id"}

	Modelmemorylimit = CatDfaColumn{"model_memory_limit"}

	Nodeaddress = CatDfaColumn{"node.address"}

	Nodeephemeralid = CatDfaColumn{"node.ephemeral_id"}

	Nodeid = CatDfaColumn{"node.id"}

	Nodename = CatDfaColumn{"node.name"}

	Progress = CatDfaColumn{"progress"}

	Sourceindex = CatDfaColumn{"source_index"}

	State = CatDfaColumn{"state"}

	Type = CatDfaColumn{"type"}

	Version = CatDfaColumn{"version"}
)

func (c CatDfaColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatDfaColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatDfaColumn) String() string { _ = "STUB: not implemented"; return "" }
