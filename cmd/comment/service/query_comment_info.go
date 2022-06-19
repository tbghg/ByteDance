package service

import (
	"ByteDance/cmd/comment"
	"ByteDance/cmd/comment/repository"
	"ByteDance/cmd/video"
	"ByteDance/dal/method"
	"github.com/araddon/dateparse"
	"sync"
)

func CommentAction(userId int32, videoId int32, commentText string, commentId int32) (commentInfo comment.TheCommentInfo, success bool) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.CommentDao.CommentUpdate(commentId, userId)

	//commentText不空，创建评论
	if IsExist == 0 && len(commentText) != 0 {
		//添加该数据
		var commentData repository.CommentInfo
		commentData, success = repository.CommentDao.CommentCreate(userId, videoId, commentText)
		followerCount, followCount, _ := method.QueryFollowCount(commentData.UserID) // 获取关注相关

		createDate, _ := dateparse.ParseAny(commentData.CreateDate) // 评论时间转换
		commentInfo = comment.TheCommentInfo{
			ID: commentData.ID,
			User: video.AuthorInfo{
				ID:            commentData.UserID,
				Name:          commentData.Username,
				FollowCount:   int(followCount),
				FollowerCount: int(followerCount),
				IsFollow:      false,
			},
			Content:    commentData.Content,
			CreateDate: createDate.Format("01-02"),
		}
		return commentInfo, success
	}

	return commentInfo, success
}

// CommentList 评论列表
func CommentList(videoId int32) (commentInfo []comment.TheCommentInfo, state int) {
	allCommentInfoData, _ := repository.CommentDao.CommentList(videoId)

	commentInfo = make([]comment.TheCommentInfo, len(allCommentInfoData))

	wg := sync.WaitGroup{}
	wg.Add(len(allCommentInfoData))

	for index, commentInfoData := range allCommentInfoData {
		go func(index int, commentInfoData repository.CommentInfo, commentInfo []comment.TheCommentInfo) {
			followerCount, followCount, _ := method.QueryFollowCount(commentInfoData.UserID) // 获取关注相关

			createDate, _ := dateparse.ParseAny(commentInfoData.CreateDate) // 评论时间转换

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
			wg.Done()
		}(index, commentInfoData, commentInfo)
	}
	wg.Wait()
	return commentInfo, 1
}
