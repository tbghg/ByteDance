package service

import (
	"ByteDance/cmd/favorite/repository"
	"ByteDance/cmd/video"
	"ByteDance/dal/method"
	"ByteDance/pkg/common"
)

//点赞列表
func FavoriteList(userId int32) (videoInfo []video.TheVideoInfo, state int) {

	allVideoInfoData, _ := repository.FavoriteDao.FavoriteList(userId)

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

	return videoInfo, 1
}
