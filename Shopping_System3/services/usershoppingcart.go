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

func OperateSizeOfGood(ReviseSize *model.ReviseSize, myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(ReviseSize.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(myclaims.Uid, int64(ReviseSize.Sid)); err != nil {
		return err
	}
	var gid int64
	var err error
	if gid, err = mysql.GetGidBySid(ReviseSize.Sid); err != nil {
		return err
	}
	if err := mysql.CheckSize(ReviseSize.NewSize, gid); err != nil {
		return err
	}
	return mysql.OperateSizeOfGood(ReviseSize.NewSize, ReviseSize.Sid)
}
func OperateNumberOfGood(ReviseNumber *model.ReviseNumber, myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(ReviseNumber.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(myclaims.Uid, int64(ReviseNumber.Sid)); err != nil {
		return err
	}
	return mysql.OperateNumberOfGood(ReviseNumber.Sid, ReviseNumber.NewNumber)
}
func DeleteGoodInShoppingCart(DeleteGood *model.DeleteAGoodFormShoppingCart, myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(DeleteGood.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(myclaims.Uid, int64(DeleteGood.Sid)); err != nil {
		return err
	}
	return mysql.DeleteGoodInShoppingCart(DeleteGood.Sid)
}
func BuyAGoodInShoppingCart(BuyA *model.BuyA, myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(BuyA.Sid); err != nil {
		return err
	}
	return mysql.BuyAGoodInShoppingCart(myclaims.Uid, BuyA.Address, BuyA.PhoneNumber, BuyA.RealName, BuyA.Sid)
}
func BuyAllInShoppingCart(BuyAll *model.BuyAll, myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	if err := mysql.BuyAllInShoppingCart(myclaims.Uid, BuyAll.Address, BuyAll.PhoneNumber, BuyAll.RealName); err != nil {
		return err
	}
	return mysql.ClearShoppingCart(myclaims.Uid)
}
func ClearShoppingCart(myclaims untils.MyClaims) error {
	if err := mysql.CheckUidInShoppingCart(myclaims.Uid); err != nil {
		return err
	}
	return mysql.ClearShoppingCart(myclaims.Uid)
}
func LoveAGood(LoveGood *model.ShoppingCart, myclaims untils.MyClaims) error {
	if err := mysql.CheckGid(int64(LoveGood.Gid)); err != nil {
		return err
	}
	if err := mysql.CheckSize(LoveGood.Size, int64(LoveGood.Gid)); err != nil {
		return err
	}
	if err := mysql.LoveAGood(LoveGood.Gid, myclaims.Uid, LoveGood.Number, LoveGood.Size); err != nil {
		return err
	}
	return mysql.AddPrice(int64(LoveGood.Gid))
}

func SearchGoodInShoppingCart(myclaims untils.MyClaims) (interface{}, error) {
	return mysql.SearchGoodInShoppingCart(myclaims.Uid)
}
func SearchGoodInShoppingCartFromRedis() (interface{}, error) {
	key := "lovelist"
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
