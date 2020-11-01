package stacklib

import "runtime"

// FullStack get full stack
func FullStack() []byte {
	buf := make([]byte, 2048)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}

		buf = make([]byte, 2*len(buf))
	}
}
