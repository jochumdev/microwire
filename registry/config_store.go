package registry

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"enabled"`
	Plugin    string   `json:"plugin" yaml:"Plugin"`
	Addresses []string `json:"addresses" yaml:"Addresses"`
}

func DefaultConfigStore() ConfigStore {
	return ConfigStore{
		Enabled:   true,
		Plugin:    "mdns",
		Addresses: []string{},
	}
}
