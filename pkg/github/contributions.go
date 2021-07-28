package github

import (
	"time"

	"github.com/shurcooL/githubv4"
	"github.com/veritem/api/pkg/graph/model"
	"github.com/veritem/api/pkg/utils"
)

var query struct {
	Viewer struct {
		ContributionsCollection struct {
			TotalCommitContributions      githubv4.Int
			ContributionYears             []githubv4.Int
			TotalIssueContributions       githubv4.Int
			TotalPullRequestContributions githubv4.Int
			StartedAt                     time.Time
		}
		PullRequests struct {
			TotalCount githubv4.Int
		}
		RepositoriesContributedTo struct {
			TotalCount githubv4.Int
		}
		Repositories struct {
			TotalCount githubv4.Int
		}
		Issues struct {
			TotalCount githubv4.Int
		}
	}
}

func Contributions() *model.OpenSource {
	err := Client().Query(Ctx, &query, nil)

	if err != nil {
		return &model.OpenSource{}
	}

	return &model.OpenSource{
		OpenedPr:                int(query.Viewer.PullRequests.TotalCount),
		IssuesSubmitted:         int(query.Viewer.Issues.TotalCount),
		Repositories:            int(query.Viewer.Repositories.TotalCount),
		RepositoriesContributed: int(query.Viewer.RepositoriesContributedTo.TotalCount),
		Started:                 utils.FormatTime(query.Viewer.ContributionsCollection.StartedAt),
	}
}
