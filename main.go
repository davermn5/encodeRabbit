package main

import (
	"encoding/json"
	"fmt"
	"log"
	"text/template"
)
import "os"

type Inventory struct {
	Material string
	Count    uint
}

type person struct {
	name string
	age  int
}

var data struct {
	ID   int
	Name string
}

type Foo struct {
	i int
	s string
}

func main() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters) //Executes to the console immediately.
	if err != nil {
		panic(err)
	}

	fmt.Println("\n")

	b, err := json.Marshal(struct {
		ID   int
		Name string
	}{42, "The answer"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)

	fmt.Println("\n")

	w := person{name: "Stats", age: 20}
	fmt.Println(w.age)

	wp := &w
	wp.age = 21
	fmt.Println(w.age)

	fmt.Println("\n")

	errA := json.Unmarshal([]byte(`{"ID": 42, "Name": "The Answer"}`), &data)
	if errA != nil {
		log.Fatal(errA)
	}
	fmt.Println(data.ID, data.Name)

	fmt.Println("\n")

	var s = []Foo{
		{6 * 9, "Question"},
		{42, "Answer"},
	}

	fmt.Println(s)

	var m = map[int]Foo{
		7: {6 * 9, "Question"},
		2: {42, "Answer"},
	}

	fmt.Println(m)

}
