package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fuzziness struct {
	v types.Fuzziness
}

func NewFuzziness() *_fuzziness { _ = "STUB: not implemented"; return nil }

func (u *_fuzziness) String(string string) *_fuzziness { _ = "STUB: not implemented"; return nil }

func (u *_fuzziness) Int(int int) *_fuzziness { _ = "STUB: not implemented"; return nil }

func (u *_fuzziness) FuzzinessCaster() *types.Fuzziness { _ = "STUB: not implemented"; return nil }
