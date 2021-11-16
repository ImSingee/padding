package padding

func LeftPadString(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	copy(b[padLen:], s)
	repeatString(b[:padLen], pad, padLen)

	return bytesToString(b)
}

func RightPadString(s, pad string, c int) string {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	l := copy(b, s)
	repeatString(b[l:], pad, padLen)

	return bytesToString(b)
}

func repeatString(b []byte, pad string, padLen int) {
	bp := copy(b[:padLen], pad)
	for bp < padLen {
		copy(b[bp:padLen], b[:bp])
		bp *= 2
	}
}
