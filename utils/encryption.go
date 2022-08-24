package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math/rand"
	"time"
)

// CreateSalt 生成密码盐
func CreateSalt(n int) string {
	rand.Seed(time.Now().UnixMicro())
	var b = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var length = len(b)
	var result = make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = b[rand.Intn(length)]
	}
	return string(result)
}

func SessionId(length int) string {
	if length == 0 {
		length = 32
	}
	var b = make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// CreateMd5 生成Md5加密
func CreateMd5(str string) string {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// GenPwd 加密 使用 bcrypt 对密码进行加密
func GenPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

// ComparePwd 解密 对比明文密码和数据库的哈希值
// hasPwd 加密过的密码（数据库存储的密码）
// pwd 用户输入的密码
func ComparePwd(hasPwd, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hasPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
