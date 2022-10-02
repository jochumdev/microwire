//go:build go1.18
// +build go1.18

package cli

import "fmt"

const (
	FlagTypeNone   = iota
	FlagTypeString = iota
	FlagTypeInt    = iota
)

type Flag struct {
	Name    string
	EnvVars []string
	Usage   string

	FlagType int

	DefaultString     string
	DestinationString *string

	DefaultInt     int
	DestinationInt *int
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
		result = append(result, Default(f.DefaultString))
	case FlagTypeInt:
		result = append(result, Default(f.DefaultInt))
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

func Default[T any](n T) FlagOption {
	return func(o *Flag) {
		switch any(n).(type) {
		case string:
			o.DefaultString = any(n).(string)
		case int:
			o.DefaultInt = any(n).(int)
		}
	}
}

func Destination[T any](n T) FlagOption {
	return func(o *Flag) {
		switch any(n).(type) {
		case *string:
			o.FlagType = FlagTypeString
			o.DestinationString = any(n).(*string)
		case *int:
			o.FlagType = FlagTypeInt
			o.DestinationInt = any(n).(*int)
		default:
			o.FlagType = FlagTypeNone
		}
	}
}

func NewFlag(opts ...FlagOption) (*Flag, error) {
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

	if options.FlagType == FlagTypeNone {
		return nil, fmt.Errorf("found flag '%s' without an Destination() option", options.Name)
	}

	return options, nil
}
