package padding

// Pkcs5Pad implements PKCS#5 padding
func Pkcs5Pad(data []byte) []byte {
	return Pkcs7Pad(data, 8)
}

// Pkcs5Unpad unpads data that was padded using Pkcs5Pad
// If data is invalid, nil is returned
func Pkcs5Unpad(data []byte) []byte {
	return Pkcs7Unpad(data, 8)
}
