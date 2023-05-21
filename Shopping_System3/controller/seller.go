package controller

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/services"
	"Shopping_System/untils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckOrder(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if orderlist, err := services.CheckOrderFromRedis(); orderlist == nil || err != nil {
		c.JSON(200, gin.H{
			"message1": "缓存中未找到相关数据！",
			"message2": "以下是您的订单列表:",
		})
	} else {
		result, err := json.MarshalIndent(orderlist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                        //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write([]byte("此数据来自缓存！"))
		c.Writer.Write([]byte("以下是您的订单列表:"))
		c.Writer.Write(result)
		return
	}
	if orderlist, err := services.CheckOrder(claims); err != nil {
		if errors.Is(err, mysql.ErrorOrderExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		result, err := json.MarshalIndent(orderlist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                        //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}
