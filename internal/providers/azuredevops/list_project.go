package azuredevops

import (
	"clone-all-the-repos/internal/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// projects
type projectResponse struct {
	Count int       `json:"count"`
	Value []project `json:"value"`
}
type project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Revision       int       `json:"revision"`
	Visibility     string    `json:"visibility"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}

type projectCliResponse struct {
	ContinuationToken string    `json:"continuationToken"`
	Value             []project `json:"value"`
}

func listProjects(token string, org string) (projects []project) {

	loggingContext := []string{
		"provider:azuredevops",
		"discovery:api",
		fmt.Sprintf("org:%s", org),
	}

	const top = 30

	// setup
	client := &http.Client{}
	uri := fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=%s&$top=%d", org, apiVersion, top)

	for {

		// request
		req, _ := http.NewRequest("GET", uri, nil)
		req.SetBasicAuth("", token)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode != 200 {
			message := fmt.Sprintf("Unexpected status code in response: %s", resp.Status)
			logger.PrintErrMessage(loggingContext, message)
		}

		// parse
		body, _ := ioutil.ReadAll(resp.Body)
		var projectResponse projectResponse
		err = json.Unmarshal(body, &projectResponse)
		if err != nil {
			log.Fatalln(err)
		}

		// append results
		projects = append(projects, projectResponse.Value...)

		// check continuationToken
		continuationToken := resp.Header["X-Ms-Continuationtoken"]
		if continuationToken == nil {
			break
		} else {
			uri = fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=%s&$top=%d&continuationToken=%s", org, apiVersion, top, continuationToken[0])
		}

	}

	return
}

func listProjectsCli(org string) (projects []project) {

	const top string = "30"
	organizationUrl := fmt.Sprintf("https://dev.azure.com/%s", org)
	cmd := exec.Command("az", "devops", "project", "list", "--organization", organizationUrl, "--top", top, "--output", "json")

	for {

		out, err := cmd.Output()
		if err != nil {
			log.Fatalf("error: %s", err.(*exec.ExitError).Stderr)
		}

		var projectCliResponse projectCliResponse
		err = json.Unmarshal(out, &projectCliResponse)
		if err != nil {
			//fmt.Println(string(out))
			//fmt.Println(strings.Join(cmd.Args, " "))
			log.Fatal("error: unable to parse az cli output")
		}

		// append results
		projects = append(projects, projectCliResponse.Value...)

		// check continuationToken
		continuationToken := projectCliResponse.ContinuationToken
		if continuationToken == "" {
			break
		} else {
			cmd = exec.Command("az", "devops", "project", "list", "--organization", organizationUrl, "--top", top, "--continuation-token", continuationToken)
		}

	}

	return
}
