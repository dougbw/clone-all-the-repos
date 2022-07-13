package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Root struct {
	AzureDevOps []AzureDevOpsConfig `yaml:"azuredevops"`
	GitHub      []GitHubConfig      `yaml:"github"`
	GitLab      []GitLabConfig      `yaml:"gitlab"`
}

type AzureDevOpsConfig struct {
	Name         string `yaml:"name" validate:"required"`
	Organization string `yaml:"organization" validate:"required"`
	CloneMethod  string `yaml:"cloneMethod" validate:"required"`
	Destination  string `yaml:"destination"`
	Include      string `yaml:"include,omitempty"`
	Exclude      string `yaml:"exclude,omitempty"`
	Cli          bool   `yaml:"cli,omitempty"`
}

type GitHubConfig struct {
	Name        string `yaml:"name" validate:"required"`
	Owner       string `yaml:"owner" validate:"required"`
	CloneMethod string `yaml:"cloneMethod" validate:"required"`
	Destination string `yaml:"destination"`
	Include     string `yaml:"include,omitempty"`
	Exclude     string `yaml:"exclude,omitempty"`
	Cli         bool   `yaml:"cli,omitempty"`
}

type GitLabConfig struct {
	Name        string `yaml:"name" validate:"required"`
	User        string `yaml:"user" validate:"required"`
	CloneMethod string `yaml:"cloneMethod" validate:"required"`
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

func Parse(configfilePath string) (config Root, err error) {

	yamlFile, err := ioutil.ReadFile(configfilePath)
	if err != nil {
		return config, err
	}

	err = yaml.UnmarshalStrict(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return config, err
	}

	return config, nil

}
