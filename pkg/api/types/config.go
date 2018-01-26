package types

type Config struct {
	Routes       []Route
	Upstreams    []Upstream
	VirtualHosts []VirtualHost
}

// spec is a generic blob object
// that will be passed to route plugins/upstream modules
type Spec map[string][]byte