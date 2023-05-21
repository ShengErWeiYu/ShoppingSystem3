package model

type ShoppingCart struct {
	Gid    int    `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	Number int    `form:"number" json:"number" binding:"min=1" msg:"商品数量不能为0及以下数字"`
	Size   string `form:"size" json:"size" binding:"required" msg:"商品规格不能为空"`
	Price  string //直接调goods表的Price过来
}
type ReviseSize struct {
	Sid     int    `form:"sid" json:"sid" binding:"required" msg:"购物车编号不能为空"`
	NewSize string `form:"newsize" json:"newsize" binding:"required" msg:"商品规格不能为空"`
}
type ReviseNumber struct {
	Sid       int `form:"sid" json:"sid" binding:"required" msg:"购物车编号不能为空"`
	NewNumber int `form:"newnumber" json:"newnumber" binding:"min=1" msg:"商品数量不能为0及以下数字"`
}
type BuyA struct {
	Address     string `form:"address" json:"address" binding:"required" msg:"收货地址不能为空"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required" msg:"电话号码不能为空"`
	RealName    string `form:"realname" json:"realname" binding:"required" msg:"请填写你的真实姓名"`
	Sid         int    `form:"sid" json:"sid" binding:"required" msg:"购物车编号不能为空"`
}
type BuyAll struct {
	Address     string `form:"address" json:"address" binding:"required" msg:"收货地址不能为空"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required" msg:"电话号码不能为空"`
	RealName    string `form:"realname" json:"realname" binding:"required" msg:"请填写你的真实姓名"`
}
type SearchShoppingCartResult struct {
	Sid    int
	Gid    int
	Number int
	Size   string
	Price  string
}
type Count struct {
	Count int
}
type GoodGids struct {
	Gid    int64
	Number int
	Size   string
}
type DeleteAGoodFormShoppingCart struct {
	Sid int `form:"sid" json:"sid" binding:"required" msg:"购物车编号不能为空"`
}
