package model

// 在评论前检查该用户Uid是否购买了这个Gid的商品，去订单表里查
type Comments struct {
	Gid     int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	Star    int    `form:"star" json:"star" binding:"min=1,max=5" msg:"评分星级在1-5之间"`
	Comment string `form:"comment" json:"comment" binding:"required" msg:"评论不能为空"`
}
type Look struct {
	Gid int64 `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
}
type LookResult struct {
	Cid     int
	Gid     int64
	Star    int
	Comment string
	Uid     string
}
