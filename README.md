# api
[![CI](https://github.com/veritem/api/actions/workflows/ci.yml/badge.svg)](https://github.com/veritem/api/actions/workflows/ci.yml)

My personal api

```
github api query


{
  viewer {
    contributionsCollection {
      totalCommitContributions
      contributionYears
      totalIssueContributions
      totalPullRequestContributions
      startedAt
    }
    pullRequests {
      totalCount
    }
    repositoriesContributedTo {
      totalCount
    }
    issues {
      totalCount
    }
  }
}

```
