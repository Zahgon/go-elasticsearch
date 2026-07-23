package types

type NodeReloadResult struct {
	KeystoreDigest *string `json:"keystore_digest,omitempty"`

	KeystoreLastModifiedTime DateTime `json:"keystore_last_modified_time,omitempty"`

	KeystorePath    *string     `json:"keystore_path,omitempty"`
	Name            string      `json:"name"`
	ReloadException *ErrorCause `json:"reload_exception,omitempty"`

	SecureSettingNames []string `json:"secure_setting_names,omitempty"`
}

func (s *NodeReloadResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeReloadResult() *NodeReloadResult { _ = "STUB: not implemented"; return nil }
