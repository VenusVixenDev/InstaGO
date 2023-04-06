package main

import (
	"github.com/Davincible/goinsta/v3"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")
	insta.SetProxy()

	// Only call Login the first time you login. Next time import your config
	if err := insta.Login(); err != nil {
		panic(err)
	}
	defer insta.Export("goinsta.txt")

	insta.Account.ExternalURL = ""
}
