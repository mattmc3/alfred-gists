package main

// Author: mattmc3
// License: MIT
// Copyright (c): 2018, mattmc3
// Website: https://github.com/mattmc3/alfred-gists

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"./gists"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
)

var (
	wf *aw.Workflow
	// ErrNoGistsFound occurs when no gists are returned
	ErrNoGistsFound = errors.New("No gists found")
	// CacheMaxAge tells us how long we keep the Alfred cache around before
	// going back out to GitHub. This prevents us from spamming GitHub with
	// our REST API calls.
	CacheMaxAge = 5 * time.Minute
	// User is the GitHub user
	User = ""
	// Token is the GitHub auth token with gist perms
	Token = ""
)

func checkErr(e error) {
	if e != nil {
		wf.FatalError(e)
	}
}

func init() {
	// the special init method is automatically called
	wf = aw.New()
}

// we need a function for LoadOrStoreJSON
func getGistsFromGitHub() (interface{}, error) {
	log.Print("[DEBUG] No current cache. Taking a trip to GitHub...")
	gists, err := gists.GetAllGists(User, Token)
	if err != nil {
		return nil, err
	}
	return &gists, nil
}

func getGistsFromCache() (*[]github.Gist, error) {
	var gists []github.Gist
	err := wf.Cache.LoadOrStoreJSON("gists", CacheMaxAge, getGistsFromGitHub, &gists)
	if err != nil {
		return nil, err
	}
	return &gists, nil
}

func loadUserAndToken() (string, string, error) {
	bytUser, e := wf.Data.Load("user")
	if e != nil {
		return "", "", errors.New("User not set. Use gist-config")
	}
	user := string(bytUser[:])

	bytToken, e := wf.Data.Load("token")
	if e != nil {
		return "", "", errors.New("Token not set. Use gist-config")
	}
	token := string(bytToken[:])
	return user, token, nil
}

func storeConfigValue(key, value string) error {
	return wf.Data.Store(key, []byte(value))
}

func execListGists() {
	// get the user and token
	var err error
	User, Token, err = loadUserAndToken()
	checkErr(err)

	// get gists
	gists, err := getGistsFromCache()
	checkErr(err)

	log.Printf("[DEBUG] Total gists found: %v\n", len(*gists))
	if len(*gists) == 0 {
		wf.WarnEmpty("Hmm... No gists found", "Check Alfred debugger if this seems wrong to you")
		return
	}

	for _, gist := range *gists {
		// assemble the list of files:
		fileCount := len(gist.Files)
		fileNames := make([]string, fileCount)
		keynum := 0
		var rawURL string
		for key, val := range gist.Files {
			if rawURL == "" {
				rawURL = *val.RawURL
			}
			fileNames[keynum] = fmt.Sprint(key)
			keynum++
		}
		fileList := strings.Join(fileNames, ",")

		// show the user the good stuff!
		title := *gist.Description
		if !*gist.Public {
			title = "🔒 " + title
		}
		subtitle := fileList
		uid := *gist.ID
		arg := *gist.ID
		htmlURL := *gist.HTMLURL

		item := wf.NewItem(title).
			Subtitle(subtitle).
			UID(uid).
			Arg(arg).
			Valid(true)

		item.NewModifier(aw.ModAlt).
			Arg(htmlURL).
			Subtitle("Open gist: " + htmlURL)

		item.NewModifier(aw.ModCtrl).
			Arg(rawURL).
			Subtitle("Open raw text: " + rawURL)
	}
}

func run() {
	var action string
	if args := wf.Args(); len(args) > 0 {
		action = args[0]
	}

	if action == "list" {
		execListGists()
	} else if action == "config" {
		query := wf.Args()[1]
		parts := strings.Split(query, "|")
		key := parts[0]
		value := parts[1]
		err := storeConfigValue(key, value)
		checkErr(err)
	} else {
		wf.Fatal("No script action specified. Ex: 'list'")
	}

	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// wrapping run means that we catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
