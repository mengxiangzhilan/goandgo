package main

import "fmt"

type ListNode struct {
	val int
	next *ListNode

}
func main() {
	/*var a ListNode
	var b ListNode
	var c ListNode
	var d ListNode
	var e ListNode
	a.val=10
	b.val=20
	c.val=30
	d.val=40
	e.val=50
	a.next=&b
	b.next=&c
	c.next=&d
	d.next=&e
	e.next=nil
	var head *ListNode
	head=&a*/

fmt.Println(family1("sss"))
	}


func family1(girlfriend string)string{
	//方法名  参数   参数类型  返回值类型
	arec,_ := family2("","","",)
	familyname := arec + girlfriend//高性能拼接字符串请用buffer.WriteString()
	return familyname


}
func family2(father, mother, yourself string)(string,int){
	//方法名  多个参数               参数类型  返回值类型
	//go支持多返回值，并且有多个返回值时用()括起来
	var sum = 3
	var s = father + mother + yourself
	return s, sum
}
func family3(facher, mother, yourself string) (a, b int) {
	a = 3
	b = 4

	return //函數返回一個命名值，不需要return指明返回值
}