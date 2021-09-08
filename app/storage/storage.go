package storage

import (
	"bytes"
	"context"
	"github.com/go-whisper/go-whisper/app/instance"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"time"
)

var _cli *cos.Client

// cli() 返回 *cos.Client
// 注意：_cli 不是并发安全的
func cli() *cos.Client {
	if _cli == nil {
		Init()
	}
	return _cli
}

func Init() {
	u, _ := url.Parse(viper.GetString("tencentyun.cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	_cli = cos.NewClient(b, &http.Client{
		Timeout: 30 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("tencentyun.cos.secretID"),
			SecretKey: viper.GetString("tencentyun.cos.secretKey"),
		},
	})
}

func Put(key string, f io.Reader) error {
	if _, err := cli().Object.Put(context.Background(), key, f, nil); err != nil {
		instance.Logger().Error("storage.Put(): cli.Object.Put() fail:", zap.Error(err))
		return err
	}
	return nil
}

func PutFromFile(key, filePath string) error {
	if _, err := cli().Object.PutFromFile(context.Background(), key, filePath, nil); err != nil {
		instance.Logger().Error("storage.PutFromFile(): cli.Object.PutFromFile() fail:", zap.Error(err))
		return err
	}
	return nil
}

func Get(key string, w io.Writer) error {
	resp, err := cli().Object.Get(context.Background(), key, nil)
	if err != nil {
		instance.Logger().Error("storage.Get(): cli.Object.Get() fail:", zap.Error(err))
		return err
	}
	defer resp.Body.Close()
	var b bytes.Buffer
	if _, err = b.ReadFrom(resp.Body); err != nil {
		instance.Logger().Error("storage.Get(): b.ReadFrom() fail:", zap.Error(err))
		return err
	}
	if _, err = b.WriteTo(w); err != nil {
		instance.Logger().Error("storage.Get(): b.WriteTo() fail:", zap.Error(err))
		return err
	}
	return nil
}
