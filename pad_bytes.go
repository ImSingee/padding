package padding

func LeftPad(s, pad []byte, c int) []byte {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	copy(b[padLen:], s)
	repeat(b[:padLen], pad, padLen)

	return b
}

func RightPad(s, pad []byte, c int) []byte {
	padLen := c - len(s)
	if padLen <= 0 {
		return s
	}

	b := make([]byte, c)
	l := copy(b, s)
	repeat(b[l:], pad, padLen)

	return b
}

func repeat(b []byte, pad []byte, padLen int) {
	bp := copy(b[:padLen], pad)
	for bp < padLen {
		copy(b[bp:padLen], b[:bp])
		bp *= 2
	}
}
