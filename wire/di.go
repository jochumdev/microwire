package wire

import (
	"github.com/go-micro/microwire/cli"
)

// InitializedCli is a marker interface which tells use that the CLI has been loaded/initialized.
// This lives here cause we would have import cycles else
type InitializedCli cli.CLI
