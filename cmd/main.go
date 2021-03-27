package main

import (
	"github.com/go-redis/redis"
	"model/Article"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Article1 := &Article{
		title: "java",
		id: "1",
		link: "java.com",
		poster: "小王",
		vote: 500,

	}
	Article1.setArticle(client)
	Article2 := &Article{
		title: "go",
		id: "2",
		link: "golang.com",
		poster: "小张",
		vote: 300,
	}
	Article2.setArticle(client)
}