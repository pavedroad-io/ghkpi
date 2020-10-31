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
    "name": "Repositories: [githubStatsKPI, clients, cockroachdb-client, frontend, go-core, integrations, pavedroad, roadctl, scripts, templates]",
    "forksCount": 2,
    "stargazersCount": 10,
    "watchersCount": 10,
    "openIssuesCount": 56,
    "subscriberCount": 0,
    "commitCount": 12,
    "pullCreatedCount": 1,
    "pullClosedCount": 0,
    "stats": {
      "lifeTimeCounts": {
        "additions": 1470647,
        "deletions": 182245,
        "commits": 599
      },
      "periodCounts": {
        "additions": 723841,
        "deletions": 8,
        "commits": 8
      }
    }
  },
  "details": [
    {
      "name": "githubStatsKPI",
      "owner": "jscharber",
      "type": "User",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 0,
      "subscriberCount": 0,
      "commitCount": 7,
      "pullCreatedCount": 1,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 723772,
          "deletions": 0,
          "commits": 4
        },
        "periodCounts": {
          "additions": 723772,
          "deletions": 0,
          "commits": 4
        }
      }
    },
    {
      "name": "clients",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 1,
      "watchersCount": 1,
      "openIssuesCount": 0,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 6064,
          "deletions": 2975,
          "commits": 8
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "cockroachdb-client",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 0,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 213,
          "deletions": 0,
          "commits": 1
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "frontend",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 0,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 55897,
          "deletions": 3757,
          "commits": 16
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "go-core",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 1,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 8408,
          "deletions": 4285,
          "commits": 79
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "integrations",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 0,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 4654,
          "deletions": 1613,
          "commits": 11
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "pavedroad",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 1,
      "stargazersCount": 7,
      "watchersCount": 7,
      "openIssuesCount": 2,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 21105,
          "deletions": 11754,
          "commits": 334
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "roadctl",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 1,
      "watchersCount": 1,
      "openIssuesCount": 49,
      "subscriberCount": 0,
      "commitCount": 0,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 595896,
          "deletions": 128887,
          "commits": 84
        },
        "periodCounts": {
          "additions": 0,
          "deletions": 0,
          "commits": 0
        }
      }
    },
    {
      "name": "scripts",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 0,
      "stargazersCount": 0,
      "watchersCount": 0,
      "openIssuesCount": 2,
      "subscriberCount": 0,
      "commitCount": 2,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 182,
          "deletions": 16,
          "commits": 4
        },
        "periodCounts": {
          "additions": 58,
          "deletions": 0,
          "commits": 1
        }
      }
    },
    {
      "name": "templates",
      "owner": "pavedroad-io",
      "type": "Organization",
      "forksCount": 1,
      "stargazersCount": 1,
      "watchersCount": 1,
      "openIssuesCount": 2,
      "subscriberCount": 0,
      "commitCount": 3,
      "pullCreatedCount": 0,
      "pullClosedCount": 0,
      "stats": {
        "lifeTimeCounts": {
          "additions": 54456,
          "deletions": 28958,
          "commits": 58
        },
        "periodCounts": {
          "additions": 11,
          "deletions": 8,
          "commits": 3
        }
      }
    }
  ]
}
```
