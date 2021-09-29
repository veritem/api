package github

import (
	"github.com/shurcooL/githubv4"
	"github.com/veritem/api/pkg/graph/model"
)

type (
	RepositoryFragment struct {
		ID          githubv4.String
		Name        githubv4.String
		URL         githubv4.String
		Description githubv4.String
		Stargazers  struct {
			TotalCount int
		}
	}
)

var rquery struct {
	Viewer struct {
		PinnedItems struct {
			Nodes []struct {
				RepositoryFragment `graphql:"... on Repository"`
			}
		} `graphql:"pinnedItems(first: 4)"`
	}
}

func LatestProjects() []*model.LatestProjects {
	err := Client().Query(Ctx, &rquery, nil)

	if err != nil {
		return nil
	}

	latestProj := make([]*model.LatestProjects, 0)

	for _, project := range rquery.Viewer.PinnedItems.Nodes {
		latestProj = append(latestProj, &model.LatestProjects{
			ID:          string(project.ID),
			Name:        string(project.Name),
			URL:         string(project.URL),
			Description: string(project.Description),
			Stars:       project.Stargazers.TotalCount,
		})
	}

	return latestProj
}
