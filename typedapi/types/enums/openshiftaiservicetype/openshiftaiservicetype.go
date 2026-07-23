package openshiftaiservicetype

type OpenShiftAiServiceType struct {
	Name string
}

var (
	Openshiftai = OpenShiftAiServiceType{"openshift_ai"}
)

func (o OpenShiftAiServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *OpenShiftAiServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (o OpenShiftAiServiceType) String() string { _ = "STUB: not implemented"; return "" }
