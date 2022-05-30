package controller

import (
	"ByteDance/cmd/video"
	"ByteDance/pkg/common"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// 获取视频返回值
type getVideoResponse struct {
	common.Response
	NextTime  int64           `json:"next_time"`
	VideoList [10]video.Video `json:"video_list"`
}

// 显示视频的接口
// 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

// 第一次访问的时候没有 latest_time 需要设置未当前时间戳，可以参考Demo的写法，token不重要

func GetVideoFeed(c *gin.Context) {
	lastTime, _ := strconv.ParseInt(c.Query("last_time"), 10, 64)
	// 获取视频信息不需要token
	if lastTime == 0 {
		lastTime = time.Now().Unix()
	}
	// 需要获取NextTime、VideoList

}
