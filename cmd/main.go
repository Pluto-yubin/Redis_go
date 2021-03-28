package main

import (
	"GolandProj/Redis/redisLearning/conf"
	"GolandProj/Redis/redisLearning/model"
)

func main() {

	client := conf.GetRedisClient()
	article1 := &model.Article{
		Title:  "java",
		Id:     "1",
		Link:   "java.com",
		Poster: "小王",
		Vote:   500,
	}
	article1.SetArticle(client)
	Article2 := &model.Article{
		Title:  "go",
		Id:     "2",
		Link:   "golang.com",
		Poster: "小张",
		Vote:   300,
	}
	Article2.SetArticle(client)
}
