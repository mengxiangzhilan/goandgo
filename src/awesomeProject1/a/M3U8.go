package a

import "container/list"

type M3U8 struct {
	basepath string
	tsList list.List
}



func (m *M3U8)Basepath()string{
	return m.basepath
}

func (m *M3U8)SetBasepath(basepath string){
	m.basepath=basepath
}
func (m *M3U8)SetTsList(tslist list.List){
	m.tsList=tslist
}
func (m *M3U8)TsList()list.List{
	return m.tsList
}
func (m *M3U8)AddTs(ts M3U8S){
	m.tsList.PushBack(ts)
}