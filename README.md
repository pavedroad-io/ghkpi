[![Go Report Card](https://goreportcard.com/badge/github.com/pavedroad-io/ghkpi)](https://goreportcard.com/report/github.com/pavedroad-io/ghkpi)

# ghkpi
Aggregate status for one or more GitHub repositories.
Scope is based on the security of the authenticated user.

You can filter repositories using the -t option and specify a
comma separated list of topics associated with your GitHub
repositories

- [Get it](#get-it)
- [Use it](#use-it)
- [Licensing](#licensing)

## Get it

```bash
go get -u github.com/pavedroad/ghkpi/ghkpi
```

Or [download the binary](https://github.com/pavedroad-io/ghkpi/releases/latest) from the releases page.

## Use it
```bash
Aggregate statistics for a group of GitHub repositories

Usage:
  ghkpi repo [flags]

Flags:
  -a, --aggregate_totals     Only output Aggregate totals
  -e, --end_date string      RFC3339 date, -e "2020-02-28T23:59:59Z"
  -h, --help                 help for repo
  -r, --range string         "current" or "prior" the current or prior month respectively
  -s, --start_date string    RFC3339 date, -s "2020-01-01T00:00:00Z"
  -t, --topics stringArray   -t topic1,topic2,topic3

Global Flags:
      --config string   config file (default is $HOME/.ghkpi.yaml)
```

### Filtering

#### By topics
Use the -t options to filter repositories based on GitHub associated topics.  You can specify more than one topic by separating them with commas.

```bash
$ ghkpi repo -t one,two,three
```

#### By simple date range
You can specify a date range with the -r option.  The two options are
**current** or **prior**, for the current or previous month.  
Failing to select a period defaults to the full repository history

```bash
$ ghkpi repo -t one,two,three -r prior
```

#### By flexible date range


```bash
$ ghkpi repo -t pr-kpi -s "2020-01-01T00:00:00Z" -e "2020-02-28T23:59:59Z"
```

### Only output totals
By default, the summary totals and details for each repository are output.
You can ask for only the aggregated totals to be output with the -a option.

```bash
$ ghkpi repo -t pr-kpi -a -s "2020-01-01T00:00:00Z" -e "2020-02-28T23:59:59Z"
```

# Output columns

## Totals
Includes the list of repositories that makeup the aggregate totals.  Along with top-level details like pull requests, issues, stars and watchers.

It also includes two status objects that hold counters for lifetime of 
all repositories and the period specified using the -r or -s/e options.  

Additions and Deletions are line counts.

It include as list of contributors and the totals number of combined commits


```json
{
  "name": "repositories: [clients, cockroachdb-client, frontend, ghkpi, go-core, integrations, pavedroad, roadctl, scripts, templates]",
  "forks_count": 2,
  "stargazers_count": 11,
  "watchers_count": 11,
  "open_issues_count": 62,
  "subscriber_count": 0,
  "commit_count": 13,
  "pull_created_count": 1,
  "pull_closed_count": 0,
  "stats": {
    "life_time_counts": {
      "period": {
        "start_date": "2019-03-10T20:02:56Z",
        "end_date": "2020-11-12T13:09:16.04317166-08:00"
      },
      "lines_added": 1123504,
      "lines_deleted": 287225,
      "commits": 602,
      "contributor_count": 8,
      "contributor_list": [
        "jscharber",
        "MarkGreenPR",
        "CandaceScharber",
        "jscharbervs",
        "agstaunton",
        "capgar",
        "rick4106t",
        "shgayle"
      ]
    },
    "period_counts": {
      "period": {
        "start_date": "2020-10-01T00:00:00-07:00",
        "end_date": "2020-10-31T00:00:00-07:00"
      },
      "lines_added": 376688,
      "lines_deleted": 104988,
      "commits": 10,
      "contributor_count": 1,
      "contributor_list": [
        "jscharber"
      ]
    }
  }
}
```

## Details
By default, the output also includes the same information for each 
repository in the details array.

```json
  "details": [
    {
      "name": "clients",
      "owner": "pavedroad-io",
....
},
{....}
]
```


## Licensing

- Any original code is licensed under the [Apache 2 License](./LICENSE).
