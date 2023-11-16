package main

import (
	"log"
	"os"
	"rayslack/utils"
	"strings"

	"github.com/slack-go/slack"
)

func main() {
	// Set your api token in evnironment variable
	content, err := os.ReadFile("./my-token.secret")
	if err != nil {
		log.Fatal(err)
	}
	// Read a token from environment variable
	token := strings.Trim(string(content), "\n")
	api := slack.New(token, slack.OptionDebug(true))

	users := utils.GetAllUsers(api)
	utils.SaveDeletedUsers(users)
}
