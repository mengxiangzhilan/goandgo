package main

import (
	"fmt"
	"net"
	"os"
)

func Keudp(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage: %s host:post",os.Args[0])
		os.Exit(1)
	}
	service:=os.Args[1]
	udpaddr,err:=net.ResolveTCPAddr("udp4",service)
	checkError(err)
	conn,err:=net.DialTCP("udp",nil,udpaddr)
	checkError(err)
	_,err=conn.Write([]byte("anything"))
	checkError(err)
	var buf[512]byte
	n,err:=conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

