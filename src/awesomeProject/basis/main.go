package main

import (
	"fmt"
	"unsafe"
	)

func main() {
	/**
                变量
内建变量
bool、string （u）int,（u）int8/16/32/64，uintptr，byte，rune，float34，float64，complex64，complex128
加u无符号，，uintptr指针，rune相当于java的char但长度32位，complex为复数
go只有强制类型转换
 */
	var age int             //变量声明,不赋值，默认为0
	var name string = "add" //声明时赋值
	var xiao = 20           //自动推断
	var w, h int = 100, 20
	var width, height int
	var ( //可以定义不同类型的变量
		name1 = ""
		age1  = 1
		h1    int
	)
	name2, age2 := "", 90 //赋值,:=左边至少要有一个变量是尚未声明的
	width = 900
	height = 600
	a := true
	b := false
	//%T格式说明符，%d用于打印字节大小
	fmt.Println(age, name, xiao, w, h, width, height, name1, age1, h1, name2, age2, a && b)
	fmt.Printf("type of width is %T,size of width is %d", width, unsafe.Sizeof(width))
	c1 := complex(5, 7) //复数
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum", cadd)
	cmul := c1 * c2
	fmt.Println("prodduct", cmul)
	//go的类型转换非常严格
	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println(sum)
	const aaa = 453 //常量,不允许将函数返回值赋值给常量
	//aaa=89
	fmt.Println(aaa)
	const hel = "hello world" //字符串常量没有任何类型
	//无类型常量有默认类型，当变量被赋值为常量时将默认类型提供给他
	const yypeH string = "ss  sss" //有类型常量
	//混合类型不允许
	var defaultName = "Sam" //允许
	type myString string
	var customString myString = "Sam" //允许
	//customString=defaultName  不允许
	fmt.Println(defaultName, customString)
	//数字只有使用时才需要类型
	const aw = 5
	var intVar int = aw
	var int32Var int32 = aw
	var float64Var float64 = aw
	var complex64Var complex64 = aw
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}

/*
函数
_  表示任何类型的任何值，可以接收你不用或则不关心的值
 */
func funs(girlfriend string) string {
	//     参数 类型   返回值类型
	arec, _ := family("", "", "")
	fmt.Println(arec)
	//var familyname string
	familyname:=arec+girlfriend
	return familyname
	//如果多个参数类型一致，在最后添加类型
}
func family(facher, mother, yourself string) (string, int) {
	//go支持多返回值
	//注意必须用()将多个返回值kuoqil
	var sum = 3
	var s = facher + mother + yourself
	return s, sum
}
func family2(facher, mother, yourself string) (a, b int) {
	//go支持多返回值
	//注意必须用()将多个返回值kuoqil
	a = 3
	b = 4
	//函數返回一個命名值，不需要return指明返回值
	return
}

/*
				包
所有可执行的go都必须包含一个main函数。为程序运行的入口。
package packagename指明文件属于那个包
一般某个包的源文件应该单独放在一个文件夹
惯例，包名即文件夹名
首字母大写变量或函数即视为public否则private
所有包都可以包含一个init()函数
 */
func init() {
	//不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它
	//init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性
	//一个包导入了另一个包，会先初始化被导入的包
	//一个包被导入多次，但是它只会被初始化一次
}

// import 包
// var _=包名.Area  //错误屏蔽器，防止包暂时不用的情况下报错
//_包名  有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量

//if else
func iff() {
	condition := true
	if condition {
	} else {

	}
	if statement := 10; condition { //statement 可选部分,在判断之前运行
		fmt.Println(statement)
	} else { //else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过

	}
}

/*
		循环
 */
func forf(){
	//for initialisation; condition; post{}
	for i:=1;i<=10;i++{
		fmt.Printf("%d",i)
	}
	//break终止for，执行for循环后的下一行代码
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	//continue跳出for中的当前循环，continue后的语句不在本次循环执行，接着执行循环体
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
//在 for 循环中可以声明和操作多个变量
		for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
			fmt.Printf("%d * %d = %d\n", no, i, no*i)
		}
	}
	//无限循环
	for {
		fmt.Println("Hello World")
	}
}
         /*
			switch
		 */
 func switch1(){

	 switch abce:=8;abce{//abce作用域仅限switch内
	 case 1:
	 	fmt.Println("")
	 	fallthrough//执行完case后不跳出执行下一个
	 case 2:
	 fmt.Println("")
	 default://默认情况，以上条件不匹配时运行
		 fmt.Println("")
	 }
	 num := 75
	 switch { // 表达式被省略了
	 case num >= 0 && num <= 50:
		 fmt.Println("num is greater than 0 and less than 50")
	 case num >= 51 && num <= 100:
		 fmt.Println("num is greater than 51 and less than 100")
	 case num >= 101:
		 fmt.Println("num is greater than 100")
	 }
 }
	 /*
	 数组
	 值类型
	  */
  func array(){
  	var a[3] int
  	a[0]=12
  	a[1]=23
  	a[2]=58
  	fmt.Println(a)
  	b:=[3]int{1,1,1}
  	fmt.Println(b)
  	c:=[...]int{}
  	fmt.Println(c)
  	var sum=0
  	for i,v:= range  a{//rang返回索引和值
  		fmt.Printf("%d\n",i,v)
  		sum+=v
	}
  	fmt.Println(sum)
  	d:=[3][4]string{
  		{"q","q"},
		{"q","q"},
		{"q","q"},
	}
  	fmt.Println(d)

  }
