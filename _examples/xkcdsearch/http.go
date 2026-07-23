package xkcdsearch

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

var httpLog zerolog.Logger

func init() {
	httpLog = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

}

func (s *Store) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
