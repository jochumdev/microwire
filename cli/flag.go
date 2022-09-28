//go:build go1.18
// +build go1.18

package cli

const (
	FlagTypeNone   int = 0
	FlagTypeString int = 1
	FlagTypeInt    int = 2
)

type Flag struct {
	Name    string
	EnvVars []string
	Usage   string

	FlagType      int
	DefaultString string
	DefaultInt    int
}

type FlagOption func(*Flag)

func (f *Flag) AsOptions() []FlagOption {
	result := []FlagOption{
		Name(f.Name),
		EnvVars(f.EnvVars...),
		Usage(f.Usage),
	}

	switch f.FlagType {
	case FlagTypeString:
		result = append(result, DefaultValue(f.DefaultString))
	case FlagTypeInt:
		result = append(result, DefaultValue(f.DefaultInt))
	}

	return result
}

func Name(n string) FlagOption {
	return func(o *Flag) {
		o.Name = n
	}
}

func EnvVars(n ...string) FlagOption {
	return func(o *Flag) {
		o.EnvVars = n
	}
}

func Usage(n string) FlagOption {
	return func(o *Flag) {
		o.Usage = n
	}
}

func DefaultValue[T any](n T) FlagOption {
	return func(o *Flag) {
		switch any(n).(type) {
		case string:
			o.FlagType = FlagTypeString
			o.DefaultString = any(n).(string)
		case int:
			o.FlagType = FlagTypeInt
			o.DefaultInt = any(n).(int)
		}
	}
}

func NewFlag(opts ...FlagOption) *Flag {
	options := &Flag{
		Name:          "",
		EnvVars:       []string{},
		Usage:         "",
		FlagType:      FlagTypeNone,
		DefaultString: "",
		DefaultInt:    0,
	}

	for _, o := range opts {
		o(options)
	}

	return options
}
