package clusterstatus

import (
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
)

var ES *elasticsearch.Client

func init() {
	log.SetFlags(0)

	var err error
	ES, err = elasticsearch.New()
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func Health(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
