package heap

import (
	"container/heap"
)

type Tweet struct {
	id   int
	time int
}

type Twitter struct {
	time      int
	tweets    map[int][]Tweet
	followers map[int]map[int]bool
}

func ConstructorTwitter() Twitter {
	return Twitter{
		time:      0,
		tweets:    make(map[int][]Tweet),
		followers: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.time++
	this.tweets[userId] = append(this.tweets[userId], Tweet{tweetId, this.time})
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time } // Max-heap based on time
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)

	// Add user's own tweets
	for _, tweet := range this.tweets[userId] {
		heap.Push(h, tweet)
	}

	// Add followees' tweets
	for followeeId := range this.followers[userId] {
		for _, tweet := range this.tweets[followeeId] {
			heap.Push(h, tweet)
		}
	}

	res := []int{}
	for h.Len() > 0 && len(res) < 10 {
		res = append(res, heap.Pop(h).(Tweet).id)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.followers[followerId] == nil {
		this.followers[followerId] = make(map[int]bool)
	}
	this.followers[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followers[followerId] != nil {
		delete(this.followers[followerId], followeeId)
	}
}
