package service

import (
	"ByteDance/cmd/video"
	"ByteDance/cmd/video/repository"
	"ByteDance/dal/method"
	"ByteDance/pkg/common"
	"ByteDance/utils"
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func GetVideoFeed(lastTime int64) (nextTime int64, videoInfo []video.TheVideoInfo, state int) {
	// state 0:已经没有视频了  1:获取成功  -1:获取失败
	stringLastTime := time.Unix(lastTime, 0).Format("2006-01-02 15:04:05")
	//formatLastTime, _ := time.Parse("2006-01-02 15:04:05", stringLastTime)
	formatLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", stringLastTime, time.Local)

	allVideoInfoData, isExist := repository.VideoDao.GetVideoFeed(formatLastTime)
	if !isExist {
		// 已经没有视频了
		return nextTime, videoInfo, 0
	}

	nextTime = allVideoInfoData[len(allVideoInfoData)-1].Time.Unix()
	videoInfo = make([]video.TheVideoInfo, len(allVideoInfoData))

	for index, videoInfoData := range allVideoInfoData {
		followerCount, followCount, _ := method.QueryFollowCount(videoInfoData.UserID)
		commentCount := method.QueryCommentCountByVideoID(videoInfoData.VideoID)
		favoriteCount := method.QueryFavoriteCount(videoInfoData.UserID)

		videoInfo[index] = video.TheVideoInfo{
			ID: videoInfoData.VideoID,
			Author: video.AuthorInfo{
				ID:            videoInfoData.UserID,
				Name:          videoInfoData.Username,
				FollowCount:   int(followCount),
				FollowerCount: int(followerCount),
				IsFollow:      false,
			},
			PlayURL:       common.OSSPreURL + videoInfoData.PlayURL + ".mp4",
			CoverURL:      common.OSSPreURL + videoInfoData.CoverURL + ".jpg",
			FavoriteCount: int(favoriteCount),
			CommentCount:  int(commentCount),
			IsFavorite:    false,
			Title:         videoInfoData.Title,
		}
	}

	return nextTime, videoInfo, 1
}

func PublishVideo(userID int, title string, fileBytes []byte) bool {
	node, _ := utils.NewWorker(1)
	randomId := node.GetId()
	fileID := fmt.Sprintf("%v", randomId)

	if !utils.UploadFile(fileBytes, fileID, "video") {
		return false
	}
	// 通过ffmpeg截取视频第一帧为视频封面
	videoURL := common.OSSPreURL + fileID + ".mp4"
	cmd := exec.Command("ffmpeg", "-i", videoURL, "-vframes", "1", "-f", "singlejpeg", "-")
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	err := cmd.Run()
	if !utils.CatchErr("ffmpeg运行错误", err) {
		return false
	}
	// 将视频封面上传至OSS中
	if !utils.UploadFile(buf.Bytes(), fileID, "picture") {
		return false
	}
	// 存入数据库中
	success := repository.VideoDao.PublishVideo(userID, title, fileID)
	if success {
		return true
	} else {
		return false
	}

}
