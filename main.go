package main

import (
	"fmt"

	"github.com/nikhilreddypasunoor/AspirationGoMapper/api"
)

func main() {
	fmt.Println("result for question1: ", api.CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
	s := api.NewSkipString(3, "Aspiration.com")
	api.MapString(&s)
	fmt.Println("result for question2: ", s)
}
