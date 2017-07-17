package utils

//短信相关模块
import (
	"github.com/astaxie/beego/logs"
	"strconv"
)

//发送普通短信
func SendCommonSms(phone, context string) bool {
	return true

}

//发送短信验证码
func SendValidCode(phone, context string) bool {
	logs.Info("begin sends:%s to phone:%s", context, phone)
	validCode := "888888"

	key := ValidCodeKey + phone
	RedisSet(key, validCode, false, "ex", strconv.FormatInt(60*2, 10))
	return true
}

func CheckValidCode(phone, validCode string) bool {
	key := ValidCodeKey + phone

	code := redisGetJsonStr(key)
	logs.Debug("redis.code:%s,validCode:%s", code, validCode)
	return validCode == code
}
