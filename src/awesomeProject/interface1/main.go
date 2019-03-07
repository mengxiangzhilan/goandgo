package main

import (
	"fmt"
	"awesomeProject/interface1/Reativer"
	"time"
)

type Lianxi interface {
	Get(url string)string
}
type Poster interface {
	Post(url string,form map[string]string)string
}

func s(i Lianxi)string {
	return i.Get("https://user.qzone.qq.com/1481643256")
}
func post(poster Poster){
	poster.Post("http://www.bilibili.com", map[string]string{
		"name":"ccmouse",
		"course":"golang",

	})
}

type RetriverPoster interface {
	Lianxi
	Poster
}

const url  ="http://www.bilibili.com"
func session(s RetriverPoster)string{
	s.Post(url, map[string]string{
		"contents":"another zhuhaia",
	})
	return s.Get(url)
}
func main() {
	var i Lianxi
	retriever :=Reativer.Lianxi{"假的url"}
fmt.Println("Try a session")
	fmt.Println(session(retriever))
	//i=Reativer.Lianxi{"假的url地址"}
	i=&Reativer.Lianxi{//可以是值也可以是指针
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n",i,i)
	switch  v:=i.(type) {
	case *Reativer.Lianxi:
		fmt.Println("UserAgentt",v.UserAgent)

	}
	//Type assertion
	realRetiever:=i.(*Reativer.Lianxi)
	fmt.Println(realRetiever)
	fmt.Println("********************************************8888")
	fmt.Println(s(i))
}
