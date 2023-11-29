package util

import (
	"math/rand"
	"regexp"
	"time"
)

const (
	VERIFY_EXP_EMAIL = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
)

// VerifyFormat 正则验证
func VerifyFormat(exp, str string) bool {
	reg := regexp.MustCompile(exp)
	return reg.MatchString(str)
}

// VerifyEmail 验证邮箱
func VerifyEmail(email string) bool {
	return VerifyFormat(VERIFY_EXP_EMAIL, email)
}

// RandomString 随机n长度的字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLXCVBNMQWERTYUIOP")
	result := make([]byte, n)
	rand.New(rand.NewSource(time.Now().Unix()))
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
