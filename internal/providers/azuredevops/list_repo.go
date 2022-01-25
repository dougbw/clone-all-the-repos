package azuredevops

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// repos
type repoResponse struct {
	Value []repo `json:"value"`
	Count int    `json:"count"`
}
type repo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Project struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		State          string    `json:"state"`
		Revision       int       `json:"revision"`
		Visibility     string    `json:"visibility"`
		LastUpdateTime time.Time `json:"lastUpdateTime"`
	} `json:"project"`
	DefaultBranch string `json:"defaultBranch,omitempty"`
	Size          int    `json:"size"`
	RemoteURL     string `json:"remoteUrl"`
	SSHURL        string `json:"sshUrl"`
	WebURL        string `json:"webUrl"`
	IsDisabled    bool   `json:"isDisabled"`
}

func listRepos(token string, org string, project string) (repos []repo) {
	// request
	client := &http.Client{}
	uri := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/git/repositories?api-version=%s", org, project, apiVersion)
	req, _ := http.NewRequest("GET", uri, nil)
	req.SetBasicAuth("", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// parse
	body, _ := ioutil.ReadAll(resp.Body)
	var repoResponse repoResponse
	err = json.Unmarshal(body, &repoResponse)
	if err != nil {
		log.Fatalln(err)
	}

	repos = repoResponse.Value
	return
}

func listReposCli(org string, project string) (repos []repo) {

	organizationUrl := fmt.Sprintf("https://dev.azure.com/%s", org)

	out, err := exec.Command("az", "repos", "list", "--organization", organizationUrl, "--project", project).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = json.Unmarshal(out, &repos)
	if err != nil {
		fmt.Println(err)
	}

	return
}
