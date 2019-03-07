package main

import (
	"net"
	"time"
)

func Uservice(){
	service:=":1200"
	udpaddr,err:=net.ResolveTCPAddr("up4",service)
	checkError(err)
	listener,err:=net.ListenTCP("udp",udpaddr)
	checkError(err)
	for{
		handleClient1(conn)
		}


}
func handleClient1(conn *net.UDPConn){
	var buf[512]byte
	_,addr,err:=conn.ReadFromUDP(buf[0:])
	if err!=nil{
		return
	}
	daytime:=time.Now().String()
	conn.WriteToUDP([]byte(daytime),addr)
}