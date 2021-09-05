package storage

import (
	"bytes"
	"os"
	"testing"
)

func TestPut(t *testing.T) {
	key := "test.txt"
	b := bytes.NewBufferString("this-is-test")
	if err := Put(key, b); err != nil {
		t.Log("error:", err)
		t.Fail()
	}

	f, err := os.Create("./TEST-FILE.txt")
	if err != nil {
		t.Log("os.Create() fail:", err)
		t.Fail()
	}
	if err := Get(key, f); err != nil {
		t.Log("Get() fail:", err)
		t.Fail()
	}
}
