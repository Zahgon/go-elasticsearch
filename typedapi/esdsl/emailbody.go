package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _emailBody struct {
	v *types.EmailBody
}

func NewEmailBody() *_emailBody { _ = "STUB: not implemented"; return nil }

func (s *_emailBody) Html(html string) *_emailBody { _ = "STUB: not implemented"; return nil }

func (s *_emailBody) Text(text string) *_emailBody { _ = "STUB: not implemented"; return nil }

func (s *_emailBody) EmailBodyCaster() *types.EmailBody { _ = "STUB: not implemented"; return nil }
