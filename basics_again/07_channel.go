package basicsagain

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE = 5
var MAX_TOFU_PRICE = 10

type Deal struct {
	Name  string
	Price float32
}

func CheckPricesOnVariousSites() {
	var chickenChannel = make(chan Deal)
	var tofuChannel = make(chan Deal)

	websites := []string{
		"Walmart", "KFC", "Burger King"}
	for i := range websites {
		go checkPrices(websites[i], chickenChannel, float32(MAX_CHICKEN_PRICE))
		go checkPrices(websites[i], tofuChannel, float32(MAX_TOFU_PRICE))

	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkPrices(websiteName string, channelName chan Deal, maxPrice float32) {
	for {
		time.Sleep(time.Second * 1)
		var dealPrice = rand.Float32() * 20
		if dealPrice <= maxPrice {
			channelName <- Deal{Name: websiteName, Price: dealPrice}
		}
	}
}

func sendMessage(chickenChannel, tofuChannel chan Deal) {
	select {
	case deal := <-chickenChannel:
		fmt.Printf("\nFound a deal on Chicken at %s for %v\n", deal.Name, deal.Price)
	case deal := <-tofuChannel:
		fmt.Printf("\nFound a deal on Tofu at %s for %v\n", deal.Name, deal.Price)
	}
}
