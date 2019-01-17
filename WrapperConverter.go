package imgui

// #include "imguiWrapperTypes.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func castBool(value bool) (cast C.IggBool) {
	if value {
		cast = 1
	}
	return
}

func wrapBool(goValue *bool) (wrapped *C.IggBool, finisher func()) {
	if goValue != nil {
		var cValue C.IggBool
		if *goValue {
			cValue = 1
		}
		wrapped = &cValue
		finisher = func() {
			*goValue = cValue != 0
		}
	} else {
		finisher = func() {}
	}
	return
}

func wrapInt32(goValue *int32) (wrapped *C.int, finisher func()) {
	if goValue != nil {
		cValue := C.int(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = int32(cValue)
		}
	} else {
		finisher = func() {}
	}
	return
}

func wrapFloat(goValue *float32) (wrapped *C.float, finisher func()) {
	if goValue != nil {
		cValue := C.float(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = float32(cValue)
		}
	} else {
		finisher = func() {}
	}
	return
}

func wrapString(value string) (wrapped *C.char, finisher func()) {
	wrapped = C.CString(value)
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}

func wrapStringPointer(value *string) (wrapped *C.char, finisher func()) {

	if value != nil {
		//fmt.Printf("wrapStringPointer: %v - %v\n", value, *value)
		wrapped = C.CString(*value)
		//fmt.Printf("C wrapped: %v\n", wrapped)
		finisher = func() {
			fmt.Printf("gonna free this shit: %v - the value is: %v\n", unsafe.Pointer(wrapped), *value)
			*value = C.GoString(wrapped)
			C.free(unsafe.Pointer(wrapped))
		} // nolint: gas
	} else {
		finisher = func() {}
	}
	return
}
