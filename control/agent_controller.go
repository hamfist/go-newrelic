package control

import (
	"github.com/meatballhat/go-newrelic/config"
)

// AgentController implements Controller and has an agent.  Crazy!
type AgentController struct {
}

// NewAgentController creates an AgentController. Bonkers!
func NewAgentController(cfg *config.Config) Controller {
	return nil
}
