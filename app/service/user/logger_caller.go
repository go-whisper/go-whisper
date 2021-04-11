package user

import "strings"

func caller(fn string, path ...string) string {
	return "service/user/" + fn + "():" + strings.Join(path, "->")
}
