package main

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
)

var p = fmt.Println

func main() {
	s := "https://user.2323/path?q=123#f"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	p(s)
	p(bs)
	fmt.Printf("%x\n", bs)

	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}
