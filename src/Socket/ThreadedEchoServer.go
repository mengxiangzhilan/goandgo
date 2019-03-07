package main

import (
	"fmt"
	"net"
)

func Threaded(){
	service:=":1201"
	tcpaddr,err:=net.ResolveTCPAddr("ip4",service)
	checkError(err)
	listener,err:=net.ListenTCP("tcp",tcpaddr)
	checkError(err)
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			continue
		}
		go handleClient(conn)

	}
}
func handleClient(conn net.Conn){
	defer conn.Close()
	var buf [512]byte
	for{
		n,err:=conn.Read(buf[0:])
		if err!=nil{
			return
		}
		fmt.Println(string(buf[0:]))
		_,err2:=conn.Write(buf[0:n])
		if err2 !=nil{
			return
		}
	}
}
