package types

type RequestCounts struct {
	GetBlob *int64 `json:"GetBlob,omitempty"`

	GetBlobProperties *int64 `json:"GetBlobProperties,omitempty"`

	GetObject *int64 `json:"GetObject,omitempty"`

	InsertObject *int64 `json:"InsertObject,omitempty"`

	ListBlobs *int64 `json:"ListBlobs,omitempty"`

	ListObjects *int64 `json:"ListObjects,omitempty"`

	PutBlob *int64 `json:"PutBlob,omitempty"`

	PutBlock *int64 `json:"PutBlock,omitempty"`

	PutBlockList *int64 `json:"PutBlockList,omitempty"`

	PutMultipartObject *int64 `json:"PutMultipartObject,omitempty"`

	PutObject *int64 `json:"PutObject,omitempty"`
}

func (s *RequestCounts) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRequestCounts() *RequestCounts { _ = "STUB: not implemented"; return nil }
