package analytics

import (
	"cli_conv/internal/model"
	"fmt"
)

func GetAnalytics(users []model.User) {
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

	fmt.Println("Analytics:")
	fmt.Println("Total users:", len(users))
	fmt.Println("Average age:", averageAge, "years")
	fmt.Println("Average tags:", averageTags, "tags")
}
