package types

type DocumentRating struct {
	Id_ string `json:"_id"`

	Index_ string `json:"_index"`

	Rating int `json:"rating"`
}

func (s *DocumentRating) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDocumentRating() *DocumentRating { _ = "STUB: not implemented"; return nil }

type DocumentRatingVariant interface {
	DocumentRatingCaster() *DocumentRating
}

func (s *DocumentRating) DocumentRatingCaster() *DocumentRating {
	_ = "STUB: not implemented"
	return nil
}
