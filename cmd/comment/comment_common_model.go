package comment

import "ByteDance/cmd/video"

//评论信息
type TheCommentInfo struct {
	ID         int32            `json:"id"`
	User       video.AuthorInfo `json:"user"`
	Content    string           `json:"content"`
	CreateDate string           `json:"create_date"`
}
