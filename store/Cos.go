package store

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"io/ioutil"
	"log"
	"os"
	"shared-clipboard/runcommand"
)

type CosStore struct {
	cosClient cos.Client
	target    *runcommand.Target
}

func NewCosStore(client cos.Client, target *runcommand.Target) *CosStore {
	return &CosStore{
		cosClient: client,
		target:    target,
	}
}

func (C *CosStore) Write(content io.Reader) error {
	ok, err := C.cosClient.Bucket.IsExist(context.Background())
	if err != nil {
		return err
	}
	if ok == false {
		opt := &cos.BucketPutOptions{
			XCosACL: "private",
		}
		_, err := C.cosClient.Bucket.Put(context.Background(), opt)
		if err != nil {
			return err
		}
	}
	name := "default-bucket"
	// 1.通过字符串上传对象
	_, err = C.cosClient.Object.Put(context.Background(), name, content, nil)
	if err != nil {
		return err
	}
	return nil
}

func (C *CosStore) Read() (string, error) {
	tempFile, err := ioutil.TempFile("/tmp", "tempCosObjectFile")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()
	defer func() {
		err := os.Remove(tempFile.Name())
		if err != nil {
			log.Fatal(err)
		}
	}()

	name := "default-bucket"
	resp, err := C.cosClient.Object.Get(context.Background(), name, nil)
	if err != nil {
		return "", err
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(bs), nil
}
