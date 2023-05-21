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

func LookComment(Look *model.Look) (interface{}, error) {
	if err := mysql.CheckGid(Look.Gid); err != nil {
		return nil, err
	}
	if err := mysql.CheckComment(Look.Gid); err != nil {
		return nil, err
	}
	return mysql.LookComment(Look.Gid)
}
func PublishComment(PB *model.Comments, myclaims untils.MyClaims) error {
	if err := mysql.CheckGoodOrderExist(PB.Gid); err != nil { //查询在order表里有没有购买这个商品的记录
		return err
	}
	if err := mysql.CheckOrderWithUid(myclaims.Uid, PB.Gid); err != nil {
		return err
	} //查看该用户是否购买该商品
	return mysql.PubulishComment(myclaims.Uid, PB.Gid, PB.Star, PB.Comment)
}
func LookCommentFromRedis() (interface{}, error) {
	key := "commentlist"
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
