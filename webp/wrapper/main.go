//go:generate go run ./generate/

package wrapper

/*
#cgo CFLAGS: -I./../../libwebp/
#cgo CFLAGS: -Wno-pointer-sign -DWEBP_USE_THREAD
#cgo !windows LDFLAGS: -lm

#include <stdlib.h>
*/
import "C"
