package basicsagain

import (
	"fmt"
)

func ChannelsMainFunction() {
	BasicChannel()
}

func BasicChannel() {
	var c = make(chan user)
	go sendUsersToChannel(c)
	for u := range c {
		fmt.Println("Received user:", u.name)
	}
}

func sendUsersToChannel(c chan user) {
	defer close(c) // Close channel after sending all users

	users := []user{
		{name: "Arjun"},
		{name: "Suchitra"},
		{name: "Ravi"},
		{name: "Meena"},
		{name: "Kiran"},
	}

	for _, u := range users {
		c <- u // Send user struct to channel
	}
}

// Struct definition
type user struct {
	name string
}
