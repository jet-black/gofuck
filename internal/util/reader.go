package util

import (
	"io"
)

type MyBufferedReader struct {
	Buf        []rune
	Underlying io.RuneReader
	Pos        int
}

func (r *MyBufferedReader) ReadToken() (rune, error) {
	if r.Pos < len(r.Buf) {
		result := r.Buf[r.Pos]
		r.Pos++
		return result, nil
	}
	token, _, err := r.Underlying.ReadRune()
	if err != nil {
		return 0, err
	}
	r.Pos++
	r.Buf = append(r.Buf, token)
	return token, nil
}
