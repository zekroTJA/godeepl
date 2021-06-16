package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zekrotja/godeepl"
)

func main() {
	godotenv.Load()

	c := godeepl.New(godeepl.ClientOptions{
		Endpoint:  godeepl.EndpointPro,
		Email:     os.Getenv("EMAIL"),
		Password:  os.Getenv("PASSWORD"),
		SessionID: os.Getenv("SESSION"),
	})

	res, err := c.Translate(godeepl.LangGerman, godeepl.LangEnglish, "Jo, was geht ab du alter Sack! Dauert das noch lange?")
	fmt.Println(res.Translation(0).Beam(0), err)
}
