package service

import (
	"ByteDance/cmd/favorite/repository"
	"ByteDance/cmd/video"
	"ByteDance/dal/method"
	"ByteDance/pkg/common"
	"ByteDance/utils"
	"fmt"
)

type FavoriteRequest struct {
	UserId     int64
	Token      string
	VideoId    int64
	ActionType int32
}

//点赞操作
func RelationAction(userId int32, videoId int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.FavoriteDao.RelationUpdate(userId, videoId)

	if IsExist == 0 {
		//添加该数据
		err = repository.FavoriteDao.RelationCreate(userId, videoId)
		utils.CatchErr("添加失败", err)
	}

	return err
}

//点赞列表
func RelationList(userId int32) (videoInfo []video.TheVideoInfo, state int) {

	allVideoInfoData, _ := repository.FavoriteDao.RelationSelect(userId)
	fmt.Println(".....................................")
	fmt.Println(allVideoInfoData)

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
