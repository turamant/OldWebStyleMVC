package hash


import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// func NewHMAC() создает и возвращает новый объект HMAC.
func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
	hmac: h,
	}
	}

// HMAC — это оболочка для создания пакета crypto/hmac.
// его немного проще использовать в нашем коде.
// hash - стандартная библиотека
type HMAC struct {
	hmac hash.Hash
}

// func Hash() будет хешировать предоставленную входную строку, используя HMAC с
// секретный ключ, указанный при создании объекта HMAC
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
	}