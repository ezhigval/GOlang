package analytics

import (
	"cli_conv/internal/model"
	"errors"
	"fmt"
	"sync"
)

type UserScore struct {
	User  model.User
	Score int
}

func calculateScore(u model.User) int {
	return u.Age * len(u.Tags)
}

func worker(id int, jobs <-chan model.User, results chan<- UserScore, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("worker", id, "recovered from panic:", r)
		}
	}()

	for u := range jobs {
		score := calculateScore(u)
		results <- UserScore{User: u, Score: score}
	}
}

func GetConcurrentScores(users []model.User) []UserScore {
	jobs := make(chan model.User)
	results := make(chan UserScore)

	var wg sync.WaitGroup
	nWorkers := 3

	// поднимаем воркеры
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// отправляем задачи
	go func() {
		for _, u := range users {
			jobs <- u
		}
		close(jobs)
	}()

	// ждём воркеры и закрываем results
	go func() {
		wg.Wait()
		close(results)
	}()

	var scores []UserScore
	for r := range results {
		scores = append(scores, r)
	}

	return scores
}

func GetAnalytics(users []model.User) {
	if len(users) == 0 {
		err := errors.New("no users")
		panic(err)
	}

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
