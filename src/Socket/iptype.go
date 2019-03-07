package main

import (
	"fmt"
	"net"
	"os"
)

func Iptype() {
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"Usage: %s ip-addr\n",os.Args[0])
		os.Exit(1)
	}
	dotAddr:=os.Args[1]
	addr:=net.ParseIP(dotAddr) //net包用于网路编程，ip类型被定义成一个字节数组
								//ParseIP获取，分隔的ipv4和：分隔的ipv6地址
	addr1,err:=net.ResolveIPAddr("ip",dotAddr)//通过ip主机名执行DNS查找
	addrs,err2:=net.LookupHost(dotAddr)

	if err!=nil{
		fmt.Println("Resolution error",err.Error())
		os.Exit(1)
	}
	if err!=nil{
		fmt.Println(" error",err2.Error())
		os.Exit(2)
	}
	for _,s:=range addrs{
		fmt.Println(s)
	}
	fmt.Println("Resolved address is",addr1.String())
	if addr==nil{
		fmt.Println("invalid address")
		os.Exit(1)
	}
	mask:=addr.DefaultMask() //返回默认ip掩码
	network:=addr.Mask(mask) //通过掩码找到该ip的网络
	one,bits:=mask.Size()
	fmt.Println("address is",addr.String(),
		"Default maks length is",bits,
		"Leading ones count is",one,
		"Mask is (hex)",mask.String(),
		"Network is ",network.String())

	os.Exit(0)
}
