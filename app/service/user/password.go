package user

import (
	"github.com/go-whisper/go-whisper/app/utils"
	"strings"
)

// Encrypt 加密密码
func Encrypt(pwd string) string {
	salt := utils.RandStr(16)
	return salt + "." + utils.Sha256([]byte(salt+pwd))
}

// Verify 校验密码 `pwd` 是不是明文 `plaintext` 的密码
func Verify(plaintext, pwd string) bool {
	s := strings.Split(pwd, ".")
	if len(s) != 2 {
		return false
	}
	return s[1] == utils.Sha256([]byte(s[0]+plaintext))
}
