package densevectorindexoptionstype

type DenseVectorIndexOptionsType struct {
	Name string
}

var (
	Bbqflat = DenseVectorIndexOptionsType{"bbq_flat"}

	Bbqhnsw = DenseVectorIndexOptionsType{"bbq_hnsw"}

	Bbqdisk = DenseVectorIndexOptionsType{"bbq_disk"}

	Flat = DenseVectorIndexOptionsType{"flat"}

	Hnsw = DenseVectorIndexOptionsType{"hnsw"}

	Int4flat = DenseVectorIndexOptionsType{"int4_flat"}

	Int4hnsw = DenseVectorIndexOptionsType{"int4_hnsw"}

	Int8flat = DenseVectorIndexOptionsType{"int8_flat"}

	Int8hnsw = DenseVectorIndexOptionsType{"int8_hnsw"}
)

func (d DenseVectorIndexOptionsType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DenseVectorIndexOptionsType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d DenseVectorIndexOptionsType) String() string { _ = "STUB: not implemented"; return "" }
