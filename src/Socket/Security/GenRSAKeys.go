package Security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

//生成RSA 公钥和私钥的程序
func GenRsaKmain() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)
	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())
	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)
	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)
	savePEMKey("private.pem", key)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		//os.Exit(1)
	}
}
func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}
func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		/**公钥基础架构（PKI）是一个公钥集合框架，它连同附加信息，如所有者名称和位置，以及
		它们之间的联系提供了一些审批机制。*/
		Bytes: x509.MarshalPKCS1PrivateKey(key)} //目前使用的pki是基于X.509证书的
	pem.Encode(outFile, privateKey)
	outFile.Close()
}
