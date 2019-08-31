package PhHelper

import (
	"bytes"
	"testing"
)

func TestOss(t *testing.T) {
	endpoint := "oss-cn-beijing.aliyuncs.com"
	accessKeyId := "LTAIEoXgk4DOHDGi"
	accessKeySecret := "x75sK6191dPGiu9wBMtKE6YcBBh8EI"
	bucketName := "pharbers-resources"
	objectKey := "TMCal.json"

	r, err := PhOssHelper{}.New(endpoint, accessKeyId, accessKeySecret).GetObject(bucketName, objectKey)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r)
	newStr := buf.String()

	println(newStr)
}
