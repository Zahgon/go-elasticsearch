package esutil

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

var ErrIndexerClosed = errors.New("bulk indexer is closed")

type BulkIndexer interface {
	Add(context.Context, BulkIndexerItem) error

	Close(context.Context) error

	Flush(context.Context) error

	Stats() BulkIndexerStats
}

type BulkIndexerConfig struct {
	NumWorkers          int
	FlushBytes          int
	FlushInterval       time.Duration
	FlushJitter         time.Duration
	QueueSizeMultiplier int

	Client      esapi.Transport
	Decoder     BulkResponseJSONDecoder
	DebugLogger BulkIndexerDebugLogger

	OnError      func(context.Context, error)
	OnFlushStart func(context.Context) context.Context
	OnFlushEnd   func(context.Context)

	Index               string
	ErrorTrace          bool
	FilterPath          []string
	Header              http.Header
	Human               bool
	Pipeline            string
	Pretty              bool
	Refresh             string
	Routing             string
	RequireAlias        bool
	Source              []string
	SourceExcludes      []string
	SourceIncludes      []string
	Timeout             time.Duration
	WaitForActiveShards string
}

type BulkIndexerStats struct {
	NumAdded     uint64
	NumFlushed   uint64
	NumFailed    uint64
	NumIndexed   uint64
	NumCreated   uint64
	NumUpdated   uint64
	NumDeleted   uint64
	NumRequests  uint64
	FlushedBytes uint64
	FlushedMs    uint64
}

type BulkIndexerItem struct {
	Index           string
	Action          string
	DocumentID      string
	Routing         string
	RequireAlias    bool
	Version         *int64
	VersionType     string
	Body            io.ReadSeeker
	RetryOnConflict *int
	IfSeqNo         *int64
	IfPrimaryTerm   *int64
	ctx             context.Context
	meta            bytes.Buffer
	payloadLength   int
	flushDone       chan struct{}

	OnSuccess func(context.Context, BulkIndexerItem, BulkIndexerResponseItem)
	OnFailure func(context.Context, BulkIndexerItem, BulkIndexerResponseItem, error)
}

//nolint:gocyclo
func (item *BulkIndexerItem) marshallMeta() { _ = "STUB: not implemented"; return }

func (item *BulkIndexerItem) computeLength() error { _ = "STUB: not implemented"; return nil }

type BulkIndexerResponse struct {
	Took      int                                  `json:"took"`
	HasErrors bool                                 `json:"errors"`
	Items     []map[string]BulkIndexerResponseItem `json:"items,omitempty"`
}

type BulkIndexerResponseItem struct {
	Index        string `json:"_index"`
	DocumentID   string `json:"_id"`
	Version      int64  `json:"_version"`
	Result       string `json:"result"`
	Status       int    `json:"status"`
	SeqNo        int64  `json:"_seq_no"`
	PrimTerm     int64  `json:"_primary_term"`
	FailureStore string `json:"failure_store,omitempty"`

	Shards struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`

	Error struct {
		Type   string `json:"type"`
		Reason string `json:"reason"`
		Cause  struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"caused_by"`
	} `json:"error,omitempty"`
}

type BulkResponseJSONDecoder interface {
	UnmarshalFromReader(io.Reader, *BulkIndexerResponse) error
}

type BulkIndexerDebugLogger interface {
	Printf(string, ...interface{})
}

type bulkIndexer struct {
	wg      sync.WaitGroup
	workers []*worker
	stats   *bulkIndexerStats

	config      BulkIndexerConfig
	ownedClient *elasticsearch.BaseClient
	addCounter  atomic.Uint64
	flushMu     sync.Mutex
	closed      atomic.Bool
}

type bulkIndexerStats struct {
	numAdded     uint64
	numFlushed   uint64
	numFailed    uint64
	numIndexed   uint64
	numCreated   uint64
	numUpdated   uint64
	numDeleted   uint64
	numRequests  uint64
	flushedBytes uint64
	flushedMs    uint64
}

func NewBulkIndexer(cfg BulkIndexerConfig) (BulkIndexer, error) {
	_ = "STUB: not implemented"
	return *new(BulkIndexer), nil
}

func (bi *bulkIndexer) Add(ctx context.Context, item BulkIndexerItem) error {
	_ = "STUB: not implemented"
	return nil
}

func (bi *bulkIndexer) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (bi *bulkIndexer) Flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (bi *bulkIndexer) Stats() BulkIndexerStats {
	_ = "STUB: not implemented"
	return *new(BulkIndexerStats)
}

func (bi *bulkIndexer) init() {
	for i := 1; i <= bi.config.NumWorkers; i++ {
		bi.wg.Add(1)
		w := worker{
			id:     i,
			ch:     make(chan BulkIndexerItem, bi.config.QueueSizeMultiplier),
			bi:     bi,
			buf:    bytes.NewBuffer(make([]byte, 0, bi.config.FlushBytes)),
			ticker: time.NewTicker(bi.nextFlushInterval()),
		}
		w.run()
		bi.workers = append(bi.workers, &w)
	}
}

func (bi *bulkIndexer) nextFlushInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

//nolint:gosec // G404: non-crypto randomness is intentional for flush-interval jitter

type worker struct {
	id     int
	ch     chan BulkIndexerItem
	bi     *bulkIndexer
	buf    *bytes.Buffer
	items  []BulkIndexerItem
	ticker *time.Ticker
}

func (w *worker) run() { _ = "STUB: not implemented"; return }

func (w *worker) writeMeta(item *BulkIndexerItem) error { _ = "STUB: not implemented"; return nil }

func (w *worker) writeBody(item *BulkIndexerItem) error { _ = "STUB: not implemented"; return nil }

func (w *worker) flush(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

//nolint:gocyclo
func (w *worker) flushBuffer(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:gosec // elapsed is guaranteed positive

func (w *worker) notifyItemsOnError(err error) { _ = "STUB: not implemented"; return }

func (w *worker) handleError(ctx context.Context, err error) { _ = "STUB: not implemented"; return }

type defaultJSONDecoder struct{}

func (d defaultJSONDecoder) UnmarshalFromReader(r io.Reader, blk *BulkIndexerResponse) error {
	_ = "STUB: not implemented"
	return nil
}
