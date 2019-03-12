package main

import (
	"bytes"
	"fmt"
	"image/color"
	"math"
)

/*函数前放一个变量，即是一个方法。将该函数附加到这种类型上，相当于为该类型定义了一个独占的方法*/
//要封装对象必须定义成一个struct
type Point struct{ X, Y float64 }
type Path []Point //命名的slice类型
type IntList struct {
	Value int
	Tail  *IntList
}
type ColoredPoint struct { //嵌入结构体来扩展类型
	//Point
	*Point
	Color color.RGBA
}

/*一个bit数组通常会用一个无符号数或者称之为“字”的slice或者来表示，每一个元素的每一位都
表示集合里的一个值。当集合的第i位被设置时，我们才说这个集合包含元素i
*/
type IntSet struct {
	words []uint64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p Point) Distance(q Point) float64 { //p 方法接受器
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

/*调用函会对器每一个参数值进行拷贝，如果函数需要更新变量，或者函数中莫一个参数实在太大，我们希望能避免进行这种默认的
拷贝。使用指针，接受者变量本身比较大时可以使用指针而不是对象
*/
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func (list *IntList) sum() int { //使用nil指针作为接受器
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.sum()
}
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] = 1 << bit
}
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword

		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func Fangfamain() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)
	//可以把ColoredPoint类型当作接收器来调用Point里的方法
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	/*var pp=ColoredPoint{Point{1,1},red}
	var qq=ColoredPoint{Point{5,4},blue}
	fmt.Println(pp.Distance(qq.Point))
	pp.ScaleBy(2)
	qq.ScaleBy(2)
	fmt.Println(pp.Distance(qq.Point))*/
	pp := ColoredPoint{&Point{1, 1}, red}
	qq := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(pp.Distance(*qq.Point))
	qq.Point = pp.Point
	pp.ScaleBy(2)
	fmt.Println(*pp.Point, *qq.Point)
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}
