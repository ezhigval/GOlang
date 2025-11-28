package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"cli_conv/internal/analytics"
	"cli_conv/internal/model"
	"cli_conv/internal/stats"
)

var actions = map[string]func([]model.User){
	"stats": func(users []model.User) {
		stats.GetStats(users)
	},
	"users_list": func(users []model.User) {
		stats.GetUsersList(users)
	},
	"user_with_tag": func(users []model.User) {
		stats.UserWithTag(users)
	},
	"analytics": func(users []model.User) {
		analytics.GetAnalytics(users)
	},
	"concurrent_scores": func(users []model.User) {
		scores := analytics.GetConcurrentScores(users)
		fmt.Println("User scores (concurrent):")
		for _, s := range scores {
			fmt.Println("-", s.User.Name, "->", s.Score)
		}
	},
}

func Jsonconv() {
	fmt.Println("Select a file to convert:")
	var filePath string
	fmt.Scanln(&filePath)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
	}

	if len(jsonData) == 0 {
		fmt.Println("File is empty")
		panic(errors.New("file is empty"))
	}

	fmt.Println("File read successfully")
	var users []model.User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		panic(err)
	}

	fmt.Println("Select a function:")
	for name := range actions {
		fmt.Println("-", name)
	}

	var functionName string
	fmt.Scanln(&functionName)
	fmt.Println("--------------------------------")

	handler, ok := actions[functionName]
	if !ok {
		fmt.Println("Invalid function")
		panic(errors.New("invalid function"))
	}
	handler(users)
	fmt.Println("-------------EXIT-------------")
}
