package a

import "strconv"

type M3U8S struct {
	file string
	seconds float64
}

func NewM3U8S(file string,seconds float64)*M3U8S{
	return &M3U8S{file:file,seconds:seconds}
}
func (m *M3U8S)File()string{
	return m.file
}

func (m *M3U8S)SetFile(file string){
	m.file=file
}
func (m *M3U8S)SetSeconds(seconds float64){
	m.seconds=seconds
}
func (m *M3U8S)Seconds() float64{
	return m.seconds
}
func (m *M3U8S)ToString()string{

	return m.file+"("+strconv.FormatFloat(m.seconds,'E',-1,64)+"sec)"
}