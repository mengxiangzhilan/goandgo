package main

import (
	"net"
	"time"
)

func Daytimeserver(){
	service:=":1200"
tcpaddr,err:=net.ResolveTCPAddr("ip4",service)
checkError(err)
listener,err:=net.ListenTCP("tcp",tcpaddr)
checkError(err)
for{
	conn,err:=listener.Accept()
	if err!=nil{
		continue
	}
	daytime :=time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}
}
