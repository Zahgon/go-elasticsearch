package storageoption

type StorageOption struct {
	Name string
}

var (
	Fullcopy = StorageOption{"full_copy"}

	Sharedcache = StorageOption{"shared_cache"}
)

func (s StorageOption) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StorageOption) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StorageOption) String() string { _ = "STUB: not implemented"; return "" }
