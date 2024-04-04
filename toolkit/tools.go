package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module. Any variableof this type
// will access all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a string of random characters of length n, using
// randomStringSource as the source for the string
func (t *Tools) RandomString(n int) string {
	s := make([]rune, n)
	r := []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
