package cutil

import (
	"unsafe"
)

// Convert C-style `const char*` to Go string without cgo.
func ToGoString(cstr *byte) string {
	if cstr == nil {
		return ""
	}

	// Find the length by scanning for the null terminator.
	length := 0
	for ptr := cstr; *ptr != 0; ptr = (*byte)(unsafe.Add(unsafe.Pointer(ptr), 1)) {
		length++
	}

	// Create a Go byte slice from the C string
	byteSlice := unsafe.Slice(cstr, length)

	return string(byteSlice)
}

// Converts Go string to a *const char (C null-terminated string).
func ToCString(str string) *byte {
	if str == "" {
		zeroByte := [1]byte{0}
		return (*byte)(unsafe.Pointer(&zeroByte[0])) // Return pointer to single null byte
	}

	// Allocate a byte slice with an extra null terminator.
	cstr := make([]byte, len(str)+1)
	copy(cstr, str)
	cstr[len(str)] = 0 // Null terminator

	// Return a pointer to the beginning of the byte slice.
	return &cstr[0]
}
