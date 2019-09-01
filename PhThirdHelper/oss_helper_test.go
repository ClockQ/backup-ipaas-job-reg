package PhThirdHelper

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPhOssHelper(t *testing.T) {
	endpoint := "oss-cn-beijing.aliyuncs.com"
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := "pharbers-resources"
	objectKey := "TMCal.json"

	oh := PhOssHelper{}.New(endpoint, accessKeyId, accessKeySecret)
	Convey("测试读取 OSS 文件", t, func() {
		str, err := oh.GetObject(bucketName, objectKey)
		if err != nil {
			panic(err)
		}

		So(str, ShouldNotBeNil)
	})
}
