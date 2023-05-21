package controller

import (
	"Shopping_System/dao/mysql"
	g "Shopping_System/global"
	"Shopping_System/model"
	"Shopping_System/services"
	"Shopping_System/untils"
	"Shopping_System/untils/http"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

func GetGoodDetail(c *gin.Context) {
	GoodDetail := new(model.GetGoodDetail)
	if err := c.ShouldBind(GoodDetail); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, GoodDetail),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if gooddetail, err := services.GetGoodDetail(GoodDetail); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"您查询的Gid为" + strconv.FormatInt(GoodDetail.Gid, 10) + "的商品详情为": gooddetail,
		})
	}
}
func GetAllGoods(c *gin.Context) {
	var get []model.GetAllGoods
	key := "goods"
	sqlstr := "select gid,goodname,price,introduction,size,uid from goods where gid>0;"
	cmd2 := g.RedisClient.ZRevRangeWithScores(context.Background(), key, 0, -1)
	if cmd2.Err() != nil {
		log.Fatalln("2")
	}
	if len(cmd2.Val()) == 0 {
		c.JSON(200, gin.H{
			"message": "在缓存中未找到数据",
		})
	} else {
		var results []interface{}
		for _, z := range cmd2.Val() {
			var data interface{}
			if err := json.Unmarshal([]byte(z.Member.(string)), &data); err != nil {
				log.Fatalln(err)
			}
			results = append(results, data)
		}
		c.JSON(200, gin.H{
			"message":   "该查询结果来自缓存!",
			"以下是您的查询结果": results,
		})
		return
	}
	if err := g.Xdb.Select(&get, sqlstr); err != nil {
		log.Fatalln(err)
	} else {
		ttl := time.Second * 10
		for _, good := range get {
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
		result, err := json.MarshalIndent(interface{}(get), "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                               //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}

func ChangeTheSizeOfGood(c *gin.Context) {
	ChangeSizeGood := new(model.ChangeSizeGood)
	if err := c.ShouldBind(ChangeSizeGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeSizeGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ChangeTheSizeOfGood(ChangeSizeGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品规格成功!",
		})
	}
}
func ChangeTheIntroductionOfGood(c *gin.Context) {
	ChangeIntroductionGood := new(model.ChangeIntroductionGood)
	if err := c.ShouldBind(ChangeIntroductionGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeIntroductionGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ChangeTheIntroductionOfGood(ChangeIntroductionGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品简介成功!",
		})
	}
}
func ChangeThePriceOfGood(c *gin.Context) {
	ChangePriceGood := new(model.ChangePriceGood)
	if err := c.ShouldBind(ChangePriceGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangePriceGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ChangeThePriceOfGood(ChangePriceGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品价格成功!",
		})
	}
}
func ChangeTheNameOfGood(c *gin.Context) {
	ChangeNameGood := new(model.ChangeNameGood)
	if err := c.ShouldBind(ChangeNameGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeNameGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ChangeTheNameOfGood(ChangeNameGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品名称成功!",
		})
	}
}
func DeleteAGood(c *gin.Context) {
	DeleteGood := new(model.GoodsDelete)
	if err := c.ShouldBind(DeleteGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, DeleteGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.DeleteAGood(DeleteGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "删除商品成功!",
		})
	}
}
func SearchMyGood(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if goodlist, err := services.SearchMyGoodsFromRedis(); goodlist == nil || err != nil {
		c.JSON(200, gin.H{
			"message1": "缓存中未找到相关数据！",
			"message2": "以下是您在售商品",
		})
	} else {
		result, err := json.MarshalIndent(goodlist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write([]byte("此数据来自缓存！"))
		c.Writer.Write([]byte("以下是您在售商品"))
		c.Writer.Write(result)
		return
	}
	if goodlist, err := services.SearchMyGoods(claims); err != nil {
		log.Fatalln(err)
	} else {
		result, err := json.MarshalIndent(goodlist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}
func PublishAGood(c *gin.Context) {
	Good := new(model.GoodsPublish)
	if err := c.ShouldBind(Good); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, Good),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.PublishAGood(Good, claims); err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorPassword) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message":         "发布商品成功!",
			"GoodInformation": Good,
		})
	}
}
