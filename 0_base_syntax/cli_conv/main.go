package main

import (
	"cli_conv/cmd"
	"errors"
	"fmt"
	"os"
)

func main() {
	safeMain()
}

func safeMain() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	runMain()
}

func runMain() {
	fmt.Println("Hello, user!")

	if len(os.Args) < 2 {
		fmt.Println("Select a function: jsonconv, csvconv, xmlconv")
		panic(errors.New("no function selected"))
	}

	functionList := []string{"jsonconv", "csvconv", "xmlconv"}

	if !contains(functionList, os.Args[1]) {
		fmt.Println("Invalid function")
		panic(errors.New("invalid function"))
	}

	sliFunc := os.Args[1]

	switch sliFunc {
	case "jsonconv":
		cmd.Jsonconv()
	// case "csvconv":
	// 	cmd.Csvconv()
	// case "xmlconv":
	// 	cmd.Xmlconv()
	default:
		fmt.Println("Invalid function")
		panic(errors.New("invalid function"))
	}
}

// contains reports whether value is present in list.
func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
