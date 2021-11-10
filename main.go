package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")

	number := 5
	number2 := 81

	hasil := number * number2
	var tampil string
	if hasil%2 == 0 {
		tampil = "hasilnya adalah genap"
	} else {
		tampil = "hasilnya adalah ganjil"
	}
	fmt.Println(tampil)

	// =====================
	// data type
	// =====================

	// === atomic data types ===

	// string
	huruf := "string"
	fmt.Println("ini adalah huruf", huruf)

	var iniAbjad string
	iniAbjad = "huruf ini adalah sebuah abjad"
	fmt.Println("ini adalah string", iniAbjad)
	fmt.Println(strings.Contains(huruf, iniAbjad))

	fmt.Println(strings.Split(iniAbjad, " "))

	// int

	angka := 40
	fmt.Println(angka)
	// int32
	// int64
	// uint
	// uint32
	// uint64

	// === unsage ===

	// Pointers

	// === abstract data type ===

	//map[] or arrays
	todo()

	pointer()

	structure("Firli", 19, 2)

	interfaceStudy()

	printAnything()

	controlStatement()

	switchCase("Sun")
	//struct{}
	//interface{}
}

func todo() {

	array := []int{1, 2, 3, 4}

	arr2 := []string{"hi", "my", "hello", "world"}

	arr2 = append(arr2, "test", "ini", "test", "append")
	fmt.Println("angka =", array, "string =", arr2)
	coba := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		coba = append(coba, array[i])
	}
	fmt.Println("reverse array", coba)

}

func pointer() {
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}

func swap(m1, m2 *int) {
	var ubah int
	ubah = *m2
	*m2 = *m1
	*m1 = ubah
}

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) GetName() {
	fmt.Println("the car name is " + c.Name + " and its age is")
}

type Mobil interface {
	Drive()
	Stop()
}
type Lambo struct {
	LamboModel string
}
type Pajero struct {
	PajeroModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo sedang mengemudi")
	fmt.Println(l.LamboModel)
}

func (p *Pajero) Drive() {
	fmt.Println("Pajero sedang mengemudi")
	fmt.Println(p.PajeroModel)
}

func interfaceStudy() {
	l := Lambo{"M221"}
	p := Pajero{"C92K"}
	l.Drive()

	p.Drive()
}

func structure(name string, age int, modelno int) {
	c := Car{
		Name:    name,
		Age:     age,
		ModelNo: modelno,
	}
	fmt.Println(c)
	c.GetName()

}

func Anything(param interface{}) {
	fmt.Println(param)
}

func printAnything() {
	Anything("ini string")
	Anything(123123)
	Anything([]string{"hi", "my", "hello", "world"})
	Anything(struct{}{})
	Anything([]interface{}{1, 2, 3, "test", "ini", struct{}{}})

	mymap := make(map[string]interface{})
	mymap["name"] = "123123"
	mymap["age"] = 123
}
func switchCase(Value string) {
	switch Value {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("hari kerja")
	case "Sun", "Sat":
		fmt.Println("hari libur")
	}
}

func controlStatement() {
	fmt.Println("Hello world")

	//if else, for, switch case, break continue
	lampu := "red"
	fmt.Println(lampu)
	if lampu == "red" {
		lampu = "yellow"
		fmt.Println(lampu)
	} else if lampu == "yellow" {
		lampu = "green"
		fmt.Println(lampu)
	} else {
		fmt.Println(lampu)
	}

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("value is il")
	} else if *flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	arr := []string{"satu", "dua", "tiga", "empat"}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	fmt.Println("=======")

	mymap2 := make(map[string]interface{})
	mymap2["name"] = "angad"
	mymap2["age"] = 21
	for k, v := range mymap2 {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}

	numberArr := []int{1, 2, 3, 4}
	for _, value := range numberArr {
		fmt.Println(value)
	}

}
