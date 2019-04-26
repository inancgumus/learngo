package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// Film represents a single Star Wars film
type Film struct {
	// All fields should be exported to be decoded
	Title     string   `json:"title"`
	Opening   string   `json:"opening_crawl"`
	Starships []string `json:"starships"`
}

func (f Film) String() string {
	var buf strings.Builder

	fmt.Fprintln(&buf, f.Title)
	fmt.Fprintln(&buf, strings.Repeat("-", len(f.Title)))
	buf.WriteByte('\n')
	fmt.Fprintln(&buf, f.Opening)

	return buf.String()
}

// fetchFilm returns a Film resource by its id
func fetchFilm(ctx context.Context, id int) (film Film, err error) {
	fmt.Println("requesting", fmt.Sprintf(swapi+"films/%d/", id))
	c, err := request(ctx, fmt.Sprintf(swapi+"films/%d/", id))
	if err != nil {
		return film, err
	}

	return film, json.Unmarshal(c, &film)
}
