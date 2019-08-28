// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"os"
	"strings"

	"github.com/inancgumus/learngo/logparser/v5/pipe"
	"github.com/inancgumus/learngo/logparser/v5/pipe/group"
	"github.com/inancgumus/learngo/logparser/v5/pipe/parse"
	"github.com/inancgumus/learngo/logparser/v5/pipe/report"
)

// fromFile generates a default pipeline.
// Detects the correct parser by the file extension.
// Uses a TextReport and groups by domain.
func fromFile(path string) (*pipe.Pipeline, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var src pipe.Iterator
	switch {
	case strings.HasSuffix(path, ".txt"):
		src = parse.FromText(f)
	case strings.HasSuffix(path, ".jsonl"):
		src = parse.FromJSON(f)
	}

	return pipe.New(
		src,
		report.AsText(os.Stdout),
		group.By(group.Domain),
	), nil
}
