package service

import (
	"ByteDance/cmd/comment"
	"ByteDance/cmd/comment/repository"
	"ByteDance/cmd/video"
	"ByteDance/dal/method"
	"github.com/araddon/dateparse"
	"sync"
)

func CommentAction(userId int32, videoId int32, commentText string, commentId int32) (success bool) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.CommentDao.CommentUpdate(commentId)

	//commentText不空，创建评论
	if IsExist == 0 && len(commentText) != 0 {
		//添加该数据
		success = repository.CommentDao.CommentCreate(userId, videoId, commentText)
	}

	return success
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
