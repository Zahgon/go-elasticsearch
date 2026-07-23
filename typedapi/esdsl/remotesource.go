package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _remoteSource struct {
	v *types.RemoteSource
}

func NewRemoteSource() *_remoteSource { _ = "STUB: not implemented"; return nil }

func (s *_remoteSource) ApiKey(apikey string) *_remoteSource { _ = "STUB: not implemented"; return nil }

func (s *_remoteSource) ConnectTimeout(duration types.DurationVariant) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) Headers(headers map[string]string) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) AddHeader(key string, value string) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) Host(host string) *_remoteSource { _ = "STUB: not implemented"; return nil }

func (s *_remoteSource) Password(password string) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) SocketTimeout(duration types.DurationVariant) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) Username(username string) *_remoteSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_remoteSource) RemoteSourceCaster() *types.RemoteSource {
	_ = "STUB: not implemented"
	return nil
}
