package agent

// Logger is the logging interface used inside an agent.  Neato!
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}
