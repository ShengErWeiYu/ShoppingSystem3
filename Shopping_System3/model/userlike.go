package model

type UserLike struct {
	Gid int64 `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
}

type MyLike struct {
	Lid int
	Gid int
}
