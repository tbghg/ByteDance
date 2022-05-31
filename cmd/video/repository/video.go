package repository

import (
	"ByteDance/dal"
	"ByteDance/utils"
	"sync"
	"time"
)

type videoInfo struct {
	VideoID       int32
	UserID        int32
	Username      string
	PlayURL       string
	CoverURL      string
	FavoriteCount int
	IsFavorite    bool
	Time          time.Time
	Title         string
}

type VideoDaoStruct struct {
}

var (
	VideoDao  *VideoDaoStruct
	videoOnce sync.Once
)

func init() {
	videoOnce.Do(func() {
		VideoDao = &VideoDaoStruct{}
	})
}

func (*VideoDaoStruct) GetVideoFeed(lastTime time.Time) ([]videoInfo, bool) {
	v := dal.ConnQuery.Video
	u := dal.ConnQuery.User
	var result []videoInfo
	// 内联查询
	err := v.Select(u.ID.As("UserID"), u.Username, v.ID.As("VideoID"), v.PlayURL, v.CoverURL, v.Time, v.Title).Where(v.Time.Lt(lastTime), v.Removed.Eq(0), v.Deleted.Eq(0)).Join(u, u.ID.EqCol(v.AuthorID)).Order(v.Time.Desc()).Limit(10).Scan(&result)
	utils.CatchErr("获取视频信息错误", err)
	if result == nil {
		return nil, false
	}
	return result, true
}
