package storagetype

type StorageType struct {
	Name string
}

var (
	Fs = StorageType{"fs"}

	Niofs = StorageType{"niofs"}

	Mmapfs = StorageType{"mmapfs"}

	Hybridfs = StorageType{"hybridfs"}
)

func (s StorageType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StorageType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s StorageType) String() string { _ = "STUB: not implemented"; return "" }
