package Reativer

import (
		"net/http"
	"net/http/httputil"
	"time"
)

type Lianxi struct {
	Contents string
	UserAgent string
	TimeOut time.Duration
}
func (i Lianxi)Post(url string,form map[string]string)string{
	i.Contents= form["contents"]
	return "ok"
}
func (i *Lianxi)Get(url string)string  {
resp,err:=http.Get(url)
if err!=nil{
	panic(err)
}
result,err:=httputil.DumpResponse(resp,true)
resp.Body.Close()
if err!=nil{
	panic(err)
}
return string(result)
}
//func (i Lianxi)Get(url string)string  {
	//return i.Contents
//}