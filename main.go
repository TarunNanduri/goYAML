package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

//Sample structure to demo yaml encoding and decoding
type Sample struct{
	ID int
	Vars []Sample1
}
//Sample1 to check heirarchy in yaml
type Sample1 struct{
	Variable int
}
func main() {
	s1:=&Sample{
		ID: 1,
		Vars: []Sample1{Sample1{Variable:1},Sample1{Variable:2}},
	}
	file,err:=os.Create("sample.yaml")
	yam,_:=yaml.Marshal(s1)
	fmt.Println(string(yam))
	if(err!=nil){
		fmt.Println("Error wiriting to file")
		return
	}
	file.WriteString(string(yam))
	file.Close()

	data,er:=ioutil.ReadFile("sample.yaml")
	var content Sample
	if(er!=nil){
		fmt.Println("Error reading file")
		return
	}
	er1:=yaml.Unmarshal(data,&content)
	if(er1!=nil){
		fmt.Println("Error unmarshaling yaml")
		return
	}
	fmt.Println(content)
}