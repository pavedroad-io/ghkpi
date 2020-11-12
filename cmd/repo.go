// Package cmd A Cobra CLI for generating  GitHum
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
var oldestRepo time.Time = time.Now()

// Command line option holders
var endDateArgument string
var startDateArgument string
var totalsOnly bool
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
	Period           Period   `json:"period"`
	Additions        int      `json:"lines_added"`
	Deletions        int      `json:"lines_deleted"`
	Commits          int      `json:"commits"`
	ContributorCount int      `json:"contributor_count"`
	ContributorList  []string `json:"contributor_list"`
}

type contributorsStats struct {
	LifeTimeCounts contributorsCounters `json:"life_time_counts"`
	PeriodCounts   contributorsCounters `json:"period_counts"`
}

type repoItem struct {
	Name             string            `json:"name"`
	Owner            string            `json:"owner,omitempty"`
	Type             string            `json:"type,omitempty"`
	ForksCount       int               `json:"forks_count"`
	StargazersCount  int               `json:"stargazers_count"`
	WatchersCount    int               `json:"watchers_count"`
	OpenIssuesCount  int               `json:"open_issues_count"`
	SubscriberCount  int               `json:"subscriber_count"`
	CommitCount      int               `json:"commit_count"`
	PullCreatedCount int               `json:"pull_created_count"`
	PullClosedCount  int               `json:"pull_closed_count"`
	Stats            contributorsStats `json:"stats"`
}

// Period contains the start and end dates for this KPI period
type Period struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type repoSummary struct {
	Totals  repoItem   `json:"totals"`
	Details []repoItem `json:"details"`
}

