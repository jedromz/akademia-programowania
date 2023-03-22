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
	Url     string
	Payload response
}

func (m *MyFetcher) Fetch(ctx context.Context) error {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		return errors.New("logger not found in context")
	}
	resp, err := makeRequest(ctx, m.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&m.Payload)
	if err != nil {
		logger.Println(err)
		return err
	}
	return nil
}
func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		return nil, errors.New("logger not found in context")
	}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	logger.Printf("Making request: %s %s\n", req.Method, req.URL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	return resp, nil
}
func (m *MyFetcher) Save(ctx context.Context, writer io.Writer) error {
	logger, ok := ctx.Value("logger").(*log.Logger)
	if !ok {
		return errors.New("logger not found in context")
	}
	logger.Printf("Saving data from %s", m.Url)
	err := writeToFile(m, writer, logger)
	if err != nil {
		return err
	}
	return nil
}

func writeToFile(m *MyFetcher, writer io.Writer, logger *log.Logger) error {
	for _, resp := range m.Payload.Data.Children {
		_, err := fmt.Fprintf(writer, "%s\n%s\n\n", resp.Data.Title, resp.Data.URL)
		if err != nil {
			logger.Println(err)
			return err
		}
	}
	return nil
}
