package service

import (
	"ByteDance/cmd/comment"
	"ByteDance/cmd/comment/repository"
	"ByteDance/cmd/video"
	"ByteDance/dal/method"
	"ByteDance/utils"
	"github.com/araddon/dateparse"
)

func CommentAction(userId int32, videoId int32, commentText string, commentId int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.CommentDao.CommentUpdate(commentId)

	//commentText不空，创建评论
	if IsExist == 0 && len(commentText) != 0 {
		//添加该数据
		err = repository.CommentDao.CommentCreate(userId, videoId, commentText)
		utils.CatchErr("添加失败", err)
	}

	return err
}

//评论列表
func CommentList(videoId int32) (commentInfo []comment.TheCommentInfo, state int) {
	allCommentInfoData, _ := repository.CommentDao.CommentList(videoId)

	commentInfo = make([]comment.TheCommentInfo, len(allCommentInfoData))

	for index, commentInfoData := range allCommentInfoData {
		//获取关注相关
		followerCount, followCount, _ := method.QueryFollowCount(commentInfoData.UserID)
		//评论时间转换
		createDate, _ := dateparse.ParseAny(commentInfoData.CreateDate)

		commentInfo[index] = comment.TheCommentInfo{
			ID: commentInfoData.ID,
			User: video.AuthorInfo{
				ID:            commentInfoData.UserID,
				Name:          commentInfoData.Username,
				FollowCount:   int(followCount),
				FollowerCount: int(followerCount),
				IsFollow:      false,
			},
			Content:    commentInfoData.Content,
			CreateDate: createDate.Format("01-02"),
		}
	}

	return commentInfo, 1
}
