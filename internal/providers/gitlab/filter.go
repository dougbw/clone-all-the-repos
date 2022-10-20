package gitlab

import (
	"regexp"
)

func filterIncludeProjects(pattern string, repos []gitLabUserProject) (ret []gitLabUserProject) {

	for _, repo := range repos {
		match, _ := regexp.MatchString(pattern, repo.PathWithNamespace)
		if match {
			ret = append(ret, repo)
		}
	}
	return
}

func filterExcludeProjects(pattern string, repos []gitLabUserProject) (ret []gitLabUserProject) {

	for _, repo := range repos {
		match, _ := regexp.MatchString(pattern, repo.PathWithNamespace)
		if !match {
			ret = append(ret, repo)
		}
	}
	return
}
