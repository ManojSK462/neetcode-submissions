type tweet struct{
	id int
	time int
	user int
	index int
}

type maxheap []tweet

func (h maxheap) Len() int           { return len(h) }
func (h maxheap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h maxheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxheap) Push(x any) {
	*h = append(*h, x.(tweet))
}

func (h *maxheap) Pop() any {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}



type Twitter struct {
    time int
	tweets    map[int][]tweet
	following map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter{
		tweets: make(map[int][]tweet),
		following: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.time++
	tt := tweet{
		id: tweetId,
		time: this.time,
		user: userId,
	}

	this.tweets[userId] = append(this.tweets[userId], tt)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    h := &maxheap{}

	rec_users := map[int]bool{userId:true}

	for followeeID := range this.following[userId]{
		rec_users[followeeID] = true
	}

	for id:= range rec_users{
		user_tt := this.tweets[id]

		if len(user_tt)>0{
			index := len(user_tt)-1
			t := user_tt[index]
			t.index = index
			heap.Push(h, t)
		}
	}
		feed:=make([]int, 0, 10)

		for h.Len()>0 && len(feed)<10{
			pt := heap.Pop(h).(tweet)
			feed = append(feed, pt.id)

			nextIndex := pt.index - 1

			if nextIndex >= 0 {
			nextTweet := this.tweets[pt.user][nextIndex]
			nextTweet.index = nextIndex
			heap.Push(h, nextTweet)
			}
		}

	

return feed
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.following[followerId] == nil {
		this.following[followerId] = make(map[int]bool)
	}
	this.following[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.following[followerId], followeeId)
}