var repoFilterList []repoFilter
var repoList []repoItem

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Aggregate statistics for a group of GitHub repositories",
	Long:  `Aggregate statistics for a group of GitHub repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		var ghClient *github.Client
		var err error

		if dateRange != "" {
			setDateRange()
		}

		if startDateArgument != "" {
			if endDateArgument == "" {
				fmt.Println("Usage: both a start and end date are required, -s date -e date")
			}
			setDynamicDateRange(startDateArgument, endDateArgument)
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

			for range commits {
				filteredRepos[i].CommitCount++
			}

			// Get pull request activity
			pr, _, _ := ghClient.PullRequests.List(context.Background(), r.Owner, r.Name, &plo)
			for i, v := range pr {
				if startDate.IsZero() {
					if v.CreatedAt != nil {
						filteredRepos[i].PullCreatedCount++
					}
					if v.ClosedAt != nil {
						filteredRepos[i].PullClosedCount++
					}
				} else {
					if v.CreatedAt.After(startDate) && v.CreatedAt.Before(endDate) {
						filteredRepos[i].PullCreatedCount++
					}
					if v.ClosedAt != nil && v.ClosedAt.After(startDate) && v.ClosedAt.Before(endDate) {
						filteredRepos[i].PullClosedCount++
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
					incPeriodContributors(&filteredRepos[i], w,
						v.GetAuthor().GetLogin())
					incActivityContributors(&filteredRepos[i], w,
						v.GetAuthor().GetLogin())
				}
			}

		}

		aggregateResults := summarize(filteredRepos)

		var data []byte
		if totalsOnly {
			data, err = json.Marshal(aggregateResults.Totals)
		} else {
			data, err = json.Marshal(aggregateResults)
		}

		if err != nil {
			fmt.Println("JSON marshal failed: ", err)
		}
		fmt.Println(string(data))

	},
}

func incActivityContributors(ri *repoItem, ws github.WeeklyStats, contributor string) {
	if !listContains(contributor, ri.Stats.LifeTimeCounts.ContributorList) && *ws.Commits > 0 {
		ri.Stats.LifeTimeCounts.ContributorList = append(
			ri.Stats.LifeTimeCounts.ContributorList, contributor)
		ri.Stats.LifeTimeCounts.ContributorCount++
	}

	return
}

func incPeriodContributors(ri *repoItem, ws github.WeeklyStats, contributor string) {
	if (ws.Week.After(startDate) && ws.Week.Before(endDate)) ||
		(startDate.IsZero() && endDate.IsZero()) {
		if !listContains(contributor, ri.Stats.PeriodCounts.ContributorList) && *ws.Commits > 0 {

			//fmt.Println(contributor, ri.Stats.PeriodCounts.ContributorList)
			ri.Stats.PeriodCounts.ContributorList = append(
				ri.Stats.PeriodCounts.ContributorList, contributor)
			ri.Stats.PeriodCounts.ContributorCount++
			//fmt.Println(contributor, ri.Stats.PeriodCounts.ContributorList)
			//fmt.Println(ws)
		}
	}
	return
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

func setDynamicDateRange(start, end string) {
	dt, err := time.Parse(time.RFC3339, start)
	if err != nil {
		fmt.Println("Start date error: ", err)
		os.Exit(-1)
	}
	startDate = dt

	dt, err = time.Parse(time.RFC3339, end)
	if err != nil {
		fmt.Println("End date error: ", err)
		os.Exit(-1)
	}
	endDate = dt
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
	sum.Totals.Name = "repositories: ["

	sum.Totals.Stats.PeriodCounts.Period.StartDate = startDate
	if endDate.IsZero() {
		sum.Totals.Stats.PeriodCounts.Period.EndDate = time.Now()
	} else {
		sum.Totals.Stats.PeriodCounts.Period.EndDate = endDate
	}
	sum.Totals.Stats.LifeTimeCounts.Period.StartDate = oldestRepo
	sum.Totals.Stats.LifeTimeCounts.Period.EndDate = time.Now()

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

		for _, c := range repo.Stats.PeriodCounts.ContributorList {
			if !listContains(c, sum.Totals.Stats.PeriodCounts.ContributorList) {
				sum.Totals.Stats.PeriodCounts.ContributorList = append(
					sum.Totals.Stats.PeriodCounts.ContributorList, c)
			}
			sum.Totals.Stats.PeriodCounts.ContributorCount = len(sum.Totals.Stats.PeriodCounts.ContributorList)
		}

		for _, c := range repo.Stats.LifeTimeCounts.ContributorList {
			if !listContains(c, sum.Totals.Stats.LifeTimeCounts.ContributorList) {
				sum.Totals.Stats.LifeTimeCounts.ContributorList = append(
					sum.Totals.Stats.LifeTimeCounts.ContributorList, c)
			}
			sum.Totals.Stats.LifeTimeCounts.ContributorCount = len(sum.Totals.Stats.LifeTimeCounts.ContributorList)
		}

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
						if listContains(t, topics) {
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

							//fmt.Println("oldest: ", oldestRepo, "repo date: ", v.CreatedAt.Time)
							if v.CreatedAt.Time.Before(oldestRepo) {
								oldestRepo = v.CreatedAt.Time
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
		f.function = listContains
		fl = append(fl, f)
	}

	return fl
}

func listContains(t string, l []string) bool {
	for _, i := range l {
		if i == t {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(repoCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "f", "A help for foo")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	repoCmd.Flags().StringVarP(&dateRange, "range", "r", "", "\"current\" or \"prior\" the current or prior month respectively")
	repoCmd.Flags().StringArrayVarP(&topics, "topics", "t", topicsDefault, "-t topic1,topic2,topic3")
	repoCmd.Flags().StringVarP(&startDateArgument, "start_date", "s", "", "RFC3339 date, -s \"2020-01-01T00:00:00Z\"")
	repoCmd.Flags().StringVarP(&endDateArgument, "end_date", "e", "", "RFC3339 date, -e \"2020-02-28T23:59:59Z\"")
	repoCmd.Flags().BoolVarP(&totalsOnly, "aggregate_totals", "a", false, "Only output Aggregate totals")

}
