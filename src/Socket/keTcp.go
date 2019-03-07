package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

)

func Ketcp(){
	if len(os.Args)!=2{
	fmt.Fprintf(os.Stderr,"Usage: %s host:post",os.Args[0])
	os.Exit(1)
	}
	service:=os.Args[1]
	tcpaddr,err:=net.ResolveTCPAddr("tcp4",service)
	checkError(err)
	conn,err:=net.DialTCP("tcp",nil,tcpaddr)
	checkError(err)
	_,err=conn.Write([]byte("HEAD/HTTP/1.0\r\n\r\n"))
	result,err:=ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}