package utils

import "github.com/go-whisper/go-whisper/app/utils/randstring"

func RandStr(n int) string {
	return randstring.RandString(n)
}
