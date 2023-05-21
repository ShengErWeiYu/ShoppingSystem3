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
	FindOrderDetailByGidStr = "select oid,uid,gid,number,size,address,phonenumber,realname from orders where gid = ?"
	FindOrderByGidStr       = "select count(oid) from orders where gid = ?"
)

func CheckOrderExist(uid string) error {
	var a []string
	if err := g.Xdb.Select(&a, "select gid from goods where uid=?", uid); err != nil {
		return err
	} else {
		var count int
		for _, v := range a {
			var c int
			if err := g.Xdb.Get(&c, FindOrderByGidStr, v); err != nil {
				return err
			}
			count += c
		}
		if count == 0 {
			return ErrorOrderExist
		} else {
			return nil
		}
	}
}
func CheckOrder(uid string) ([]model.CheckOrderResult, error) {
	var Order []model.CheckOrderResult
	var a []string
	key := "orderlist"
	if err := g.Xdb.Select(&a, "select gid from goods where uid = ?", uid); err != nil {
		return nil, err
	} else {
		for _, v := range a {
			var o []model.CheckOrderResult
			if err := g.Xdb.Select(&o, FindOrderDetailByGidStr, v); err != nil {
				return nil, err
			}
			Order = append(Order, o...)
		}
		ttl := time.Second * 10
		for _, good := range Order {
			value, err := json.Marshal(good)
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
			err = g.RedisClient.Expire(context.Background(), key, ttl).Err()
			if err != nil {
				log.Fatalln(err)
			}
		}
		return Order, nil
	}
}
