package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

const (
	AddLikeStr             = "insert into likes (gid,uid) value (?,?)"
	FindGoodInLikeByUidStr = "select lid,gid from likes where lid>0 and uid=?"
)

func SearchGoodInLike(uid string) (interface{}, error) {
	var MyLikeGoods []model.MyLike
	key := "likelist"
	if err := g.Xdb.Select(&MyLikeGoods, FindGoodInLikeByUidStr, uid); err != nil {
		return nil, err
	} else {
		ttl := time.Second * 10
		for _, like := range MyLikeGoods {
			value, err := json.Marshal(like)
			if err != nil {
				log.Fatalln(err)
			}
			err = g.RedisClient.ZAdd(context.Background(), key, &redis.Z{
				Score:  float64(time.Now().Unix()),
				Member: value,
			}).Err()
			if err != nil {
				log.Fatalln(err)
			}
			err = g.RedisClient.Expire(context.Background(), "goods", ttl).Err()
			if err != nil {
				log.Fatalln(err)
			}
		}
		return MyLikeGoods, err
	}
}

func LikeAGood(gid int64, uid string) error {
	if _, err := g.Xdb.Exec(AddLikeStr, gid, uid); err != nil {
		return err
	}
	return nil
}
