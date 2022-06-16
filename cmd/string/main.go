package main

/*

#cgo CFLAGS: -I../../include
#cgo LDFLAGS: ${SRCDIR}/../../static/libfoo.a
#include "foo.h"
#include <stdlib.h>

*/
import "C"

import (
	"fmt"
	"unsafe"
)

func hasLetter(str string, letter byte) bool {
	mystr := C.CString(str)
	defer C.free(unsafe.Pointer(mystr))

	return C.hasLetter(mystr, C.char(letter)) != 0
}

func main() {
	fmt.Printf("Strings\n\n")

	for letter, str := range map[byte]string{
		'l': "Hello world",
		'y': "Hello world",
	} {
		if hasLetter(str, letter) {
			fmt.Printf("[%c] found in [%s]\n", letter, str)
		} else {
			fmt.Printf("[%c] not found in [%s]\n", letter, str)
		}
	}

}
