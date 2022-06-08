package controller

import (
	"ByteDance/cmd/video"
	"ByteDance/cmd/video/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

// 获取视频返回值
type getVideoResponse struct {
	common.Response
	NextTime  int64                `json:"next_time"`
	VideoList []video.TheVideoInfo `json:"video_list"`
}

// 获取视频返回值
type videoListResponse struct {
	common.Response
	VideoList []video.TheVideoInfo `json:"video_list"`
}

// GetVideoFeed 获取视频流信息
func GetVideoFeed(c *gin.Context) {
	lastTime, _ := strconv.ParseInt(c.Query("last_time"), 10, 32)
	// 获取视频信息不需要token
	if lastTime == 0 {
		lastTime = time.Now().Unix()
	}
	// 需要获取NextTime、VideoList
	nextTime, videoInfo, state := service.GetVideoFeed(lastTime)
	if state == 0 {
		c.JSON(http.StatusOK, &getVideoResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.HasNoVideoMsg,
			}, NextTime: lastTime,
		})
	} else if state == 1 {
		c.JSON(http.StatusOK, &getVideoResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.GetVideoInfoSuccessMsg,
			}, NextTime: nextTime,
			VideoList: videoInfo,
		})
	}
}

func PublishVideo(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	userID, success := c.Get("user_id")
	if !success {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.TokenParameterAcquisitionError,
			})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.PublishVideoFailedMsg,
			})
		return
	}

	fileHandle, err1 := data.Open() //打开上传文件
	utils.CatchErr("打开文件失败", err1)

	// 闭包处理错误
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		utils.CatchErr("关闭文件错误", err)
	}(fileHandle)

	fileByte, err2 := ioutil.ReadAll(fileHandle)
	utils.CatchErr("读取文件错误", err2)

	if service.PublishVideo(userID.(int), title, fileByte) {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: 0,
				StatusMsg:  msg.PublishVideoSuccessMsg,
			})
	} else {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.PublishVideoFailedMsg,
			})
	}
}

// PublicList 登录用户的视频发布列表，直接列出用户所有投稿过的视频
func PublicList(c *gin.Context) {

	userID, success := c.Get("user_id")
	if !success {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.TokenParameterAcquisitionError,
			})
		return
	}
	videoInfo, success2 := service.PublishList(userID.(int))
	if success2 {
		c.JSON(http.StatusOK, &videoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.GetPublishListSuccessMsg,
			},
			VideoList: videoInfo,
		})
	} else {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.GetPublishListFailedMsg,
			})
	}
}
