package github

type githubRepo struct {
	Name      string `graphql:"name"`
	Url       string `graphql:"url"`
	SSHUrl    string `graphql:"sshUrl"`
	IsPrivate bool   `graphql:"isPrivate"`
	IsFork    bool   `graphql:"isFork"`
}
