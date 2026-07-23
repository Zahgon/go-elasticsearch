package types

type NodesCredentials struct {
	FileTokens map[string]NodesCredentialsFileToken `json:"file_tokens"`

	NodeStats NodeStatistics `json:"_nodes"`
}

func NewNodesCredentials() *NodesCredentials { _ = "STUB: not implemented"; return nil }
