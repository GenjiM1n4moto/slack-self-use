package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
)

func SaveDeletedUsers(users []slack.User) {
	deletedUsers := DeletedUsers(users)
	currentTimestamp := time.Now()
	timeDateString := currentTimestamp.Format("2006-01-02-15:04:05")
	if _, err := os.Stat("./deleted-users/"); os.IsNotExist(err) {
		os.Mkdir("./deleted-users/", 0644)
	}
	// create a file with timestamp
	filePath := "./deleted-users/" + "deleted-users-" + timeDateString + ".txt"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, deletedUser := range deletedUsers {
		fmt.Fprintf(file, "ID: %s", deletedUser.ID)
		fmt.Fprintf(file, " Name: %s", deletedUser.Name)
		fmt.Fprintf(file, " DisplayName: %s", deletedUser.Profile.DisplayName)
		fmt.Fprintf(file, " Deleted: %v \n", deletedUser.Deleted)
	}
}
