package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _suggestFuzziness struct {
	v *types.SuggestFuzziness
}

func NewSuggestFuzziness() *_suggestFuzziness { _ = "STUB: not implemented"; return nil }

func (s *_suggestFuzziness) Fuzziness(fuzziness types.FuzzinessVariant) *_suggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestFuzziness) MinLength(minlength int) *_suggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestFuzziness) PrefixLength(prefixlength int) *_suggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestFuzziness) Transpositions(transpositions bool) *_suggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestFuzziness) UnicodeAware(unicodeaware bool) *_suggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}

func (s *_suggestFuzziness) SuggestFuzzinessCaster() *types.SuggestFuzziness {
	_ = "STUB: not implemented"
	return nil
}
