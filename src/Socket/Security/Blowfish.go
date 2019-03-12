package Security

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/blowfish"
)

func Blowmain() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	src := []byte("hello\n\n\n")
	var enc [512]byte
	cipher.Encrypt(enc[0:], src)
	var decrype [8]byte
	cipher.Decrypt(decrype[0:], enc[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrype[0:8])
	fmt.Println(string(result.Bytes()))
}
