package github

import (
	"clone-all-the-repos/internal/execwrapper"
	"clone-all-the-repos/internal/logger"
	"encoding/json"
	"fmt"
)

func findReposCli(owner string) (githubrepos []githubRepo) {

	// find repos
	output, err := execwrapper.Exec("gh", "repo", "list", owner, "--limit", "500", "--json", "name,nameWithOwner,url,sshUrl,defaultBranchRef,isFork,isPrivate")
	if err != nil {
		logger.PrintErr(err.Error())
	}

	err = json.Unmarshal(output.Stdout, &githubrepos)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
