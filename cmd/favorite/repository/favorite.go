package repository

import (
	videoRepository "ByteDance/cmd/video/repository"
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

type FavoriteStruct struct {
}

var (
	FavoriteDao  *FavoriteStruct
	favoriteOnce sync.Once
)

// 单例模式
func init() {
	favoriteOnce.Do(func() {
		FavoriteDao = &FavoriteStruct{}
	})
}

// FavoriteAction 更新点赞操作
func (*FavoriteStruct) FavoriteAction(userId int32, videoId int32, actionType int32) (success bool) {
	f := dal.ConnQuery.Favorite
	var removed int32
	if actionType == 1 {
		removed = -1
	} else {
		removed = 1
	}

	count, _ := f.Where(f.UserID.Eq(userId), f.VideoID.Eq(videoId), f.Deleted.Eq(0)).Count()

	if count == 0 {
		// 不存在相关记录，需要进行创建
		favorite := &model.Favorite{UserID: userId, VideoID: videoId}
		err := f.Create(favorite)
		if !utils.CatchErr("添加favorite错误", err) {
			return false
		}
	} else {
		_, err := f.Where(f.UserID.Eq(userId), f.VideoID.Eq(videoId), f.Deleted.Eq(0)).Update(f.Removed, removed)
		if err != nil {
			return false
		}
	}
	return true
}

// FavoriteList 点赞列表
func (*FavoriteStruct) FavoriteList(userId int32) ([]videoRepository.VideoInfo, bool) {
	v := dal.ConnQuery.Video
	u := dal.ConnQuery.User
	f := dal.ConnQuery.Favorite

	var result []videoRepository.VideoInfo
	// 内联查询
	err := f.Select(u.ID.As("UserID"), u.Username, v.ID.As("VideoID"), v.PlayURL, v.CoverURL, v.Time, v.Title).Where(f.UserID.Eq(userId), f.Removed.Eq(-1), f.Deleted.Eq(0)).Join(v, v.ID.EqCol(f.VideoID)).Join(u, u.ID.EqCol(v.AuthorID)).Scan(&result)
	utils.CatchErr("获取视频信息错误", err)
	if result == nil {
		return nil, false
	}
	return result, true
}
