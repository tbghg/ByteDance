package method

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
)

func QueryFollowCount(userID int32) (followerCount int64, followCount int64, success bool) {
	// follower_count 粉丝数	follow_count 关注数
	var err error
	f := dal.ConnQuery.Follow

	followerCount, err = f.Select(f.UserID).Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.UserID.Eq(userID)).Count()
	utils.CatchErr("查询粉丝数错误", err)
	if err != nil {
		return 0, 0, false
	}

	followCount, err = f.Select(f.FunID).Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.FunID.Eq(userID)).Count()
	utils.CatchErr("查询关注数错误", err)
	if err != nil {
		return 0, 0, false
	}

	return followerCount, followCount, true
}

func QueryCommentCountByVideoID(videoID int32) int64 {
	c := dal.ConnQuery.Comment
	commentCount, err := c.Select(c.VideoID).Where(c.Deleted.Eq(0), c.Removed.Eq(0), c.VideoID.Eq(videoID)).Count()
	utils.CatchErr("查询评论数错误", err)
	return commentCount
}

func QueryFavoriteCount(videoID int32) int64 {
	f := dal.ConnQuery.Favorite
	favoriteCount, err := f.Select(f.VideoID).Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.VideoID.Eq(videoID)).Count()
	utils.CatchErr("查询点赞数错误", err)
	return favoriteCount
}

func QueryUserById(userId int32) (user *model.User) {
	u := dal.ConnQuery.User
	user, _ = u.Where(u.ID.Eq(userId)).First()
	return user
}
