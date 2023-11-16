package utils

import "github.com/slack-go/slack"

func DeletedUsers(users []slack.User) []slack.User {
	deletedUsers := make([]slack.User, 0)
	for _, user := range users {
		if user.Deleted && user.IsBot == false {
			deletedUsers = append(deletedUsers, user)
		}
	}
	return deletedUsers
}
