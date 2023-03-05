package gosms

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

// Md5String md5加密
func Md5String(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1String sha1加密
func Sha1String(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
func Sha256String(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

// URLEncode urlencode()
func URLEncode(text string) string {
	return url.QueryEscape(text)
}

// Base64Encode  base64
func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
