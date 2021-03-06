package main

import (
	"encoding/base64"

	"github.com/lemon-mint/m"
)

func main() {
	orignal := []byte("Hello, World!")
	data := ResultBase64DecodeString(base64.StdEncoding.EncodeToString(orignal))
	data.Match(func(v []byte) {
		println(string(v))
		if string(v) != string(orignal) {
			panic("not equal")
		}
	}, func(err error) {
		panic(err)
	})

}

func ResultBase64DecodeString(s string) m.Result[[]byte] {
	return m.RValue(base64.StdEncoding.DecodeString(s))
}
