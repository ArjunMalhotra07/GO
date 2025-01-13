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
type PurchaseInformation struct {
	BuyerName  string
	Price      float32
	ItemBought string
}

func loadData[T ContactInfo | PurchaseInformation](jsonFilePath string) []T {
	data, _ := ioutil.ReadFile(jsonFilePath)
	loaded := []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}

func UnMarshalJsons() {
	var contact []ContactInfo = loadData[ContactInfo]("./jsons/contact.json")
	var purchases []PurchaseInformation = loadData[PurchaseInformation]("./jsons/purchase.json")
	fmt.Printf("\n%+v", contact)
	fmt.Printf("\n%+v", purchases)

}
