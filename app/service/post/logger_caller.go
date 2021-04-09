package post

import "strings"

func caller(fn string, path ...string) string {
	return "service/post/" + fn + "():" + strings.Join(path, "->")
}
