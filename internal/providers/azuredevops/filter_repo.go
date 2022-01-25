package azuredevops

import "regexp"

func filterIncludeRepos(pattern string, repos []repo) (ret []repo) {

	for _, repo := range repos {
		match, _ := regexp.MatchString(pattern, repo.Name)
		if match {
			ret = append(ret, repo)
		}
	}
	return
}

func filterExcludeRepos(pattern string, repos []repo) (ret []repo) {

	for _, repo := range repos {
		match, _ := regexp.MatchString(pattern, repo.Name)
		if !match {
			ret = append(ret, repo)
		}
	}
	return
}
