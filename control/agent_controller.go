package control

import (
	"github.com/meatballhat/go-newrelic/agent"
	"github.com/meatballhat/go-newrelic/config"
)

// AgentController implements Controller and has an agent.  Crazy!
type AgentController struct {
	Agent *agent.Agent
}

// NewAgentController creates an AgentController. Bonkers!
func NewAgentController(cfg *config.Config) Controller {
	return &AgentController{
		Agent: agent.NewAgent(agent.NewOptionsFromConfig(cfg)),
	}
}
