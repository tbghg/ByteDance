package utils

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var bucket *oss.Bucket
var once sync.Once

// OSSInit 初始化，将ConnQuery与数据库绑定
func OSSInit() {
	once.Do(func() {
		// 连接OSS账户
		client, err1 := oss.New("http://oss-cn-shanghai.aliyuncs.com", "LTAI5tPgZiMVJY6bNxQFs9MW", "pd1aiM8jh6gNdZCxRscesOJ1Hif3aE")
		if err1 != nil {
			Log.Error("连接OSS账户失败" + err1.Error())
		} else { // OSS账户连接成功
			var err2 error
			// 连接存储空间
			bucket, err2 = client.Bucket("byte-dance-01")
			if err2 != nil {
				Log.Error("连接存储空间失败" + err2.Error())
			} else { // 存储空间连接成功
				Log.Info("OSS初始化完成")
			}
		}
	})
}

func UploadFile(file []byte, filename string, fileType string) bool {
	var fileSuffix string
	if fileType == "video" {
		fileSuffix = ".mp4"
	} else if fileType == "picture" {
		fileSuffix = ".jpg"
	} else {
		Log.Error("无法上传" + fileType + "类型文件")
		return false
	}
	err := bucket.PutObject("video/"+filename+fileSuffix, bytes.NewReader(file))
	if err != nil {
		Log.Error("上传文件失败" + err.Error())
		return false
	} else {
		return true
	}
}
