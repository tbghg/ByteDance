package video

// TheVideoInfo 视频信息
type TheVideoInfo struct {
	ID            int32      `json:"id"`
	Author        AuthorInfo `json:"author"`
	PlayURL       string     `json:"play_url"`
	CoverURL      string     `json:"cover_url"`
	FavoriteCount int        `json:"favorite_count"`
	CommentCount  int        `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
}

// AuthorInfo 作者信息
type AuthorInfo struct {
	ID            int32  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
