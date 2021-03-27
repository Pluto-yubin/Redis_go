package cmd

import (
	
	"model/Article"
	"github.com/go-redis/redis"
)

const ONE_WEEK_IN_SECONDS = 7 * 86400
const VOTE_SCORE = 432

func (a *Article)setArticle(client *redis.Client) (err error) {
	err = client.HSet("Article:"+a.id, "title", a.title).Err()
	err = client.HSet("Article:"+a.id, "link", a.link).Err()
	err = client.HSet("Article:"+a.id, "poster", "user:"+a.poster).Err()
	err = client.HSet("Article:"+a.id, "vote", a.vote).Err()
	err = client.HSet("Article:"+a.id, "publishTime", a.publishTime).Err()
	return
}

