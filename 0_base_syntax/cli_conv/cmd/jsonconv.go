package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"cli_conv/internal/model"
)

func Jsonconv() {
	fmt.Println("Select a file to convert:")
	var filePath string
	fmt.Scanln(&filePath)

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(jsonData) == 0 {
		fmt.Println("File is empty")
		return
	} else {
		fmt.Println("File read successfully")
		var user model.User
		err = json.Unmarshal(jsonData, &user)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		fmt.Println("User:")
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.Name)
		fmt.Println("Email:", user.Email)
		fmt.Println("Age:", user.Age)
		fmt.Println("Tags:", user.Tags)

	}
}
