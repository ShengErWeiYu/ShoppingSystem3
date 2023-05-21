package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/global"
	"Shopping_System/model"
	"Shopping_System/untils"
	"context"
	"log"
	"time"
)

func DisRegister(myclaims untils.MyClaims) error {
	return mysql.DisRegister(myclaims.Uid)
}
func GetSecurityQuestionFromRedis(GetSecurityQuestion *model.GetSecurityQuestion) (interface{}, error) {
	redisKey := GetSecurityQuestion.Uid
	value, err := global.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}
func GetSecurityQuestion(GetSecurityQuestion *model.GetSecurityQuestion) (interface{}, error, error) {
	redisKey := GetSecurityQuestion.Uid
	if err1 := mysql.CheckUid2(GetSecurityQuestion.Uid); err1 != nil {

		return nil, err1, nil //err2返回nil，所以猜controller想着&&的条件为err2==nil
	}
	if err1 := mysql.CheckUid2(GetSecurityQuestion.Uid); err1 == nil {
		if SecurityQuestion, err2 := mysql.GetSecurityQuestion(GetSecurityQuestion.Uid); err2 == nil {
			ttl := time.Hour * 24
			if err := global.RedisClient.Set(context.Background(), redisKey, SecurityQuestion, ttl).Err(); err != nil {
				log.Fatalln(err)
			}
			return SecurityQuestion, nil, nil //err2返回nil
		}
	}
	return nil, nil, nil //err2返回nil
}
func GetUsernameByUid(uid string) string {
	username := mysql.GetUsernameByUid(uid)
	return username
}
func ReviseSex(ReviseSexUser *model.ReviseSexUser, myclaims untils.MyClaims) error {

	return mysql.ReviseSex(myclaims.Uid, ReviseSexUser.NewSex)
}
func RevisePassword(RevisePasswordUser *model.RevisePasswordUser, myclaims untils.MyClaims) error {

	if err := mysql.CheckAnswer(myclaims.Uid, RevisePasswordUser.Answer); err != nil {
		return err
	}
	return mysql.RevisePassword(myclaims.Uid, RevisePasswordUser.NewPassword)
}
func ReviseUsername(ReviseNameUser *model.ReviseNameUser, myclaims untils.MyClaims) error {
	return mysql.ReviseUsername(myclaims.Uid, ReviseNameUser.NewUsername)
}
func ForgetPassword(FgtPswUser *model.FgtPswUser) error {
	if err := mysql.CheckUid2(FgtPswUser.Uid); err != nil {
		return err
	}
	if err := mysql.CheckAnswer(FgtPswUser.Uid, FgtPswUser.Answer); err != nil {
		return err
	}
	if err := mysql.ResetPassword(FgtPswUser.Uid, FgtPswUser.RePassword); err != nil {
		return err
	}
	return nil
}

func Register(User *model.User) error {
	if err := mysql.CheckUsername(User.Username); err != nil {
		return err
	}
	if err := mysql.CheckUid1(User.Uid); err != nil {
		return err
	}

	user := &model.User{
		Uid:              User.Uid,
		Username:         User.Username,
		Password:         untils.ComplexPassword(User.Password),
		Sex:              User.Sex,
		SecurityQuestion: User.SecurityQuestion,
		Answer:           User.Answer,
	}
	return mysql.AddUser(user)
}

func Login(LoginUser *model.LoginUser) error {
	if err := mysql.CheckUid2(LoginUser.Uid); err != nil { //错误UID
		return err
	}
	if err := mysql.CheckPassword(LoginUser.Password, LoginUser.Uid); err != nil { //正确UID正确密码
		return err
	}
	return nil
}
