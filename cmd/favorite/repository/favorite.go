package repository

import (
	videoRepository "ByteDance/cmd/video/repository"
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/pkg/common"
	"ByteDance/utils"
	"sync"
)

//点赞操作
type Favorite struct {
	model.Favorite
}

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

//取消点赞
func (*FavoriteStruct) FavoriteUpdate(userId int32, videoId int32) (RowsAffected int64) {
	f := dal.ConnQuery.Favorite

	favorite := &model.Favorite{UserID: userId, VideoID: videoId}

	row, err := f.Where(f.UserID.Eq(favorite.UserID), f.VideoID.Eq(favorite.VideoID)).Update(f.Removed, common.Removed)

	utils.CatchErr("更新错误", err)

	return row.RowsAffected
}

//点赞
func (*FavoriteStruct) FavoriteCreate(userId int32, videoId int32) (err error) {
	f := dal.ConnQuery.Favorite

	favorite := &model.Favorite{UserID: userId, VideoID: videoId}

	err = f.Create(favorite)

	return err
}

//点赞列表
func (*FavoriteStruct) FavoriteList(userId int32) ([]videoRepository.VideoInfo, bool) {
	v := dal.ConnQuery.Video
	u := dal.ConnQuery.User
	f := dal.ConnQuery.Favorite

	var result []videoRepository.VideoInfo
	// 内联查询
	err := f.Select(u.ID.As("UserID"), u.Username, v.ID.As("VideoID"), v.PlayURL, v.CoverURL, v.Time, v.Title).Where(f.UserID.Eq(userId), f.Removed.Eq(0), f.Deleted.Eq(0)).Join(v, v.ID.EqCol(f.VideoID)).Join(u, u.ID.EqCol(v.AuthorID)).Scan(&result)
	utils.CatchErr("获取视频信息错误", err)
	if result == nil {
		return nil, false

	}
	return result, true
}
