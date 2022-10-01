package cli

type Config struct {
	Name        string `json:"name" yaml:"Name"`
	Version     string `json:"version" yaml:"Version"`
	Description string `json:"description" yaml:"Description"`
	Usage       string `json:"usage" yaml:"Usage"`
	NoFlags     bool   `json:"no_flags" yaml:"NoFlags"`
	ArgPrefix   string `json:"arg_prefix" yaml:"ArgPrefix"`
	Plugin      string `json:"plugin" yaml:"Plugin"`
	ConfigFile  string `json:"config_file" yaml:"ConfigFile"`
	Flags       []Flag `json:"-" yaml:"-"`
}

func NewConfig() *Config {
	return &Config{
		NoFlags:    false,
		ArgPrefix:  "",
		Plugin:     "urfave",
		ConfigFile: "",
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

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
