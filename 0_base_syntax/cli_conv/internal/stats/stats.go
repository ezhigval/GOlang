package stats

import (
	"cli_conv/internal/model"
	"errors"
	"fmt"
	"slices"
)

func GetStats(users []model.User) {
	averageAge := 0
	for _, user := range users {
		averageAge += user.Age
	}
	averageAge /= len(users)

	averageTags := 0
	for _, user := range users {
		averageTags += len(user.Tags)
	}
	averageTags /= len(users)

	fmt.Println("--------------------------------")
	fmt.Println("Stats:")
	fmt.Println("Total users:", len(users))
	fmt.Println("Average age:", averageAge, "years")
	fmt.Println("Average tags:", averageTags, "tags")
	fmt.Println("--------------------------------")
}

func GetUsersList(users []model.User) {
	for _, user := range users {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Age: ", user.Age)
		fmt.Println("Tags: ", user.Tags)
		fmt.Println("--------------------------------")
	}
}

func UserWithTag(users []model.User) ([]model.User, error) {
	var tag string
	fmt.Println("Enter a tag:")
	fmt.Scanln(&tag)
	if tag == "" {
		fmt.Println("Invalid tag")
		return nil, errors.New("invalid tag")
	}
	var userWithTag []model.User
	for _, user := range users {
		if slices.Contains(user.Tags, tag) {
			userWithTag = append(userWithTag, user)
		}
	}
	if len(userWithTag) == 0 {
		return nil, errors.New("no users with tag")
	}
	fmt.Println("Users with tag:", tag)
	for _, user := range userWithTag {
		fmt.Println("ID:", user.ID)
		fmt.Println("Name: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Age: ", user.Age)
		fmt.Println("Tags: ", user.Tags)
		fmt.Println("--------------------------------")
	}
	return userWithTag, nil
}
