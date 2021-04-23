//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream) (tweets []*Tweet) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(ch)
			return tweets
		}
		ch <- tweet
		tweets = append(tweets, tweet)
	}
}

func consumer() {
	defer wg.Done()
	for {
		t, ok := <-ch
		if !ok {
			return
		}
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

var wg sync.WaitGroup
var ch = make(chan *Tweet)

func main() {
	wg.Add(2)
	start := time.Now()
	stream := GetMockStream()

	// Producer
	go producer(stream)

	// Consumer
	go consumer()
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
