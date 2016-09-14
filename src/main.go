package main

import "fmt"
import "reflect"

type MagicInterface interface {
	Error() string
}

type MagicObject struct{}

func (e *MagicObject) Error() string {
	fmt.Println("Error")
	return "Magic!!!"
}

func Generate() *MagicObject {
	fmt.Println("Generate")
	return nil
}

func TestMagicInterface() MagicInterface {
	fmt.Println("Test magic interface")
	return Generate()
}

func TestMagicObject() *MagicObject {
	fmt.Println("Test magic object")
	return Generate()
}

func main() {
	fmt.Println(reflect.TypeOf(TestMagicInterface()))
	fmt.Println("---------")
	if TestMagicInterface() != nil {
		fmt.Println("Hello, magic interface!")
	}
	fmt.Println("---------")
	fmt.Println(reflect.TypeOf(TestMagicObject()))
	fmt.Println("---------")
	if TestMagicObject() != nil {
		fmt.Println("Hello, magic oject!")
	}
}
