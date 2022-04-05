package store

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"shared-clipboard/runcommand"
)

type Store interface {
	Write(content io.Reader) error
	Read() (string, error)
}

func NewStore(target *runcommand.Target) (Store, error) {
	client, err := NewClient(target)
	if err != nil {
		return nil, fmt.Errorf("Failed initializing client: %s", err.Error())
	}
	return NewCosStore(*client, target), nil
}

func NewClient(target *runcommand.Target) (*cos.Client, error) {
	secretId := target.SecretId
	secretKey := target.SecretKey

	u, _ := url.Parse("https://default-bucket-1306768814.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: secretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			//AKIDf9OO9ette9ON6apoY3IH0UihQt75nFTS
			//lHYZLr78tOeeOt88tviGZcPwEG2YqsfC
		},
	})

	_, err := c.Bucket.Put(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}
