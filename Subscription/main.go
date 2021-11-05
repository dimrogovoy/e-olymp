package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Item struct {
	Tile    string
	Channel string
	GUID    string
}

// Fetcher fetches Items and returns the time when the next fetch
// should be attempted. On failure, fetch returns non-nil error.
type Fetcher interface {
	Fetch() (item []Item, next time.Time, err error)
}
type fetch struct {
	rnd      *rand.Rand
	maxcount int
}

func NewFetcher() Fetcher {
	rand.Seed(time.Now().UnixNano())
	return &fetch{
		rnd:      rand.New(rand.NewSource(time.Now().UnixNano())),
		maxcount: rand.Intn(10),
	}
}

func (f *fetch) Fetch() (items []Item, next time.Time, err error) {
	if f.maxcount < 0 {
		return nil, time.Time{}, fmt.Errorf("no more fetch")
	}

	for i := 0; i < f.rnd.Intn(10); i++ {
		var item Item
		item.Tile = fmt.Sprintf("tile %d", i)
		item.Channel = fmt.Sprintf("channel %d", i)
		item.GUID = strconv.Itoa(f.rnd.Intn(100))
		items = append(items, item)
	}

	next = time.Now().Add(time.Duration(f.rnd.Intn(1e3)) * time.Millisecond)

	f.maxcount--

	return
}

// A Subscription delivers Items over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type Subscription interface {
	Updates() <-chan Item
	Close() error
}
type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	return s
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	// STOPCLOSESIG OMIT
	errc := make(chan error)
	s.closing <- errc // HLchan
	return <-errc     // HLchan
}

func (s *sub) subscribe() {
	c := <-s.closing
	for {
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			fmt.Println(err)
			c <- err
			return
		}
		for _, item := range items {
			s.updates <- item
		}
		time.Sleep(next.Sub(time.Now()))
	}
}

func main() {
	f := NewFetcher()

	s := Subscribe(f)
	defer s.Close()

	if sub, ok := s.(*sub); ok {
		go sub.subscribe()
	}

	go func() {
		ch := s.Updates()
		for item := range ch {
			fmt.Println(item)
		}
	}()

}
