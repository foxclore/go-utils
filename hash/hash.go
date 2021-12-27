package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Bytes(income []byte) string {
	hash := sha256.Sum256(income)
	return hex.EncodeToString(hash[:])
}

func String(income string) string {
	return Bytes([]byte(income))
}

func Hash(incomes ...interface{}) string {
	a, err := json.Marshal(incomes)
	if err != nil {
		return ""
	}
	return Bytes(a)
}
