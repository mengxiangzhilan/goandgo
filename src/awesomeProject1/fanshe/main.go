package main

import (
	"strconv"
	"reflect"
	"fmt"
	"io"
	"os"
)

//为了处理一类不能满足普通接口的类型的值（有可能该类型还没有）
func main() {
	//反射由reflect支持
	t:= reflect.TypeOf(3)//reflect.TypeOf 接受任意的 interface{} 类型
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer=os.Stdout
	fmt.Println(reflect.TypeOf(w))// "*os.File"
	fmt.Printf("%T\n",3)//fmt.Printf 提供了一个简短的 %T 标志参数, 内部使用 reflect.TypeOf 的结果输出:
v:=reflect.ValueOf(3)
fmt.Println(v)
fmt.Printf("%v\n",v)
fmt.Println(v.String())
	/*reflect.Value 也满足 fmt.Stringer 接口, 但是除非 Value 持有的是字符串,
		否则 String 只是返回具体的类型. 相同, 使用 fmt 包的 %v 标志参数, 将使用 reflect.Values 的
	结果格式化.*/
b:=v.Type()
fmt.Println(b.String())
q:=reflect.ValueOf(3)
e:=q.Interface()
r:=e.(int)
fmt.Println("%d\n",r)
	}
func Sprint(x interface{})string{
	type stringer interface {
		String()string
	}
	switch x:=x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x{
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}