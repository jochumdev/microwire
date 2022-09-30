package cli

type ConfigStore struct {
	NoFlags    bool   `json:"no_flags" yaml:"NoFlags"`
	ArgPrefix  string `json:"arg_prefix" yaml:"ArgPrefix"`
	Plugin     string `json:"plugin" yaml:"Plugin"`
	ConfigFile string `json:"config_file" yaml:"ConfigFile"`
}

func NewConfigStore() ConfigStore {
	return ConfigStore{
		NoFlags:    false,
		ArgPrefix:  "",
		Plugin:     "urfave",
		ConfigFile: "",
	}
}

func (d *ConfigStore) Merge(src *ConfigStore) error {
	def := NewConfigStore()

	if src.NoFlags != def.NoFlags {
		d.NoFlags = src.NoFlags
	}
	if src.ArgPrefix != def.ArgPrefix {
		d.ArgPrefix = src.ArgPrefix
	}
	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
	}
	if src.ConfigFile != def.ConfigFile {
		d.ConfigFile = src.ConfigFile
	}

	return nil
}
