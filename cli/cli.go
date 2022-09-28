package cli

type CLI interface {
	// Add adds a Int Flag to CLI
	Add(opts ...FlagOption) error

	// Init parses flags from args you MUST Add Flags first
	Init(args []string, opts ...Option) error

	// StringValue returns the flag "name" as String
	StringValue(name string) string

	// IntValue returns the flag "name" as Int
	IntValue(name string) int

	// String returns the name of the current implementation
	String() string
}
