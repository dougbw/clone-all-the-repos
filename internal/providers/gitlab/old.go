package gitlab

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// todo:
// * pagination
// * filtering

func getUserId(bearer string, username string) (userId int) {

	// request
	client := &http.Client{}
	uri := fmt.Sprintf("https://gitlab.com/api/v4/users?username=%s", username)
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", bearer)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// parse
	body, _ := ioutil.ReadAll(resp.Body)
	var user GitLabUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalln(err)
	}
	userId = user[0].ID
	return
}

func getGroups(bearer string) (groups GitLabGroup) {

	client := &http.Client{}
	uri := "https://gitlab.com/api/v4/groups"
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &groups)
	if err != nil {
		log.Fatalln(err)
	}

	return

}

func getGroupProjects(bearer string, groupId int) (projects GitLabGroupProject) {

	client := &http.Client{}
	uri := fmt.Sprintf("https://gitlab.com/api/v4/groups/%d/projects", groupId)
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", bearer)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func getUserProjects(bearer string, userId int) (projects []gitLabUserProject) {

	client := &http.Client{}
	uri := fmt.Sprintf("https://gitlab.com/api/v4/users/%d/projects/", userId)
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", bearer)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))

	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetRepos1(GitLabConfig config.GitLabConfig) (Repos []config.GitRepo) {

	logger.Context = []string{
		"provider:gitlab",
		fmt.Sprintf("config:%s", GitLabConfig.Name),
		fmt.Sprintf("user:%s", GitLabConfig.User),
	}
	logger.Print("🔍 Finding projects")

	token, present := os.LookupEnv("GITLAB_TOKEN")
	if !present {
		logger.Print("⛔ required environment variable not set: GITLAB_TOKEN")
		os.Exit(2)
	}

	bearer := fmt.Sprintf("Bearer %s", token)

	allProjects := getProjects(bearer)
	fmt.Print(allProjects)

	// get user ID
	userId := getUserId(bearer, GitLabConfig.User)

	// find user projects
	userProjects := getUserProjects(bearer, userId)
	for _, project := range userProjects {

		var httpsUrl string
		if GitLabConfig.CloneMethod == "https" {
			authHttps := fmt.Sprintf("https://%s:%s@", GitLabConfig.User, token)
			httpsUrl = strings.Replace(project.HTTPURLToRepo, "https://", authHttps, -1)
		} else {
			httpsUrl = project.HTTPURLToRepo
		}

		var GitRepo = config.GitRepo{
			Context:     logger.Context,
			Name:        project.Name,
			HttpsUrl:    httpsUrl,
			SshUrl:      project.SSHURLToRepo,
			CloneMethod: GitLabConfig.CloneMethod,
			Destination: GitLabConfig.Destination,
		}
		Repos = append(Repos, GitRepo)
	}

	// find groups
	logger.Context = []string{
		"provider:gitlab",
		fmt.Sprintf("config:%s", GitLabConfig.Name),
		fmt.Sprintf("user:%s", GitLabConfig.User),
	}
	logger.Print("🔍 Finding groups")
	groups := getGroups(bearer)
	for _, group := range groups {

		logger.Context = []string{
			"provider:gitlab",
			fmt.Sprintf("config:%s", GitLabConfig.Name),
			fmt.Sprintf("user:%s", GitLabConfig.User),
			fmt.Sprintf("group:%s", group.Name),
		}
		logger.Print("🔍 Finding projects in group")

		// find group projects
		groupProjects := getGroupProjects(bearer, group.ID)
		for _, project := range groupProjects {

			var httpsUrl string
			if GitLabConfig.CloneMethod == "https" {
				authHttps := fmt.Sprintf("https://%s:%s@", GitLabConfig.User, token)
				httpsUrl = strings.Replace(project.HTTPURLToRepo, "https://", authHttps, -1)
			} else {
				httpsUrl = project.HTTPURLToRepo
			}

			fullName := strings.Replace(project.NameWithNamespace, " ", "", -1)
			destination := path.Join(GitLabConfig.Destination, fullName)
			destinationDir, _ := filepath.Split(destination)

			var GitRepo = config.GitRepo{
				Context:     logger.Context,
				Name:        project.Name,
				HttpsUrl:    httpsUrl,
				SshUrl:      project.SSHURLToRepo,
				CloneMethod: GitLabConfig.CloneMethod,
				Destination: destinationDir,
			}
			Repos = append(Repos, GitRepo)

		}

	}

	return
}
