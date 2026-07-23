package types

type CertificateInformation struct {
	Alias *string `json:"alias,omitempty"`

	Expiry DateTime `json:"expiry"`

	Format string `json:"format"`

	HasPrivateKey bool `json:"has_private_key"`

	Issuer *string `json:"issuer,omitempty"`

	Path string `json:"path"`

	SerialNumber string `json:"serial_number"`

	SubjectDn string `json:"subject_dn"`
}

func (s *CertificateInformation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCertificateInformation() *CertificateInformation { _ = "STUB: not implemented"; return nil }
