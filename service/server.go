package service

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"redisLearning/model"
	"time"
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

func ArticleVote(client *redis.Client, userId, articleId string) {
	cutoff := time.Now().Unix() - model.ONE_WEEK_IN_SECONDS
	if client.ZScore("time", "Article:"+articleId).Val() < float64(cutoff) {
		fmt.Sprintf("超出时间，无法投票")
		return
	}
	if client.SAdd("voted:" + articleId, userId).Val() != 1 {
		fmt.Sprintf("该用户已投过票")
		return
	}
	client.ZIncrBy("score", model.VOTE_SCORE, "Article:" + articleId)
	client.HIncrBy("Article：" + articleId, "vote", 1)
}