// Package gists is a wrapper library for interacting with GitHub gists
package gists

// Author: mattmc3
// License: MIT
// Copyright (c): 2018, mattmc3
// Website: https://github.com/mattmc3/alfred-gists

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// gistListResult holds the result of gist list api call
type gistListResult struct {
	Gists []*github.Gist
	Resp  *github.Response
}

func getGistList(user, token string, pageNum int) (*gistListResult, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.GistListOptions{
		ListOptions: github.ListOptions{Page: pageNum},
	}

	gists, resp, err := client.Gists.List(ctx, user, opt)
	if err != nil {
		return nil, err
	}
	result := &gistListResult{
		Gists: gists,
		Resp:  resp,
	}
	return result, nil
}

func getGistListConcurrent(user, token string, pageNum int, ch chan<- *gistListResult) {
	gresult, _ := getGistList(user, token, pageNum)
	ch <- gresult
}

// GetAllGists returns every page of gists
func GetAllGists(user, token string) ([]*github.Gist, error) {
	var allGists []*github.Gist

	// the first page can't be concurrent b/c we need to know how many
	// pages we have in total
	pageOneGistResult, err := getGistList(user, token, 1)
	if err != nil {
		return nil, err
	}

	// add page one gists to result
	allGists = append(allGists, pageOneGistResult.Gists...)
	lastPage := pageOneGistResult.Resp.LastPage

	// the rest of the pages can be nabbed concurrently
	ch := make(chan *gistListResult, lastPage-1)
	for pageNum := 2; pageNum <= lastPage; pageNum++ {
		go getGistListConcurrent(user, token, pageNum, ch)
	}

	for pageNum := 2; pageNum <= lastPage; pageNum++ {
		var gistResult = <-ch
		allGists = append(allGists, gistResult.Gists...)
	}

	return allGists, nil
}
