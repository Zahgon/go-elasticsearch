package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _documentRating struct {
	v *types.DocumentRating
}

func NewDocumentRating(rating int) *_documentRating { _ = "STUB: not implemented"; return nil }

func (s *_documentRating) Id_(id string) *_documentRating { _ = "STUB: not implemented"; return nil }

func (s *_documentRating) Index_(indexname string) *_documentRating {
	_ = "STUB: not implemented"
	return nil
}

func (s *_documentRating) Rating(rating int) *_documentRating {
	_ = "STUB: not implemented"
	return nil
}

func (s *_documentRating) DocumentRatingCaster() *types.DocumentRating {
	_ = "STUB: not implemented"
	return nil
}
