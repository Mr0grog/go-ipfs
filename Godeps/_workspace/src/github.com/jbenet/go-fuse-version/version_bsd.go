// +build dragonfly freebsd netbsd openbsd

package fuseversion

import (
	"runtime"
)

func getLocalFuseSystems() (*Systems, error) {
	return nil, fmt.Sprintf(notImplYet, runtime.GOARCH())
}