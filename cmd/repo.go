/*
 */
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

var topics = []string{}
var topicsDefault = []string{}
var startDate time.Time
var endDate time.Time
var dateRange string

type repoQuery struct {
	filters []repoFilter
}

type repoFilter struct {
	filter   string
	values   []string
	function interface{}
}

type contributorsCounters struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Commits   int `json:"commits"`
}

type contributorsStats struct {
	LifeTimeCounts contributorsCounters `json:"lifeTimeCounts"`
	PeriodCounts   contributorsCounters `json:"periodCounts"`
}

type repoItem struct {
	Name             string            `json:"name"`
	Owner            string            `json:"owner,omitempty"`
	Type             string            `json:"type,omitempty"`
	ForksCount       int               `json:"forksCount"`
	StargazersCount  int               `json:"stargazersCount"`
	WatchersCount    int               `json:"watchersCount"`
	OpenIssuesCount  int               `json:"openIssuesCount"`
	SubscriberCount  int               `json:"subscriberCount"`
	CommitCount      int               `json:"commitCount"`
	PullCreatedCount int               `json:"pullCreatedCount"`
	PullClosedCount  int               `json:"pullClosedCount"`
	Stats            contributorsStats `json:"stats"`
}

type repoSummary struct {
	Totals          repoItem   `json:"totals"`
	Details         []repoItem `json:"details"`
	name            string
	forksCount      int
	stargazersCount int
	watchersCount   int
}

var repoFilterList []repoFilter
var repoList []repoItem

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Generate statistics for a single repository",
	Long: `Generate statistics for a single repository
	githubStatsKPI repository name, i.e., pavedroad-io`,
	Run: func(cmd *cobra.Command, args []string) {
		var ghClient *github.Client
		var err error

		if dateRange != "" {
			setDateRange()
		}

		f := &repoFilter{}

		repoFilterList = f.init(repoFilterList)

		conf := gitClientConfig{authType: oauth}
		if ghClient, err = getClient(conf); err != nil {
			fmt.Println(err)
		}

		// only lists ones user owns or is a member of in an organization
		// maybe a filter
		lo := github.RepositoryListOptions{Affiliation: "organization_member,owner"}

		// (context, user, options)
		// leaving user ""/blank lists repositories for the currently
		// authenticated user
		repos, _, err := ghClient.Repositories.List(
			context.Background(), "", &lo)
		if err != nil {
			os.Exit(-1)
		}

		filteredRepos := f.filterRepos(repos, repoFilterList)

		// PR don't support date ranges, include those created or closed in the specified date range
		plo := github.PullRequestListOptions{}

		// Commit options support date ranges
		clo := github.CommitsListOptions{}
		if !startDate.IsZero() && !endDate.IsZero() {
			clo.Since = startDate
			clo.Until = endDate
		}
		for i, r := range filteredRepos {
			// Get commit activity
			commits, _, _ := ghClient.Repositories.ListCommits(context.Background(), r.Owner, r.Name, &clo)

			for _, _ = range commits {
				filteredRepos[i].CommitCount += 1
			}

			// Get pull request activity
			pr, _, _ := ghClient.PullRequests.List(context.Background(), r.Owner, r.Name, &plo)
			for i, v := range pr {
				if startDate.IsZero() {
					if v.CreatedAt != nil {
						filteredRepos[i].PullCreatedCount += 1
					}
					if v.ClosedAt != nil {
						filteredRepos[i].PullClosedCount += 1
					}
				} else {
					if v.CreatedAt.After(startDate) && v.CreatedAt.Before(endDate) {
						filteredRepos[i].PullCreatedCount += 1
					}
					if v.ClosedAt != nil && v.ClosedAt.After(startDate) && v.ClosedAt.Before(endDate) {
						filteredRepos[i].PullClosedCount += 1
					}
				}
			}

			// Look at contributors
			contributorActivity, _, err := ghClient.Repositories.ListContributorsStats(context.Background(), r.Owner, r.Name)
			if err != nil {
				fmt.Println("Contributor stats failed", err)
			}

			for _, v := range contributorActivity {
				for _, w := range v.Weeks {
					incActivity(&filteredRepos[i], w)
					incPeriodActivity(&filteredRepos[i], w)
				}
			}

		}

		aggregateResults := summarize(filteredRepos)
		data, err := json.Marshal(aggregateResults)
		if err != nil {
			fmt.Println("JSON marshal failed: ", err)
		}
		fmt.Println(string(data))

	},
}

func incActivity(ri *repoItem, ws github.WeeklyStats) {
	ri.Stats.LifeTimeCounts.Additions += *ws.Additions
	ri.Stats.LifeTimeCounts.Commits += *ws.Commits
	ri.Stats.LifeTimeCounts.Deletions += *ws.Deletions
	return
}

