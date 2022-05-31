package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

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

func (*FavoriteStruct) RelationUpdate(userId int32, videoId int32, actionType int32) (RowsAffected int64) {
	f := dal.ConnQuery.Favorite

	favorite := &model.Favorite{UserID: userId, VideoID: videoId, Removed: actionType}

	row, err := f.Where(f.UserID.Eq(favorite.UserID), f.VideoID.Eq(favorite.VideoID)).Update(f.Removed, favorite.Removed)

	utils.CatchErr("更新错误", err)

	return row.RowsAffected
}

func (*FavoriteStruct) RelationCreate(userId int32, videoId int32) (err error) {
	f := dal.ConnQuery.Favorite

	favorite := &model.Favorite{UserID: userId, VideoID: videoId}

	err = f.Create(favorite)

	return err
}
