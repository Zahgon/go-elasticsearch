package types

type DenseVectorIndexOptionsRescoreVector struct {
	Oversample float32 `json:"oversample"`
}

func (s *DenseVectorIndexOptionsRescoreVector) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDenseVectorIndexOptionsRescoreVector() *DenseVectorIndexOptionsRescoreVector {
	_ = "STUB: not implemented"
	return nil
}

type DenseVectorIndexOptionsRescoreVectorVariant interface {
	DenseVectorIndexOptionsRescoreVectorCaster() *DenseVectorIndexOptionsRescoreVector
}

func (s *DenseVectorIndexOptionsRescoreVector) DenseVectorIndexOptionsRescoreVectorCaster() *DenseVectorIndexOptionsRescoreVector {
	_ = "STUB: not implemented"
	return nil
}
