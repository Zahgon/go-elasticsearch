package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _like struct {
	v types.Like
}

func NewLike() *_like { _ = "STUB: not implemented"; return nil }

func (u *_like) String(string string) *_like { _ = "STUB: not implemented"; return nil }

func (u *_like) LikeDocument(likedocument types.LikeDocumentVariant) *_like {
	_ = "STUB: not implemented"
	return nil
}

func (u *_likeDocument) LikeCaster() *types.Like { _ = "STUB: not implemented"; return nil }

func (u *_like) LikeCaster() *types.Like { _ = "STUB: not implemented"; return nil }
