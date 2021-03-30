package service

import (
	"github.com/go-redis/redis"
	"math"
	"time"
	"web/model"
)

func checkToken(client redis.Client, token string) string {
	return client.HGet("login:", token).Val()
}

func updateToken(client redis.Client, token, user, item string, )  {
	timeStamp := time.Now()
	client.HSet("login:", token, user)
	client.ZAdd("recent:", redis.Z{
		Score: float64(timeStamp.Unix()),
		Member: token,
	})
	if item != "" {
		client.ZAdd("view:"+token, redis.Z{
			Score: float64(timeStamp.Unix()),
			Member: item,
		} )
		// 这里0表示第一个，-26表示倒数第26个， 假设redis中共有len个字段，第0个也是第-len个，这里只保留25条数据
		client.ZRemRangeByRank("view:"+token, 0, -26)
	}
}

func cleanFullSession(client redis.Client) {
	for {
		size := client.ZCard("recent")
		if size.Val() < model.LIMIT {
			time.Sleep(1000 * time.Millisecond)
		}
		endIndex := math.Min(float64(size.Val()-model.LIMIT), 100)
		tokens := client.ZRange("recent:", 0, int64(endIndex-1))
		sessionsKeys := []string{}
		for _, sess := range tokens.Val() {
			sessionsKeys = append(sessionsKeys, "viewed:" + sess)
			sessionsKeys = append(sessionsKeys, "cart:" + sess)
		}
		client.Del(sessionsKeys...)
		client.HDel("login:", sessionsKeys...)
		client.ZRem("recent:", sessionsKeys)
	}


}
