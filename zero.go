package padding

func RightPadZero(s []byte, c int) []byte {
	return RightPad(s, []byte{0}, c)
}

func ZeroPad(data []byte, c int) []byte {
	return RightPadZero(data, c)
}
