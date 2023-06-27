package rand

import (
	"crypto/rand"
	"encoding/base64"
)


const RememberTokenBytes = 32


// func Bytes() сгенерирует n случайных байтов или
// вернуть ошибку, если она была. Используем пакет
// crypto/rand, его можно безопасно использовать для
// того что бы запомнить токен.

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// func String() будет генерировать байтовый фрагмент размером nBytes, а затем
// вернёт строку, которая является версией URL в кодировке base64

func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}


// func RememberToken() запомнить токены заданного размера в байтах
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}