package git

import (
	"bytes"
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

func CloneGitRepo(baseDirectory string, repo config.GitRepo) {

	repoPath := filepath.Join(baseDirectory, repo.Destination, repo.Name)

	logger.Context = []string{
		"clone",
	}
	logger.Context = append(logger.Context, repo.Context...)

	if _, err := os.Stat(repoPath); !os.IsNotExist(err) {

		status := checkGitStatus(repoPath)
		if status == 0 {
			logger.Printf("✅ %s", repo.Name)
			return
		} else {
			logger.Printf("⚠️ git status returned non-zero exit code: %d", status)
			return
		}

	}

	var cmd *exec.Cmd
	switch repo.CloneMethod {

	case "https":
		cmd = exec.Command("git", "clone", repo.HttpsUrl, repoPath)
		logger.Print("⌛ Cloning repo (https)")

	case "ssh":
		cmd = exec.Command("git", "clone", repo.SshUrl, repoPath)
		logger.Print("⌛ Cloning repo (ssh)")
	}

	//  git writing directly to stdout causes an issue on windows where ansi colors would stop working
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stdout
	// using a multi writer seems to work around the issue above for some reason
	var out bytes.Buffer
	multi := io.MultiWriter(os.Stdout, &out)
	cmd.Stdout = multi
	cmd.Stderr = multi

	err := cmd.Run()
	if err != nil {
		fmt.Println(strings.Join(cmd.Args, " "))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	logger.Printf("✅ %s", repo.Name)

}
