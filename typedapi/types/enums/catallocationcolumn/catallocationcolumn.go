package catallocationcolumn

type CatAllocationColumn struct {
	Name string
}

var (
	Shards = CatAllocationColumn{"shards"}

	Shardsundesired = CatAllocationColumn{"shards.undesired"}

	Writeloadforecast = CatAllocationColumn{"write_load.forecast"}

	Diskindicesforecast = CatAllocationColumn{"disk.indices.forecast"}

	Diskindices = CatAllocationColumn{"disk.indices"}

	Diskused = CatAllocationColumn{"disk.used"}

	Diskavail = CatAllocationColumn{"disk.avail"}

	Disktotal = CatAllocationColumn{"disk.total"}

	Diskpercent = CatAllocationColumn{"disk.percent"}

	Host = CatAllocationColumn{"host"}

	Ip = CatAllocationColumn{"ip"}

	Node = CatAllocationColumn{"node"}

	Noderole = CatAllocationColumn{"node.role"}
)

func (c CatAllocationColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatAllocationColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatAllocationColumn) String() string { _ = "STUB: not implemented"; return "" }
