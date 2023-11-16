package utils

import "github.com/slack-go/slack"

func GetAllUsers(api *slack.Client) []slack.User {
	users, err := api.GetUsers()
	if err != nil {
		panic(err)
	}
	return users
}
