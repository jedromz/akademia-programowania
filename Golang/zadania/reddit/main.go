package main

import (
	"context"
	"errors"
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

func processSubreddit(ctx context.Context, subreddit string) error {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		return errors.New("logger not found in context")
	}
	filename := strings.TrimSuffix(path.Base(subreddit), ".json")
	myFetcher := fetcher.SimpleRedditFetcher{Url: Path + subreddit}
	file, err := os.Create(TargetDir + filename)
	if err != nil {
		logger.Printf("failed to create file: %v", err)
		return err
	}
	defer file.Close()

	err = myFetcher.Fetch(ctx)
	if err != nil {
		logger.Printf("failed to fetch data for %s: %v", subreddit, err)
		return err
	}
	err = myFetcher.Save(ctx, file)
	if err != nil {
		logger.Printf("failed to save data for %s: %v", subreddit, err)
		return err
	}
	return nil
}
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
			err := processSubreddit(ctx, subreddit)
			if err != nil {
				log.Println(err)
			}
		}(v)
	}
	wg.Wait()
}