// incPeriodActivity if not date range set or it is within the request date range
func incPeriodActivity(ri *repoItem, ws github.WeeklyStats) {
	if (ws.Week.After(startDate) && ws.Week.Before(endDate)) ||
		(startDate.IsZero() && endDate.IsZero()) {
		ri.Stats.PeriodCounts.Additions += *ws.Additions
		ri.Stats.PeriodCounts.Commits += *ws.Commits
		ri.Stats.PeriodCounts.Deletions += *ws.Deletions
	}
	return
}

func setDateRange() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	switch dateRange {
	case "current":
		startDate = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		endDate = startDate.AddDate(0, 1, -1)
		break
	case "prior":
		startDate = time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
		endDate = startDate.AddDate(0, 1, -1)
		break
	default:
		fmt.Printf("Invalid date range options %v\n", dateRange)
		os.Exit(-1)
	}
}

func summarize(r []repoItem) repoSummary {
	sum := repoSummary{}
	sum.Totals.Name = "Repositories: ["
	sum.Totals.ForksCount = 0
	sum.Totals.OpenIssuesCount = 0
	sum.Totals.StargazersCount = 0
	sum.Totals.WatchersCount = 0
	sum.Totals.SubscriberCount = 0
	sum.Totals.PullClosedCount = 0
	sum.Totals.PullCreatedCount = 0
	sum.Totals.CommitCount = 0

	for i, repo := range r {
		if i == 0 {
			sum.Totals.Name += repo.Name
		} else {
			sum.Totals.Name += ", " + repo.Name
		}
		sum.Totals.ForksCount += repo.ForksCount
		sum.Totals.OpenIssuesCount += repo.OpenIssuesCount
		sum.Totals.StargazersCount += repo.StargazersCount
		sum.Totals.WatchersCount += repo.WatchersCount
		sum.Totals.SubscriberCount += repo.SubscriberCount
		sum.Totals.PullClosedCount += repo.PullClosedCount
		sum.Totals.PullCreatedCount += repo.PullCreatedCount
		sum.Totals.CommitCount += repo.CommitCount
		sum.Totals.Stats.LifeTimeCounts.Additions += repo.Stats.LifeTimeCounts.Additions
		sum.Totals.Stats.LifeTimeCounts.Deletions += repo.Stats.LifeTimeCounts.Deletions
		sum.Totals.Stats.LifeTimeCounts.Commits += repo.Stats.LifeTimeCounts.Commits
		sum.Totals.Stats.PeriodCounts.Additions += repo.Stats.PeriodCounts.Additions
		sum.Totals.Stats.PeriodCounts.Deletions += repo.Stats.PeriodCounts.Deletions
		sum.Totals.Stats.PeriodCounts.Commits += repo.Stats.PeriodCounts.Commits
	}

	sum.Totals.Name += "]"

	sum.Details = r

	return sum
}

func (f *repoFilter) filterRepos(r []*github.Repository, fl []repoFilter) []repoItem {
	items := []repoItem{}

	if len(r) <= 0 {
		fmt.Println("no repos")
		// No repositories passed in return empty list
	} else if len(fl) > 0 {
		for _, v := range r {
			// filter the list
			for _, f := range fl {

				switch f.filter {
				case "topic":
					for _, t := range v.Topics {
						if containsTopic(t, topics) {
							i := repoItem{}
							i.Name = *v.Name
							i.Owner = *v.Owner.Login
							i.Type = *v.Owner.Type
							i.ForksCount = *v.ForksCount
							i.StargazersCount = *v.StargazersCount
							i.WatchersCount = *v.WatchersCount
							i.OpenIssuesCount = *v.OpenIssuesCount
							if v.SubscribersCount != nil {
								i.SubscriberCount = *v.SubscribersCount
							} else {
								i.SubscriberCount = 0
							}

							items = append(items, i)
							break
						}
					}
				}

			}
		}
	} else {
		// return all
		for _, v := range r {
			i := repoItem{}
			i.Name = *v.Name
			i.Owner = *v.Owner.Login
			i.Type = *v.Owner.Type
			i.ForksCount = *v.ForksCount
			i.StargazersCount = *v.StargazersCount
			i.WatchersCount = *v.WatchersCount
			i.OpenIssuesCount = *v.OpenIssuesCount
			if v.SubscribersCount != nil {
				i.SubscriberCount = *v.SubscribersCount
			} else {
				i.SubscriberCount = 0
			}
			items = append(items, i)
		}
	}

	return items
}

func (f *repoFilter) init(fl []repoFilter) []repoFilter {
	// Look for topic filter
	if len(topics) > 0 {
		f := repoFilter{}
		f.filter = "topic"
		f.values = topics
		f.function = containsTopic
		fl = append(fl, f)
	}

	return fl
}

func containsTopic(t string, l []string) bool {
	for _, i := range l {
		if i == t {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(repoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	repoCmd.PersistentFlags().String("foo", "f", "A help for foo")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	repoCmd.Flags().StringVarP(&dateRange, "range", "r", "", "\"current\" or \"prior\" month")
	repoCmd.Flags().StringArrayVarP(&topics, "topics", "t", topicsDefault, "-t help")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
