package goversion

import "runtime"

// GetVersion returns go version
//
func GetVersion() string {
	return runtime.Version()
}
