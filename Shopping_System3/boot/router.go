package boot

import (
	"Shopping_System/controller"
	"Shopping_System/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouters() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/registration", controller.Register)                                                       //Register
		user.POST("/login", controller.Login)                                                                 //Login
		user.POST("/password/forget", controller.ForgetPassword)                                              //ForgetPassword,把密保问题写出来
		user.POST("/username/revise", middleware.JWTAuthMiddleware(), controller.ReviseUsername, homeHandler) //ReviseUsername
		user.POST("/password/revise", middleware.JWTAuthMiddleware(), controller.RevisePassword, homeHandler) //RevisePassword
		user.POST("/sex/revise", middleware.JWTAuthMiddleware(), controller.ReviseSex, homeHandler)           //ReviseSex
		user.POST("/securityquestion/forget", controller.GetSecurityQuestion)
		user.GET("/deregistration", middleware.JWTAuthMiddleware(), controller.DisRegister, homeHandler) //DisRegister
		goods := r.Group("/user/goods")
		{
			//买家用户的商品操作
			goods.GET("/get", controller.GetAllGoods)                                                           //GetAllGoods
			goods.POST("/search", controller.GetGoodDetail)                                                     //GetGoodDetail
			goods.POST("/like", middleware.JWTAuthMiddleware(), controller.LikeAGood, homeHandler)              //LikeAGood
			goods.GET("/like/search", middleware.JWTAuthMiddleware(), controller.SearchGoodInLike, homeHandler) //SearchGoodInLike1111111111111111111111
			//商家用户的商品操作
			goods.POST("/publish", middleware.JWTAuthMiddleware(), controller.PublishAGood, homeHandler)     //PublishAGood
			goods.GET("/searchmygood", middleware.JWTAuthMiddleware(), controller.SearchMyGood, homeHandler) //SearchMyGood
			goods.POST("/delete", middleware.JWTAuthMiddleware(), controller.DeleteAGood, homeHandler)       //DeleteAGood
			operation := r.Group("/user/goods/operation")
			{
				operation.POST("/goodname", middleware.JWTAuthMiddleware(), controller.ChangeTheNameOfGood, homeHandler)             //ChangeTheNameOfGood
				operation.POST("/price", middleware.JWTAuthMiddleware(), controller.ChangeThePriceOfGood, homeHandler)               //ChangeThePriceOfGood
				operation.POST("/introduction", middleware.JWTAuthMiddleware(), controller.ChangeTheIntroductionOfGood, homeHandler) //ChangeTheIntroductionOfGood
				operation.POST("/size", middleware.JWTAuthMiddleware(), controller.ChangeTheSizeOfGood, homeHandler)                 //ChangeTheSizeOfGood
			}
		}
		ShoppingCart := r.Group("/user/shoppingcart")
		{
			//用户对购物车的操作
			ShoppingCart.POST("/love", middleware.JWTAuthMiddleware(), controller.LoveAGood, homeHandler)                  //LoveAGood
			ShoppingCart.GET("/search", middleware.JWTAuthMiddleware(), controller.SearchGoodInShoppingCart, homeHandler)  //SearchGoodInShoppingCart
			ShoppingCart.GET("/clear", middleware.JWTAuthMiddleware(), controller.ClearShoppingCart, homeHandler)          //ClearShoppingCart
			ShoppingCart.POST("/allbuy", middleware.JWTAuthMiddleware(), controller.BuyAllInShoppingCart, homeHandler)     //BuyAllInShoppingCart
			ShoppingCart.POST("/buyagood", middleware.JWTAuthMiddleware(), controller.BuyAGoodInShoppingCart, homeHandler) //
			operation := r.Group("user/shoppingcart/operation")
			{
				operation.POST("/operatenum", middleware.JWTAuthMiddleware(), controller.OperateNumberOfGood, homeHandler)      //OperateNumberOfGood
				operation.POST("/operatesize", middleware.JWTAuthMiddleware(), controller.OperateSizeOfGood, homeHandler)       //OperateSizeOfGood
				operation.POST("/deletegood", middleware.JWTAuthMiddleware(), controller.DeleteGoodInShoppingCart, homeHandler) //DeleteGoodInShoppingCart
			}
		}
		Comment := r.Group("/user/comment")
		{
			Comment.POST("/publish", middleware.JWTAuthMiddleware(), controller.PublishComment, homeHandler) //PublishComment
			Comment.POST("/look", controller.LookComment)
		}
		user.GET("/checkorder", middleware.JWTAuthMiddleware(), controller.CheckOrder, homeHandler) //CheckOrder
	}
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Router initialization successful")
}

func homeHandler(c *gin.Context) {
	uid := c.MustGet("uid").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"uid": uid},
	})
}
