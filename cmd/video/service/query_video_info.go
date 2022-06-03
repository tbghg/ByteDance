package service

import (
	"ByteDance/cmd/video"
	"ByteDance/cmd/video/repository"
	"ByteDance/dal/method"
	"ByteDance/pkg/common"
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
			PlayURL:       common.OSSPreURL + videoInfoData.PlayURL,
			CoverURL:      common.OSSPreURL + videoInfoData.CoverURL,
			FavoriteCount: int(favoriteCount),
			CommentCount:  int(commentCount),
			IsFavorite:    false,
			Title:         videoInfoData.Title,
		}
	}

	return nextTime, videoInfo, 1
}
