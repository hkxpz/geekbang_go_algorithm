package Week04

//130. 被围绕的区域
var n, m int

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m = len(board), len(board[0])
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}

//355. 设计推特
type Twitter struct {
	Tweets     []int
	UserTweets map[int][]int
	Follows    map[int][]int
	IsFollowMy map[int]bool
}

func Constructor() Twitter {
	// 每一次实例化的时候，都重新分配一次，这样不会造成示例重复

	var Tweets []int

	// 某用户发的某条推特
	var UserTweets = make(map[int][]int)

	// 某用户关注了哪些用户
	var Follows = make(map[int][]int)

	var IsFollowMy = make(map[int]bool)

	t := Twitter{
		Tweets:     Tweets,
		UserTweets: UserTweets,
		Follows:    Follows,
		IsFollowMy: IsFollowMy,
	}
	return t
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	// 每个人每次发推特，都记录到一个地方
	t.Tweets = append(t.Tweets, tweetId)
	// 某个用户发了推特，存到自己推特列表里
	t.UserTweets[userId] = append(t.UserTweets[userId], tweetId)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	fs := t.Follows[userId] // 先获取该用户的关注列表
	var allTweets []int
	for _, v := range fs {
		// 把关注列表的人的所有推特都集中起来
		allTweets = append(allTweets, t.UserTweets[v]...)
	}
	if !t.IsFollowMy[userId] {
		// 如果自己没有关注自己，那么也需要把自己发的推特加到一起
		allTweets = append(allTweets, t.UserTweets[userId]...)
	}
	var sortTweets []int
	aTLen := len(t.Tweets)
	s := 0
	// 按照发的推特顺序进行倒序排序
	for i := aTLen - 1; i >= 0; i-- {
		if s >= 10 {
			break
		}
		for _, n := range allTweets {

			// 只取 10条数据
			if t.Tweets[i] == n && s < 10 {
				s++
				sortTweets = append(sortTweets, n)
			}
		}
	}

	return sortTweets
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	// 如果自己关注了自己，标记一下
	if followerId == followeeId {
		t.IsFollowMy[followerId] = true
	}

	// 下面是判断这人是否关注了，如果已经关注了，那么就不再关注了
	var isFed bool
	for _, v := range t.Follows[followerId] {
		if v == followeeId {
			isFed = true
		}
	}
	if !isFed {
		t.Follows[followerId] = append(t.Follows[followerId], followeeId)
	}
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	// 如果自己取关了自己，标记一下
	if followeeId == followerId {
		t.IsFollowMy[followerId] = false
	}

	// 去掉自己关注列表里那个被关注的人
	var temp []int
	for _, v := range t.Follows[followerId] {
		if v != followeeId {
			temp = append(temp, v)
		}
	}
	t.Follows[followerId] = temp
}
