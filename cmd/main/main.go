package main

import (
	"encoding/base64"

	"github.com/lemon-mint/m"
)

func main() {
	orignal := []byte("Hello, World!")
	data := m.Unwrap(base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(orignal)))
	if string(data) != string(orignal) {
		panic("failed")
	}
}
