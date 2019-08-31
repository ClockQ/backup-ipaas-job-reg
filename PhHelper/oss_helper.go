package PhHelper

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type PhOssHelper struct {
	Oss             *oss.Client
	endpoint        string
	accessKeyId     string
	accessKeySecret string
}

func (helper PhOssHelper) New(endpoint, accessKeyId, accessKeySecret string) *PhOssHelper {
	helper.endpoint = endpoint
	helper.accessKeyId = accessKeyId
	helper.accessKeySecret = accessKeySecret

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	helper.Oss = client

	return &helper
}

func (helper PhOssHelper) GetObject(bucketName, objectKey string, options ...oss.Option) (io.ReadCloser, error) {
	bucket, err := helper.Oss.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket.GetObject(objectKey)
}
