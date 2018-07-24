package webos

import "math/rand"

func requestID() string {
	rs := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = rs[rand.Intn(len(rs))]
	}
	return string(b)
}
