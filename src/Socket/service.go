package main

import (
	"fmt"
	"net"
	"os"
)

func Service(){
	if len(os.Args)!=3{
		fmt.Fprintf(os.Stderr,"Usage: %s network-type\n",os.Args[0])
		os.Exit(1)
	}
	networkType:=os.Args[1]
	service:=os.Args[2]
	port,err:=net.LookupPort(networkType,service)//unix下获取常用端口
	if err!=nil{
		fmt.Println(" error",err.Error())
		os.Exit(2)
	}
	fmt.Println("service port",port)
	os.Exit(0)
}