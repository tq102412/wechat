package encryptor

import "math/rand"

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

//StrRandom 随机数
func StrRandom(n int) []byte {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return b
}
