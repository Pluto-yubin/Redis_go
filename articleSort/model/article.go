package model

import (
	"github.com/go-redis/redis"

)

type Article struct {
	Title       string
	Id          string
	Link        string
	Poster      string
	Vote        int
	PublishTime int64
}

func (a *Article) SetArticleIntoRedis(client *redis.Client) (err error) {
	err = client.HSet( "Article:"+a.Id, "title", a.Title).Err()
	err = client.HSet( "Article:"+a.Id, "link", a.Link).Err()
	err = client.HSet( "Article:"+a.Id, "poster", "user:"+a.Poster).Err()
	err = client.HSet( "Article:"+a.Id, "vote", a.Vote).Err()
	err = client.HSet("Article:"+a.Id, "publishTime", a.PublishTime).Err()

	return
}

func (a *Article) SetArticleScore(client *redis.Client) (err error) {
	err = client.ZAdd("score", redis.Z{
		Score: float64(a.PublishTime + VOTE_SCORE),
		Member: "Article:" + a.Id,
	} ).Err()
	return
}

func (a *Article) SetArticleTime(client *redis.Client) (err error) {
	err = client.ZAdd("time", redis.Z{
		Score: float64(a.PublishTime),
		Member: "Article:" + a.Id,
	}).Err()
	return
}

