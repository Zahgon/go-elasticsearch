package consumer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/elastic/go-elasticsearch/v9/esutil"
)

type Consumer struct {
	BrokerURL string
	TopicName string

	Indexer esutil.BulkIndexer
	reader  *kafka.Reader

	startTime     time.Time
	totalMessages int64
	totalErrors   int64
	totalBytes    int64
}

func (c *Consumer) Run(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

type Stats struct {
	Duration      time.Duration
	TotalLag      int64
	TotalMessages int64
	TotalErrors   int64
	TotalBytes    int64
	Throughput    float64
}

func (c *Consumer) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }
