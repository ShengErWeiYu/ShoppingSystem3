package services

import (
	"Shopping_System/dao/mysql"
	g "Shopping_System/global"
	"Shopping_System/model"
	"Shopping_System/untils"
	"context"
	"encoding/json"
	"log"
)

func SearchGoodInLikeFromRedis() (interface{}, error) {
	key := "likelist"
	cmd2 := g.RedisClient.ZRevRangeWithScores(context.Background(), key, 0, -1)
	if cmd2.Err() != nil {
		log.Fatalln("2")
	}
	if len(cmd2.Val()) == 0 {
		return nil, nil
	} else {
		var results []interface{}
		for _, z := range cmd2.Val() {
			var data interface{}
			if err := json.Unmarshal([]byte(z.Member.(string)), &data); err != nil {
				log.Fatalln(err)
			}
			results = append(results, data)
		}
		return results, nil
	}
}
func SearchGoodInLike(myclaims untils.MyClaims) (interface{}, error) {

	return mysql.SearchGoodInLike(myclaims.Uid)
}
func LikeAGood(LikeGood *model.UserLike, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(LikeGood.Gid); err != nil {
		return err
	}
	return mysql.LikeAGood(LikeGood.Gid, myclaims.Uid)
}
