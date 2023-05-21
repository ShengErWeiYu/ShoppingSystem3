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

func DeleteGoodInShoppingCart(c *gin.Context) {
	DeleteGood := new(model.DeleteAGoodFormShoppingCart)
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
	if err := services.DeleteGoodInShoppingCart(DeleteGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
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
			"message": "您已成功从购物车删除本商品",
		})
	}
}

func BuyAGoodInShoppingCart(c *gin.Context) {
	BuyA := new(model.BuyA)
	if err := c.ShouldBind(BuyA); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, BuyA),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.BuyAGoodInShoppingCart(BuyA, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您已成功购买本商品",
		})
	}
}
func OperateSizeOfGood(c *gin.Context) {
	ReviseSize := new(model.ReviseSize)
	if err := c.ShouldBind(ReviseSize); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseSize),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.OperateSizeOfGood(ReviseSize, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorSizeExist) {
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
			"message": "已修改您购物车内选中的商品的规格",
		})
	}
}
func OperateNumberOfGood(c *gin.Context) {
	ReviseNumber := new(model.ReviseNumber)
	if err := c.ShouldBind(ReviseNumber); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ReviseNumber),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.OperateNumberOfGood(ReviseNumber, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorThisGoodInShoppingCart) {
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
			"message": "已修改您购物车内选中的商品的数量",
		})
	}
}
func BuyAllInShoppingCart(c *gin.Context) {
	BuyAll := new(model.BuyAll)
	if err := c.ShouldBind(BuyAll); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, BuyAll),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.BuyAllInShoppingCart(BuyAll, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您购物车内的商品已全部购买",
		})
	}
}
func ClearShoppingCart(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.ClearShoppingCart(claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExistInShoppingCart) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "您的购物车已清空",
		})
	}
}
func LoveAGood(c *gin.Context) {
	LoveGood := new(model.ShoppingCart)
	if err := c.ShouldBind(LoveGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, LoveGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if err := services.LoveAGood(LoveGood, claims); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorSizeExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message1": "已将Gid为" + strconv.FormatInt(int64(LoveGood.Gid), 10) + "的商品添加到您的收藏中",
			"message2": "如若想查看商品详情，请移步用户商品查询页:http://127.0.0.1:8080/user/shoppingcart/search",
		})
	}
}

func SearchGoodInShoppingCart(c *gin.Context) {
	uid := c.GetString("uid")
	claims := untils.MyClaims{Uid: uid}
	if lovelist, err := services.SearchGoodInShoppingCartFromRedis(); lovelist == nil || err != nil {
		c.JSON(200, gin.H{
			"message1": "缓存中未找到相关数据！",
			"message2": "以下是您购物车内的商品",
		})
	} else {
		result, err := json.MarshalIndent(lovelist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write([]byte("此数据来自缓存！"))
		c.Writer.Write([]byte("以下是您购物车内的商品"))
		c.Writer.Write(result)
		return
	}
	if lovelist, err := services.SearchGoodInShoppingCart(claims); err != nil {
		log.Fatalln(err)
	} else {
		result, err := json.MarshalIndent(lovelist, "", "  ") //这是因为在输出数组时，c.JSON 方法会将所有元素紧密地排列在一起，而在 Redis 缓存中，每个元素都是独立的 key-value 对形式，所以在输出缓存值时，每个元素都是独占一行的，更加整齐。
		if err != nil {                                       //json.MarshalIndent 方法会将 get 对象序列化为一个 JSON 字符串，并且通过第二个参数设置每一行前面需要添加的缩进字符串（这里我们设置为空格），最后一个参数则指定不同层级之间的分隔符（这里我们设置为两个空格）。这样输出的 JSON 字符串就会更加整齐、易于阅读。
			log.Fatalln(err)
		}
		c.Writer.Write(result)
	}
}
