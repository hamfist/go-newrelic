package nr

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/meatballhat/go-newrelic/config"
	"github.com/meatballhat/go-newrelic/control"
)

// Configure creates a default newrelic client and adds middleware and such
func Configure(m *martini.ClassicMartini) error {
	cfg, err := config.LoadEnv()
	if err != nil {
		return err
	}

	ctl := control.NewAgentController(cfg)

	m.MapTo(ctl, (*control.Controller)(nil))
	m.Use(NewRelicMiddleware(cfg))
	return nil
}

// NewRelicMiddleware builds a martini handler for use in middleware, eh
func NewRelicMiddleware(c *config.Config) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger, ctl control.Controller) {
		log.Printf("ctl=%v\n", ctl)
	}
}
