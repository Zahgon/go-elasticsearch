package main

import (
	_ "github.com/dustin/go-humanize"
	_ "github.com/elastic/go-elasticsearch/v9"
	_ "github.com/elastic/go-elasticsearch/v9/esutil"
	_ "go.elastic.co/apm"
	_ "go.elastic.co/apm/module/apmelasticsearch"
)
