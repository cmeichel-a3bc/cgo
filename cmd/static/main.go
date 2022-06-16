package main

/*

#cgo CFLAGS: -I../../include
#cgo LDFLAGS: ${SRCDIR}/../../static/libfoo.a -lm
#include "foo.h"

*/
import "C"

import "fmt"

func main() {
	fmt.Printf("Using static library\n\n")

	fmt.Printf("%f\n", C.squareRoot(10.0))
}
