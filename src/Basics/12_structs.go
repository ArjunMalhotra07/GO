//Structs: Tags

package basics

import (
	"fmt"
	"reflect"
)

type Animal02 struct {
	Name string `required max:"100"`
	//Backticks as delimiters of the tag + space delimited key value pairs
	Origin string
}

func Structs12() {
	t := reflect.TypeOf(Animal02{})
	//We need to get the type of object we are working with
	//Reflect Package ka TypeOf func. with object as arg
	field, _ := t.FieldByName("Name")
	//Getting a field from that type by using the Types's FieldByName method
	fmt.Println(field.Tag)
	//Tags can be added to struct fields to describe field

}
