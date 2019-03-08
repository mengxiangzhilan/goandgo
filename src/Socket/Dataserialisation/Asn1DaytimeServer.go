package Dataserialisation

import (
	"encoding/asn1"
	"net"
	"time"
)

func as1() {
	service := ":1200"
	tcpaddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpaddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}
