package web

import "strings"

func caller(fn string, path ...string) string {
	return "controller/default/" + fn + "():" + strings.Join(path, "->")
}
