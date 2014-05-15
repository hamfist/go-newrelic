package config

// Configger defines a simple gettable/settable type for use in configuring everything
type Configger interface {
	Get(string) string
	Set(string, string)
}
