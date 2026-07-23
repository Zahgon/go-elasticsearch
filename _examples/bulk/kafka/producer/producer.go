package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	sides    = []string{"BUY", "SELL"}
	symbols  = []string{"KBCU", "KBCU", "KBCU", "KJPR", "KJPR", "KSJD", "KXCV", "WRHV", "WTJB", "WMLU"}
	accounts = []string{"ABC123", "ABC123", "ABC123", "LMN456", "LMN456", "STU789"}
)

func init() {
	kafka.DefaultClientID = "go-elasticsearch-kafka-demo"
}

type Producer struct {
	BrokerURL   string
	TopicName   string
	TopicParts  int
	MessageRate int

	writer *kafka.Writer

	startTime     time.Time
	totalMessages int64
	totalErrors   int64
	totalBytes    int64
}

func (p *Producer) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *Producer) CreateTopic(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *Producer) generateMessage(t time.Time) []byte { _ = "STUB: not implemented"; return nil }

type Stats struct {
	Duration      time.Duration
	TotalMessages int64
	TotalErrors   int64
	TotalBytes    int64
	Throughput    float64
}

func (p *Producer) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }
