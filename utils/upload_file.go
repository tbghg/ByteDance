package utils

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var bucket *oss.Bucket
var once sync.Once

// 初始化，将ConnQuery与数据库绑定
func init() {
	once.Do(func() {
		// 连接OSS账户
		client, err := oss.New("http://oss-cn-shanghai.aliyuncs.com", "LTAI5tPgZiMVJY6bNxQFs9MW", "pd1aiM8jh6gNdZCxRscesOJ1Hif3aE")
		CatchErr("连接OSS账户失败", err)

		// 连接存储空间
		bucket, err = client.Bucket("byte-dance-01")
		CatchErr("连接存储空间失败", err)
	})
}

func UploadFile(file []byte, filename string, fileType string) bool {
	var fileSuffix string
	if fileType == "video" {
		fileSuffix = ".mp4"
	} else if fileType == "picture" {
		fileSuffix = ".jpg"
	} else {
		return false
	}
	err := bucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	CatchErr("上传文件失败", err)
	if err != nil {
		return false
	} else {
		return true
	}
}
