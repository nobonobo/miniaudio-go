package cutil

import "bytes"

func String(cstr []byte) string {
	return string(bytes.TrimRight(cstr, "\x00"))
}
