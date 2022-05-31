package user

type RegUserData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

type LoginData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

type GetUserInfoData struct {
	ID            int32  `json:"id"`
	UseName       string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
