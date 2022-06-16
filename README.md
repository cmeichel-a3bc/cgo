# CGO test

To build and run the project:

```
make
./inlineTest
./sharedTest
./staticTest
./stringTest
```


This presents 3 type of CGO implementation:
* `C` code inline into the `GO` code
* `C`code compiled as a shared library. The shared library must be present on the machine to run the executable
* `C`code compiled as a static library. The executable is standalone

It also presents an implementation with strings.

## cmd/inline

1. In cgo comments, just put th c code.
2. `import "C"` must be just after the cgo comments without any space/

## cmd/shared

1. `import "C"` must be just after the cgo comments without any space/
2. `#cgo CFLAGS: -I../../include` specify where to find header `.h`
3. `#cgo LDFLAGS: -L../../shared -lfoo -lm -Wl,-rpath=./shared` specify where to find the lib for compilation, and where to find the lib at runtime
4. `#include "foo.h"` import the library, so that it can be used in the code

## cmd/static

1. `import "C"` must be just after the cgo comments without any space/
2. `#cgo CFLAGS: -I../../include` specify where to find header `.h`
3. `#cgo LDFLAGS: ${SRCDIR}/../../static/libfoo.a -lm` specify the lib path and the use of the math lib. When the cgo directives are parsed, any occurrence of the string ${SRCDIR} will be replaced by the absolute path to the directory containing the source file
4. `#include "foo.h"` import the library, so that it can be used in the code

## cmd/string

1. `import "C"` must be just after the cgo comments without any space/
2. `#cgo CFLAGS: -I../../include` specify where to find header `.h`
3. `#cgo LDFLAGS: ${SRCDIR}/../../static/libfoo.a` specify the lib path. When the cgo directives are parsed, any occurrence of the string ${SRCDIR} will be replaced by the absolute path to the directory containing the source file
4. `#include "foo.h"` import the library, so that it can be used in the code
5. `#include <stdlib.h>` We are going to be using the C function ```*free()```, and this is contained in stdlib, so we'll need the stdlib headers.
