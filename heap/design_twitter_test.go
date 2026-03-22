package heap

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := ConstructorTwitter()
	twitter.PostTweet(1, 5)
	if !reflect.DeepEqual(twitter.GetNewsFeed(1), []int{5}) {
		t.Errorf("Newsfeed for user 1 should be [5]")
	}

	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	if !reflect.DeepEqual(twitter.GetNewsFeed(1), []int{6, 5}) {
		t.Errorf("Newsfeed for user 1 should be [6, 5]")
	}

	twitter.Unfollow(1, 2)
	if !reflect.DeepEqual(twitter.GetNewsFeed(1), []int{5}) {
		t.Errorf("Newsfeed for user 1 should be [5] after unfollowing user 2")
	}
}
