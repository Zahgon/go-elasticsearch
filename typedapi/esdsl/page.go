package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _page struct {
	v *types.Page
}

func NewPage() *_page { _ = "STUB: not implemented"; return nil }

func (s *_page) From(from int) *_page { _ = "STUB: not implemented"; return nil }

func (s *_page) Size(size int) *_page { _ = "STUB: not implemented"; return nil }

func (s *_page) PageCaster() *types.Page { _ = "STUB: not implemented"; return nil }
