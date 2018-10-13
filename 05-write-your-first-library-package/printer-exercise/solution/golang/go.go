// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package golang

import (
	"runtime"
)

// Version returns the current Go version
func Version() string {
	return runtime.Version()
}
