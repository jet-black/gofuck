package interpreter

import (
	"io"
)

type myBufferedReader struct {
	buf        []rune
	underlying io.RuneReader
	pos        int
}

func (r *myBufferedReader) ReadToken() (rune, error) {
	if r.pos < len(r.buf) {
		result := r.buf[r.pos]
		r.pos++
		return result, nil
	}
	token, _, err := r.underlying.ReadRune()
	if err != nil {
		return 0, err
	}
	r.pos++
	r.buf = append(r.buf, token)
	return token, nil
}
