package cli

type CLI interface {
	// AddString adds a String Flag to CLI
	AddString(opts ...FlagOption) error

	// AddInt adds a Int Flag to CLI
	AddInt(opts ...FlagOption) error

	// Init parses flags from args you MUST Add Flags first
	Init(args []string, opts ...Option) error

	// String returns the string value of a flag
	String(name string) string

	// Int returns the integer value of a flag
	Int(name string) int
}
