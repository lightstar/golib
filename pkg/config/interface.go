package config

// Interface of configuration services that must be used by any typical external applications.
type Interface interface {
	Get(interface{}) error
	GetByKey(string, interface{}) error
}
