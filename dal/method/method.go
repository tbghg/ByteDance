package method

import "gorm.io/gen"

// UserMethod Method 指定自定义查询方法
type FollowMethod interface {
	// 查询用户粉丝数
	//
	//select count(1) from follow Where user_id = @userID and removed = 0 and deleted = 0
	QueryFollowerCount(userID int32) int64

	// 查询粉丝数目
	//
	//select count(1) from follow Where fun_id = @funID and removed = 0 and deleted = 0
	QueryFollowCount(funID int32) int64
}

type FavoriteMethod interface {
	// 查询视频点赞数目
	//
	// select count(1) from favorite where video_id = @videoID and removed = 0 and deleted = 0
	QueryFavoriteCount(videoID int32) int64

	// removed置反
	//
	// update favorite set removed = -removed where user_id = @userID and video_id = @videoID and deleted = 0
	UpdateFavoriteRemoved(userID int32, videoID int32) (gen.RowsAffected, error)
}

type CommentMethod interface {
	// 查询视频评论数目
	//
	// select count(1) from comment where video_id = @videoID and removed = 0 and deleted = 0
	QueryCommentCount(videoID int32) int64
}
