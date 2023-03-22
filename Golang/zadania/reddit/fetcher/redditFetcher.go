package fetcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch(ctx context.Context) error
	Save(context.Context, io.Writer) error
}

type MyFetcher struct {
	Url  string
	Data response
}

func (m *MyFetcher) Fetch(ctx context.Context) error {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		// Handle missing logger in context
		return errors.New("logger not found in context")
	}
	req, err := http.NewRequestWithContext(ctx, "GET", m.Url, nil)
	if err != nil {
		logger.Println(err)
		return err
	}
	logger.Printf("Making request: %s %s\n", req.Method, req.URL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Println(err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&m.Data)
	if err != nil {
		logger.Println(err)
		return err
	}
	return nil
}

func (m *MyFetcher) Save(ctx context.Context, writer io.Writer) error {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		// Handle missing logger in context
		return errors.New("logger not found in context")
	}
	logger.Printf("Saving data...")
	for _, dto := range m.Data.Data.Children {
		_, err := fmt.Fprintf(writer, "%s\n%s\n\n", dto.Data.Title, dto.Data.URL)
		if err != nil {
			logger.Println(err)
			return err
		}
	}
	return nil
}
