// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

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
