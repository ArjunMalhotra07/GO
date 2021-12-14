//Structs: Tags

package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string `required max:"100"`
	//Backticks as delimiters of the tag + space delimited key value pairs
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	//We need to get the type of object we are working with
	//Reflect Package ka TypeOf func. with object as arg
	field, _ := t.FieldByName("Name")
	//Getting a field from that type by using the Types's FieldByName method
	fmt.Println(field.Tag)

}
