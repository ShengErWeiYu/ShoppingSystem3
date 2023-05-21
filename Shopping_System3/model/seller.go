package model

// 这里是订单,是买家下订单的地方，也是卖家查看订单的地方

type CheckOrderResult struct {
	Oid         int
	Uid         string
	Gid         int64
	Number      int
	Size        string
	Address     string
	PhoneNumber string
	RealName    string
}
