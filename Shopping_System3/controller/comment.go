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
)

func LookComment(c *gin.Context) {
	Look := new(model.Look)
	if err := c.ShouldBind(Look); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, Look),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if commentList, err := services.LookCommentFromRedis(); commentList == nil || err != nil {
		c.JSON(200, gin.H{
			"message1": "缓存中未找到相关数据！",
			"message2": "以下是该商品的评论:",
		})
	} else {
		result, err := json.MarshalIndent(commentList, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                          //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write([]byte("此数据来自缓存！"))
		c.Writer.Write([]byte("以下是该商品的评论:"))
		c.Writer.Write(result)
		return
	}
	if comments, err := services.LookComment(Look); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.CommentNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		result, err := json.MarshalIndent(comments, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}
func PublishComment(c *gin.Context) {
	PB := new(model.Comments)
	if err := c.ShouldBind(PB); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, PB),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.PublishComment(PB, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInOrders) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorCommentPublishUser) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "评论成功",
		})
	}
}
