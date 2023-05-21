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
	FindGoodInOrdersStr    = "select count(oid) from orders where gid = ?"
	GetOidOfOrder          = "select count(oid) from orders where uid=?  and gid = ?"
	PubulishCommentStr     = "insert into comments (gid, uid, star, comment) values (?,?,?,?)"
	FindCommentStr         = "select cid, gid, uid, star, comment from comments where gid=?"
	CheckCommentExistOrNot = "select count(star) from comments where gid =?"
)

func LookComment(gid int64) (interface{}, error) {
	var comment []model.LookResult
	key := "commentlist"
	if err := g.Xdb.Select(&comment, FindCommentStr, gid); err != nil {
		return nil, err
	} else {
		ttl := time.Second * 10
		for _, good := range comment {
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
		return comment, err
	}
}
func PubulishComment(uid string, gid int64, star int, comment string) error {
	if _, err := g.Xdb.Exec(PubulishCommentStr, gid, uid, star, comment); err != nil {
		return err
	}
	return nil
}
func CheckGoodOrderExist(gid int64) error {
	var count int
	if err := g.Xdb.Get(&count, FindGoodInOrdersStr, gid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorGoodExistInOrders
	}
	return nil
}

func CheckOrderWithUid(uid string, gid int64) error { //买了对应GID就行
	var count int
	if err := g.Xdb.Get(&count, GetOidOfOrder, uid, gid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorCommentPublishUser
	}
	return nil
}
func CheckComment(gid int64) error {
	var count int
	if err := g.Xdb.Get(&count, CheckCommentExistOrNot, gid); err != nil {
		return err
	}
	if count == 0 {
		return CommentNotExist
	}
	return nil
}
