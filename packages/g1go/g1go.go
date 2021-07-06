package g1go

import "runtime"

func Version() string {
	return runtime.Version()
}
