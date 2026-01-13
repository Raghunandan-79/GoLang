package main

import (
	"fmt"

	"github.com/Raghunandan-79/Podcast/auth"
	"github.com/Raghunandan-79/Podcast/user"
)

func main() {
	auth.LoginWithCredentials("raghu79", "password")
	fmt.Println(auth.GetSession())

	u := user.New("raghu", "raghu@gmail.com")
	
	u.SetDetails("raghu12", "raghu12@gmail.com")
	u.GetDetails()
}