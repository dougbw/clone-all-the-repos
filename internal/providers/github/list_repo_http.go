package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func findRepos(owner string, token string) (githubrepos []githubRepo) {

	// setup client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	// setup graphql query
	var repoQuery struct {
		RepositoryOwner struct {
			Login githubv4.String

			Repositories struct {
				Nodes []githubRepo

				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"repositories(first: 100, after: $reposCursor, ownerAffiliations: OWNER)"`
		} `graphql:"repositoryOwner(login: $owner)"`
	}

	// setup graphql variables
	variables := map[string]interface{}{
		"owner":       githubv4.String(owner),
		"reposCursor": (*githubv4.String)(nil),
	}

	// execute query
	for {

		err := client.Query(context.Background(), &repoQuery, variables)
		if err != nil {
			fmt.Println(err)
		}

		githubrepos = append(githubrepos, repoQuery.RepositoryOwner.Repositories.Nodes...)

		if !repoQuery.RepositoryOwner.Repositories.PageInfo.HasNextPage {
			break
		}

		variables["reposCursor"] = githubv4.NewString(repoQuery.RepositoryOwner.Repositories.PageInfo.EndCursor)

	}

	return

}
