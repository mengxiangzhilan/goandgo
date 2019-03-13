package Security

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func Securmain() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)
	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())
	var publicKey rsa.PublicKey
	loadKey("public.key", &publicKey)
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

}
func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile) //通过gob序列化保存正式或取回证书
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()

}
