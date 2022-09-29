package cli

type ConfigStore struct {
	NoFlags    bool   `json:"no_flags" yaml:"NoFlags"`
	ArgPrefix  string `json:"arg_prefix" yaml:"ArgPrefix"`
	Plugin     string `json:"plugin" yaml:"Plugin"`
	ConfigFile string `json:"config_file" yaml:"ConfigFile"`
}

func DefaultConfigStore() ConfigStore {
	return ConfigStore{
		NoFlags:    false,
		ArgPrefix:  "",
		Plugin:     "urfave",
		ConfigFile: "",
	}
}
