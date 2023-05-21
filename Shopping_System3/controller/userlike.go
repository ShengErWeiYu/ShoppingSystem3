package controller

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/services"
	"Shopping_System/untils"
	"Shopping_System/untils/http"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func SearchGoodInLike(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if likelist, err := services.SearchGoodInLikeFromRedis(); likelist == nil || err != nil {
		c.JSON(200, gin.H{
			"message": "缓存中未找到相关数据！",
		})
	} else {
		result, err := json.MarshalIndent(likelist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write([]byte("此数据来自缓存！"))
		c.Writer.Write([]byte("以下是您收藏中的商品:"))
		c.Writer.Write(result)
		return
	}
	if likelist, err := services.SearchGoodInLike(claims); err != nil {
		log.Fatalln(err)
	} else {
		result, err := json.MarshalIndent(likelist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}
func LikeAGood(c *gin.Context) {
	LikeGood := new(model.UserLike)
	if err := c.ShouldBind(LikeGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, LikeGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.LikeAGood(LikeGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message1": "已将Gid为" + strconv.FormatInt(LikeGood.Gid, 10) + "的商品添加到您的收藏中",
			"message2": "如若想查看商品详情，请移步用户商品查询页:http://127.0.0.1:8080/user/goods/search",
		})
	}
}
