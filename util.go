package webos

import "math/rand"

// requestID returns a random 8 character string. Requests and Responses sent to and from
// the TV are linked by this ID.
func requestID() string {
	rs := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = rs[rand.Intn(len(rs))]
	}
	return string(b)
}
