package Dataserialisation

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}
type Name struct {
	Family   string
	Personal string
}
type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + "" + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address
	}
	return s
}
func lj() {
	var person Person
	loadJson("person.json", &person)
	fmt.Println("person", person.String())
}
func loadJson(fileName string, key interface{}) {
	infile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(infile)
	err = decoder.Decode(key)
	checkError(err)
	infile.Close()
}
func sj() {
	person := Person{
		Name: Name{
			Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@box.edu.au"}}}
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}
	os.Exit(0)
	//saveJson("person.json",person)
}
func saveJson(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}
