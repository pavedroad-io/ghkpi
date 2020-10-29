// Package cmd from cobra
package cmd

/*
Copyright © 2019,2020 PavedRoad <info@pavedroad.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// getClient: return a githubClient with:
//   TokenAuth, BasicAuth, or no authentication
//

const (
	unathenticated = iota
	oauth
	basic
)

type gitClientConfig struct {
	authType     int
	accessToken  string
	userName     string
	userPassword string
}

func getClient(thisConfiguration gitClientConfig) (client *github.Client, err error) {

	if thisConfiguration.authType == oauth {
		// Use OAUTH
		if thisConfiguration.accessToken == "" {
			//Look for it in the environment
			thisConfiguration.accessToken = getAccessToken()

			if thisConfiguration.accessToken == "" {
				return nil, errors.New("No access token found, try setting GITHUB_TOKEN")
			}
		}

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: thisConfiguration.accessToken},
		)
		// Construct *http.Client to pass to github.NewClient
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else if thisConfiguration.authType == basic {
		// Use basic authentication
		// If not set look in environment
		if thisConfiguration.userName == "" {
			if thisConfiguration.userName = getUser(); thisConfiguration.userName == "" {
				return nil, errors.New("No user name found, try setting GITHUB_USER")
			}
		}

		if thisConfiguration.userPassword == "" {
			if thisConfiguration.userPassword = getUser(); thisConfiguration.userPassword == "" {
				return nil, errors.New("No password found, try setting GITHUB_PASSWORD")
			}
		}
		bat := github.BasicAuthTransport{
			Username: strings.TrimSpace(thisConfiguration.userName),
			Password: strings.TrimSpace(thisConfiguration.userPassword),
		}
		client = github.NewClient(bat.Client())
	} else {
		// Use an unauthenticated client
		client = github.NewClient(nil)
	}

	return client, nil
}

// getGitHubTokenn looks for GITHUB_TOKEN environment variable
// 	This follows the conventions used by GitHub GH CLI
//  and returns the empty string if not found
func getAccessToken() (token string) {
	//Prefix for environment name with service (GH GitHub)
	return os.Getenv("GITHUB_TOKEN")
}

// NOTE: GitHub is ending support for user name and password
// authentication

// getUser looks for GITHUB_USER environment variable
func getUser() (token string) {
	//Prefix for environment name with service (GH GitHub)
	return os.Getenv("GITHUB_USER")
}

// getPassword looks for GITHUB_PASSWORD environment variable
func getPassword() (token string) {
	//Prefix for environment name with service (GH GitHub)
	return os.Getenv("GITHUB_PASSWORD")
}
