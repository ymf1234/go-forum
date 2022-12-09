package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432 // 每一票的值432分
	PostPerAge               = 20
)

/*
投票算法：http://www.ruanyifeng.com/blog/2012/03/ranking_algorithm_reddit.html
本项目使用简化版的投票分数
投一票就加432分 86400/200 -> 200张赞成票就可以给帖子在首页续天  -> 《redis实战》
*/

/*
	PostVote 为帖子投票

投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况

	1.之前没投过票，现在要投赞成票		 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过反对票，现在要改为赞成票	 --> 更新分数和投票记录	差值的绝对值：2	+432*2

v=0时，有两种情况

	1.之前投过反对票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	+432
	2.之前投过赞成票，现在要取消			 --> 更新分数和投票记录	差值的绝对值：1	-432

v=-1时，有两种情况

	1.之前没投过票，现在要投反对票		 --> 更新分数和投票记录	差值的绝对值：1	-432
	2.之前投过赞成票，现在要改为反对票	 --> 更新分数和投票记录	差值的绝对值：2	-432*2

投票的限制：
每个帖子子发表之日起一个星期之内允许用户投票，超过一个星期就不允许投票了

	1、到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2、到期之后删除那个 KeyPostVotedZSetPrefix
*/
func VoteForPost(userID string, postID string, v float64) (err error) {
	return
}

func CreatePost(postID, userID uint64, title, summary string, CommunityID uint64) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	communityKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(CommunityID))

	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  summary,
		"post:id":  postID,
		"user:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}

	// 事务操作
	pipeline := client.TxPipeline()
	pipeline.ZAdd(votedKey, redis.Z{
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*OneWeekInSeconds) // 一周时间

	pipeline.HMSet(KeyPostVotedZSetPrefix+strconv.Itoa(int(postID)), postInfo)

	// 添加到分数的ZSet
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})

	// 添加到时间的ZSet
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})

	// 添加到对应版块  把帖子添加到社区的set
	pipeline.SAdd(communityKey, postID)

	_, err = pipeline.Exec()
	return

}
