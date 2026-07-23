package fingerprintdigest

type FingerprintDigest struct {
	Name string
}

var (
	Md5 = FingerprintDigest{"MD5"}

	Sha1 = FingerprintDigest{"SHA-1"}

	Sha256 = FingerprintDigest{"SHA-256"}

	Sha512 = FingerprintDigest{"SHA-512"}

	MurmurHash3 = FingerprintDigest{"MurmurHash3"}
)

func (f FingerprintDigest) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FingerprintDigest) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f FingerprintDigest) String() string { _ = "STUB: not implemented"; return "" }
