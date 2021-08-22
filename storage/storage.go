package storage

import (
	"GoProject/types"
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	setValues()
}

func setValues() {
	hackers := map[string]float64{
		"Alan Turing":        1912,
		"Claude Shannon":     1916,
		"Alan Kay":           1940,
		"Richard Stallman":   1953,
		"Yukihiro Matsumoto": 1965,
		"Linus Torvalds":     1969,
	}

	for name, birthday := range hackers {
		_, err := client.ZAdd(string(types.Hackers), redis.Z{birthday, name}).Result()
		if err != nil {
			log.Fatalf("\nZADD Error: %+s\nInner params: {name: %+s, birthday: %+v}\n", err.Error(), name, birthday)
			continue
		}
	}
}

func GetValues() ([]*types.Hacker, error) {
	var hackersSlice []*types.Hacker
	hackers, err := client.ZRangeWithScores(string(types.Hackers), 0, -1).Result()
	if err != nil {
		log.Fatalf("Ошибка при попытке получить данные: %+s", err)
		return nil, err
	}
	for _, value := range hackers {
		hackersSlice = append(hackersSlice, &types.Hacker{Name: value.Member.(string), Birthday: int64(value.Score)})
	}

	return hackersSlice, nil
}
