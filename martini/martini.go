package martini

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/meatballhat/go-newrelic/config"
)

// NewRelicMiddleware builds a martini handler for use in middleware, eh
func NewRelicMiddleware(c config.Configger) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
	}
}
