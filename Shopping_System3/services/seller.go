package services

import (
	"Shopping_System/dao/mysql"
	g "Shopping_System/global"
	"Shopping_System/untils"
	"context"
	"encoding/json"
	"log"
)

func CheckOrderFromRedis() (interface{}, error) {
	key := "orderlist"
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
func CheckOrder(myclaims untils.MyClaims) (interface{}, error) {
	if err := mysql.CheckOrderExist(myclaims.Uid); err != nil {
		return nil, err
	}
	return mysql.CheckOrder(myclaims.Uid)
}
