package service

import (
	"github.com/go-redis/redis"
	"log"
	"redisLearning/model"
)

func PostArticle(client *redis.Client, article *model.Article,title, id, link, poster string, vote int, publishTime int64)  (err error){
	article = &model.Article{
		Title: title,
		Id: id,
		Poster: poster,
		Vote: vote,
		PublishTime: publishTime,
	}
	err = article.SetArticleIntoRedis(client)
	err = article.SetArticleScore(client)
	err = article.SetArticleTime(client)
	if err != nil {
		log.Fatal("PostArticle error")
		return
	}
	return
}

func GetArticle(client *redis.Client, page int, order string)  []map[string]string{
	var articles	[]map[string]string
	var start = page + model.ARTICLES_PER_PAGE
	end := -1
	articleIds := client.ZRevRange(order, int64(start), int64(end))
	for _, articleId := range articleIds.Val() {
		articleMap := client.HGetAll(articleId)
		articles = append(articles, articleMap.Val())
	}
	return articles
}