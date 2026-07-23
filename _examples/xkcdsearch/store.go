package xkcdsearch

import (
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type SearchResults struct {
	Total int    `json:"total"`
	Hits  []*Hit `json:"hits"`
}

type Hit struct {
	Document
	URL        string        `json:"url"`
	Sort       []interface{} `json:"sort"`
	Highlights *struct {
		Title      []string `json:"title"`
		Alt        []string `json:"alt"`
		Transcript []string `json:"transcript"`
	} `json:"highlights,omitempty"`
}

type StoreConfig struct {
	Client    *elasticsearch.TypedClient
	IndexName string
}

type Store struct {
	es        *elasticsearch.TypedClient
	indexName string
}

func NewStore(c StoreConfig) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Store) CreateIndex(mapping string) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Create(item *Document) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Exists(id string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Store) Search(query string, after ...string) (*SearchResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSearchAfter(s string) []types.FieldValue { _ = "STUB: not implemented"; return nil }
