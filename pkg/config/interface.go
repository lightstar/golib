package config

// Interface of configuration service that should be used by any typical external applications.
type Interface interface {
	Get(interface{}) error
	GetByKey(string, interface{}) error
}
