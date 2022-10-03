package codec

import (
	"io"

	"github.com/go-micro/microwire/v5/util/generic"
)

var Plugins = generic.NewContainer(func(rwc io.ReadWriteCloser) Codec { return nil })
