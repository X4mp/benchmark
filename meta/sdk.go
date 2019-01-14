package meta

import (
	"github.com/xmnservices/xmnsuite/blockchains/core/meta"
)

// SDKFunc represents the Meta SDK func
var SDKFunc = struct {
	Create func() meta.Meta
}{
	Create: func() meta.Meta {
		out, outErr := createMeta()
		if outErr != nil {
			panic(outErr)
		}

		return out
	},
}
