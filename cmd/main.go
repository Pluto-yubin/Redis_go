package main

import (
	"fmt"
	"redisLearning/conf"
	"redisLearning/model"
	server "redisLearning/service"
	"time"
)

func main() {

	client := conf.GetRedisClient()
	var article1, article2 *model.Article
	article2 = &model.Article{}
	article1 = &model.Article{}
	server.PostArticle(client, article1, "java", "1", "java.com", "小王",1,time.Now().Unix() - model.ONE_WEEK_IN_SECONDS - 1)
	server.PostArticle(client, article2, "go", "2", "golang.com", "小红", 1, time.Now().Unix())
	article := server.GetArticle(client, 0, "time")
	fmt.Println(article)

}
