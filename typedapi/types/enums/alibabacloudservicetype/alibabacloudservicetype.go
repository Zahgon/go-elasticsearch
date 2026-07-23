package alibabacloudservicetype

type AlibabaCloudServiceType struct {
	Name string
}

var (
	AlibabacloudAiSearch = AlibabaCloudServiceType{"alibabacloud-ai-search"}
)

func (a AlibabaCloudServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AlibabaCloudServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (a AlibabaCloudServiceType) String() string { _ = "STUB: not implemented"; return "" }
