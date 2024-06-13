package basicsagain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GenericsMainFunc() {
	UnMarshalJsons()
}

type ContactInfo struct {
	Name  string
	Phone int
}

func loadData[T ContactInfo](jsonFilePath string) []T {
	data, _ := ioutil.ReadFile(jsonFilePath)
	loaded := []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}

func UnMarshalJsons() {
	var contact = loadData("./jsons/contact.json")
	fmt.Printf("\n%+v", contact)
}
