package model

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

type Article struct {
	Title       string
	Id          string
	Link        string
	Poster      string
	Vote        int
	PublishTime time.Time
}

func (a *Article) SetArticle(client *redis.Client) (err error) {
	err = client.HSet(context.Background(), "Article:"+a.Id, "title", a.Title).Err()
	err = client.HSet(context.Background(), "Article:"+a.Id, "link", a.Link).Err()
	err = client.HSet(context.Background(), "Article:"+a.Id, "poster", "user:"+a.Poster).Err()
	err = client.HSet(context.Background(), "Article:"+a.Id, "vote", a.Vote).Err()
	err = client.HSet(context.Background(), "Article:"+a.Id, "publishTime", a.PublishTime).Err()

	return
}
