package Dataserialisation

import (
	"encoding/asn1"
	"fmt"
	"os"
)
/*
func Marshal(val interface{}) ([]byte, os.Error)
func Unmarshal(val interface{}, b []byte) (rest []byte, err os.Error)
对数据编/解组，需要对interface类型的参数进行更多的类型检查
编组时，我们只需要传递某个类型的变量的值即可，解组它，则需
要一个与被序列化过的数据匹配的确定类型的变量，我们将在后面讨论这部分的细节。除了
有确定类型的变量外，我们同时需要保证那个变量的内存已经被分配，以使被解组后的数据
能有实际被写入的地址。
*/
func asn11(){
	madata,err:=asn1.Marshal(13)
	checkError(err)
	var  n int
	_,err1:=asn1.Unmarshal(madata,&n)
	checkError(err1)
	fmt.Println("After marshal/unmarshal:",n)
}
func checkError(err error){
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}