package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Root struct {
	AzureDevOps []AzureDevOpsConfig `json:"azuredevops"`
	GitHub      []GitHubConfig      `json:"github"`
	GitLab      []GitLabConfig      `json:"gitlab"`
}

type AzureDevOpsConfig struct {
	Name         string `yaml:"name"`
	Organization string `yaml:"organization"`
	CloneMethod  string `yaml:"cloneMethod"`
	Destination  string `yaml:"destination"`
	Include      string `yaml:"include,omitempty"`
	Exclude      string `yaml:"exclude,omitempty"`
	Cli          bool   `yaml:"cli,omitempty"`
}

type GitHubConfig struct {
	Name        string `yaml:"name"`
	Owner       string `yaml:"owner"`
	CloneMethod string `yaml:"cloneMethod"`
	Destination string `yaml:"destination"`
	Include     string `yaml:"include,omitempty"`
	Exclude     string `yaml:"exclude,omitempty"`
	Cli         bool   `yaml:"cli,omitempty"`
}

type GitLabConfig struct {
	Name        string `yaml:"name"`
	User        string `yaml:"user"`
	CloneMethod string `yaml:"cloneMethod"`
	Destination string `yaml:"destination"`
	Include     string `yaml:"include,omitempty"`
	Exclude     string `yaml:"exclude,omitempty"`
}

type GitRepo struct {
	Name        string   `json:"name"`
	HttpsUrl    string   `json:"httpsUrl"`
	SshUrl      string   `json:"sshUrl"`
	CloneMethod string   `json:"cloneMethod"`
	Destination string   `json:"destination"`
	Context     []string `json:"context,omitempty"`
}

func Open(configfilePath string) Root {

	jsonFile, err := os.Open(configfilePath)
	defer jsonFile.Close()

	if err != nil {
		//fmt.Println(err)
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Root
	json.Unmarshal(byteValue, &config)
	return config

}

func OpenYaml(configfilePath string) (config Root) {

	yamlFile, err := ioutil.ReadFile(configfilePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config

}
