package PhThirdHelper

import (
	"github.com/PharbersDeveloper/ipaas-job-reg/PhEnv"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestPhOssHelper(t *testing.T) {
	PhEnv.SetEnv()

	endpoint := os.Getenv("ENDPOINT")
	accessKeyId := os.Getenv("ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ACCESS_KEY_SECRET")
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
