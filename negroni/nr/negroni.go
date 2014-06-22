package nr

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/meatballhat/go-newrelic/config"
	"github.com/meatballhat/go-newrelic/control"
)

// Configure creates a default newrelic client and adds middleware and such
func Configure(n *negroni.Negroni) error {
	cfg, err := config.LoadEnv()
	if err != nil {
		return err
	}

	ctl := control.NewAgentController(cfg)

	n.Use(NewRelicMiddleware(ctl, cfg))
	return nil
}

// NewRelicMiddleware builds a negroni HandlerFunc for use in middleware, eh
func NewRelicMiddleware(ctl control.Controller, c *config.Config) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		log.Printf("whatever: %#v\n", ctl)
	}
}
