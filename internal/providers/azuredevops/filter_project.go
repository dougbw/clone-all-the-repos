package azuredevops

import "regexp"

func filterIncludeProjects(pattern string, projects []project) (ret []project) {

	for _, project := range projects {
		match, _ := regexp.MatchString(pattern, project.Name)
		if match {
			ret = append(ret, project)
		}
	}
	return
}

func filterExcludeProjects(pattern string, projects []project) (ret []project) {

	for _, project := range projects {
		match, _ := regexp.MatchString(pattern, project.Name)
		if !match {
			ret = append(ret, project)
		}
	}
	return
}
