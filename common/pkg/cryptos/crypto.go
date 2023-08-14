// Package cryptos 加密相关操作
package cryptos

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

// HMACsha256 return base64 format.
func HMACsha256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)

	if _, err := h.Write([]byte(message)); err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// MD5 md5 摘要算法
func MD5(text string) string {
	algorithm := md5.New()
	return stringHasher(algorithm, text)
}

// SHA256 sha256 算法
func SHA256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

// stringHasher 算法，文本
func stringHasher(algorithm hash.Hash, text string) string {
	if _, err := algorithm.Write([]byte(text)); err != nil {
		return ""
	}
	return hex.EncodeToString(algorithm.Sum(nil))
}

// DoubleHashH double sha256
func DoubleHashH(b []byte) [32]byte {
	first := sha256.Sum256(b)
	return sha256.Sum256(first[:])
}
