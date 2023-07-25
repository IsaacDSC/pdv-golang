package main

import (
	"crypto/subtle"
)

// CompareSecurely compara duas fatias de bytes de forma segura.
func CompareSecurely(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

func main() {

}
