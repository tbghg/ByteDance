package service

import (
	"ByteDance/cmd/favorite/repository"
	"ByteDance/cmd/video"
	videoRepository "ByteDance/cmd/video/repository"
	"ByteDance/pkg/common"
	"sync"
)

// FavoriteList 点赞列表
func FavoriteList(userId int32) (videoInfo []video.TheVideoInfo, state int) {

	allVideoInfoData, _ := repository.FavoriteDao.FavoriteList(userId)

	videoInfo = make([]video.TheVideoInfo, len(allVideoInfoData))

	wg := sync.WaitGroup{}
	wg.Add(len(allVideoInfoData))

	for index, videoInfoData := range allVideoInfoData {
		go func(index int, videoInfo []video.TheVideoInfo, videoInfoData videoRepository.VideoInfo, userID int32) {
			followerCount, followCount, commentCount, favoriteCount := videoRepository.VideoDao.GetVideoInfo(videoInfoData.UserID, videoInfoData.VideoID)
			isFollow := videoRepository.VideoDao.QueryIsFollow(userID, videoInfoData.UserID)

			videoInfo[index] = video.TheVideoInfo{
				ID: videoInfoData.VideoID,
				Author: video.AuthorInfo{
					ID:            videoInfoData.UserID,
					Name:          videoInfoData.Username,
					FollowCount:   int(followCount),
					FollowerCount: int(followerCount),
					IsFollow:      isFollow,
				},
				PlayURL:       common.OSSPreURL + videoInfoData.PlayURL + ".mp4",
				CoverURL:      common.OSSPreURL + videoInfoData.CoverURL + ".jpg",
				FavoriteCount: int(favoriteCount),
				CommentCount:  int(commentCount),
				IsFavorite:    true,
				Title:         videoInfoData.Title,
			}
			wg.Done()
		}(index, videoInfo, videoInfoData, userId)
	}
	wg.Wait()
	return videoInfo, 1
}
