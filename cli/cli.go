package cli

type CLI interface {
	// Add adds a Int Flag to CLI
	Add(opts ...FlagOption) error

	// Init parses flags from args you MUST Add Flags first
	Init(args []string, opts ...Option) error

	// String returns the name of the current implementation
	String() string
}
