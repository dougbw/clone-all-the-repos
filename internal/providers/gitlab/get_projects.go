package gitlab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func getProjects(bearer string) (projects []gitLabUserProject) {

	perPage := 100
	client := &http.Client{}
	var pageProjects []gitLabUserProject
	uri := "https://gitlab.com/api/v4/projects"
	// uri := fmt.Sprintf("https://gitlab.com/api/v4/projects?pagination=keyset&per_page=%d&order_by=id&membership=true&simple=true", perPage)
	page := 0

	// regex to parse the link header
	var pattern = regexp.MustCompile(`\<(.*)\>;\s*rel="next"`)

	for {

		// build request
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Add("Authorization", bearer)
		q := req.URL.Query()
		q.Add("pagination", "keyset")
		q.Add("order_by", "id")
		q.Add("pagination", "keyset")
		q.Add("per_page", fmt.Sprintf("%d", perPage))
		q.Add("membership", "true")
		q.Add("simple", "true")
		req.URL.RawQuery = q.Encode()

		// perform request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		// parse
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &pageProjects)
		if err != nil {
			log.Fatalln(err)
		}

		// append
		projects = append(projects, pageProjects...)
		fmt.Printf("debug: pageIndex=%d, total=%d, pageCount=%d\n", page, len(projects), len(pageProjects))
		if link, ok := resp.Header["Link"]; ok {
			match := pattern.FindStringSubmatch(link[0])
			uri = match[1]
		} else {
			break
		}
		page++
	}

	return
}
