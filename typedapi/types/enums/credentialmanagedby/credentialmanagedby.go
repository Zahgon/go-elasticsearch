package credentialmanagedby

type CredentialManagedBy struct {
	Name string
}

var (
	Cloud = CredentialManagedBy{"cloud"}

	Elasticsearch = CredentialManagedBy{"elasticsearch"}
)

func (c CredentialManagedBy) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CredentialManagedBy) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CredentialManagedBy) String() string { _ = "STUB: not implemented"; return "" }
