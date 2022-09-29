package cli

type CLI interface {
	// Add adds a Flag to CLI
	Add(opts ...FlagOption) error

	// Parse parses flags from args you MUST Add Flags first
	Parse(args []string, opts ...Option) error

	// String returns the name of the current implementation
	String() string
}
