package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

type VideoInfo struct {
	VideoID       int32
	UserID        int32
	Username      string
	PlayURL       string
	CoverURL      string
	FavoriteCount int
	IsFavorite    bool
	Time          int32
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

func (*VideoDaoStruct) GetVideoFeed(lastTime int32) ([]VideoInfo, bool) {
	v := dal.ConnQuery.Video
	u := dal.ConnQuery.User
	var result []VideoInfo
	// 内联查询
	err := v.Select(u.ID.As("UserID"), u.Username, v.ID.As("VideoID"), v.PlayURL, v.CoverURL, v.Time, v.Title).Where(v.Time.Lt(lastTime), v.Removed.Eq(0), v.Deleted.Eq(0)).Join(u, u.ID.EqCol(v.AuthorID)).Order(v.Time.Desc()).Limit(10).Scan(&result)
	if !utils.CatchErr("获取视频信息错误", err) {
		return nil, false
	}
	if result == nil {
		return nil, false
	}
	return result, true
}

func (*VideoDaoStruct) GetVideoInfo(userID int32, videoID int32) (followerCount int64, followCount int64, commentCount int64, favoriteCount int64) {
	f := dal.ConnQuery.Follow
	c := dal.ConnQuery.Comment
	favorite := dal.ConnQuery.Favorite
	followerCount = f.QueryFollowerCount(userID)
	followCount = f.QueryFollowCount(userID)
	commentCount = c.QueryCommentCount(videoID)
	favoriteCount = favorite.QueryFavoriteCount(videoID)
	return followerCount, followCount, commentCount, favoriteCount
}

func (*VideoDaoStruct) QueryIsFavorite(userID int32, videoID int32) bool {
	f := dal.ConnQuery.Favorite
	count, _ := f.Where(f.UserID.Eq(userID), f.VideoID.Eq(videoID), f.Removed.Eq(0), f.Deleted.Eq(0)).Count()
	if count == 0 {
		return false
	} else {
		return true
	}
}

func (*VideoDaoStruct) QueryIsFollow(userID int32, authorID int32) bool {
	f := dal.ConnQuery.Follow
	count, _ := f.Where(f.UserID.Eq(authorID), f.Removed.Eq(0), f.Deleted.Eq(0), f.FunID.Eq(userID)).Count()
	if count == 0 {
		return false
	} else {
		return true
	}
}

func (*VideoDaoStruct) GetVideoList(userID int32) ([]VideoInfo, bool) {
	v := dal.ConnQuery.Video
	u := dal.ConnQuery.User
	var result []VideoInfo
	// 内联查询
	err := v.Select(u.ID.As("UserID"), u.Username, v.ID.As("VideoID"), v.PlayURL, v.CoverURL, v.Time, v.Title).Where(v.AuthorID.Eq(userID), v.Removed.Eq(0), v.Deleted.Eq(0)).Join(u, u.ID.EqCol(v.AuthorID)).Order(v.Time.Desc()).Scan(&result)
	if !utils.CatchErr("获取视频信息错误", err) {
		return nil, false
	}
	if result == nil {
		return nil, false
	}
	return result, true
}

func (*VideoDaoStruct) PublishVideo(userID int, title string, videoNumID string) bool {
	v := dal.ConnQuery.Video
	video := model.Video{
		AuthorID: int32(userID),
		PlayURL:  videoNumID,
		CoverURL: videoNumID,
		Title:    title,
	}
	err := v.Create(&video)
	if !utils.CatchErr("video插入数据错误", err) {
		return false
	}
	return true
}
