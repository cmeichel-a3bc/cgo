package main

/*

#cgo CFLAGS: -I../../include
#cgo LDFLAGS: -L../../shared -lfoo -lm -Wl,-rpath=./shared
#include "foo.h"

*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Printf("Using shared library\n\n")

	fmt.Printf("%f\n", C.squareRoot(10.0))
}
