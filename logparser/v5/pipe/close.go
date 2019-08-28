// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package pipe

import (
	"io"
)

// readClose the reader if it's a io.Closer.
func readClose(r io.Reader) {
	if rc, ok := r.(io.Closer); ok {
		rc.Close()
	}
}
