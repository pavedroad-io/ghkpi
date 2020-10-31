# ghkpi
Aggregate status for one or more GitHub repositories.
Scope is based on the security of the authenticated user.

You can filter repositories using the -t option and specify a
comma separated list of topics associated with your GitHub
repositories

```bash
ghkpi repo -t one,two,three
```

You can specify a date range with the -r option.  The two options are
**current** or **prior**, for the current or previous month.  Failing to select a 
period defaults to the full repository history

The stats objects hold counters for lifetime and period specified.  Additions and Deletions are line counts.

```bash
ghkpi repo -t one,two,three -r prior
```



```json
{
  "totals": {
    "name": "repositories: [githubStatsKPI, clients, cockroachdb-client, frontend, go-core, integrations, pavedroad, roadctl, scripts, templates]",
    "forks_count": 2,
    "stargazers_count": 10,
    "watchers_count": 10,
    "open_issues_count": 56,
    "subscriber_count": 0,
    "commit_count": 12,
    "pull_created_count": 1,
    "pull_closed_count": 0,
    "stats": {
      "life_time_counts": {
        "lines_added": 1470647,
        "lines_deleted": 182245,
        "commits": 599
      },
      "period_counts": {
        "lines_added": 723841,
        "lines_deleted": 8,
        "commits": 8
      }
    }
  },
  "period": {
    "start_date": "2020-10-01T00:00:00-07:00",
    "end_date": "2020-10-31T00:00:00-07:00"
  },
  "details": [
    {
      "name": "githubStatsKPI",
      "owner": "jscharber",
      "type": "User",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 0,
      "subscriber_count": 0,
      "commit_count": 7,
      "pull_created_count": 1,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 723772,
          "lines_deleted": 0,
          "commits": 4
        },
        "period_counts": {
          "lines_added": 723772,
          "lines_deleted": 0,
          "commits": 4
        }
      }
    },
    {
      "name": "clients",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 1,
      "watchers_count": 1,
      "open_issues_count": 0,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 6064,
          "lines_deleted": 2975,
          "commits": 8
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "cockroachdb-client",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 0,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 213,
          "lines_deleted": 0,
          "commits": 1
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "frontend",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 0,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 55897,
          "lines_deleted": 3757,
          "commits": 16
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "go-core",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 1,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 8408,
          "lines_deleted": 4285,
          "commits": 79
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "integrations",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 0,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 4654,
          "lines_deleted": 1613,
          "commits": 11
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "pavedroad",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 1,
      "stargazers_count": 7,
      "watchers_count": 7,
      "open_issues_count": 2,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 21105,
          "lines_deleted": 11754,
          "commits": 334
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "roadctl",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 1,
      "watchers_count": 1,
      "open_issues_count": 49,
      "subscriber_count": 0,
      "commit_count": 0,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 595896,
          "lines_deleted": 128887,
          "commits": 84
        },
        "period_counts": {
          "lines_added": 0,
          "lines_deleted": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "scripts",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 0,
      "stargazers_count": 0,
      "watchers_count": 0,
      "open_issues_count": 2,
      "subscriber_count": 0,
      "commit_count": 2,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 182,
          "lines_deleted": 16,
          "commits": 4
        },
        "period_counts": {
          "lines_added": 58,
          "lines_deleted": 0,
          "commits": 1
        }
      }
    },
    {
      "name": "templates",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forks_count": 1,
      "stargazers_count": 1,
      "watchers_count": 1,
      "open_issues_count": 2,
      "subscriber_count": 0,
      "commit_count": 3,
      "pull_created_count": 0,
      "pull_closed_count": 0,
      "stats": {
        "life_time_counts": {
          "lines_added": 54456,
          "lines_deleted": 28958,
          "commits": 58
        },
        "period_counts": {
          "lines_added": 11,
          "lines_deleted": 8,
          "commits": 3
        }
      }
    }
  ]
}
```
