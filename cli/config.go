package cli

type Config struct {
	Cli CliConfig `json:"cli" yaml:"Cli"`
}

type CliConfig struct {
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
		Cli: CliConfig{
			NoFlags:    false,
			ArgPrefix:  "",
			Plugin:     "",
			ConfigFile: "",
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Cli.NoFlags != def.Cli.NoFlags {
		d.Cli.NoFlags = src.Cli.NoFlags
	}
	if src.Cli.ArgPrefix != def.Cli.ArgPrefix {
		d.Cli.ArgPrefix = src.Cli.ArgPrefix
	}
	if src.Cli.Plugin != def.Cli.Plugin {
		d.Cli.Plugin = src.Cli.Plugin
	}
	if src.Cli.ConfigFile != def.Cli.ConfigFile {
		d.Cli.ConfigFile = src.Cli.ConfigFile
	}

	return nil
}
