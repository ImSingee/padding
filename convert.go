package padding

import "unsafe"

func bytesToString(s []byte) string {
	if s == nil {
		return ""
	}

	return *(*string)(unsafe.Pointer(&s))
}
