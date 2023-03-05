package logic

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
)

const NumChars = 8

func EncodeWithSalt(url string, salt int) string {
	var withSalt = url
	if salt > 0 {
		withSalt = fmt.Sprintf("%d%s", salt, url)
	}
	return hashToKChars(withSalt, NumChars)
}

func hashToKChars(url string, k int) string {
	// 1. md5 hex  32个字符
	tiny := getMD5String(url)
	// 2. base64  43个字符
	tiny = base64.URLEncoding.EncodeToString([]byte(tiny))
	// 3. random get k chars, ignore last char "="
	start := rand.Intn(len(tiny) - k - 1)
	return tiny[start : start+k]
}

func getMD5String(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
