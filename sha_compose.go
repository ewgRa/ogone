package ogone

import (
	"strings"
	"sort"
	"crypto/sha1"
	"encoding/hex"
)

func ShaCompose(params map[string]string, secret string) (string) {
	keys := make([]string, len(params))
	values := make(map[string]string, len(params))

	keyCounter := 0

	for k, v := range params {
		key := strings.ToUpper(k)
		keys[keyCounter] = key
		values[key] = v
		keyCounter++
	}

	sort.Strings(keys)

	sign := ""

	for _, key := range keys {
		sign += key + "=" + values[key] + secret
	}

	sum := sha1.Sum([]byte(sign))

	return strings.ToUpper(hex.EncodeToString(sum[:]))
}