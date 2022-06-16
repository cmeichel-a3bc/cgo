package main

/*

double identity(double i) {
	return i;
}

*/
import "C"

import (
	"fmt"
)

func main() {
	fmt.Printf("Using inline C\n\n")

	fmt.Printf("%f\n", C.identity(10.0))
}
