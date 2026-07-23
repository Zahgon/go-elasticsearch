package catsegmentscolumn

type CatSegmentsColumn struct {
	Name string
}

var (
	Index = CatSegmentsColumn{"index"}

	Shard = CatSegmentsColumn{"shard"}

	Prirep = CatSegmentsColumn{"prirep"}

	Ip = CatSegmentsColumn{"ip"}

	Segment = CatSegmentsColumn{"segment"}

	Generation = CatSegmentsColumn{"generation"}

	Docscount = CatSegmentsColumn{"docs.count"}

	Docsdeleted = CatSegmentsColumn{"docs.deleted"}

	Size = CatSegmentsColumn{"size"}

	Sizememory = CatSegmentsColumn{"size.memory"}

	Committed = CatSegmentsColumn{"committed"}

	Searchable = CatSegmentsColumn{"searchable"}

	Version = CatSegmentsColumn{"version"}

	Compound = CatSegmentsColumn{"compound"}

	Id = CatSegmentsColumn{"id"}
)

func (c CatSegmentsColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatSegmentsColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatSegmentsColumn) String() string { _ = "STUB: not implemented"; return "" }
