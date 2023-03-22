package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"reddit/fetcher"
	"strings"
	"sync"
	"time"
)

const (
	Path      = "https://www.reddit.com/"
	TargetDir = "target/"
)

func main() {
	fmt.Println("hello r/golang!")
	ctx := context.Background()
	logger := log.New(os.Stdout, "[log] ", log.LstdFlags)
	ctx = context.WithValue(ctx, "logger", logger)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subreddits := []string{
		"r/oddlysatisfying.json",
		"r/misleadingthumbnails.json",
		"r/explainlikeimfive.json",
		"r/golang.json",
		"r/java.json",
	}

	var wg sync.WaitGroup
	wg.Add(len(subreddits))

	for _, v := range subreddits {
		go func(subreddit string) {
			defer wg.Done()
			filename := strings.TrimSuffix(path.Base(subreddit), ".json")
			myFetcher := fetcher.MyFetcher{Url: Path + subreddit}
			file, err := os.Create(TargetDir + filename)
			if err != nil {
				logger.Printf("failed to create file: %v", err)
				return
			}
			defer file.Close()

			err = myFetcher.Fetch(ctx)
			if err != nil {
				logger.Printf("failed to fetch data for %s: %v", subreddit, err)
				return
			}
			err = myFetcher.Save(ctx, file)
			if err != nil {
				logger.Printf("failed to save data for %s: %v", subreddit, err)
				return
			}
		}(v)
	}
	wg.Wait()
}
