package padding

import (
	"bytes"
)

// Pkcs7Pad implements PKCS#7 (RFC 5652) padding
// blockSize must be less than 256, (or else the behavior is undefined)
func Pkcs7Pad(data []byte, blockSize int) []byte {
	if blockSize <= 0 || blockSize >= 256 {
		return nil
	}

	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// Pkcs7Unpad unpads data that was padded using Pkcs7Pad
// If data is invalid, nil is returned
func Pkcs7Unpad(data []byte, blockSize int) []byte {
	length := len(data)

	if length == 0 || length%blockSize != 0 {
		return nil
	}

	padLen := int(data[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	if padLen > blockSize || padLen == 0 || !bytes.HasSuffix(data, ref) {
		return nil
	}
	return data[:length-padLen]
}
