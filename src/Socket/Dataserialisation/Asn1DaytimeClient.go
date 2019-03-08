package Dataserialisation

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func asc() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:post", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)
	fmt.Println("after marshal/unmarshal:", newtime.String())
	os.Exit(0)
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
