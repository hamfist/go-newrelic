package agent

import (
	"github.com/Sirupsen/logrus"
	"github.com/meatballhat/go-newrelic/config"
)

// Options is used to initialize an agent.
type Options struct {
	Env string
	Log Logger
}

// NewOptionsFromConfig extracts bits from a config to make an agent options object
func NewOptionsFromConfig(cfg *config.Config) *Options {
	log := logrus.New()
	opts := &Options{
		Env: cfg.Env,
		Log: logrus.New(),
	}

	if cfg.LogFormat == "json" {
		log.Formatter = &logrus.JSONFormatter{}
	}

	opts.Log = Logger(log)
	return opts
}
