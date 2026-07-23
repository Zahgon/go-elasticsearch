package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoBounds struct {
	v types.GeoBounds
}

func NewGeoBounds() *_geoBounds { _ = "STUB: not implemented"; return nil }

func (u *_geoBounds) CoordsGeoBounds(coordsgeobounds types.CoordsGeoBoundsVariant) *_geoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_coordsGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoBounds) TopLeftBottomRightGeoBounds(topleftbottomrightgeobounds types.TopLeftBottomRightGeoBoundsVariant) *_geoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_topLeftBottomRightGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoBounds) TopRightBottomLeftGeoBounds(toprightbottomleftgeobounds types.TopRightBottomLeftGeoBoundsVariant) *_geoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_topRightBottomLeftGeoBounds) GeoBoundsCaster() *types.GeoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoBounds) WktGeoBounds(wktgeobounds types.WktGeoBoundsVariant) *_geoBounds {
	_ = "STUB: not implemented"
	return nil
}

func (u *_wktGeoBounds) GeoBoundsCaster() *types.GeoBounds { _ = "STUB: not implemented"; return nil }

func (u *_geoBounds) GeoBoundsCaster() *types.GeoBounds { _ = "STUB: not implemented"; return nil }
