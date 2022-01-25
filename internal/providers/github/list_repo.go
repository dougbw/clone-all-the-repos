package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

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

func findReposCli(owner string) (githubrepos []githubRepo) {

	out, err := exec.Command("gh", "repo", "list", owner, "--limit", "500", "--json", "name,nameWithOwner,url,sshUrl,defaultBranchRef,isFork,isPrivate").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = json.Unmarshal(out, &githubrepos)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
